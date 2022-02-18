// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy/deploy.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"
	time "time"

	cloudformation "github.com/aws/copilot-cli/internal/pkg/aws/cloudformation"
	s3 "github.com/aws/copilot-cli/internal/pkg/aws/s3"
	deploy "github.com/aws/copilot-cli/internal/pkg/deploy"
	cloudformation0 "github.com/aws/copilot-cli/internal/pkg/deploy/cloudformation"
	dockerengine "github.com/aws/copilot-cli/internal/pkg/docker/dockerengine"
	repository "github.com/aws/copilot-cli/internal/pkg/repository"
	progress "github.com/aws/copilot-cli/internal/pkg/term/progress"
	gomock "github.com/golang/mock/gomock"
)

// MockActionRecommender is a mock of ActionRecommender interface.
type MockActionRecommender struct {
	ctrl     *gomock.Controller
	recorder *MockActionRecommenderMockRecorder
}

// MockActionRecommenderMockRecorder is the mock recorder for MockActionRecommender.
type MockActionRecommenderMockRecorder struct {
	mock *MockActionRecommender
}

// NewMockActionRecommender creates a new mock instance.
func NewMockActionRecommender(ctrl *gomock.Controller) *MockActionRecommender {
	mock := &MockActionRecommender{ctrl: ctrl}
	mock.recorder = &MockActionRecommenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionRecommender) EXPECT() *MockActionRecommenderMockRecorder {
	return m.recorder
}

// RecommendedActions mocks base method.
func (m *MockActionRecommender) RecommendedActions() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecommendedActions")
	ret0, _ := ret[0].([]string)
	return ret0
}

// RecommendedActions indicates an expected call of RecommendedActions.
func (mr *MockActionRecommenderMockRecorder) RecommendedActions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecommendedActions", reflect.TypeOf((*MockActionRecommender)(nil).RecommendedActions))
}

// MockimageBuilderPusher is a mock of imageBuilderPusher interface.
type MockimageBuilderPusher struct {
	ctrl     *gomock.Controller
	recorder *MockimageBuilderPusherMockRecorder
}

// MockimageBuilderPusherMockRecorder is the mock recorder for MockimageBuilderPusher.
type MockimageBuilderPusherMockRecorder struct {
	mock *MockimageBuilderPusher
}

// NewMockimageBuilderPusher creates a new mock instance.
func NewMockimageBuilderPusher(ctrl *gomock.Controller) *MockimageBuilderPusher {
	mock := &MockimageBuilderPusher{ctrl: ctrl}
	mock.recorder = &MockimageBuilderPusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockimageBuilderPusher) EXPECT() *MockimageBuilderPusherMockRecorder {
	return m.recorder
}

// BuildAndPush mocks base method.
func (m *MockimageBuilderPusher) BuildAndPush(docker repository.ContainerLoginBuildPusher, args *dockerengine.BuildArguments) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAndPush", docker, args)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildAndPush indicates an expected call of BuildAndPush.
func (mr *MockimageBuilderPusherMockRecorder) BuildAndPush(docker, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndPush", reflect.TypeOf((*MockimageBuilderPusher)(nil).BuildAndPush), docker, args)
}

// Mockuploader is a mock of uploader interface.
type Mockuploader struct {
	ctrl     *gomock.Controller
	recorder *MockuploaderMockRecorder
}

// MockuploaderMockRecorder is the mock recorder for Mockuploader.
type MockuploaderMockRecorder struct {
	mock *Mockuploader
}

// NewMockuploader creates a new mock instance.
func NewMockuploader(ctrl *gomock.Controller) *Mockuploader {
	mock := &Mockuploader{ctrl: ctrl}
	mock.recorder = &MockuploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockuploader) EXPECT() *MockuploaderMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *Mockuploader) Upload(bucket, key string, data io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", bucket, key, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockuploaderMockRecorder) Upload(bucket, key, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*Mockuploader)(nil).Upload), bucket, key, data)
}

