package expression

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/expression/interpolater"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg"
	"github.com/pkg/errors"
)

var _ = Context("EvalToObject", func() {
	var _ = Context("EvalToObject", func() {
		Context("expression is map[string]interface{}", func() {
			It("should return expected result", func() {
				/* arrange */
				providedExpression := map[string]interface{}{"dummyName": "dummyValue"}

				objectUnderTest := _objectEvaluator{}

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToObject(
					map[string]*model.Value{},
					providedExpression,
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(*actualValue).To(Equal(model.Value{Object: providedExpression}))
				Expect(actualErr).To(BeNil())
			})
		})
		Context("expression is string", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				providedScope := map[string]*model.Value{"dummyName": {}}
				providedExpression := "dummyExpression"
				providedPkgRef := new(pkg.FakeHandle)

				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				objectUnderTest := _objectEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToObject(
					providedScope,
					providedExpression,
					providedPkgRef,
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(providedExpression))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgRef).To(Equal(providedPkgRef))

			})
			Context("interpolater.Interpolate errs", func() {
				It("should return expected err", func() {
					/* arrange */
					fakeInterpolater := new(interpolater.Fake)
					interpolateErr := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", interpolateErr)

					objectUnderTest := _objectEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					_, actualErr := objectUnderTest.EvalToObject(
						map[string]*model.Value{},
						"dummyExpression",
						new(pkg.FakeHandle),
					)

					/* assert */
					Expect(actualErr).To(Equal(interpolateErr))

				})
			})
			Context("interpolater.Interpolate doesn't err", func() {
				It("should call data.CoerceToObject w/ expected args & return result", func() {
					/* arrange */
					fakeInterpolater := new(interpolater.Fake)

					interpolatedValue := "dummyString"
					fakeInterpolater.InterpolateReturns(interpolatedValue, nil)

					fakeData := new(data.Fake)

					coercedValue := model.Value{Object: map[string]interface{}{"dummyName": "dummyValue"}}
					fakeData.CoerceToObjectReturns(&coercedValue, nil)

					objectUnderTest := _objectEvaluator{
						data:         fakeData,
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToObject(
						map[string]*model.Value{},
						"dummyExpression",
						new(pkg.FakeHandle),
					)

					/* assert */
					Expect(*fakeData.CoerceToObjectArgsForCall(0)).To(Equal(model.Value{String: &interpolatedValue}))

					Expect(*actualValue).To(Equal(coercedValue))
					Expect(actualErr).To(BeNil())
				})
			})
		})
	})
})