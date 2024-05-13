// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/s-urbaniak/kube-atlas-cli/internal/store (interfaces: RestoreJobsLister,RestoreJobsDescriber,RestoreJobsCreator,SnapshotsLister,SnapshotsCreator,SnapshotsDescriber,SnapshotsDeleter,ExportJobsLister,ExportJobsDescriber,ExportJobsCreator,ExportBucketsLister,ExportBucketsCreator,ExportBucketsDeleter,ExportBucketsDescriber,ScheduleDescriber,ScheduleDescriberUpdater,ScheduleDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115013/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockRestoreJobsLister is a mock of RestoreJobsLister interface.
type MockRestoreJobsLister struct {
	ctrl     *gomock.Controller
	recorder *MockRestoreJobsListerMockRecorder
}

// MockRestoreJobsListerMockRecorder is the mock recorder for MockRestoreJobsLister.
type MockRestoreJobsListerMockRecorder struct {
	mock *MockRestoreJobsLister
}

// NewMockRestoreJobsLister creates a new mock instance.
func NewMockRestoreJobsLister(ctrl *gomock.Controller) *MockRestoreJobsLister {
	mock := &MockRestoreJobsLister{ctrl: ctrl}
	mock.recorder = &MockRestoreJobsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestoreJobsLister) EXPECT() *MockRestoreJobsListerMockRecorder {
	return m.recorder
}

// RestoreJobs mocks base method.
func (m *MockRestoreJobsLister) RestoreJobs(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*admin.PaginatedCloudBackupRestoreJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreJobs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedCloudBackupRestoreJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreJobs indicates an expected call of RestoreJobs.
func (mr *MockRestoreJobsListerMockRecorder) RestoreJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreJobs", reflect.TypeOf((*MockRestoreJobsLister)(nil).RestoreJobs), arg0, arg1, arg2)
}

// MockRestoreJobsDescriber is a mock of RestoreJobsDescriber interface.
type MockRestoreJobsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockRestoreJobsDescriberMockRecorder
}

// MockRestoreJobsDescriberMockRecorder is the mock recorder for MockRestoreJobsDescriber.
type MockRestoreJobsDescriberMockRecorder struct {
	mock *MockRestoreJobsDescriber
}

// NewMockRestoreJobsDescriber creates a new mock instance.
func NewMockRestoreJobsDescriber(ctrl *gomock.Controller) *MockRestoreJobsDescriber {
	mock := &MockRestoreJobsDescriber{ctrl: ctrl}
	mock.recorder = &MockRestoreJobsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestoreJobsDescriber) EXPECT() *MockRestoreJobsDescriberMockRecorder {
	return m.recorder
}

// RestoreJob mocks base method.
func (m *MockRestoreJobsDescriber) RestoreJob(arg0, arg1, arg2 string) (*admin.DiskBackupSnapshotRestoreJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotRestoreJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreJob indicates an expected call of RestoreJob.
func (mr *MockRestoreJobsDescriberMockRecorder) RestoreJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreJob", reflect.TypeOf((*MockRestoreJobsDescriber)(nil).RestoreJob), arg0, arg1, arg2)
}

// MockRestoreJobsCreator is a mock of RestoreJobsCreator interface.
type MockRestoreJobsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockRestoreJobsCreatorMockRecorder
}

// MockRestoreJobsCreatorMockRecorder is the mock recorder for MockRestoreJobsCreator.
type MockRestoreJobsCreatorMockRecorder struct {
	mock *MockRestoreJobsCreator
}

// NewMockRestoreJobsCreator creates a new mock instance.
func NewMockRestoreJobsCreator(ctrl *gomock.Controller) *MockRestoreJobsCreator {
	mock := &MockRestoreJobsCreator{ctrl: ctrl}
	mock.recorder = &MockRestoreJobsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestoreJobsCreator) EXPECT() *MockRestoreJobsCreatorMockRecorder {
	return m.recorder
}

// CreateRestoreJobs mocks base method.
func (m *MockRestoreJobsCreator) CreateRestoreJobs(arg0, arg1 string, arg2 *admin.DiskBackupSnapshotRestoreJob) (*admin.DiskBackupSnapshotRestoreJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRestoreJobs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotRestoreJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRestoreJobs indicates an expected call of CreateRestoreJobs.
func (mr *MockRestoreJobsCreatorMockRecorder) CreateRestoreJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestoreJobs", reflect.TypeOf((*MockRestoreJobsCreator)(nil).CreateRestoreJobs), arg0, arg1, arg2)
}

