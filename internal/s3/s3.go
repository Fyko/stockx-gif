package s3

import (
	"bytes"
	"fmt"
	"os"
	"stockx-gif-next/pkg/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var bucket = os.Getenv("AWS_S3_BUCKET")
var region = os.Getenv("AWS_S3_BUCKET_REGION")
var awsId = os.Getenv("AWS_ID")
var awsSecret = os.Getenv("AWS_SECRET")
var config = &aws.Config{
	Region:      aws.String(region),
	Credentials: credentials.NewStaticCredentials(awsId, awsSecret, ""),
}

var sess = session.Must(session.NewSession(config))
var log = logger.CreateLogger()

func FetchExisting(id string, preview bool) string {
	var key string
	if preview {
		key = fmt.Sprintf("%s_preview.gif", id)
	} else {
		key = fmt.Sprintf("%s.gif", id)
	}

	svc := s3.New(sess)
	_, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return ""
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
}

func UploadGIF(id string, preview bool, gif *bytes.Buffer) string {
	var key string
	if preview {
		key = fmt.Sprintf("%s_preview.gif", id)
	} else {
		key = fmt.Sprintf("%s.gif", id)
	}

	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             &bucket,
		Body:               gif,
		Key:                &key,
		ACL:                aws.String("public-read"),
		ContentType:        aws.String("image/gif"),
		ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		log.Errorf("Error uploading %s to S3 %v", key, err)
		return ""
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
}
