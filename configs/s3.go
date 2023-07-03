package configs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	//"s3_blog/internal/config"
)

var (
	awsAccessKeyID     = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsBucketRegion    = os.Getenv("AWS_REGION")
	awsBucketName      = os.Getenv("BUCKET_NAME")
)

var sess = CreateSession()
var s3Session = CreateS3Session(sess)

// This function creates session and requires AWS credentials
func CreateSession() *session.Session {

	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(awsBucketRegion),
			Credentials: credentials.NewStaticCredentials(
				awsAccessKeyID,
				awsSecretAccessKey,
				"",
			),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}