// MockSnapshotsLister is a mock of SnapshotsLister interface.
type MockSnapshotsLister struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsListerMockRecorder
}

// MockSnapshotsListerMockRecorder is the mock recorder for MockSnapshotsLister.
type MockSnapshotsListerMockRecorder struct {
	mock *MockSnapshotsLister
}

// NewMockSnapshotsLister creates a new mock instance.
func NewMockSnapshotsLister(ctrl *gomock.Controller) *MockSnapshotsLister {
	mock := &MockSnapshotsLister{ctrl: ctrl}
	mock.recorder = &MockSnapshotsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotsLister) EXPECT() *MockSnapshotsListerMockRecorder {
	return m.recorder
}

// Snapshots mocks base method.
func (m *MockSnapshotsLister) Snapshots(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*admin.PaginatedCloudBackupReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedCloudBackupReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshots indicates an expected call of Snapshots.
func (mr *MockSnapshotsListerMockRecorder) Snapshots(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockSnapshotsLister)(nil).Snapshots), arg0, arg1, arg2)
}

// MockSnapshotsCreator is a mock of SnapshotsCreator interface.
type MockSnapshotsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsCreatorMockRecorder
}

// MockSnapshotsCreatorMockRecorder is the mock recorder for MockSnapshotsCreator.
type MockSnapshotsCreatorMockRecorder struct {
	mock *MockSnapshotsCreator
}

// NewMockSnapshotsCreator creates a new mock instance.
func NewMockSnapshotsCreator(ctrl *gomock.Controller) *MockSnapshotsCreator {
	mock := &MockSnapshotsCreator{ctrl: ctrl}
	mock.recorder = &MockSnapshotsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotsCreator) EXPECT() *MockSnapshotsCreatorMockRecorder {
	return m.recorder
}

// CreateSnapshot mocks base method.
func (m *MockSnapshotsCreator) CreateSnapshot(arg0, arg1 string, arg2 *admin.DiskBackupOnDemandSnapshotRequest) (*admin.DiskBackupSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockSnapshotsCreatorMockRecorder) CreateSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockSnapshotsCreator)(nil).CreateSnapshot), arg0, arg1, arg2)
}

// MockSnapshotsDescriber is a mock of SnapshotsDescriber interface.
type MockSnapshotsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsDescriberMockRecorder
}

// MockSnapshotsDescriberMockRecorder is the mock recorder for MockSnapshotsDescriber.
type MockSnapshotsDescriberMockRecorder struct {
	mock *MockSnapshotsDescriber
}

// NewMockSnapshotsDescriber creates a new mock instance.
func NewMockSnapshotsDescriber(ctrl *gomock.Controller) *MockSnapshotsDescriber {
	mock := &MockSnapshotsDescriber{ctrl: ctrl}
	mock.recorder = &MockSnapshotsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotsDescriber) EXPECT() *MockSnapshotsDescriberMockRecorder {
	return m.recorder
}

// Snapshot mocks base method.
func (m *MockSnapshotsDescriber) Snapshot(arg0, arg1, arg2 string) (*admin.DiskBackupReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockSnapshotsDescriberMockRecorder) Snapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockSnapshotsDescriber)(nil).Snapshot), arg0, arg1, arg2)
}

// MockSnapshotsDeleter is a mock of SnapshotsDeleter interface.
type MockSnapshotsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsDeleterMockRecorder
}

// MockSnapshotsDeleterMockRecorder is the mock recorder for MockSnapshotsDeleter.
type MockSnapshotsDeleterMockRecorder struct {
	mock *MockSnapshotsDeleter
}

// NewMockSnapshotsDeleter creates a new mock instance.
func NewMockSnapshotsDeleter(ctrl *gomock.Controller) *MockSnapshotsDeleter {
	mock := &MockSnapshotsDeleter{ctrl: ctrl}
	mock.recorder = &MockSnapshotsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotsDeleter) EXPECT() *MockSnapshotsDeleterMockRecorder {
	return m.recorder
}

