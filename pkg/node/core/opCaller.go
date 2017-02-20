package core

//go:generate counterfeiter -o ./fakeOpCaller.go --fake-name fakeOpCaller ./ opCaller

import (
	"bytes"
	"fmt"
	"github.com/opspec-io/opctl/util/pubsub"
	"github.com/opspec-io/opctl/util/uniquestring"
	"github.com/opspec-io/sdk-golang/pkg/bundle"
	"github.com/opspec-io/sdk-golang/pkg/model"
	"github.com/opspec-io/sdk-golang/pkg/validate"
	"time"
)

type opCaller interface {
	// Executes an op call
	Call(
		inboundScope map[string]*model.Data,
		opId string,
		opRef string,
		opGraphId string,
	) (
		outboundScope map[string]*model.Data,
		err error,
	)
}

func newOpCaller(
	bundle bundle.Bundle,
	pubSub pubsub.PubSub,
	dcgNodeRepo dcgNodeRepo,
	caller caller,
	uniqueStringFactory uniquestring.UniqueStringFactory,
	validate validate.Validate,
) opCaller {
	return _opCaller{
		bundle:              bundle,
		pubSub:              pubSub,
		dcgNodeRepo:         dcgNodeRepo,
		caller:              caller,
		uniqueStringFactory: uniqueStringFactory,
		validate:            validate,
	}
}

type _opCaller struct {
	bundle              bundle.Bundle
	pubSub              pubsub.PubSub
	dcgNodeRepo         dcgNodeRepo
	caller              caller
	uniqueStringFactory uniquestring.UniqueStringFactory
	validate            validate.Validate
}

func (this _opCaller) Call(
	inboundScope map[string]*model.Data,
	opId string,
	opRef string,
	opGraphId string,
) (
	outboundScope map[string]*model.Data,
	err error,
) {
	defer func() {
		// defer must be defined before conditional return statements so it always runs

		if nil == this.dcgNodeRepo.GetIfExists(opGraphId) {
			// guard: op killed (we got preempted)
			this.pubSub.Publish(
				&model.Event{
					Timestamp: time.Now().UTC(),
					OpEnded: &model.OpEndedEvent{
						OpId:      opId,
						Outcome:   model.OpOutcomeKilled,
						OpGraphId: opGraphId,
						OpRef:     opRef,
					},
				},
			)
			return
		}

		this.dcgNodeRepo.DeleteIfExists(opId)

		var opOutcome string
		if nil != err {
			this.pubSub.Publish(
				&model.Event{
					Timestamp: time.Now().UTC(),
					OpEncounteredError: &model.OpEncounteredErrorEvent{
						Msg:       err.Error(),
						OpId:      opId,
						OpRef:     opRef,
						OpGraphId: opGraphId,
					},
				},
			)
			opOutcome = model.OpOutcomeFailed
		} else {
			opOutcome = model.OpOutcomeSucceeded
		}

		this.pubSub.Publish(
			&model.Event{
				Timestamp: time.Now().UTC(),
				OpEnded: &model.OpEndedEvent{
					OpId:      opId,
					OpRef:     opRef,
					Outcome:   opOutcome,
					OpGraphId: opGraphId,
				},
			},
		)

	}()

	this.dcgNodeRepo.Add(
		&dcgNodeDescriptor{
			Id:        opId,
			OpRef:     opRef,
			OpGraphId: opGraphId,
			Op:        &dcgOpDescriptor{},
		},
	)

	op, err := this.bundle.GetOp(
		opRef,
	)
	if nil != err {
		return
	}

	// validate inputs
	err = this.validateScope("input", inboundScope, op.Inputs)
	if nil != err {
		return
	}

	this.pubSub.Publish(
		&model.Event{
			Timestamp: time.Now().UTC(),
			OpStarted: &model.OpStartedEvent{
				OpId:      opId,
				OpRef:     opRef,
				OpGraphId: opGraphId,
			},
		},
	)

	outboundScope, err = this.caller.Call(
		this.uniqueStringFactory.Construct(),
		inboundScope,
		op.Run,
		opRef,
		opGraphId,
	)
	if nil != err {
		return
	}

	// validate outputs
	err = this.validateScope("output", outboundScope, op.Outputs)

	return

}

func (this _opCaller) validateScope(
	scopeType string,
	scope map[string]*model.Data,
	params map[string]*model.Param,
) error {

	messageBuffer := bytes.NewBufferString(``)
	for paramName, param := range params {
		var (
			arg             *model.Data
			argDisplayValue string
		)

		// resolve var for param
		var ok bool
		switch {
		case nil != param.Dir:
			if arg, ok = scope[paramName]; ok {
				argDisplayValue = arg.Dir
			}
		case nil != param.File:
			if arg, ok = scope[paramName]; ok {
				argDisplayValue = arg.File
			}
		case nil != param.Socket:
			if arg, ok = scope[paramName]; ok {
				argDisplayValue = arg.Socket
			}
		case nil != param.String:
			if arg, ok = scope[paramName]; ok {
				argDisplayValue = arg.String
			} else {
				// fallback to default
				arg = &model.Data{String: param.String.Default}
				argDisplayValue = arg.String
			}
			if param.String.IsSecret {
				argDisplayValue = "************"
			}
		}

		// validate
		validationErrors := this.validate.Param(arg, param)

		if len(validationErrors) > 0 {
			messageBuffer.WriteString(fmt.Sprintf(`
  Name: %v
  Value: %v
  Error(s):`, paramName, argDisplayValue),
			)
			for _, validationError := range validationErrors {
				messageBuffer.WriteString(fmt.Sprintf(`
    - %v`, validationError.Error()))
			}
			messageBuffer.WriteString(`
`)
		}
	}

	if messageBuffer.Len() > 0 {
		return fmt.Errorf(`
-
  validation of the following op %v(s) failed:
%v
-`, scopeType, messageBuffer.String())
	}
	return nil
}
