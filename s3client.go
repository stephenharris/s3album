package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	svc    *s3.S3
	bucket string
}

type S3ListObjectsOut struct {
	Objects []S3Object `json:"objects"`
	Folders []S3Folder `json:"folders"`
}

type S3Object struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type S3Folder struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func (c *S3Client) listObjects(path string) S3ListObjectsOut {

	bucket := c.bucket
	prefix := path
	resp, err := c.svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}

	var objects []S3Object
	for _, item := range resp.Contents {
		objects = append(objects, S3Object{
			Key:  *item.Key,
			Name: *item.Key,
			Url:  c.generateSignedURL(*item.Key),
		})
	}

	var folders []S3Folder
	for _, item := range resp.CommonPrefixes {
		print(*item.Prefix)
		folders = append(folders, S3Folder{
			Key:  *item.Prefix,
			Name: strings.TrimSuffix(strings.TrimPrefix(*item.Prefix, path), "/"),
		})
	}

	return S3ListObjectsOut{
		Objects: objects,
		Folders: folders,
	}

}

func (c *S3Client) bulkDeleteKeys(keys []string) {

	objectsToDelete := make([]*s3.ObjectIdentifier, 0, len(keys))
	for _, key := range keys {
		print(key + "\n")
		obj := s3.ObjectIdentifier{
			Key: aws.String(key),
		}
		objectsToDelete = append(objectsToDelete, &obj)
	}

	deleteArray := s3.Delete{Objects: objectsToDelete}
	deleteParams := &s3.DeleteObjectsInput{
		Bucket: aws.String(c.bucket),
		Delete: &deleteArray,
	}

	result, err := c.svc.DeleteObjects(deleteParams)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

	if err != nil {
		print("Failed to sign request", err)
	}

}

func (c *S3Client) generateSignedURL(key string) string {

	req, _ := c.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})

	signedURL, err := req.Presign(24 * time.Hour)

	if err != nil {
		print("Failed to sign request", err)
	}

	return signedURL
}