// ZipAndUpload mocks base method.
func (m *Mockuploader) ZipAndUpload(bucket, key string, files ...s3.NamedBinary) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{bucket, key}
	for _, a := range files {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZipAndUpload", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZipAndUpload indicates an expected call of ZipAndUpload.
func (mr *MockuploaderMockRecorder) ZipAndUpload(bucket, key interface{}, files ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{bucket, key}, files...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZipAndUpload", reflect.TypeOf((*Mockuploader)(nil).ZipAndUpload), varargs...)
}

// Mocktemplater is a mock of templater interface.
type Mocktemplater struct {
	ctrl     *gomock.Controller
	recorder *MocktemplaterMockRecorder
}

// MocktemplaterMockRecorder is the mock recorder for Mocktemplater.
type MocktemplaterMockRecorder struct {
	mock *Mocktemplater
}

// NewMocktemplater creates a new mock instance.
func NewMocktemplater(ctrl *gomock.Controller) *Mocktemplater {
	mock := &Mocktemplater{ctrl: ctrl}
	mock.recorder = &MocktemplaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktemplater) EXPECT() *MocktemplaterMockRecorder {
	return m.recorder
}

// Template mocks base method.
func (m *Mocktemplater) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MocktemplaterMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*Mocktemplater)(nil).Template))
}

// MockstackSerializer is a mock of stackSerializer interface.
type MockstackSerializer struct {
	ctrl     *gomock.Controller
	recorder *MockstackSerializerMockRecorder
}

// MockstackSerializerMockRecorder is the mock recorder for MockstackSerializer.
type MockstackSerializerMockRecorder struct {
	mock *MockstackSerializer
}

// NewMockstackSerializer creates a new mock instance.
func NewMockstackSerializer(ctrl *gomock.Controller) *MockstackSerializer {
	mock := &MockstackSerializer{ctrl: ctrl}
	mock.recorder = &MockstackSerializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstackSerializer) EXPECT() *MockstackSerializerMockRecorder {
	return m.recorder
}

// SerializedParameters mocks base method.
func (m *MockstackSerializer) SerializedParameters() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SerializedParameters")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SerializedParameters indicates an expected call of SerializedParameters.
func (mr *MockstackSerializerMockRecorder) SerializedParameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SerializedParameters", reflect.TypeOf((*MockstackSerializer)(nil).SerializedParameters))
}

// Template mocks base method.
func (m *MockstackSerializer) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockstackSerializerMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockstackSerializer)(nil).Template))
}

// MockendpointGetter is a mock of endpointGetter interface.
type MockendpointGetter struct {
	ctrl     *gomock.Controller
	recorder *MockendpointGetterMockRecorder
}

// MockendpointGetterMockRecorder is the mock recorder for MockendpointGetter.
type MockendpointGetterMockRecorder struct {
	mock *MockendpointGetter
}

// NewMockendpointGetter creates a new mock instance.
func NewMockendpointGetter(ctrl *gomock.Controller) *MockendpointGetter {
	mock := &MockendpointGetter{ctrl: ctrl}
	mock.recorder = &MockendpointGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockendpointGetter) EXPECT() *MockendpointGetterMockRecorder {
	return m.recorder
}

// ServiceDiscoveryEndpoint mocks base method.
func (m *MockendpointGetter) ServiceDiscoveryEndpoint() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceDiscoveryEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceDiscoveryEndpoint indicates an expected call of ServiceDiscoveryEndpoint.
func (mr *MockendpointGetterMockRecorder) ServiceDiscoveryEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceDiscoveryEndpoint", reflect.TypeOf((*MockendpointGetter)(nil).ServiceDiscoveryEndpoint))
}

// MockversionGetter is a mock of versionGetter interface.
type MockversionGetter struct {
	ctrl     *gomock.Controller
	recorder *MockversionGetterMockRecorder
}

// MockversionGetterMockRecorder is the mock recorder for MockversionGetter.
type MockversionGetterMockRecorder struct {
	mock *MockversionGetter
}

