package core

import (
	"context"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/outputs"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
)

//counterfeiter:generate -o internal/fakes/opCaller.go . opCaller
type opCaller interface {
	// Executes an op call
	Call(
		ctx context.Context,
		opCall *model.OpCall,
		inboundScope map[string]*model.Value,
		parentCallID *string,
		rootCallID string,
		opCallSpec *model.OpCallSpec,
	) (
		map[string]*model.Value,
		error,
	)
}

func newOpCaller(
	caller caller,
	dataDirPath string,
) opCaller {
	return _opCaller{
		caller:         caller,
		callScratchDir: filepath.Join(dataDirPath, "call"),
	}
}

type _opCaller struct {
	caller         caller
	callScratchDir string
}

func (oc _opCaller) Call(
	ctx context.Context,
	opCall *model.OpCall,
	inboundScope map[string]*model.Value,
	parentCallID *string,
	rootCallID string,
	opCallSpec *model.OpCallSpec,
) (
	map[string]*model.Value,
	error,
) {
	var err error
	outboundScope := map[string]*model.Value{}

	// form scope for op call by combining defined inputs & op dir
	opCallScope := map[string]*model.Value{}
	for varName, varData := range opCall.Inputs {
		opCallScope[varName] = varData
	}
	// add deprecated absolute path to scope
	opCallScope["/"] = &model.Value{
		Link: &opCall.OpPath,
	}
	// add current directory to scope
	opCallScope["./"] = &model.Value{
		Link: &opCall.OpPath,
	}

	// add parent directory to scope
	parentDirPath := filepath.Dir(opCall.OpPath)
	opCallScope["../"] = &model.Value{
		Link: &parentDirPath,
	}

	opOutputs, err := oc.caller.Call(
		ctx,
		opCall.ChildCallID,
		opCallScope,
		opCall.ChildCallCallSpec,
		opCall.OpPath,
		&opCall.OpID,
		rootCallID,
	)
	if nil != err {
		return outboundScope, err
	}

	var opFile *model.OpSpec
	opFile, err = opfile.Get(
		ctx,
		opCall.OpPath,
	)
	if nil != err {
		return outboundScope, err
	}
	opOutputs, err = outputs.Interpret(
		opOutputs,
		opFile.Outputs,
		opCall.OpPath,
		filepath.Join(oc.callScratchDir, opCall.OpID),
	)

	// filter op outboundScope to bound call outboundScope
	for boundName, boundValue := range opCallSpec.Outputs {
		// return bound outboundScope
		if "" == boundValue {
			// implicit value
			boundValue = boundName
		} else if !regexp.MustCompile("^\\$\\(.+\\)$").MatchString(boundValue) {
			// handle obsolete syntax by swapping order
			prevBoundName := boundName
			boundName = boundValue
			boundValue = prevBoundName
		} else {
			boundValue = strings.TrimSuffix(strings.TrimPrefix(boundValue, "$("), ")")
		}
		for opOutputName, opOutputValue := range opOutputs {
			if boundName == opOutputName {
				outboundScope[boundValue] = opOutputValue
			}
		}
	}

	return outboundScope, err
}
