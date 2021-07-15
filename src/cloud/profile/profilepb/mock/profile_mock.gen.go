// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package mock_profilepb is a generated GoMock package.
package mock_profilepb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	uuidpb "px.dev/pixie/src/api/proto/uuidpb"
	profilepb "px.dev/pixie/src/cloud/profile/profilepb"
)

// MockProfileServiceClient is a mock of ProfileServiceClient interface.
type MockProfileServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceClientMockRecorder
}

// MockProfileServiceClientMockRecorder is the mock recorder for MockProfileServiceClient.
type MockProfileServiceClientMockRecorder struct {
	mock *MockProfileServiceClient
}

// NewMockProfileServiceClient creates a new mock instance.
func NewMockProfileServiceClient(ctrl *gomock.Controller) *MockProfileServiceClient {
	mock := &MockProfileServiceClient{ctrl: ctrl}
	mock.recorder = &MockProfileServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileServiceClient) EXPECT() *MockProfileServiceClientMockRecorder {
	return m.recorder
}

// CreateOrgAndUser mocks base method.
func (m *MockProfileServiceClient) CreateOrgAndUser(ctx context.Context, in *profilepb.CreateOrgAndUserRequest, opts ...grpc.CallOption) (*profilepb.CreateOrgAndUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrgAndUser", varargs...)
	ret0, _ := ret[0].(*profilepb.CreateOrgAndUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndUser indicates an expected call of CreateOrgAndUser.
func (mr *MockProfileServiceClientMockRecorder) CreateOrgAndUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndUser", reflect.TypeOf((*MockProfileServiceClient)(nil).CreateOrgAndUser), varargs...)
}

// CreateUser mocks base method.
func (m *MockProfileServiceClient) CreateUser(ctx context.Context, in *profilepb.CreateUserRequest, opts ...grpc.CallOption) (*uuidpb.UUID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*uuidpb.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockProfileServiceClientMockRecorder) CreateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProfileServiceClient)(nil).CreateUser), varargs...)
}

// GetOrg mocks base method.
func (m *MockProfileServiceClient) GetOrg(ctx context.Context, in *uuidpb.UUID, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg.
func (mr *MockProfileServiceClientMockRecorder) GetOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockProfileServiceClient)(nil).GetOrg), varargs...)
}

// GetOrgByDomain mocks base method.
func (m *MockProfileServiceClient) GetOrgByDomain(ctx context.Context, in *profilepb.GetOrgByDomainRequest, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgByDomain", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByDomain indicates an expected call of GetOrgByDomain.
func (mr *MockProfileServiceClientMockRecorder) GetOrgByDomain(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByDomain", reflect.TypeOf((*MockProfileServiceClient)(nil).GetOrgByDomain), varargs...)
}

// GetOrgs mocks base method.
func (m *MockProfileServiceClient) GetOrgs(ctx context.Context, in *profilepb.GetOrgsRequest, opts ...grpc.CallOption) (*profilepb.GetOrgsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgs", varargs...)
	ret0, _ := ret[0].(*profilepb.GetOrgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockProfileServiceClientMockRecorder) GetOrgs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockProfileServiceClient)(nil).GetOrgs), varargs...)
}

// GetUser mocks base method.
func (m *MockProfileServiceClient) GetUser(ctx context.Context, in *uuidpb.UUID, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockProfileServiceClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUser), varargs...)
}