// DeleteSnapshot mocks base method.
func (m *MockSnapshotsDeleter) DeleteSnapshot(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockSnapshotsDeleterMockRecorder) DeleteSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockSnapshotsDeleter)(nil).DeleteSnapshot), arg0, arg1, arg2)
}

// MockExportJobsLister is a mock of ExportJobsLister interface.
type MockExportJobsLister struct {
	ctrl     *gomock.Controller
	recorder *MockExportJobsListerMockRecorder
}

// MockExportJobsListerMockRecorder is the mock recorder for MockExportJobsLister.
type MockExportJobsListerMockRecorder struct {
	mock *MockExportJobsLister
}

// NewMockExportJobsLister creates a new mock instance.
func NewMockExportJobsLister(ctrl *gomock.Controller) *MockExportJobsLister {
	mock := &MockExportJobsLister{ctrl: ctrl}
	mock.recorder = &MockExportJobsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportJobsLister) EXPECT() *MockExportJobsListerMockRecorder {
	return m.recorder
}

// ExportJobs mocks base method.
func (m *MockExportJobsLister) ExportJobs(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*admin.PaginatedApiAtlasDiskBackupExportJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportJobs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.PaginatedApiAtlasDiskBackupExportJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportJobs indicates an expected call of ExportJobs.
func (mr *MockExportJobsListerMockRecorder) ExportJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportJobs", reflect.TypeOf((*MockExportJobsLister)(nil).ExportJobs), arg0, arg1, arg2)
}

// MockExportJobsDescriber is a mock of ExportJobsDescriber interface.
type MockExportJobsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockExportJobsDescriberMockRecorder
}

// MockExportJobsDescriberMockRecorder is the mock recorder for MockExportJobsDescriber.
type MockExportJobsDescriberMockRecorder struct {
	mock *MockExportJobsDescriber
}

// NewMockExportJobsDescriber creates a new mock instance.
func NewMockExportJobsDescriber(ctrl *gomock.Controller) *MockExportJobsDescriber {
	mock := &MockExportJobsDescriber{ctrl: ctrl}
	mock.recorder = &MockExportJobsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportJobsDescriber) EXPECT() *MockExportJobsDescriberMockRecorder {
	return m.recorder
}

// ExportJob mocks base method.
func (m *MockExportJobsDescriber) ExportJob(arg0, arg1, arg2 string) (*admin.DiskBackupExportJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupExportJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportJob indicates an expected call of ExportJob.
func (mr *MockExportJobsDescriberMockRecorder) ExportJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportJob", reflect.TypeOf((*MockExportJobsDescriber)(nil).ExportJob), arg0, arg1, arg2)
}

// MockExportJobsCreator is a mock of ExportJobsCreator interface.
type MockExportJobsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockExportJobsCreatorMockRecorder
}

// MockExportJobsCreatorMockRecorder is the mock recorder for MockExportJobsCreator.
type MockExportJobsCreatorMockRecorder struct {
	mock *MockExportJobsCreator
}

// NewMockExportJobsCreator creates a new mock instance.
func NewMockExportJobsCreator(ctrl *gomock.Controller) *MockExportJobsCreator {
	mock := &MockExportJobsCreator{ctrl: ctrl}
	mock.recorder = &MockExportJobsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportJobsCreator) EXPECT() *MockExportJobsCreatorMockRecorder {
	return m.recorder
}

// CreateExportJob mocks base method.
func (m *MockExportJobsCreator) CreateExportJob(arg0, arg1 string, arg2 *admin.DiskBackupExportJobRequest) (*admin.DiskBackupExportJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExportJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupExportJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExportJob indicates an expected call of CreateExportJob.
func (mr *MockExportJobsCreatorMockRecorder) CreateExportJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExportJob", reflect.TypeOf((*MockExportJobsCreator)(nil).CreateExportJob), arg0, arg1, arg2)
}

// MockExportBucketsLister is a mock of ExportBucketsLister interface.
type MockExportBucketsLister struct {
	ctrl     *gomock.Controller
	recorder *MockExportBucketsListerMockRecorder
}

// MockExportBucketsListerMockRecorder is the mock recorder for MockExportBucketsLister.
type MockExportBucketsListerMockRecorder struct {
	mock *MockExportBucketsLister
}

