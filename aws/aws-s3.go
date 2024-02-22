//API: https://docs.aws.amazon.com/sdk-for-go/api/service/s3/
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var AWS_BUCKET *string = aws.String(os.Getenv("AWS_BUCKET"))
var AWS_REGION *string = aws.String(os.Getenv("AWS_REGION"))

func main() {
	var sess = initSession()
	objects, err := fetchBucketObjects(sess,"")
	if err != nil {
		return
	}
	for _, object := range objects {
		fmt.Println(*object)
	}
	// _ = downloadObject(sess, aws.String("Tumblr_oivwukbROc1v0jfsto3_1280.webp"), AWS_BUCKET)
}

func initSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: AWS_REGION},
	)
	if err != nil {
		fmt.Println(err)
	}
	return sess
}

func fetchBucketObjects(sess *session.Session, search_string string) ([]*s3.Object, error) {
	svc := s3.New(sess)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: AWS_BUCKET,
		Prefix: aws.String(search_string)})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp.Contents, err
}

func downloadObject(sess *session.Session, key *string) error {
	file, err := os.Create(*key)
	if err != nil {
		return err
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: AWS_BUCKET,
			Key:    key,
		})
	if err != nil {
		fmt.Println("failure downloading file")
		return err
	}

	fmt.Println("file downloaded")
	return nil
}