// GetUserAttributes mocks base method.
func (m *MockProfileServiceClient) GetUserAttributes(ctx context.Context, in *profilepb.GetUserAttributesRequest, opts ...grpc.CallOption) (*profilepb.GetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserAttributes", varargs...)
	ret0, _ := ret[0].(*profilepb.GetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAttributes indicates an expected call of GetUserAttributes.
func (mr *MockProfileServiceClientMockRecorder) GetUserAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAttributes", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserAttributes), varargs...)
}

// GetUserByAuthProviderID mocks base method.
func (m *MockProfileServiceClient) GetUserByAuthProviderID(ctx context.Context, in *profilepb.GetUserByAuthProviderIDRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByAuthProviderID", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuthProviderID indicates an expected call of GetUserByAuthProviderID.
func (mr *MockProfileServiceClientMockRecorder) GetUserByAuthProviderID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuthProviderID", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserByAuthProviderID), varargs...)
}

// GetUserByEmail mocks base method.
func (m *MockProfileServiceClient) GetUserByEmail(ctx context.Context, in *profilepb.GetUserByEmailRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByEmail", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockProfileServiceClientMockRecorder) GetUserByEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserByEmail), varargs...)
}

// GetUserSettings mocks base method.
func (m *MockProfileServiceClient) GetUserSettings(ctx context.Context, in *profilepb.GetUserSettingsRequest, opts ...grpc.CallOption) (*profilepb.GetUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserSettings", varargs...)
	ret0, _ := ret[0].(*profilepb.GetUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSettings indicates an expected call of GetUserSettings.
func (mr *MockProfileServiceClientMockRecorder) GetUserSettings(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSettings", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserSettings), varargs...)
}

// GetUsersInOrg mocks base method.
func (m *MockProfileServiceClient) GetUsersInOrg(ctx context.Context, in *profilepb.GetUsersInOrgRequest, opts ...grpc.CallOption) (*profilepb.GetUsersInOrgResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersInOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.GetUsersInOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersInOrg indicates an expected call of GetUsersInOrg.
func (mr *MockProfileServiceClientMockRecorder) GetUsersInOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersInOrg", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUsersInOrg), varargs...)
}

// InviteUser mocks base method.
func (m *MockProfileServiceClient) InviteUser(ctx context.Context, in *profilepb.InviteUserRequest, opts ...grpc.CallOption) (*profilepb.InviteUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InviteUser", varargs...)
	ret0, _ := ret[0].(*profilepb.InviteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUser indicates an expected call of InviteUser.
func (mr *MockProfileServiceClientMockRecorder) InviteUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUser", reflect.TypeOf((*MockProfileServiceClient)(nil).InviteUser), varargs...)
}

// SetUserAttributes mocks base method.
func (m *MockProfileServiceClient) SetUserAttributes(ctx context.Context, in *profilepb.SetUserAttributesRequest, opts ...grpc.CallOption) (*profilepb.SetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetUserAttributes", varargs...)
	ret0, _ := ret[0].(*profilepb.SetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserAttributes indicates an expected call of SetUserAttributes.
func (mr *MockProfileServiceClientMockRecorder) SetUserAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAttributes", reflect.TypeOf((*MockProfileServiceClient)(nil).SetUserAttributes), varargs...)
}

// UpdateOrg mocks base method.
func (m *MockProfileServiceClient) UpdateOrg(ctx context.Context, in *profilepb.UpdateOrgRequest, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrg indicates an expected call of UpdateOrg.
func (mr *MockProfileServiceClientMockRecorder) UpdateOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrg", reflect.TypeOf((*MockProfileServiceClient)(nil).UpdateOrg), varargs...)
}

// UpdateUser mocks base method.
func (m *MockProfileServiceClient) UpdateUser(ctx context.Context, in *profilepb.UpdateUserRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUser", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockProfileServiceClientMockRecorder) UpdateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockProfileServiceClient)(nil).UpdateUser), varargs...)
}

// UpdateUserSettings mocks base method.
func (m *MockProfileServiceClient) UpdateUserSettings(ctx context.Context, in *profilepb.UpdateUserSettingsRequest, opts ...grpc.CallOption) (*profilepb.UpdateUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserSettings", varargs...)
	ret0, _ := ret[0].(*profilepb.UpdateUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserSettings indicates an expected call of UpdateUserSettings.
func (mr *MockProfileServiceClientMockRecorder) UpdateUserSettings(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSettings", reflect.TypeOf((*MockProfileServiceClient)(nil).UpdateUserSettings), varargs...)
}

// MockProfileServiceServer is a mock of ProfileServiceServer interface.
type MockProfileServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceServerMockRecorder
}

// MockProfileServiceServerMockRecorder is the mock recorder for MockProfileServiceServer.
type MockProfileServiceServerMockRecorder struct {
	mock *MockProfileServiceServer
}

// NewMockProfileServiceServer creates a new mock instance.
func NewMockProfileServiceServer(ctrl *gomock.Controller) *MockProfileServiceServer {
	mock := &MockProfileServiceServer{ctrl: ctrl}
	mock.recorder = &MockProfileServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileServiceServer) EXPECT() *MockProfileServiceServerMockRecorder {
	return m.recorder
}

// CreateOrgAndUser mocks base method.
func (m *MockProfileServiceServer) CreateOrgAndUser(arg0 context.Context, arg1 *profilepb.CreateOrgAndUserRequest) (*profilepb.CreateOrgAndUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrgAndUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.CreateOrgAndUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndUser indicates an expected call of CreateOrgAndUser.
func (mr *MockProfileServiceServerMockRecorder) CreateOrgAndUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndUser", reflect.TypeOf((*MockProfileServiceServer)(nil).CreateOrgAndUser), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockProfileServiceServer) CreateUser(arg0 context.Context, arg1 *profilepb.CreateUserRequest) (*uuidpb.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*uuidpb.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockProfileServiceServerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProfileServiceServer)(nil).CreateUser), arg0, arg1)
}

// GetOrg mocks base method.
func (m *MockProfileServiceServer) GetOrg(arg0 context.Context, arg1 *uuidpb.UUID) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg.
func (mr *MockProfileServiceServerMockRecorder) GetOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockProfileServiceServer)(nil).GetOrg), arg0, arg1)
}

// GetOrgByDomain mocks base method.
func (m *MockProfileServiceServer) GetOrgByDomain(arg0 context.Context, arg1 *profilepb.GetOrgByDomainRequest) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgByDomain", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByDomain indicates an expected call of GetOrgByDomain.
func (mr *MockProfileServiceServerMockRecorder) GetOrgByDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByDomain", reflect.TypeOf((*MockProfileServiceServer)(nil).GetOrgByDomain), arg0, arg1)
}

// GetOrgs mocks base method.
func (m *MockProfileServiceServer) GetOrgs(arg0 context.Context, arg1 *profilepb.GetOrgsRequest) (*profilepb.GetOrgsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgs", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetOrgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockProfileServiceServerMockRecorder) GetOrgs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockProfileServiceServer)(nil).GetOrgs), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockProfileServiceServer) GetUser(arg0 context.Context, arg1 *uuidpb.UUID) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockProfileServiceServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUser), arg0, arg1)
}

// GetUserAttributes mocks base method.
func (m *MockProfileServiceServer) GetUserAttributes(arg0 context.Context, arg1 *profilepb.GetUserAttributesRequest) (*profilepb.GetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAttributes", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAttributes indicates an expected call of GetUserAttributes.
func (mr *MockProfileServiceServerMockRecorder) GetUserAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAttributes", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserAttributes), arg0, arg1)
}

// GetUserByAuthProviderID mocks base method.
func (m *MockProfileServiceServer) GetUserByAuthProviderID(arg0 context.Context, arg1 *profilepb.GetUserByAuthProviderIDRequest) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuthProviderID", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuthProviderID indicates an expected call of GetUserByAuthProviderID.
func (mr *MockProfileServiceServerMockRecorder) GetUserByAuthProviderID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuthProviderID", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserByAuthProviderID), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockProfileServiceServer) GetUserByEmail(arg0 context.Context, arg1 *profilepb.GetUserByEmailRequest) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockProfileServiceServerMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserSettings mocks base method.
func (m *MockProfileServiceServer) GetUserSettings(arg0 context.Context, arg1 *profilepb.GetUserSettingsRequest) (*profilepb.GetUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSettings", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSettings indicates an expected call of GetUserSettings.
func (mr *MockProfileServiceServerMockRecorder) GetUserSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSettings", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserSettings), arg0, arg1)
}

// GetUsersInOrg mocks base method.
func (m *MockProfileServiceServer) GetUsersInOrg(arg0 context.Context, arg1 *profilepb.GetUsersInOrgRequest) (*profilepb.GetUsersInOrgResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersInOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetUsersInOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersInOrg indicates an expected call of GetUsersInOrg.
func (mr *MockProfileServiceServerMockRecorder) GetUsersInOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersInOrg", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUsersInOrg), arg0, arg1)
}

// InviteUser mocks base method.
func (m *MockProfileServiceServer) InviteUser(arg0 context.Context, arg1 *profilepb.InviteUserRequest) (*profilepb.InviteUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.InviteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InviteUser indicates an expected call of InviteUser.
func (mr *MockProfileServiceServerMockRecorder) InviteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUser", reflect.TypeOf((*MockProfileServiceServer)(nil).InviteUser), arg0, arg1)
}

// SetUserAttributes mocks base method.
func (m *MockProfileServiceServer) SetUserAttributes(arg0 context.Context, arg1 *profilepb.SetUserAttributesRequest) (*profilepb.SetUserAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserAttributes", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.SetUserAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserAttributes indicates an expected call of SetUserAttributes.
func (mr *MockProfileServiceServerMockRecorder) SetUserAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAttributes", reflect.TypeOf((*MockProfileServiceServer)(nil).SetUserAttributes), arg0, arg1)
}

// UpdateOrg mocks base method.
func (m *MockProfileServiceServer) UpdateOrg(arg0 context.Context, arg1 *profilepb.UpdateOrgRequest) (*profilepb.OrgInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrg indicates an expected call of UpdateOrg.
func (mr *MockProfileServiceServerMockRecorder) UpdateOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrg", reflect.TypeOf((*MockProfileServiceServer)(nil).UpdateOrg), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockProfileServiceServer) UpdateUser(arg0 context.Context, arg1 *profilepb.UpdateUserRequest) (*profilepb.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockProfileServiceServerMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockProfileServiceServer)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserSettings mocks base method.
func (m *MockProfileServiceServer) UpdateUserSettings(arg0 context.Context, arg1 *profilepb.UpdateUserSettingsRequest) (*profilepb.UpdateUserSettingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserSettings", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UpdateUserSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserSettings indicates an expected call of UpdateUserSettings.
func (mr *MockProfileServiceServerMockRecorder) UpdateUserSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSettings", reflect.TypeOf((*MockProfileServiceServer)(nil).UpdateUserSettings), arg0, arg1)
}