// NewMockExportBucketsLister creates a new mock instance.
func NewMockExportBucketsLister(ctrl *gomock.Controller) *MockExportBucketsLister {
	mock := &MockExportBucketsLister{ctrl: ctrl}
	mock.recorder = &MockExportBucketsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportBucketsLister) EXPECT() *MockExportBucketsListerMockRecorder {
	return m.recorder
}

// ExportBuckets mocks base method.
func (m *MockExportBucketsLister) ExportBuckets(arg0 string, arg1 *mongodbatlas.ListOptions) (*admin.PaginatedBackupSnapshotExportBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportBuckets", arg0, arg1)
	ret0, _ := ret[0].(*admin.PaginatedBackupSnapshotExportBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportBuckets indicates an expected call of ExportBuckets.
func (mr *MockExportBucketsListerMockRecorder) ExportBuckets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportBuckets", reflect.TypeOf((*MockExportBucketsLister)(nil).ExportBuckets), arg0, arg1)
}

// MockExportBucketsCreator is a mock of ExportBucketsCreator interface.
type MockExportBucketsCreator struct {
	ctrl     *gomock.Controller
	recorder *MockExportBucketsCreatorMockRecorder
}

// MockExportBucketsCreatorMockRecorder is the mock recorder for MockExportBucketsCreator.
type MockExportBucketsCreatorMockRecorder struct {
	mock *MockExportBucketsCreator
}

// NewMockExportBucketsCreator creates a new mock instance.
func NewMockExportBucketsCreator(ctrl *gomock.Controller) *MockExportBucketsCreator {
	mock := &MockExportBucketsCreator{ctrl: ctrl}
	mock.recorder = &MockExportBucketsCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportBucketsCreator) EXPECT() *MockExportBucketsCreatorMockRecorder {
	return m.recorder
}

// CreateExportBucket mocks base method.
func (m *MockExportBucketsCreator) CreateExportBucket(arg0 string, arg1 *admin.DiskBackupSnapshotAWSExportBucket) (*admin.DiskBackupSnapshotAWSExportBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExportBucket", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotAWSExportBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExportBucket indicates an expected call of CreateExportBucket.
func (mr *MockExportBucketsCreatorMockRecorder) CreateExportBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExportBucket", reflect.TypeOf((*MockExportBucketsCreator)(nil).CreateExportBucket), arg0, arg1)
}

// MockExportBucketsDeleter is a mock of ExportBucketsDeleter interface.
type MockExportBucketsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockExportBucketsDeleterMockRecorder
}

// MockExportBucketsDeleterMockRecorder is the mock recorder for MockExportBucketsDeleter.
type MockExportBucketsDeleterMockRecorder struct {
	mock *MockExportBucketsDeleter
}

// NewMockExportBucketsDeleter creates a new mock instance.
func NewMockExportBucketsDeleter(ctrl *gomock.Controller) *MockExportBucketsDeleter {
	mock := &MockExportBucketsDeleter{ctrl: ctrl}
	mock.recorder = &MockExportBucketsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportBucketsDeleter) EXPECT() *MockExportBucketsDeleterMockRecorder {
	return m.recorder
}

// DeleteExportBucket mocks base method.
func (m *MockExportBucketsDeleter) DeleteExportBucket(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExportBucket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExportBucket indicates an expected call of DeleteExportBucket.
func (mr *MockExportBucketsDeleterMockRecorder) DeleteExportBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExportBucket", reflect.TypeOf((*MockExportBucketsDeleter)(nil).DeleteExportBucket), arg0, arg1)
}

// MockExportBucketsDescriber is a mock of ExportBucketsDescriber interface.
type MockExportBucketsDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockExportBucketsDescriberMockRecorder
}

// MockExportBucketsDescriberMockRecorder is the mock recorder for MockExportBucketsDescriber.
type MockExportBucketsDescriberMockRecorder struct {
	mock *MockExportBucketsDescriber
}

// NewMockExportBucketsDescriber creates a new mock instance.
func NewMockExportBucketsDescriber(ctrl *gomock.Controller) *MockExportBucketsDescriber {
	mock := &MockExportBucketsDescriber{ctrl: ctrl}
	mock.recorder = &MockExportBucketsDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportBucketsDescriber) EXPECT() *MockExportBucketsDescriberMockRecorder {
	return m.recorder
}

