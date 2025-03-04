// Code generated by MockGen. DO NOT EDIT.
// Source: ../ocm/ocm.go
//
// Generated by this command:
//
//	mockgen-v0.5.0 -source=../ocm/ocm.go -destination=ocm.go -package mocks github.com/Azure/ARO-HCP/internal/ocm ClusterServiceClientSpec
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ocm "github.com/Azure/ARO-HCP/internal/ocm"
	v1alpha1 "github.com/openshift-online/ocm-sdk-go/arohcp/v1alpha1"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockClusterServiceClientSpec is a mock of ClusterServiceClientSpec interface.
type MockClusterServiceClientSpec struct {
	ctrl     *gomock.Controller
	recorder *MockClusterServiceClientSpecMockRecorder
	isgomock struct{}
}

// MockClusterServiceClientSpecMockRecorder is the mock recorder for MockClusterServiceClientSpec.
type MockClusterServiceClientSpecMockRecorder struct {
	mock *MockClusterServiceClientSpec
}

// NewMockClusterServiceClientSpec creates a new mock instance.
func NewMockClusterServiceClientSpec(ctrl *gomock.Controller) *MockClusterServiceClientSpec {
	mock := &MockClusterServiceClientSpec{ctrl: ctrl}
	mock.recorder = &MockClusterServiceClientSpecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterServiceClientSpec) EXPECT() *MockClusterServiceClientSpecMockRecorder {
	return m.recorder
}

// AddProperties mocks base method.
func (m *MockClusterServiceClientSpec) AddProperties(builder *v1alpha1.ClusterBuilder) *v1alpha1.ClusterBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProperties", builder)
	ret0, _ := ret[0].(*v1alpha1.ClusterBuilder)
	return ret0
}

// AddProperties indicates an expected call of AddProperties.
func (mr *MockClusterServiceClientSpecMockRecorder) AddProperties(builder any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProperties", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).AddProperties), builder)
}

// DeleteBreakGlassCredentials mocks base method.
func (m *MockClusterServiceClientSpec) DeleteBreakGlassCredentials(ctx context.Context, clusterInternalID ocm.InternalID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBreakGlassCredentials", ctx, clusterInternalID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBreakGlassCredentials indicates an expected call of DeleteBreakGlassCredentials.
func (mr *MockClusterServiceClientSpecMockRecorder) DeleteBreakGlassCredentials(ctx, clusterInternalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBreakGlassCredentials", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).DeleteBreakGlassCredentials), ctx, clusterInternalID)
}

// DeleteCluster mocks base method.
func (m *MockClusterServiceClientSpec) DeleteCluster(ctx context.Context, internalID ocm.InternalID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, internalID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockClusterServiceClientSpecMockRecorder) DeleteCluster(ctx, internalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).DeleteCluster), ctx, internalID)
}

// DeleteNodePool mocks base method.
func (m *MockClusterServiceClientSpec) DeleteNodePool(ctx context.Context, internalID ocm.InternalID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNodePool", ctx, internalID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodePool indicates an expected call of DeleteNodePool.
func (mr *MockClusterServiceClientSpecMockRecorder) DeleteNodePool(ctx, internalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodePool", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).DeleteNodePool), ctx, internalID)
}

// GetBreakGlassCredential mocks base method.
func (m *MockClusterServiceClientSpec) GetBreakGlassCredential(ctx context.Context, internalID ocm.InternalID) (*v1.BreakGlassCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBreakGlassCredential", ctx, internalID)
	ret0, _ := ret[0].(*v1.BreakGlassCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBreakGlassCredential indicates an expected call of GetBreakGlassCredential.
func (mr *MockClusterServiceClientSpecMockRecorder) GetBreakGlassCredential(ctx, internalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBreakGlassCredential", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).GetBreakGlassCredential), ctx, internalID)
}

// GetCluster mocks base method.
func (m *MockClusterServiceClientSpec) GetCluster(ctx context.Context, internalID ocm.InternalID) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx, internalID)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockClusterServiceClientSpecMockRecorder) GetCluster(ctx, internalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).GetCluster), ctx, internalID)
}

