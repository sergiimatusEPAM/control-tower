package iaas

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"io/ioutil"

	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	// New versions of github.com/aws/aws-sdk-go/aws have these consts
	// but the version currently pinned by bosh-cli v2 does not

	// ErrCodeNoSuchBucket for service response error code
	// "NoSuchBucket".
	//
	// The specified bucket does not exist.
	awsErrCodeNoSuchBucket = "NoSuchBucket"

	// ErrCodeNoSuchKey for service response error code
	// "NoSuchKey".
	//
	// The specified key does not exist.
	awsErrCodeNoSuchKey = "NoSuchKey"

	// Returned when calling HEAD on non-existant bucket or object
	awsErrCodeNotFound = "NotFound"
)

// DeleteVersionedBucket deletes and empties a versioned bucket
func (client *AWSProvider) DeleteVersionedBucket(name string) error {

	s3Client := s3.New(client.sess)

	// Delete all objects
	objects := []*s3.ObjectVersion{}
	err := s3Client.ListObjectVersionsPages(&s3.ListObjectVersionsInput{Bucket: &name},
		func(output *s3.ListObjectVersionsOutput, _ bool) bool {
			objects = append(objects, output.Versions...)

			return true
		})
	if err != nil {
		return err
	}

	for _, object := range objects {
		_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: &name,
			Key:    object.Key,
			VersionId: object.VersionId,
		})
		if err != nil {
			return nil
		}
	}

	time.Sleep(time.Second)

	_, err = s3Client.DeleteBucket(&s3.DeleteBucketInput{Bucket: &name})
	return err
}

// CreateBucket checks if the named bucket exists and creates it if it doesn't
func (client *AWSProvider) CreateBucket(name string) error {

	s3Client := s3.New(client.sess)

	bucketInput := &s3.CreateBucketInput{
		Bucket: &name,
	}
	// NOTE the location constraint should only be set if using a bucket OTHER than us-east-1
	// http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUT.html
	if *client.sess.Config.Region != "us-east-1" {
		bucketInput.CreateBucketConfiguration = &s3.CreateBucketConfiguration{
			LocationConstraint: client.sess.Config.Region,
		}
	}

	_, err := s3Client.CreateBucket(bucketInput)

	if err != nil {
		return fmt.Errorf("error creating S3 bucket [%v]: [%v]", name, err)
	}

	versioningConfig := &s3.VersioningConfiguration{
		Status: aws.String(s3.BucketVersioningStatusEnabled),
	}

	versioningInput := &s3.PutBucketVersioningInput{
		Bucket: &name,
		VersioningConfiguration: versioningConfig,
	}

	_, err = s3Client.PutBucketVersioning(versioningInput)
	if err != nil {
		return fmt.Errorf("error enabling versioning on S3 bucket [%v]: [%v]", name, err)
	}

	return nil
}

// BucketExists checks if the named bucket exists
func (client *AWSProvider) BucketExists(name string) (bool, error) {

	s3Client := s3.New(client.sess)

	_, err := s3Client.HeadBucket(&s3.HeadBucketInput{Bucket: &name})
	if err == nil {
		return true, nil
	}

	awsErrCode := err.(awserr.Error).Code()
	if awsErrCode != awsErrCodeNotFound && awsErrCode != awsErrCodeNoSuchBucket {
		return false, err
	}

	return false, nil
}

// WriteFile writes the specified S3 object
func (client *AWSProvider) WriteFile(bucket, path string, contents []byte) error {
	s3Client := s3.New(client.sess)

	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &path,
		Body:   bytes.NewReader(contents),
	})
	return err
}

// HasFile returns true if the specified S3 object exists
func (client *AWSProvider) HasFile(bucket, path string) (bool, error) {
	s3Client := s3.New(client.sess)

	_, err := s3Client.HeadObject(&s3.HeadObjectInput{Bucket: &bucket, Key: &path})
	if err != nil {
		awsErrCode := err.(awserr.Error).Code()
		if awsErrCode == awsErrCodeNotFound || awsErrCode == awsErrCodeNoSuchKey {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// EnsureFileExists checks for the named file in S3 and creates it if it doesn't exist
// Second argument is true if new file was created
func (client *AWSProvider) EnsureFileExists(bucket, path string, defaultContents []byte) ([]byte, bool, error) {

	s3Client := s3.New(client.sess)

	// Trying to get the Object
	output, err := s3Client.GetObject(&s3.GetObjectInput{Bucket: &bucket, Key: &path})
	if err == nil {
		var contents []byte
		contents, err = ioutil.ReadAll(output.Body)
		if err != nil {
			return nil, false, err
		}

		// Found an Object and returns its contents
		return contents, false, nil
	}

	// Bubble up the error if it was irelevant of NotFound
	awsErrCode := err.(awserr.Error).Code()
	if awsErrCode != awsErrCodeNoSuchKey && awsErrCode != awsErrCodeNotFound {
		return nil, false, err
	}

	// Create the file (path) in the bucket with defaultContents
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &path,
		Body:   bytes.NewReader(defaultContents),
	})
	if err != nil {
		return nil, false, err
	}

	// The file was created (new) and contains the defaultContents
	return defaultContents, true, nil
}

// LoadFile loads a file from S3
func (client *AWSProvider) LoadFile(bucket, path string) ([]byte, error) {

	s3Client := s3.New(client.sess)

	output, err := s3Client.GetObject(&s3.GetObjectInput{Bucket: &bucket, Key: &path})
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(output.Body)
}

// DeleteFile deletes a file from S3
func (client *AWSProvider) DeleteFile(bucket, path string) error {

	s3Client := s3.New(client.sess)
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &path,
	})

	return err
}
