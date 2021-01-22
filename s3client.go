package main

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
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

	var objects []*s3.ObjectIdentifier
	for _, key := range keys {
		objects = append(objects, &s3.ObjectIdentifier{
			Key: &key,
		})
	}

	_, err := c.svc.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(c.bucket),
		Delete: &s3.Delete{
			Objects: objects,
		},
	})
	//
	//signedURL, err := req.Presign(24 * time.Hour)
	//
	if err != nil {
		print("Failed to sign request", err)
	}
	/**
	curl -X DELETE localhost:8090/api/objects -H "Accept: application/json" -d "[\"helllow\", \"world\"]"
	*/

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

	print("The URL is", signedURL)

	return signedURL
}