// GetClusterStatus mocks base method.
func (m *MockClusterServiceClientSpec) GetClusterStatus(ctx context.Context, internalID ocm.InternalID) (*v1alpha1.ClusterStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterStatus", ctx, internalID)
	ret0, _ := ret[0].(*v1alpha1.ClusterStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterStatus indicates an expected call of GetClusterStatus.
func (mr *MockClusterServiceClientSpecMockRecorder) GetClusterStatus(ctx, internalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterStatus", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).GetClusterStatus), ctx, internalID)
}

// GetNodePool mocks base method.
func (m *MockClusterServiceClientSpec) GetNodePool(ctx context.Context, internalID ocm.InternalID) (*v1.NodePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodePool", ctx, internalID)
	ret0, _ := ret[0].(*v1.NodePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodePool indicates an expected call of GetNodePool.
func (mr *MockClusterServiceClientSpecMockRecorder) GetNodePool(ctx, internalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodePool", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).GetNodePool), ctx, internalID)
}

// ListBreakGlassCredentials mocks base method.
func (m *MockClusterServiceClientSpec) ListBreakGlassCredentials(clusterInternalID ocm.InternalID, searchExpression string) ocm.BreakGlassCredentialListIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBreakGlassCredentials", clusterInternalID, searchExpression)
	ret0, _ := ret[0].(ocm.BreakGlassCredentialListIterator)
	return ret0
}

// ListBreakGlassCredentials indicates an expected call of ListBreakGlassCredentials.
func (mr *MockClusterServiceClientSpecMockRecorder) ListBreakGlassCredentials(clusterInternalID, searchExpression any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBreakGlassCredentials", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).ListBreakGlassCredentials), clusterInternalID, searchExpression)
}

// ListClusters mocks base method.
func (m *MockClusterServiceClientSpec) ListClusters(searchExpression string) ocm.ClusterListIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", searchExpression)
	ret0, _ := ret[0].(ocm.ClusterListIterator)
	return ret0
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockClusterServiceClientSpecMockRecorder) ListClusters(searchExpression any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).ListClusters), searchExpression)
}

// ListNodePools mocks base method.
func (m *MockClusterServiceClientSpec) ListNodePools(clusterInternalID ocm.InternalID, searchExpression string) ocm.NodePoolListIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodePools", clusterInternalID, searchExpression)
	ret0, _ := ret[0].(ocm.NodePoolListIterator)
	return ret0
}

// ListNodePools indicates an expected call of ListNodePools.
func (mr *MockClusterServiceClientSpecMockRecorder) ListNodePools(clusterInternalID, searchExpression any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodePools", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).ListNodePools), clusterInternalID, searchExpression)
}

// PostBreakGlassCredential mocks base method.
func (m *MockClusterServiceClientSpec) PostBreakGlassCredential(ctx context.Context, clusterInternalID ocm.InternalID) (*v1.BreakGlassCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBreakGlassCredential", ctx, clusterInternalID)
	ret0, _ := ret[0].(*v1.BreakGlassCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBreakGlassCredential indicates an expected call of PostBreakGlassCredential.
func (mr *MockClusterServiceClientSpecMockRecorder) PostBreakGlassCredential(ctx, clusterInternalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBreakGlassCredential", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).PostBreakGlassCredential), ctx, clusterInternalID)
}

// PostCluster mocks base method.
func (m *MockClusterServiceClientSpec) PostCluster(ctx context.Context, cluster *v1alpha1.Cluster) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCluster", ctx, cluster)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostCluster indicates an expected call of PostCluster.
func (mr *MockClusterServiceClientSpecMockRecorder) PostCluster(ctx, cluster any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCluster", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).PostCluster), ctx, cluster)
}

// PostNodePool mocks base method.
func (m *MockClusterServiceClientSpec) PostNodePool(ctx context.Context, clusterInternalID ocm.InternalID, nodePool *v1.NodePool) (*v1.NodePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostNodePool", ctx, clusterInternalID, nodePool)
	ret0, _ := ret[0].(*v1.NodePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostNodePool indicates an expected call of PostNodePool.
func (mr *MockClusterServiceClientSpecMockRecorder) PostNodePool(ctx, clusterInternalID, nodePool any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostNodePool", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).PostNodePool), ctx, clusterInternalID, nodePool)
}

// UpdateCluster mocks base method.
func (m *MockClusterServiceClientSpec) UpdateCluster(ctx context.Context, internalID ocm.InternalID, cluster *v1alpha1.Cluster) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", ctx, internalID, cluster)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCluster indicates an expected call of UpdateCluster.
func (mr *MockClusterServiceClientSpecMockRecorder) UpdateCluster(ctx, internalID, cluster any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).UpdateCluster), ctx, internalID, cluster)
}

// UpdateNodePool mocks base method.
func (m *MockClusterServiceClientSpec) UpdateNodePool(ctx context.Context, internalID ocm.InternalID, nodePool *v1.NodePool) (*v1.NodePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNodePool", ctx, internalID, nodePool)
	ret0, _ := ret[0].(*v1.NodePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNodePool indicates an expected call of UpdateNodePool.
func (mr *MockClusterServiceClientSpecMockRecorder) UpdateNodePool(ctx, internalID, nodePool any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodePool", reflect.TypeOf((*MockClusterServiceClientSpec)(nil).UpdateNodePool), ctx, internalID, nodePool)
}
