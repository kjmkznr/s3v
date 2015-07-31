package commands

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/colorstring"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff <version-id> <version-id> <name>",
	Short: "show diff remote object",
	Run:   RunDiff,
}

func RunDiff(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Usage()
		fmt.Errorf("need parameters to be specified")
		return
	}

	v1 := args[0]
	v2 := args[1]
	fileName := args[2]

	s3c := NewS3Client()
	b1, LastModified1 := s3c.getObject(fileName, v1)
	b2, LastModified2 := s3c.getObject(fileName, v2)

	fmt.Printf("--- %s:%s %s\n", v1, fileName, LastModified1)
	fmt.Printf("+++ %s:%s %s\n", v2, fileName, LastModified2)
	printDiff(lineDiff(b1, b2))
}

func (s3c *S3Client) getObject(key, versionID string) (*string, *time.Time) {
	input := &s3.GetObjectInput{
		Bucket:    aws.String(s3c.bucketName),
		VersionID: aws.String(versionID),
		Key:       aws.String(key),
	}

	output, err := s3c.nativeClient.GetObject(input)
	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())
		return nil, nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(output.Body)

	body := buf.String()

	return &body, output.LastModified
}

func lineDiff(src1, src2 *string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(*src1, *src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	return result
}

func printDiff(diffs []diffmatchpatch.Diff) {
	for _, d := range diffs {
		lines := strings.Split(strings.TrimRight(d.Text, "\n"), "\n")

		var prefix string
		switch d.Type {
		case diffmatchpatch.DiffDelete:
			prefix = "[red]- "
		case diffmatchpatch.DiffInsert:
			prefix = "[green]+ "
		case diffmatchpatch.DiffEqual:
			prefix = "  "
		}
		for _, l := range lines {
			s := fmt.Sprintf("%s %s", prefix, l)
			fmt.Println(colorstring.Color(s))
		}
	}
}
