package commands

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List remote object list",
	Run:   RunList,
}

func RunList(cmd *cobra.Command, args []string) {
	s3c := NewS3Client()
	s3c.listObjects()
}

func (s3c *S3Client) listObjects() {
	input := &s3.ListObjectsInput{
		Bucket: aws.String(s3c.bucketName),
	}
	output, err := s3c.nativeClient.ListObjects(input)
	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())
	}
	for _, obj := range output.Contents {
		fmt.Printf("%s\t\t%s\n", *obj.LastModified, *obj.Key)
	}
}
