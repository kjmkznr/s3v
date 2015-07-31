package commands

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log <name>",
	Short: "Show history remote object",
	Run:   RunLog,
}

func RunLog(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		fmt.Errorf("need names to be specified")
		return
	}
	fileName := args[0]
	fmt.Printf("[%s]\n", fileName)
	s3c := NewS3Client()
	s3c.listObjectVersions(fileName)
}

func (s3c *S3Client) listObjectVersions(prefix string) {
	input := &s3.ListObjectVersionsInput{
		Bucket: aws.String(s3c.bucketName),
		Prefix: aws.String(prefix),
	}
	output, err := s3c.nativeClient.ListObjectVersions(input)

	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())
	}
	for _, obj := range output.Versions {
		fmt.Printf("%s\t\t%s", *obj.LastModified, *obj.VersionID)
		if *obj.IsLatest {
			fmt.Println(" [LATEST]")
		} else {
			fmt.Println("")
		}
	}

}
