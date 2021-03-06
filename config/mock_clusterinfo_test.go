// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/config (interfaces: ClusterInfo)

package config_test

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockClusterInfo struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClusterInfo() *MockClusterInfo {
	return &MockClusterInfo{fail: pegomock.GlobalFailHandler}
}

func (mock *MockClusterInfo) ActiveClusterOrDie() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterInfo().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ActiveClusterOrDie", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterInfo) AllClustersOrDie() []string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterInfo().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AllClustersOrDie", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem()})
	var ret0 []string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
	}
	return ret0
}

func (mock *MockClusterInfo) AllNamespacesOrDie() []string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterInfo().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AllNamespacesOrDie", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem()})
	var ret0 []string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
	}
	return ret0
}

func (mock *MockClusterInfo) VerifyWasCalledOnce() *VerifierClusterInfo {
	return &VerifierClusterInfo{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClusterInfo) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierClusterInfo {
	return &VerifierClusterInfo{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClusterInfo) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierClusterInfo {
	return &VerifierClusterInfo{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClusterInfo) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierClusterInfo {
	return &VerifierClusterInfo{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierClusterInfo struct {
	mock                   *MockClusterInfo
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierClusterInfo) ActiveClusterOrDie() *ClusterInfo_ActiveClusterOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ActiveClusterOrDie", params, verifier.timeout)
	return &ClusterInfo_ActiveClusterOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterInfo_ActiveClusterOrDie_OngoingVerification struct {
	mock              *MockClusterInfo
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterInfo_ActiveClusterOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterInfo_ActiveClusterOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterInfo) AllClustersOrDie() *ClusterInfo_AllClustersOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AllClustersOrDie", params, verifier.timeout)
	return &ClusterInfo_AllClustersOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterInfo_AllClustersOrDie_OngoingVerification struct {
	mock              *MockClusterInfo
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterInfo_AllClustersOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterInfo_AllClustersOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterInfo) AllNamespacesOrDie() *ClusterInfo_AllNamespacesOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AllNamespacesOrDie", params, verifier.timeout)
	return &ClusterInfo_AllNamespacesOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterInfo_AllNamespacesOrDie_OngoingVerification struct {
	mock              *MockClusterInfo
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterInfo_AllNamespacesOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterInfo_AllNamespacesOrDie_OngoingVerification) GetAllCapturedArguments() {
}
