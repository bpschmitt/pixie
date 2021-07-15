// Code generated by MockGen. DO NOT EDIT.
// Source: auth.pb.go

// Package mock_authpb is a generated GoMock package.
package mock_authpb

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	uuidpb "px.dev/pixie/src/api/proto/uuidpb"
	authpb "px.dev/pixie/src/cloud/auth/authpb"
)

// MockAuthServiceClient is a mock of AuthServiceClient interface.
type MockAuthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceClientMockRecorder
}

// MockAuthServiceClientMockRecorder is the mock recorder for MockAuthServiceClient.
type MockAuthServiceClientMockRecorder struct {
	mock *MockAuthServiceClient
}

// NewMockAuthServiceClient creates a new mock instance.
func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
	return m.recorder
}

// CreateOrgAndInviteUser mocks base method.
func (m *MockAuthServiceClient) CreateOrgAndInviteUser(ctx context.Context, in *authpb.CreateOrgAndInviteUserRequest, opts ...grpc.CallOption) (*authpb.CreateOrgAndInviteUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrgAndInviteUser", varargs...)
	ret0, _ := ret[0].(*authpb.CreateOrgAndInviteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndInviteUser indicates an expected call of CreateOrgAndInviteUser.
func (mr *MockAuthServiceClientMockRecorder) CreateOrgAndInviteUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndInviteUser", reflect.TypeOf((*MockAuthServiceClient)(nil).CreateOrgAndInviteUser), varargs...)
}

// GetAugmentedToken mocks base method.
func (m *MockAuthServiceClient) GetAugmentedToken(ctx context.Context, in *authpb.GetAugmentedAuthTokenRequest, opts ...grpc.CallOption) (*authpb.GetAugmentedAuthTokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAugmentedToken", varargs...)
	ret0, _ := ret[0].(*authpb.GetAugmentedAuthTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAugmentedToken indicates an expected call of GetAugmentedToken.
func (mr *MockAuthServiceClientMockRecorder) GetAugmentedToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAugmentedToken", reflect.TypeOf((*MockAuthServiceClient)(nil).GetAugmentedToken), varargs...)
}

// GetAugmentedTokenForAPIKey mocks base method.
func (m *MockAuthServiceClient) GetAugmentedTokenForAPIKey(ctx context.Context, in *authpb.GetAugmentedTokenForAPIKeyRequest, opts ...grpc.CallOption) (*authpb.GetAugmentedTokenForAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAugmentedTokenForAPIKey", varargs...)
	ret0, _ := ret[0].(*authpb.GetAugmentedTokenForAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAugmentedTokenForAPIKey indicates an expected call of GetAugmentedTokenForAPIKey.
func (mr *MockAuthServiceClientMockRecorder) GetAugmentedTokenForAPIKey(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAugmentedTokenForAPIKey", reflect.TypeOf((*MockAuthServiceClient)(nil).GetAugmentedTokenForAPIKey), varargs...)
}

// InviteUser mocks base method.
func (m *MockAuthServiceClient) InviteUser(ctx context.Context, in *authpb.InviteUserRequest, opts ...grpc.CallOption) (*authpb.InviteUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InviteUser", varargs...)
	ret0, _ := ret[0].(*authpb.InviteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUser indicates an expected call of InviteUser.
func (mr *MockAuthServiceClientMockRecorder) InviteUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUser", reflect.TypeOf((*MockAuthServiceClient)(nil).InviteUser), varargs...)
}

// Login mocks base method.
func (m *MockAuthServiceClient) Login(ctx context.Context, in *authpb.LoginRequest, opts ...grpc.CallOption) (*authpb.LoginReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*authpb.LoginReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServiceClient)(nil).Login), varargs...)
}

// Signup mocks base method.
func (m *MockAuthServiceClient) Signup(ctx context.Context, in *authpb.SignupRequest, opts ...grpc.CallOption) (*authpb.SignupReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signup", varargs...)
	ret0, _ := ret[0].(*authpb.SignupReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthServiceClientMockRecorder) Signup(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuthServiceClient)(nil).Signup), varargs...)
}

// MockAuthServiceServer is a mock of AuthServiceServer interface.
type MockAuthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceServerMockRecorder
}

// MockAuthServiceServerMockRecorder is the mock recorder for MockAuthServiceServer.
type MockAuthServiceServerMockRecorder struct {
	mock *MockAuthServiceServer
}

// NewMockAuthServiceServer creates a new mock instance.
func NewMockAuthServiceServer(ctrl *gomock.Controller) *MockAuthServiceServer {
	mock := &MockAuthServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceServer) EXPECT() *MockAuthServiceServerMockRecorder {
	return m.recorder
}

