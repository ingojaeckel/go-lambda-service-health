package report

import (
	"bytes"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

func (r Reporter) GetExistingData() (*Report, error) {
	// (1) get from s3
	sess := session.New()
	svc := s3.New(sess, aws.NewConfig().WithRegion(r.Conf.Region))

	log.Printf("Loading report from bucket=%s, key=%s, region=%s...\n", r.Conf.S3Bucket, r.Conf.S3KeyData, r.Conf.Region)
	results, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(r.Conf.S3Bucket),
		Key:    aws.String(r.Conf.S3KeyData),
	})
	if err != nil {
		log.Printf("Failed to load existing data: %s\n", err.Error())
		return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	fileContentCompressed := buf.Bytes()
	// (2) uncompress - TODO
	fileContentUncompressed := string(fileContentCompressed)
	// (3) parse
	return parse(fileContentUncompressed)
}

func (r Reporter) UpdateMeasurements(prevReport *Report, newData Check) error {
	// (1) serialize new measurements, append to existing data
	prevReport.Checks = append(prevReport.Checks, newData)

	var serializedUncompressed string
	for _, c := range prevReport.Checks {
		serializedUncompressed += c.String()
		serializedUncompressed += "\n"
	}
	// (2) compress - TODO
	serializedCompressed := serializedUncompressed

	// (3) upload measurements
	log.Printf("Uploading data to s3://%s/%s [%s]\n", r.Conf.S3Bucket, r.Conf.S3KeyData, r.Conf.Region)
	sess := session.New()
	svc := s3.New(sess, aws.NewConfig().WithRegion(r.Conf.Region))
	svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(r.Conf.S3Bucket),
		Key:         aws.String(r.Conf.S3KeyData),
		Body:        strings.NewReader(serializedCompressed),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		ContentType: aws.String("text/plain"),
	})

	// (4) Generate report
	maxAge := time.Now().Add(-1 * time.Duration(r.Conf.RetainDataForSeconds) * time.Second)

	htmlStr := GenerateReport(*prevReport, maxAge)
	log.Printf("Uploading report to s3://%s/%s [%s]\n", r.Conf.S3Bucket, r.Conf.S3KeyReport, r.Conf.Region)
	// (5) Upload report
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(r.Conf.S3Bucket),
		Key:         aws.String(r.Conf.S3KeyReport),
		Body:        strings.NewReader(htmlStr),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		ContentType: aws.String("text/html"),
	})
	return err
}
