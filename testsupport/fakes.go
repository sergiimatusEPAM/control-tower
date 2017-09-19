package testsupport

import (
	"github.com/EngineerBetter/concourse-up/bosh"
	"github.com/EngineerBetter/concourse-up/config"
	"github.com/EngineerBetter/concourse-up/terraform"
)

// FakeAWSClient implements iaas.IClient for testing
type FakeAWSClient struct {
	FakeDeleteVMsInVPC                func(vpcID string) error
	FakeDeleteFile                    func(bucket, path string) error
	FakeDeleteVersionedBucket         func(name string) error
	FakeEnsureBucketExists            func(name string) error
	FakeEnsureFileExists              func(bucket, path string, defaultContents []byte) ([]byte, bool, error)
	FakeFindLongestMatchingHostedZone func(subdomain string) (string, string, error)
	FakeHasFile                       func(bucket, path string) (bool, error)
	FakeLoadFile                      func(bucket, path string) ([]byte, error)
	FakeWriteFile                     func(bucket, path string, contents []byte) error
	FakeRegion                        func() string
}

// IAAS is here to implement iaas.IClient
func (client *FakeAWSClient) IAAS() string {
	return "AWS"
}

// Region delegates to FakeRegion which is dynamically set by the tests
func (client *FakeAWSClient) Region() string {
	return client.FakeRegion()
}

// DeleteVMsInVPC delegates to FakeDeleteVMsInVPC which is dynamically set by the tests
func (client *FakeAWSClient) DeleteVMsInVPC(vpcID string) error {
	return client.FakeDeleteVMsInVPC(vpcID)
}

// DeleteFile delegates to FakeDeleteFile which is dynamically set by the tests
func (client *FakeAWSClient) DeleteFile(bucket, path string) error {
	return client.FakeDeleteFile(bucket, path)
}

// DeleteVersionedBucket delegates to FakeDeleteVersionedBucket which is dynamically set by the tests
func (client *FakeAWSClient) DeleteVersionedBucket(name string) error {
	return client.FakeDeleteVersionedBucket(name)
}

// EnsureBucketExists delegates to FakeEnsureBucketExists which is dynamically set by the tests
func (client *FakeAWSClient) EnsureBucketExists(name string) error {
	return client.FakeEnsureBucketExists(name)
}

// EnsureFileExists delegates to FakeEnsureFileExists which is dynamically set by the tests
func (client *FakeAWSClient) EnsureFileExists(bucket, path string, defaultContents []byte) ([]byte, bool, error) {
	return client.FakeEnsureFileExists(bucket, path, defaultContents)
}

// FindLongestMatchingHostedZone delegates to FakeFindLongestMatchingHostedZone which is dynamically set by the tests
func (client *FakeAWSClient) FindLongestMatchingHostedZone(subdomain string) (string, string, error) {
	return client.FakeFindLongestMatchingHostedZone(subdomain)
}

// HasFile delegates to FakeHasFile which is dynamically set by the tests
func (client *FakeAWSClient) HasFile(bucket, path string) (bool, error) {
	return client.FakeHasFile(bucket, path)
}

// LoadFile delegates to FakeLoadFile which is dynamically set by the tests
func (client *FakeAWSClient) LoadFile(bucket, path string) ([]byte, error) {
	return client.FakeLoadFile(bucket, path)
}

// WriteFile delegates to FakeWriteFile which is dynamically set by the tests
func (client *FakeAWSClient) WriteFile(bucket, path string, contents []byte) error {
	return client.FakeWriteFile(bucket, path, contents)
}

// FakeFlyClient implements fly.IClient for testing
type FakeFlyClient struct {
	FakeSetDefaultPipeline func(deployAgs *config.DeployArgs, config *config.Config) error
	FakeCleanup            func() error
	FakeCanConnect         func() (bool, error)
}

// SetDefaultPipeline delegates to FakeSetDefaultPipeline which is dynamically set by the tests
func (client *FakeFlyClient) SetDefaultPipeline(deployArgs *config.DeployArgs, config *config.Config) error {
	return client.FakeSetDefaultPipeline(deployArgs, config)
}

// Cleanup delegates to FakeCleanup which is dynamically set by the tests
func (client *FakeFlyClient) Cleanup() error {
	return client.FakeCleanup()
}

// CanConnect delegates to FakeCanConnect which is dynamically set by the tests
func (client *FakeFlyClient) CanConnect() (bool, error) {
	return client.FakeCanConnect()
}

