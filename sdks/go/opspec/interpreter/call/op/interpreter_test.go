package op

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/data"
	. "github.com/opctl/opctl/sdks/go/data/fakes"
	"github.com/opctl/opctl/sdks/go/data/provider"
	. "github.com/opctl/opctl/sdks/go/data/provider/fakes"
	uniquestringFakes "github.com/opctl/opctl/sdks/go/internal/uniquestring/fakes"
	"github.com/opctl/opctl/sdks/go/model"
	modelFakes "github.com/opctl/opctl/sdks/go/model/fakes"
	inputsFakes "github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/inputs/fakes"
	dirFakes "github.com/opctl/opctl/sdks/go/opspec/interpreter/dir/fakes"
	strFakes "github.com/opctl/opctl/sdks/go/opspec/interpreter/str/fakes"
	. "github.com/opctl/opctl/sdks/go/opspec/opfile/fakes"
)

var _ = Context("Interpreter", func() {
	Context("NewInterpreter", func() {
		It("shouldn't return nil", func() {
			/* arrange/act/assert */
			Expect(NewInterpreter("")).To(Not(BeNil()))
		})
	})
	Context("Interpret", func() {
		Context("called w/ opspec ../../test-suite scenarios", func() {
			It("should return result fulfilling scenario.call.expect", func() {
				tempDir, err := ioutil.TempDir("", "")
				if nil != err {
					panic(err)
				}
				rootPath := "../../../../../../test-suite"

				filepath.Walk(rootPath,
					func(path string, info os.FileInfo, err error) error {
						if info.IsDir() {
							scenariosOpFilePath := filepath.Join(path, "scenarios.yml")
							if _, err := os.Stat(scenariosOpFilePath); nil == err {
								/* arrange */
								absOpPath, err := filepath.Abs(path)
								if nil != err {
									panic(fmt.Errorf("error getting absOpPath %v; error was %v", path, err))
								}

								data := data.New()
								opHandle, err := data.Resolve(
									context.Background(),
									absOpPath,
									data.NewFSProvider(),
								)
								if nil != err {
									panic(fmt.Errorf("error resolving pkg for %v; error was %v", path, err))
								}

								scenariosOpFileBytes, err := ioutil.ReadFile(scenariosOpFilePath)
								if nil != err {
									panic(err)
								}

								var scenarioOpFile []struct {
									Name      string
									Interpret *struct {
										Expect string
										Scope  map[string]*model.Value
									}
								}
								if err := yaml.Unmarshal(scenariosOpFileBytes, &scenarioOpFile); nil != err {
									panic(fmt.Errorf("error unmarshalling scenario.yml for %v; error was %v", path, err))
								}

								for _, scenario := range scenarioOpFile {
									if nil != scenario.Interpret {
										scgOpCall := &model.SCGOpCall{
											Ref:    absOpPath,
											Inputs: map[string]interface{}{},
										}

										for name := range scenario.Interpret.Scope {
											// map as passed
											scgOpCall.Inputs[name] = ""
										}

										/* act */
										objectUnderTest := NewInterpreter(tempDir)
										_, actualErr := objectUnderTest.Interpret(
											scenario.Interpret.Scope,
											scgOpCall,
											"",
											*opHandle.Path(),
											"",
										)

										/* assert */
										description := fmt.Sprintf("pkg: '%v'\nscenario: '%v'", path, scenario.Name)
										switch expect := scenario.Interpret.Expect; expect {
										case "success":
											Expect(actualErr).To(BeNil(), description)
										case "failure":
											Expect(actualErr).To(Not(BeNil()), description)
										}
									}
								}
							}
						}
						return nil
					})
			})
		})
		It("should call data.NewFSProvider w/ expected args", func() {
			/* arrange */
			providedOpPath := "providedOpPath"

			fakeData := new(FakeData)
			// error to trigger immediate return
			fakeData.ResolveReturns(nil, errors.New("dummyError"))

			objectUnderTest := _interpreter{
				data: fakeData,
			}

			/* act */
			objectUnderTest.Interpret(
				map[string]*model.Value{},
				&model.SCGOpCall{
					Ref: "dummyOpRef",
				},
				"dummyOpID",
				providedOpPath,
				"dummyRootOpID",
			)

			/* assert */
			Expect(fakeData.NewFSProviderArgsForCall(0)).
				To(ConsistOf(providedOpPath, filepath.Dir(providedOpPath)))
		})
		Context("scgOpCall.PullCreds is nil", func() {
			It("should call data.NewGitProvider w/ expected args", func() {
				/* arrange */
				providedDataCachePath := "dummyDataCachePath"

				fakeData := new(FakeData)
				// error to trigger immediate return
				fakeData.ResolveReturns(nil, errors.New("dummyError"))

				objectUnderTest := _interpreter{
					data:          fakeData,
					dataCachePath: providedDataCachePath,
				}

				/* act */
				objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGOpCall{
						Ref: "dummyOpRef",
					},
					"dummyOpID",
					"dummyOpPath",
					"dummyRootOpID",
				)

				/* assert */
				actualBasePath,
					actualPullCreds := fakeData.NewGitProviderArgsForCall(0)

				Expect(actualBasePath).To(Equal(providedDataCachePath))
				Expect(actualPullCreds).To(BeNil())
			})
		})
		Context("scgOpCall.PullCreds isn't nil", func() {
			Context("stringInterpreter.Interpret errs", func() {
				It("should return expected result", func() {
					/* arrange */
					fakeStrInterpreter := new(strFakes.FakeInterpreter)
					interpretError := errors.New("dummyError")
					fakeStrInterpreter.InterpretReturns(nil, interpretError)

					objectUnderTest := _interpreter{
						stringInterpreter: fakeStrInterpreter,
					}

					/* act */
					_, actualError := objectUnderTest.Interpret(
						map[string]*model.Value{},
						&model.SCGOpCall{
							PullCreds: &model.SCGPullCreds{},
						},
						"dummyOpID",
						"dummyOpPath",
						"dummyRootOpID",
					)

					/* assert */
					Expect(actualError).To(Equal(interpretError))
				})
			})
			Context("stringInterpreter.Interpret doesn't err", func() {
				It("should call data.NewGitProvider w/ expected args", func() {
					/* arrange */
					providedDataCachePath := "dummyDataCachePath"

					fakeStrInterpreter := new(strFakes.FakeInterpreter)
					expectedPullCreds := &model.PullCreds{Username: "dummyUsername", Password: "dummyPassword"}
					fakeStrInterpreter.InterpretReturnsOnCall(0, &model.Value{String: &expectedPullCreds.Username}, nil)
					fakeStrInterpreter.InterpretReturnsOnCall(1, &model.Value{String: &expectedPullCreds.Password}, nil)

					fakeData := new(FakeData)
					// error to trigger immediate return
					fakeData.ResolveReturns(nil, errors.New("dummyError"))

					objectUnderTest := _interpreter{
						stringInterpreter: fakeStrInterpreter,
						data:              fakeData,
						dataCachePath:     providedDataCachePath,
					}

					/* act */
					objectUnderTest.Interpret(
						map[string]*model.Value{},
						&model.SCGOpCall{
							Ref:       "dummyOpRef",
							PullCreds: &model.SCGPullCreds{},
						},
						"dummyOpID",
						"dummyOpPath",
						"dummyRootOpID",
					)

					/* assert */
					actualBasePath,
						actualPullCreds := fakeData.NewGitProviderArgsForCall(0)

					Expect(actualBasePath).To(Equal(providedDataCachePath))
					Expect(actualPullCreds).To(Equal(expectedPullCreds))
				})
			})
		})
		Context("scgOpCall.Ref is value referency", func() {
			It("should call opfile.Get w/ expected args", func() {
				/* arrange */
				opPath := "opPath"

				fakeDirInterpreter := new(dirFakes.FakeInterpreter)
				fakeDirInterpreter.InterpretReturns(&model.Value{Dir: &opPath}, nil)

				fakeOpFileGetter := new(FakeGetter)
				// err to trigger immediate return
				fakeOpFileGetter.GetReturns(nil, errors.New("dummyErr"))

				objectUnderTest := _interpreter{
					dirInterpreter:      fakeDirInterpreter,
					opFileGetter:        fakeOpFileGetter,
					uniqueStringFactory: new(uniquestringFakes.FakeUniqueStringFactory),
				}

				/* act */
				objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGOpCall{
						Ref: "$(dummyVar)",
					},
					"dummyOpID",
					"dummyOpPath",
					"dummyRootOpID",
				)

				/* assert */
				actualCtx,
					actualOpPath := fakeOpFileGetter.GetArgsForCall(0)

				Expect(actualCtx).To(Equal(context.TODO()))
				Expect(actualOpPath).To(Equal(opPath))
			})
		})
		Context("scgOpCall.Ref not value referency", func() {
			It("should call data.Resolve w/ expected args", func() {
				/* arrange */
				providedOpPath := "providedOpPath"

				provideddataDirPath := "dummydataDirPath"
				providedSCGOpCall := &model.SCGOpCall{
					Ref: providedOpPath,
				}

				fakeData := new(FakeData)

				expectedPkgProviders := []provider.Provider{
					new(FakeProvider),
					new(FakeProvider),
				}
				fakeData.NewFSProviderReturns(expectedPkgProviders[0])
				fakeData.NewGitProviderReturns(expectedPkgProviders[1])

				// error to trigger immediate return
				fakeData.ResolveReturns(nil, errors.New("dummyError"))

				objectUnderTest := _interpreter{
					data:          fakeData,
					dataCachePath: filepath.Join(provideddataDirPath, "ops"),
				}

				/* act */
				objectUnderTest.Interpret(
					map[string]*model.Value{},
					providedSCGOpCall,
					"dummyOpID",
					providedOpPath,
					"dummyRootOpID",
				)

				/* assert */
				actualCtx,
					actualOpRef,
					actualPkgProviders := fakeData.ResolveArgsForCall(0)

				Expect(actualCtx).To(Equal(context.TODO()))
				Expect(actualOpRef).To(Equal(providedOpPath))
				Expect(actualPkgProviders).To(Equal(expectedPkgProviders))
			})
			Context("data.Resolve errs", func() {
				It("should return err", func() {
					/* arrange */
					expectedErr := errors.New("dummyError")
					fakeData := new(FakeData)
					fakeData.ResolveReturns(nil, expectedErr)

					objectUnderTest := _interpreter{
						data:                fakeData,
						uniqueStringFactory: new(uniquestringFakes.FakeUniqueStringFactory),
					}

					/* act */
					_, actualErr := objectUnderTest.Interpret(
						map[string]*model.Value{},
						&model.SCGOpCall{},
						"dummyOpID",
						"dummyOpPath",
						"dummyRootOpID",
					)

					/* assert */
					Expect(actualErr).To(Equal(expectedErr))
				})
			})
		})
		It("should call opfile.Get w/ expected args", func() {
			/* arrange */
			fakeDataHandle := new(modelFakes.FakeDataHandle)
			opPath := "opPath"
			fakeDataHandle.PathReturns(&opPath)

			fakeData := new(FakeData)
			fakeData.ResolveReturns(fakeDataHandle, nil)

			fakeOpFileGetter := new(FakeGetter)
			// err to trigger immediate return
			fakeOpFileGetter.GetReturns(nil, errors.New("dummyErr"))

			objectUnderTest := _interpreter{
				data:                fakeData,
				opFileGetter:        fakeOpFileGetter,
				uniqueStringFactory: new(uniquestringFakes.FakeUniqueStringFactory),
			}

			/* act */
			objectUnderTest.Interpret(
				map[string]*model.Value{},
				&model.SCGOpCall{},
				"dummyOpID",
				"dummyOpPath",
				"dummyRootOpID",
			)

			/* assert */
			actualCtx,
				actualOpPath := fakeOpFileGetter.GetArgsForCall(0)

			Expect(actualCtx).To(Equal(context.TODO()))
			Expect(actualOpPath).To(Equal(opPath))
		})
		Context("opfile.Get errs", func() {
			It("should return err", func() {
				/* arrange */
				expectedErr := errors.New("dummyError")
				fakeOpFileGetter := new(FakeGetter)
				fakeOpFileGetter.GetReturns(nil, expectedErr)

				fakeDataHandle := new(modelFakes.FakeDataHandle)
				fakeDataHandle.PathReturns(new(string))

				fakeData := new(FakeData)
				fakeData.ResolveReturns(fakeDataHandle, nil)

				objectUnderTest := _interpreter{
					data:                fakeData,
					opFileGetter:        fakeOpFileGetter,
					uniqueStringFactory: new(uniquestringFakes.FakeUniqueStringFactory),
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					&model.SCGOpCall{},
					"dummyOpID",
					"dummyOpPath",
					"dummyRootOpID",
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		It("should call inputsFakes.Interpret w/ expected inputs", func() {
			/* arrange */
			providedScope := map[string]*model.Value{
				"dummyScopeRef1Name": {String: new(string)},
			}
			expectedScope := providedScope

			expectedInputArgs := map[string]interface{}{"dummySCGOpCallInputName": "dummyScgOpCallInputValue"}

			providedSCGOpCall := &model.SCGOpCall{
				Inputs: expectedInputArgs,
			}

			providedOpID := "dummyOpID"

			providedParentOpPath := "providedParentOpPath"

			fakeDataHandle := new(modelFakes.FakeDataHandle)
			opPath := "dummyOpPath"
			fakeDataHandle.PathReturns(&opPath)

			fakeData := new(FakeData)
			fakeData.ResolveReturns(fakeDataHandle, nil)

			expectedInputParams := map[string]*model.Param{
				"dummyParam1Name": {String: &model.StringParam{}},
			}

			fakeOpFileGetter := new(FakeGetter)
			returnedManifest := &model.OpFile{
				Inputs: expectedInputParams,
			}
			fakeOpFileGetter.GetReturns(returnedManifest, nil)

			fakeInputsInterpreter := new(inputsFakes.FakeInterpreter)

			dcgScratchDir := "dummyDCGScratchDir"

			objectUnderTest := _interpreter{
				dcgScratchDir:       dcgScratchDir,
				data:                fakeData,
				opFileGetter:        fakeOpFileGetter,
				uniqueStringFactory: new(uniquestringFakes.FakeUniqueStringFactory),
				inputsInterpreter:   fakeInputsInterpreter,
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedSCGOpCall,
				providedOpID,
				providedParentOpPath,
				"dummyRootOpID",
			)

			/* assert */
			actualSCGArgs,
				actualSCGInputs,
				actualOpPath,
				actualScope,
				actualOpScratchDir := fakeInputsInterpreter.InterpretArgsForCall(0)

			Expect(actualScope).To(Equal(expectedScope))
			Expect(actualSCGArgs).To(Equal(expectedInputArgs))
			Expect(actualOpPath).To(Equal(opPath))
			Expect(actualSCGInputs).To(Equal(expectedInputParams))
			Expect(actualOpScratchDir).To(Equal(filepath.Join(dcgScratchDir, providedOpID)))

		})
	})
})
