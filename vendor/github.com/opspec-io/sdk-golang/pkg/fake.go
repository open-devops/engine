// Code generated by counterfeiter. DO NOT EDIT.
package pkg

import (
	"context"
	"net/url"
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	CreateStub        func(path, pkgName, pkgDescription string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		path           string
		pkgName        string
		pkgDescription string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	GetManifestStub        func(pkgHandle model.PkgHandle) (*model.PkgManifest, error)
	getManifestMutex       sync.RWMutex
	getManifestArgsForCall []struct {
		pkgHandle model.PkgHandle
	}
	getManifestReturns struct {
		result1 *model.PkgManifest
		result2 error
	}
	getManifestReturnsOnCall map[int]struct {
		result1 *model.PkgManifest
		result2 error
	}
	InstallStub        func(ctx context.Context, path string, pkgHandle model.PkgHandle) error
	installMutex       sync.RWMutex
	installArgsForCall []struct {
		ctx       context.Context
		path      string
		pkgHandle model.PkgHandle
	}
	installReturns struct {
		result1 error
	}
	installReturnsOnCall map[int]struct {
		result1 error
	}
	ListOpsStub        func(dirPath string) ([]*model.PkgManifest, error)
	listOpsMutex       sync.RWMutex
	listOpsArgsForCall []struct {
		dirPath string
	}
	listOpsReturns struct {
		result1 []*model.PkgManifest
		result2 error
	}
	listOpsReturnsOnCall map[int]struct {
		result1 []*model.PkgManifest
		result2 error
	}
	NewFSProviderStub        func(basePaths ...string) Provider
	newFSProviderMutex       sync.RWMutex
	newFSProviderArgsForCall []struct {
		basePaths []string
	}
	newFSProviderReturns struct {
		result1 Provider
	}
	newFSProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	NewGitProviderStub        func(basePath string, pullCreds *model.PullCreds) Provider
	newGitProviderMutex       sync.RWMutex
	newGitProviderArgsForCall []struct {
		basePath  string
		pullCreds *model.PullCreds
	}
	newGitProviderReturns struct {
		result1 Provider
	}
	newGitProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	NewNodeProviderStub        func(apiBaseURL url.URL, pullCreds *model.PullCreds) Provider
	newNodeProviderMutex       sync.RWMutex
	newNodeProviderArgsForCall []struct {
		apiBaseURL url.URL
		pullCreds  *model.PullCreds
	}
	newNodeProviderReturns struct {
		result1 Provider
	}
	newNodeProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	ResolveStub        func(ctx context.Context, pkgRef string, providers ...Provider) (model.PkgHandle, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		ctx       context.Context
		pkgRef    string
		providers []Provider
	}
	resolveReturns struct {
		result1 model.PkgHandle
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 model.PkgHandle
		result2 error
	}
	ValidateStub        func(pkgHandle model.PkgHandle) []error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		pkgHandle model.PkgHandle
	}
	validateReturns struct {
		result1 []error
	}
	validateReturnsOnCall map[int]struct {
		result1 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Create(path string, pkgName string, pkgDescription string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		path           string
		pkgName        string
		pkgDescription string
	}{path, pkgName, pkgDescription})
	fake.recordInvocation("Create", []interface{}{path, pkgName, pkgDescription})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(path, pkgName, pkgDescription)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *Fake) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Fake) CreateArgsForCall(i int) (string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].path, fake.createArgsForCall[i].pkgName, fake.createArgsForCall[i].pkgDescription
}