// DescribeExportBucket mocks base method.
func (m *MockExportBucketsDescriber) DescribeExportBucket(arg0, arg1 string) (*admin.DiskBackupSnapshotAWSExportBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeExportBucket", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotAWSExportBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeExportBucket indicates an expected call of DescribeExportBucket.
func (mr *MockExportBucketsDescriberMockRecorder) DescribeExportBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeExportBucket", reflect.TypeOf((*MockExportBucketsDescriber)(nil).DescribeExportBucket), arg0, arg1)
}

// MockScheduleDescriber is a mock of ScheduleDescriber interface.
type MockScheduleDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleDescriberMockRecorder
}

// MockScheduleDescriberMockRecorder is the mock recorder for MockScheduleDescriber.
type MockScheduleDescriberMockRecorder struct {
	mock *MockScheduleDescriber
}

// NewMockScheduleDescriber creates a new mock instance.
func NewMockScheduleDescriber(ctrl *gomock.Controller) *MockScheduleDescriber {
	mock := &MockScheduleDescriber{ctrl: ctrl}
	mock.recorder = &MockScheduleDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleDescriber) EXPECT() *MockScheduleDescriberMockRecorder {
	return m.recorder
}

// DescribeSchedule mocks base method.
func (m *MockScheduleDescriber) DescribeSchedule(arg0, arg1 string) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockScheduleDescriberMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockScheduleDescriber)(nil).DescribeSchedule), arg0, arg1)
}

// MockScheduleDescriberUpdater is a mock of ScheduleDescriberUpdater interface.
type MockScheduleDescriberUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleDescriberUpdaterMockRecorder
}

// MockScheduleDescriberUpdaterMockRecorder is the mock recorder for MockScheduleDescriberUpdater.
type MockScheduleDescriberUpdaterMockRecorder struct {
	mock *MockScheduleDescriberUpdater
}

// NewMockScheduleDescriberUpdater creates a new mock instance.
func NewMockScheduleDescriberUpdater(ctrl *gomock.Controller) *MockScheduleDescriberUpdater {
	mock := &MockScheduleDescriberUpdater{ctrl: ctrl}
	mock.recorder = &MockScheduleDescriberUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleDescriberUpdater) EXPECT() *MockScheduleDescriberUpdaterMockRecorder {
	return m.recorder
}

// DescribeSchedule mocks base method.
func (m *MockScheduleDescriberUpdater) DescribeSchedule(arg0, arg1 string) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockScheduleDescriberUpdaterMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockScheduleDescriberUpdater)(nil).DescribeSchedule), arg0, arg1)
}

// UpdateSchedule mocks base method.
func (m *MockScheduleDescriberUpdater) UpdateSchedule(arg0, arg1 string, arg2 *admin.DiskBackupSnapshotSchedule) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchedule", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchedule indicates an expected call of UpdateSchedule.
func (mr *MockScheduleDescriberUpdaterMockRecorder) UpdateSchedule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedule", reflect.TypeOf((*MockScheduleDescriberUpdater)(nil).UpdateSchedule), arg0, arg1, arg2)
}

// MockScheduleDeleter is a mock of ScheduleDeleter interface.
type MockScheduleDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleDeleterMockRecorder
}

// MockScheduleDeleterMockRecorder is the mock recorder for MockScheduleDeleter.
type MockScheduleDeleterMockRecorder struct {
	mock *MockScheduleDeleter
}

// NewMockScheduleDeleter creates a new mock instance.
func NewMockScheduleDeleter(ctrl *gomock.Controller) *MockScheduleDeleter {
	mock := &MockScheduleDeleter{ctrl: ctrl}
	mock.recorder = &MockScheduleDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleDeleter) EXPECT() *MockScheduleDeleterMockRecorder {
	return m.recorder
}

// DeleteSchedule mocks base method.
func (m *MockScheduleDeleter) DeleteSchedule(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSchedule", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSchedule indicates an expected call of DeleteSchedule.
func (mr *MockScheduleDeleterMockRecorder) DeleteSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSchedule", reflect.TypeOf((*MockScheduleDeleter)(nil).DeleteSchedule), arg0, arg1)
}