// NewMockversionGetter creates a new mock instance.
func NewMockversionGetter(ctrl *gomock.Controller) *MockversionGetter {
	mock := &MockversionGetter{ctrl: ctrl}
	mock.recorder = &MockversionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockversionGetter) EXPECT() *MockversionGetterMockRecorder {
	return m.recorder
}

// Version mocks base method.
func (m *MockversionGetter) Version() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockversionGetterMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockversionGetter)(nil).Version))
}

// MockpublicCIDRBlocksGetter is a mock of publicCIDRBlocksGetter interface.
type MockpublicCIDRBlocksGetter struct {
	ctrl     *gomock.Controller
	recorder *MockpublicCIDRBlocksGetterMockRecorder
}

// MockpublicCIDRBlocksGetterMockRecorder is the mock recorder for MockpublicCIDRBlocksGetter.
type MockpublicCIDRBlocksGetterMockRecorder struct {
	mock *MockpublicCIDRBlocksGetter
}

// NewMockpublicCIDRBlocksGetter creates a new mock instance.
func NewMockpublicCIDRBlocksGetter(ctrl *gomock.Controller) *MockpublicCIDRBlocksGetter {
	mock := &MockpublicCIDRBlocksGetter{ctrl: ctrl}
	mock.recorder = &MockpublicCIDRBlocksGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpublicCIDRBlocksGetter) EXPECT() *MockpublicCIDRBlocksGetterMockRecorder {
	return m.recorder
}

// PublicCIDRBlocks mocks base method.
func (m *MockpublicCIDRBlocksGetter) PublicCIDRBlocks() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicCIDRBlocks")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicCIDRBlocks indicates an expected call of PublicCIDRBlocks.
func (mr *MockpublicCIDRBlocksGetterMockRecorder) PublicCIDRBlocks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicCIDRBlocks", reflect.TypeOf((*MockpublicCIDRBlocksGetter)(nil).PublicCIDRBlocks))
}

// MockcustomResourcesUploader is a mock of customResourcesUploader interface.
type MockcustomResourcesUploader struct {
	ctrl     *gomock.Controller
	recorder *MockcustomResourcesUploaderMockRecorder
}

// MockcustomResourcesUploaderMockRecorder is the mock recorder for MockcustomResourcesUploader.
type MockcustomResourcesUploaderMockRecorder struct {
	mock *MockcustomResourcesUploader
}

// NewMockcustomResourcesUploader creates a new mock instance.
func NewMockcustomResourcesUploader(ctrl *gomock.Controller) *MockcustomResourcesUploader {
	mock := &MockcustomResourcesUploader{ctrl: ctrl}
	mock.recorder = &MockcustomResourcesUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcustomResourcesUploader) EXPECT() *MockcustomResourcesUploaderMockRecorder {
	return m.recorder
}

// UploadEnvironmentCustomResources mocks base method.
func (m *MockcustomResourcesUploader) UploadEnvironmentCustomResources(upload s3.CompressAndUploadFunc) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadEnvironmentCustomResources", upload)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadEnvironmentCustomResources indicates an expected call of UploadEnvironmentCustomResources.
func (mr *MockcustomResourcesUploaderMockRecorder) UploadEnvironmentCustomResources(upload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadEnvironmentCustomResources", reflect.TypeOf((*MockcustomResourcesUploader)(nil).UploadEnvironmentCustomResources), upload)
}

// UploadNetworkLoadBalancedWebServiceCustomResources mocks base method.
func (m *MockcustomResourcesUploader) UploadNetworkLoadBalancedWebServiceCustomResources(upload s3.CompressAndUploadFunc) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadNetworkLoadBalancedWebServiceCustomResources", upload)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadNetworkLoadBalancedWebServiceCustomResources indicates an expected call of UploadNetworkLoadBalancedWebServiceCustomResources.
func (mr *MockcustomResourcesUploaderMockRecorder) UploadNetworkLoadBalancedWebServiceCustomResources(upload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadNetworkLoadBalancedWebServiceCustomResources", reflect.TypeOf((*MockcustomResourcesUploader)(nil).UploadNetworkLoadBalancedWebServiceCustomResources), upload)
}

