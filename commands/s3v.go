package commands

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var S3VCmd = &cobra.Command{
	Use:   "s3v",
	Short: "tools for S3 versioning object",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var Region, Bucket string

func Execute() {
	AddCommands()

	S3VCmd.PersistentFlags().StringVarP(&Region, "region", "", "ap-northeast-1", "S3 bucket region")
	S3VCmd.PersistentFlags().StringVarP(&Bucket, "bucket", "", "", "S3 bucket")
	S3VCmd.Execute()
}

func AddCommands() {
	S3VCmd.AddCommand(lsCmd)
	S3VCmd.AddCommand(logCmd)
	S3VCmd.AddCommand(diffCmd)
}

type S3Client struct {
	nativeClient *s3.S3
	bucketName   string
}

func NewS3Client() *S3Client {
	awsConfig := &aws.Config{
		Region: &Region,
	}

	nativeClient := s3.New(awsConfig)
	return &S3Client{
		nativeClient: nativeClient,
		bucketName:   Bucket,
	}
}
