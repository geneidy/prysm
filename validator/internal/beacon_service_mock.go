// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: BeaconServiceClient,BeaconService_LatestAttestationClient,BeaconService_WaitForChainStartClient)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	v10 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockBeaconServiceClient is a mock of BeaconServiceClient interface
type MockBeaconServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconServiceClientMockRecorder
}

// MockBeaconServiceClientMockRecorder is the mock recorder for MockBeaconServiceClient
type MockBeaconServiceClientMockRecorder struct {
	mock *MockBeaconServiceClient
}

// NewMockBeaconServiceClient creates a new mock instance
func NewMockBeaconServiceClient(ctrl *gomock.Controller) *MockBeaconServiceClient {
	mock := &MockBeaconServiceClient{ctrl: ctrl}
	mock.recorder = &MockBeaconServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconServiceClient) EXPECT() *MockBeaconServiceClientMockRecorder {
	return m.recorder
}

// CanonicalHead mocks base method
func (m *MockBeaconServiceClient) CanonicalHead(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*v1.BeaconBlock, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CanonicalHead", varargs...)
	ret0, _ := ret[0].(*v1.BeaconBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanonicalHead indicates an expected call of CanonicalHead
func (mr *MockBeaconServiceClientMockRecorder) CanonicalHead(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanonicalHead", reflect.TypeOf((*MockBeaconServiceClient)(nil).CanonicalHead), varargs...)
}

// Eth1Data mocks base method
func (m *MockBeaconServiceClient) Eth1Data(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*v10.Eth1DataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Eth1Data", varargs...)
	ret0, _ := ret[0].(*v10.Eth1DataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eth1Data indicates an expected call of Eth1Data
func (mr *MockBeaconServiceClientMockRecorder) Eth1Data(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eth1Data", reflect.TypeOf((*MockBeaconServiceClient)(nil).Eth1Data), varargs...)
}

// ForkData mocks base method
func (m *MockBeaconServiceClient) ForkData(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*v1.Fork, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ForkData", varargs...)
	ret0, _ := ret[0].(*v1.Fork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForkData indicates an expected call of ForkData
func (mr *MockBeaconServiceClientMockRecorder) ForkData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForkData", reflect.TypeOf((*MockBeaconServiceClient)(nil).ForkData), varargs...)
}

// LatestAttestation mocks base method
func (m *MockBeaconServiceClient) LatestAttestation(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (v10.BeaconService_LatestAttestationClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LatestAttestation", varargs...)
	ret0, _ := ret[0].(v10.BeaconService_LatestAttestationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestAttestation indicates an expected call of LatestAttestation
func (mr *MockBeaconServiceClientMockRecorder) LatestAttestation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestAttestation", reflect.TypeOf((*MockBeaconServiceClient)(nil).LatestAttestation), varargs...)
}

// PendingDeposits mocks base method
func (m *MockBeaconServiceClient) PendingDeposits(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*v10.PendingDepositsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PendingDeposits", varargs...)
	ret0, _ := ret[0].(*v10.PendingDepositsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingDeposits indicates an expected call of PendingDeposits
func (mr *MockBeaconServiceClientMockRecorder) PendingDeposits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingDeposits", reflect.TypeOf((*MockBeaconServiceClient)(nil).PendingDeposits), varargs...)
}

// WaitForChainStart mocks base method
func (m *MockBeaconServiceClient) WaitForChainStart(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (v10.BeaconService_WaitForChainStartClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForChainStart", varargs...)
	ret0, _ := ret[0].(v10.BeaconService_WaitForChainStartClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForChainStart indicates an expected call of WaitForChainStart
func (mr *MockBeaconServiceClientMockRecorder) WaitForChainStart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForChainStart", reflect.TypeOf((*MockBeaconServiceClient)(nil).WaitForChainStart), varargs...)
}

// MockBeaconService_LatestAttestationClient is a mock of BeaconService_LatestAttestationClient interface
type MockBeaconService_LatestAttestationClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconService_LatestAttestationClientMockRecorder
}

// MockBeaconService_LatestAttestationClientMockRecorder is the mock recorder for MockBeaconService_LatestAttestationClient
type MockBeaconService_LatestAttestationClientMockRecorder struct {
	mock *MockBeaconService_LatestAttestationClient
}

// NewMockBeaconService_LatestAttestationClient creates a new mock instance
func NewMockBeaconService_LatestAttestationClient(ctrl *gomock.Controller) *MockBeaconService_LatestAttestationClient {
	mock := &MockBeaconService_LatestAttestationClient{ctrl: ctrl}
	mock.recorder = &MockBeaconService_LatestAttestationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconService_LatestAttestationClient) EXPECT() *MockBeaconService_LatestAttestationClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconService_LatestAttestationClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconService_LatestAttestationClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconService_LatestAttestationClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconService_LatestAttestationClient) Recv() (*v1.Attestation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.Attestation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconService_LatestAttestationClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconService_LatestAttestationClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconService_LatestAttestationClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Trailer))
}

// MockBeaconService_WaitForChainStartClient is a mock of BeaconService_WaitForChainStartClient interface
type MockBeaconService_WaitForChainStartClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconService_WaitForChainStartClientMockRecorder
}

// MockBeaconService_WaitForChainStartClientMockRecorder is the mock recorder for MockBeaconService_WaitForChainStartClient
type MockBeaconService_WaitForChainStartClientMockRecorder struct {
	mock *MockBeaconService_WaitForChainStartClient
}

// NewMockBeaconService_WaitForChainStartClient creates a new mock instance
func NewMockBeaconService_WaitForChainStartClient(ctrl *gomock.Controller) *MockBeaconService_WaitForChainStartClient {
	mock := &MockBeaconService_WaitForChainStartClient{ctrl: ctrl}
	mock.recorder = &MockBeaconService_WaitForChainStartClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconService_WaitForChainStartClient) EXPECT() *MockBeaconService_WaitForChainStartClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconService_WaitForChainStartClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconService_WaitForChainStartClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconService_WaitForChainStartClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconService_WaitForChainStartClient) Recv() (*v10.ChainStartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v10.ChainStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconService_WaitForChainStartClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconService_WaitForChainStartClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconService_WaitForChainStartClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconService_WaitForChainStartClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconService_WaitForChainStartClient)(nil).Trailer))
}
