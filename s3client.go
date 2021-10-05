package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"archive/zip"
	"net/http"
	"io"
	"path"
)

type FakeWriterAt struct {
    w io.Writer
}

func (fw FakeWriterAt) WriteAt(p []byte, offset int64) (n int, err error) {
    // ignore 'offset' because we forced sequential downloads
    return fw.w.Write(p)
}

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
	Thumbnail  string `json:"thumbnail"`
	Video bool `json:"video"`
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

		
		thumbnailKey := "thumbnail/" + *item.Key
		
		if (isVideo(*item.Key)) {
			thumbnailKey = thumbnailKey + ".jpg"
		}

		thumbnail := c.generateSignedURL(thumbnailKey)

		objects = append(objects, S3Object{
			Key:  *item.Key,
			Name: *item.Key,
			Url:  c.generateSignedURL(*item.Key),
			Thumbnail: thumbnail,
			Video: isVideo(*item.Key),
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

func (c *S3Client) bulkDownload(w http.ResponseWriter, keys []string) {

	downloader := s3manager.NewDownloaderWithClient(c.svc)
	downloader.Concurrency = 1

	zipWriter := zip.NewWriter(w)
	
	for _, fn := range keys {
		fmt.Printf("Processing '%s'\n", fn)

		zw, err := zipWriter.Create(path.Base(fn))

		if err != nil {
			panic(err)
		}

		_, err = downloader.Download(FakeWriterAt{zw},
			&s3.GetObjectInput{
				Bucket: aws.String(c.bucket),
				Key:    aws.String(fn),
			})

		if err != nil {
			panic(err)
		}
	}

	zipWriter.Close()

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

func isVideo(key string) bool {
	return (
		strings.HasSuffix(strings.ToLower(key), ".avi") || 
		strings.HasSuffix(strings.ToLower(key), ".3gp") || 
		strings.HasSuffix(strings.ToLower(key), ".mp4") )
}