// UploadRequestDrivenWebServiceCustomResources mocks base method.
func (m *MockcustomResourcesUploader) UploadRequestDrivenWebServiceCustomResources(upload s3.CompressAndUploadFunc) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadRequestDrivenWebServiceCustomResources", upload)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadRequestDrivenWebServiceCustomResources indicates an expected call of UploadRequestDrivenWebServiceCustomResources.
func (mr *MockcustomResourcesUploaderMockRecorder) UploadRequestDrivenWebServiceCustomResources(upload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadRequestDrivenWebServiceCustomResources", reflect.TypeOf((*MockcustomResourcesUploader)(nil).UploadRequestDrivenWebServiceCustomResources), upload)
}

// MocksnsTopicsLister is a mock of snsTopicsLister interface.
type MocksnsTopicsLister struct {
	ctrl     *gomock.Controller
	recorder *MocksnsTopicsListerMockRecorder
}

// MocksnsTopicsListerMockRecorder is the mock recorder for MocksnsTopicsLister.
type MocksnsTopicsListerMockRecorder struct {
	mock *MocksnsTopicsLister
}

// NewMocksnsTopicsLister creates a new mock instance.
func NewMocksnsTopicsLister(ctrl *gomock.Controller) *MocksnsTopicsLister {
	mock := &MocksnsTopicsLister{ctrl: ctrl}
	mock.recorder = &MocksnsTopicsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksnsTopicsLister) EXPECT() *MocksnsTopicsListerMockRecorder {
	return m.recorder
}

// ListSNSTopics mocks base method.
func (m *MocksnsTopicsLister) ListSNSTopics(appName, envName string) ([]deploy.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSNSTopics", appName, envName)
	ret0, _ := ret[0].([]deploy.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSNSTopics indicates an expected call of ListSNSTopics.
func (mr *MocksnsTopicsListerMockRecorder) ListSNSTopics(appName, envName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSNSTopics", reflect.TypeOf((*MocksnsTopicsLister)(nil).ListSNSTopics), appName, envName)
}

// MockserviceDeployer is a mock of serviceDeployer interface.
type MockserviceDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockserviceDeployerMockRecorder
}

// MockserviceDeployerMockRecorder is the mock recorder for MockserviceDeployer.
type MockserviceDeployerMockRecorder struct {
	mock *MockserviceDeployer
}

// NewMockserviceDeployer creates a new mock instance.
func NewMockserviceDeployer(ctrl *gomock.Controller) *MockserviceDeployer {
	mock := &MockserviceDeployer{ctrl: ctrl}
	mock.recorder = &MockserviceDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserviceDeployer) EXPECT() *MockserviceDeployerMockRecorder {
	return m.recorder
}

// DeployService mocks base method.
func (m *MockserviceDeployer) DeployService(out progress.FileWriter, conf cloudformation0.StackConfiguration, bucketName string, opts ...cloudformation.StackOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{out, conf, bucketName}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeployService", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployService indicates an expected call of DeployService.
func (mr *MockserviceDeployerMockRecorder) DeployService(out, conf, bucketName interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out, conf, bucketName}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployService", reflect.TypeOf((*MockserviceDeployer)(nil).DeployService), varargs...)
}

// MockserviceForceUpdater is a mock of serviceForceUpdater interface.
type MockserviceForceUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockserviceForceUpdaterMockRecorder
}

// MockserviceForceUpdaterMockRecorder is the mock recorder for MockserviceForceUpdater.
type MockserviceForceUpdaterMockRecorder struct {
	mock *MockserviceForceUpdater
}