// FakeConfigClient implements config.IClient for testing
type FakeConfigClient struct {
	FakeLoad         func() (*config.Config, error)
	FakeUpdate       func(*config.Config) error
	FakeLoadOrCreate func(deployArgs *config.DeployArgs) (*config.Config, bool, error)
	FakeStoreAsset   func(filename string, contents []byte) error
	FakeLoadAsset    func(filename string) ([]byte, error)
	FakeDeleteAsset  func(filename string) error
	FakeDeleteAll    func(config *config.Config) error
	FakeHasAsset     func(filename string) (bool, error)
}

// Load delegates to FakeLoad which is dynamically set by the tests
func (client *FakeConfigClient) Load() (*config.Config, error) {
	return client.FakeLoad()
}

// Update delegates to FakeUpdate which is dynamically set by the tests
func (client *FakeConfigClient) Update(config *config.Config) error {
	return client.FakeUpdate(config)
}

// LoadOrCreate delegates to FakeLoadOrCreate which is dynamically set by the tests
func (client *FakeConfigClient) LoadOrCreate(deployArgs *config.DeployArgs) (*config.Config, bool, error) {
	return client.FakeLoadOrCreate(deployArgs)
}

// StoreAsset delegates to FakeStoreAsset which is dynamically set by the tests
func (client *FakeConfigClient) StoreAsset(filename string, contents []byte) error {
	return client.FakeStoreAsset(filename, contents)
}

// LoadAsset delegates to FakeLoadAsset which is dynamically set by the tests
func (client *FakeConfigClient) LoadAsset(filename string) ([]byte, error) {
	return client.FakeLoadAsset(filename)
}

// DeleteAsset delegates to FakeDeleteAsset which is dynamically set by the tests
func (client *FakeConfigClient) DeleteAsset(filename string) error {
	return client.FakeDeleteAsset(filename)
}

// DeleteAll delegates to FakeDeleteAll which is dynamically set by the tests
func (client *FakeConfigClient) DeleteAll(config *config.Config) error {
	return client.FakeDeleteAll(config)
}

// HasAsset delegates to FakeHasAsset which is dynamically set by the tests
func (client *FakeConfigClient) HasAsset(filename string) (bool, error) {
	return client.FakeHasAsset(filename)
}

// FakeTerraformClient implements terraform.IClient for testing
type FakeTerraformClient struct {
	FakeOutput  func() (*terraform.Metadata, error)
	FakeApply   func(dryrun bool) error
	FakeDestroy func() error
	FakeCleanup func() error
}

// Output delegates to FakeOutput which is dynamically set by the tests
func (client *FakeTerraformClient) Output() (*terraform.Metadata, error) {
	return client.FakeOutput()
}

// Apply delegates to FakeApply which is dynamically set by the tests
func (client *FakeTerraformClient) Apply(dryrun bool) error {
	return client.FakeApply(dryrun)
}

// Destroy delegates to FakeDestroy which is dynamically set by the tests
func (client *FakeTerraformClient) Destroy() error {
	return client.FakeDestroy()
}

// Cleanup delegates to FakeCleanup which is dynamically set by the tests
func (client *FakeTerraformClient) Cleanup() error {
	return client.FakeCleanup()
}

// FakeBoshClient implements bosh.IClient for testing
type FakeBoshClient struct {
	FakeDeploy    func([]byte, bool) ([]byte, error)
	FakeDelete    func([]byte) ([]byte, error)
	FakeCleanup   func() error
	FakeInstances func() ([]bosh.Instance, error)
}

// Deploy delegates to FakeDeploy which is dynamically set by the tests
func (client *FakeBoshClient) Deploy(stateFileBytes []byte, detach bool) ([]byte, error) {
	return client.FakeDeploy(stateFileBytes, detach)
}

// Delete delegates to FakeDelete which is dynamically set by the tests
func (client *FakeBoshClient) Delete(stateFileBytes []byte) ([]byte, error) {
	return client.FakeDelete(stateFileBytes)
}

// Cleanup delegates to FakeCleanup which is dynamically set by the tests
func (client *FakeBoshClient) Cleanup() error {
	return client.FakeCleanup()
}

// Instances delegates to FakeInstances which is dynamically set by the tests
func (client *FakeBoshClient) Instances() ([]bosh.Instance, error) {
	return client.FakeInstances()
}