func (fake *Fake) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) GetManifest(pkgHandle model.PkgHandle) (*model.PkgManifest, error) {
	fake.getManifestMutex.Lock()
	ret, specificReturn := fake.getManifestReturnsOnCall[len(fake.getManifestArgsForCall)]
	fake.getManifestArgsForCall = append(fake.getManifestArgsForCall, struct {
		pkgHandle model.PkgHandle
	}{pkgHandle})
	fake.recordInvocation("GetManifest", []interface{}{pkgHandle})
	fake.getManifestMutex.Unlock()
	if fake.GetManifestStub != nil {
		return fake.GetManifestStub(pkgHandle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getManifestReturns.result1, fake.getManifestReturns.result2
}

func (fake *Fake) GetManifestCallCount() int {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return len(fake.getManifestArgsForCall)
}

func (fake *Fake) GetManifestArgsForCall(i int) model.PkgHandle {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return fake.getManifestArgsForCall[i].pkgHandle
}

func (fake *Fake) GetManifestReturns(result1 *model.PkgManifest, result2 error) {
	fake.GetManifestStub = nil
	fake.getManifestReturns = struct {
		result1 *model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) GetManifestReturnsOnCall(i int, result1 *model.PkgManifest, result2 error) {
	fake.GetManifestStub = nil
	if fake.getManifestReturnsOnCall == nil {
		fake.getManifestReturnsOnCall = make(map[int]struct {
			result1 *model.PkgManifest
			result2 error
		})
	}
	fake.getManifestReturnsOnCall[i] = struct {
		result1 *model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) Install(ctx context.Context, path string, pkgHandle model.PkgHandle) error {
	fake.installMutex.Lock()
	ret, specificReturn := fake.installReturnsOnCall[len(fake.installArgsForCall)]
	fake.installArgsForCall = append(fake.installArgsForCall, struct {
		ctx       context.Context
		path      string
		pkgHandle model.PkgHandle
	}{ctx, path, pkgHandle})
	fake.recordInvocation("Install", []interface{}{ctx, path, pkgHandle})
	fake.installMutex.Unlock()
	if fake.InstallStub != nil {
		return fake.InstallStub(ctx, path, pkgHandle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.installReturns.result1
}

func (fake *Fake) InstallCallCount() int {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return len(fake.installArgsForCall)
}

func (fake *Fake) InstallArgsForCall(i int) (context.Context, string, model.PkgHandle) {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return fake.installArgsForCall[i].ctx, fake.installArgsForCall[i].path, fake.installArgsForCall[i].pkgHandle
}

func (fake *Fake) InstallReturns(result1 error) {
	fake.InstallStub = nil
	fake.installReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) InstallReturnsOnCall(i int, result1 error) {
	fake.InstallStub = nil
	if fake.installReturnsOnCall == nil {
		fake.installReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.installReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) ListOps(dirPath string) ([]*model.PkgManifest, error) {
	fake.listOpsMutex.Lock()
	ret, specificReturn := fake.listOpsReturnsOnCall[len(fake.listOpsArgsForCall)]
	fake.listOpsArgsForCall = append(fake.listOpsArgsForCall, struct {
		dirPath string
	}{dirPath})
	fake.recordInvocation("ListOps", []interface{}{dirPath})
	fake.listOpsMutex.Unlock()
	if fake.ListOpsStub != nil {
		return fake.ListOpsStub(dirPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listOpsReturns.result1, fake.listOpsReturns.result2
}

func (fake *Fake) ListOpsCallCount() int {
	fake.listOpsMutex.RLock()
	defer fake.listOpsMutex.RUnlock()
	return len(fake.listOpsArgsForCall)
}

func (fake *Fake) ListOpsArgsForCall(i int) string {
	fake.listOpsMutex.RLock()
	defer fake.listOpsMutex.RUnlock()
	return fake.listOpsArgsForCall[i].dirPath
}

func (fake *Fake) ListOpsReturns(result1 []*model.PkgManifest, result2 error) {
	fake.ListOpsStub = nil
	fake.listOpsReturns = struct {
		result1 []*model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) ListOpsReturnsOnCall(i int, result1 []*model.PkgManifest, result2 error) {
	fake.ListOpsStub = nil
	if fake.listOpsReturnsOnCall == nil {
		fake.listOpsReturnsOnCall = make(map[int]struct {
			result1 []*model.PkgManifest
			result2 error
		})
	}
	fake.listOpsReturnsOnCall[i] = struct {
		result1 []*model.PkgManifest
		result2 error
	}{result1, result2}
}

func (fake *Fake) NewFSProvider(basePaths ...string) Provider {
	fake.newFSProviderMutex.Lock()
	ret, specificReturn := fake.newFSProviderReturnsOnCall[len(fake.newFSProviderArgsForCall)]
	fake.newFSProviderArgsForCall = append(fake.newFSProviderArgsForCall, struct {
		basePaths []string
	}{basePaths})
	fake.recordInvocation("NewFSProvider", []interface{}{basePaths})
	fake.newFSProviderMutex.Unlock()
	if fake.NewFSProviderStub != nil {
		return fake.NewFSProviderStub(basePaths...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newFSProviderReturns.result1
}

func (fake *Fake) NewFSProviderCallCount() int {
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	return len(fake.newFSProviderArgsForCall)
}

func (fake *Fake) NewFSProviderArgsForCall(i int) []string {
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	return fake.newFSProviderArgsForCall[i].basePaths
}

func (fake *Fake) NewFSProviderReturns(result1 Provider) {
	fake.NewFSProviderStub = nil
	fake.newFSProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewFSProviderReturnsOnCall(i int, result1 Provider) {
	fake.NewFSProviderStub = nil
	if fake.newFSProviderReturnsOnCall == nil {
		fake.newFSProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newFSProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewGitProvider(basePath string, pullCreds *model.PullCreds) Provider {
	fake.newGitProviderMutex.Lock()
	ret, specificReturn := fake.newGitProviderReturnsOnCall[len(fake.newGitProviderArgsForCall)]
	fake.newGitProviderArgsForCall = append(fake.newGitProviderArgsForCall, struct {
		basePath  string
		pullCreds *model.PullCreds
	}{basePath, pullCreds})
	fake.recordInvocation("NewGitProvider", []interface{}{basePath, pullCreds})
	fake.newGitProviderMutex.Unlock()
	if fake.NewGitProviderStub != nil {
		return fake.NewGitProviderStub(basePath, pullCreds)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newGitProviderReturns.result1
}

func (fake *Fake) NewGitProviderCallCount() int {
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	return len(fake.newGitProviderArgsForCall)
}

func (fake *Fake) NewGitProviderArgsForCall(i int) (string, *model.PullCreds) {
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	return fake.newGitProviderArgsForCall[i].basePath, fake.newGitProviderArgsForCall[i].pullCreds
}

func (fake *Fake) NewGitProviderReturns(result1 Provider) {
	fake.NewGitProviderStub = nil
	fake.newGitProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewGitProviderReturnsOnCall(i int, result1 Provider) {
	fake.NewGitProviderStub = nil
	if fake.newGitProviderReturnsOnCall == nil {
		fake.newGitProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newGitProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewNodeProvider(apiBaseURL url.URL, pullCreds *model.PullCreds) Provider {
	fake.newNodeProviderMutex.Lock()
	ret, specificReturn := fake.newNodeProviderReturnsOnCall[len(fake.newNodeProviderArgsForCall)]
	fake.newNodeProviderArgsForCall = append(fake.newNodeProviderArgsForCall, struct {
		apiBaseURL url.URL
		pullCreds  *model.PullCreds
	}{apiBaseURL, pullCreds})
	fake.recordInvocation("NewNodeProvider", []interface{}{apiBaseURL, pullCreds})
	fake.newNodeProviderMutex.Unlock()
	if fake.NewNodeProviderStub != nil {
		return fake.NewNodeProviderStub(apiBaseURL, pullCreds)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newNodeProviderReturns.result1
}

func (fake *Fake) NewNodeProviderCallCount() int {
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	return len(fake.newNodeProviderArgsForCall)
}

func (fake *Fake) NewNodeProviderArgsForCall(i int) (url.URL, *model.PullCreds) {
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	return fake.newNodeProviderArgsForCall[i].apiBaseURL, fake.newNodeProviderArgsForCall[i].pullCreds
}

func (fake *Fake) NewNodeProviderReturns(result1 Provider) {
	fake.NewNodeProviderStub = nil
	fake.newNodeProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewNodeProviderReturnsOnCall(i int, result1 Provider) {
	fake.NewNodeProviderStub = nil
	if fake.newNodeProviderReturnsOnCall == nil {
		fake.newNodeProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newNodeProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) Resolve(ctx context.Context, pkgRef string, providers ...Provider) (model.PkgHandle, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		ctx       context.Context
		pkgRef    string
		providers []Provider
	}{ctx, pkgRef, providers})
	fake.recordInvocation("Resolve", []interface{}{ctx, pkgRef, providers})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(ctx, pkgRef, providers...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resolveReturns.result1, fake.resolveReturns.result2
}

func (fake *Fake) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *Fake) ResolveArgsForCall(i int) (context.Context, string, []Provider) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return fake.resolveArgsForCall[i].ctx, fake.resolveArgsForCall[i].pkgRef, fake.resolveArgsForCall[i].providers
}

func (fake *Fake) ResolveReturns(result1 model.PkgHandle, result2 error) {
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 model.PkgHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) ResolveReturnsOnCall(i int, result1 model.PkgHandle, result2 error) {
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 model.PkgHandle
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 model.PkgHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) Validate(pkgHandle model.PkgHandle) []error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		pkgHandle model.PkgHandle
	}{pkgHandle})
	fake.recordInvocation("Validate", []interface{}{pkgHandle})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(pkgHandle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *Fake) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *Fake) ValidateArgsForCall(i int) model.PkgHandle {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].pkgHandle
}

func (fake *Fake) ValidateReturns(result1 []error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) ValidateReturnsOnCall(i int, result1 []error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 []error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	fake.listOpsMutex.RLock()
	defer fake.listOpsMutex.RUnlock()
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ Pkg = new(Fake)
