package report

import (
	"bytes"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (r Reporter) GetExistingData() (*Report, error) {
	// (1) get from s3
	sess := session.New()
	svc := s3.New(sess, aws.NewConfig().WithRegion(r.Conf.Region))

	results, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(r.Conf.S3Bucket),
		Key:    aws.String(r.Conf.S3KeyData),
	})
	if err != nil {
		return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	fileContentCompressed := buf.Bytes()
	// (2) uncompress - TODO
	fileContentUnompressed := string(fileContentCompressed)
	// (3) parse
	return parse(fileContentUnompressed)
}

func (r Reporter) UpdateMeasurements(prevReport *Report, newData Check) {
	// (1) serialize new measurements, append to existing data
	prevReport.checks = append(prevReport.checks, newData)

	var serializedUncompressed string
	for _, c := range prevReport.checks {
		serializedUncompressed += c.String()
		serializedUncompressed += "\n"
	}

	// (2) compress - TODO
	serializedCompressed := serializedUncompressed

	// (3) upload measurements
	sess := session.New()
	svc := s3.New(sess, aws.NewConfig().WithRegion(r.Conf.Region))
	svc.PutObject(&s3.PutObjectInput{
		Bucket: &r.Conf.S3Bucket,
		Key:    &r.Conf.S3KeyData,
		Body:   strings.NewReader(serializedCompressed),
	})

	// (4) Generate report

	// (5) Upload report
	svc.PutObject(&s3.PutObjectInput{
		Bucket: &r.Conf.S3Bucket,
		Key:    &r.Conf.S3KeyReport,
		Body:   strings.NewReader("<html><head><title>Report</title></head><body><h1>the report</h1></body></html>"),
	})
}