// NewMockserviceForceUpdater creates a new mock instance.
func NewMockserviceForceUpdater(ctrl *gomock.Controller) *MockserviceForceUpdater {
	mock := &MockserviceForceUpdater{ctrl: ctrl}
	mock.recorder = &MockserviceForceUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserviceForceUpdater) EXPECT() *MockserviceForceUpdaterMockRecorder {
	return m.recorder
}

// ForceUpdateService mocks base method.
func (m *MockserviceForceUpdater) ForceUpdateService(app, env, svc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUpdateService", app, env, svc)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceUpdateService indicates an expected call of ForceUpdateService.
func (mr *MockserviceForceUpdaterMockRecorder) ForceUpdateService(app, env, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUpdateService", reflect.TypeOf((*MockserviceForceUpdater)(nil).ForceUpdateService), app, env, svc)
}

// LastUpdatedAt mocks base method.
func (m *MockserviceForceUpdater) LastUpdatedAt(app, env, svc string) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastUpdatedAt", app, env, svc)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastUpdatedAt indicates an expected call of LastUpdatedAt.
func (mr *MockserviceForceUpdaterMockRecorder) LastUpdatedAt(app, env, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastUpdatedAt", reflect.TypeOf((*MockserviceForceUpdater)(nil).LastUpdatedAt), app, env, svc)
}

// Mockspinner is a mock of spinner interface.
type Mockspinner struct {
	ctrl     *gomock.Controller
	recorder *MockspinnerMockRecorder
}

// MockspinnerMockRecorder is the mock recorder for Mockspinner.
type MockspinnerMockRecorder struct {
	mock *Mockspinner
}

// NewMockspinner creates a new mock instance.
func NewMockspinner(ctrl *gomock.Controller) *Mockspinner {
	mock := &Mockspinner{ctrl: ctrl}
	mock.recorder = &MockspinnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockspinner) EXPECT() *MockspinnerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *Mockspinner) Start(label string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", label)
}

// Start indicates an expected call of Start.
func (mr *MockspinnerMockRecorder) Start(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockspinner)(nil).Start), label)
}

// Stop mocks base method.
func (m *Mockspinner) Stop(label string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", label)
}

// Stop indicates an expected call of Stop.
func (mr *MockspinnerMockRecorder) Stop(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Mockspinner)(nil).Stop), label)
}

// MockfileReader is a mock of fileReader interface.
type MockfileReader struct {
	ctrl     *gomock.Controller
	recorder *MockfileReaderMockRecorder
}

// MockfileReaderMockRecorder is the mock recorder for MockfileReader.
type MockfileReaderMockRecorder struct {
	mock *MockfileReader
}

// NewMockfileReader creates a new mock instance.
func NewMockfileReader(ctrl *gomock.Controller) *MockfileReader {
	mock := &MockfileReader{ctrl: ctrl}
	mock.recorder = &MockfileReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileReader) EXPECT() *MockfileReaderMockRecorder {
	return m.recorder
}

// ReadFile mocks base method.
func (m *MockfileReader) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockfileReaderMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockfileReader)(nil).ReadFile), arg0)
}

// MocktimeoutError is a mock of timeoutError interface.
type MocktimeoutError struct {
	ctrl     *gomock.Controller
	recorder *MocktimeoutErrorMockRecorder
}

// MocktimeoutErrorMockRecorder is the mock recorder for MocktimeoutError.
type MocktimeoutErrorMockRecorder struct {
	mock *MocktimeoutError
}

// NewMocktimeoutError creates a new mock instance.
func NewMocktimeoutError(ctrl *gomock.Controller) *MocktimeoutError {
	mock := &MocktimeoutError{ctrl: ctrl}
	mock.recorder = &MocktimeoutErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktimeoutError) EXPECT() *MocktimeoutErrorMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MocktimeoutError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MocktimeoutErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MocktimeoutError)(nil).Error))
}

// Timeout mocks base method.
func (m *MocktimeoutError) Timeout() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MocktimeoutErrorMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MocktimeoutError)(nil).Timeout))
}