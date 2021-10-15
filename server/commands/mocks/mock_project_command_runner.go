// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/commands (interfaces: ProjectCommandRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockProjectCommandRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectCommandRunner(options ...pegomock.Option) *MockProjectCommandRunner {
	mock := &MockProjectCommandRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectCommandRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectCommandRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectCommandRunner) Plan(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Plan", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) Apply(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Apply", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) PolicyCheck(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PolicyCheck", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) ApprovePolicies(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ApprovePolicies", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) Version(ctx models.ProjectCommandContext) models.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Version", params, []reflect.Type{reflect.TypeOf((*models.ProjectResult)(nil)).Elem()})
	var ret0 models.ProjectResult
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ProjectResult)
		}
	}
	return ret0
}

func (mock *MockProjectCommandRunner) VerifyWasCalledOnce() *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectCommandRunner struct {
	mock                   *MockProjectCommandRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectCommandRunner) Plan(ctx models.ProjectCommandContext) *MockProjectCommandRunner_Plan_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Plan", params, verifier.timeout)
	return &MockProjectCommandRunner_Plan_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Plan_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Plan_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Plan_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) Apply(ctx models.ProjectCommandContext) *MockProjectCommandRunner_Apply_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Apply", params, verifier.timeout)
	return &MockProjectCommandRunner_Apply_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Apply_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Apply_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Apply_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) PolicyCheck(ctx models.ProjectCommandContext) *MockProjectCommandRunner_PolicyCheck_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PolicyCheck", params, verifier.timeout)
	return &MockProjectCommandRunner_PolicyCheck_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_PolicyCheck_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_PolicyCheck_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_PolicyCheck_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) ApprovePolicies(ctx models.ProjectCommandContext) *MockProjectCommandRunner_ApprovePolicies_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ApprovePolicies", params, verifier.timeout)
	return &MockProjectCommandRunner_ApprovePolicies_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_ApprovePolicies_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_ApprovePolicies_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_ApprovePolicies_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) Version(ctx models.ProjectCommandContext) *MockProjectCommandRunner_Version_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Version", params, verifier.timeout)
	return &MockProjectCommandRunner_Version_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Version_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Version_OngoingVerification) GetCapturedArguments() models.ProjectCommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Version_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
	}
	return
}