// CreateOrgAndInviteUser mocks base method.
func (m *MockAuthServiceServer) CreateOrgAndInviteUser(arg0 context.Context, arg1 *authpb.CreateOrgAndInviteUserRequest) (*authpb.CreateOrgAndInviteUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgAndInviteUser", arg0, arg1)
	ret0, _ := ret[0].(*authpb.CreateOrgAndInviteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndInviteUser indicates an expected call of CreateOrgAndInviteUser.
func (mr *MockAuthServiceServerMockRecorder) CreateOrgAndInviteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndInviteUser", reflect.TypeOf((*MockAuthServiceServer)(nil).CreateOrgAndInviteUser), arg0, arg1)
}

// GetAugmentedToken mocks base method.
func (m *MockAuthServiceServer) GetAugmentedToken(arg0 context.Context, arg1 *authpb.GetAugmentedAuthTokenRequest) (*authpb.GetAugmentedAuthTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAugmentedToken", arg0, arg1)
	ret0, _ := ret[0].(*authpb.GetAugmentedAuthTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAugmentedToken indicates an expected call of GetAugmentedToken.
func (mr *MockAuthServiceServerMockRecorder) GetAugmentedToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAugmentedToken", reflect.TypeOf((*MockAuthServiceServer)(nil).GetAugmentedToken), arg0, arg1)
}

// GetAugmentedTokenForAPIKey mocks base method.
func (m *MockAuthServiceServer) GetAugmentedTokenForAPIKey(arg0 context.Context, arg1 *authpb.GetAugmentedTokenForAPIKeyRequest) (*authpb.GetAugmentedTokenForAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAugmentedTokenForAPIKey", arg0, arg1)
	ret0, _ := ret[0].(*authpb.GetAugmentedTokenForAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAugmentedTokenForAPIKey indicates an expected call of GetAugmentedTokenForAPIKey.
func (mr *MockAuthServiceServerMockRecorder) GetAugmentedTokenForAPIKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAugmentedTokenForAPIKey", reflect.TypeOf((*MockAuthServiceServer)(nil).GetAugmentedTokenForAPIKey), arg0, arg1)
}

// InviteUser mocks base method.
func (m *MockAuthServiceServer) InviteUser(arg0 context.Context, arg1 *authpb.InviteUserRequest) (*authpb.InviteUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUser", arg0, arg1)
	ret0, _ := ret[0].(*authpb.InviteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUser indicates an expected call of InviteUser.
func (mr *MockAuthServiceServerMockRecorder) InviteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUser", reflect.TypeOf((*MockAuthServiceServer)(nil).InviteUser), arg0, arg1)
}

// Login mocks base method.
func (m *MockAuthServiceServer) Login(arg0 context.Context, arg1 *authpb.LoginRequest) (*authpb.LoginReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*authpb.LoginReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServiceServer)(nil).Login), arg0, arg1)
}

// Signup mocks base method.
func (m *MockAuthServiceServer) Signup(arg0 context.Context, arg1 *authpb.SignupRequest) (*authpb.SignupReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", arg0, arg1)
	ret0, _ := ret[0].(*authpb.SignupReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthServiceServerMockRecorder) Signup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuthServiceServer)(nil).Signup), arg0, arg1)
}

// MockAPIKeyServiceClient is a mock of APIKeyServiceClient interface.
type MockAPIKeyServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyServiceClientMockRecorder
}

// MockAPIKeyServiceClientMockRecorder is the mock recorder for MockAPIKeyServiceClient.
type MockAPIKeyServiceClientMockRecorder struct {
	mock *MockAPIKeyServiceClient
}

// NewMockAPIKeyServiceClient creates a new mock instance.
func NewMockAPIKeyServiceClient(ctrl *gomock.Controller) *MockAPIKeyServiceClient {
	mock := &MockAPIKeyServiceClient{ctrl: ctrl}
	mock.recorder = &MockAPIKeyServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIKeyServiceClient) EXPECT() *MockAPIKeyServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAPIKeyServiceClient) Create(ctx context.Context, in *authpb.CreateAPIKeyRequest, opts ...grpc.CallOption) (*authpb.APIKey, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*authpb.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAPIKeyServiceClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAPIKeyServiceClient)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockAPIKeyServiceClient) Delete(ctx context.Context, in *uuidpb.UUID, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAPIKeyServiceClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAPIKeyServiceClient)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockAPIKeyServiceClient) Get(ctx context.Context, in *authpb.GetAPIKeyRequest, opts ...grpc.CallOption) (*authpb.GetAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*authpb.GetAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAPIKeyServiceClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAPIKeyServiceClient)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockAPIKeyServiceClient) List(ctx context.Context, in *authpb.ListAPIKeyRequest, opts ...grpc.CallOption) (*authpb.ListAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*authpb.ListAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAPIKeyServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAPIKeyServiceClient)(nil).List), varargs...)
}

// MockAPIKeyServiceServer is a mock of APIKeyServiceServer interface.
type MockAPIKeyServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyServiceServerMockRecorder
}

// MockAPIKeyServiceServerMockRecorder is the mock recorder for MockAPIKeyServiceServer.
type MockAPIKeyServiceServerMockRecorder struct {
	mock *MockAPIKeyServiceServer
}

// NewMockAPIKeyServiceServer creates a new mock instance.
func NewMockAPIKeyServiceServer(ctrl *gomock.Controller) *MockAPIKeyServiceServer {
	mock := &MockAPIKeyServiceServer{ctrl: ctrl}
	mock.recorder = &MockAPIKeyServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIKeyServiceServer) EXPECT() *MockAPIKeyServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAPIKeyServiceServer) Create(arg0 context.Context, arg1 *authpb.CreateAPIKeyRequest) (*authpb.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*authpb.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAPIKeyServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAPIKeyServiceServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockAPIKeyServiceServer) Delete(arg0 context.Context, arg1 *uuidpb.UUID) (*types.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAPIKeyServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAPIKeyServiceServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockAPIKeyServiceServer) Get(arg0 context.Context, arg1 *authpb.GetAPIKeyRequest) (*authpb.GetAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*authpb.GetAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAPIKeyServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAPIKeyServiceServer)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockAPIKeyServiceServer) List(arg0 context.Context, arg1 *authpb.ListAPIKeyRequest) (*authpb.ListAPIKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*authpb.ListAPIKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAPIKeyServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAPIKeyServiceServer)(nil).List), arg0, arg1)
}
