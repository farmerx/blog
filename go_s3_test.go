package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	bucket := aws.String("hatlonely")
	key := aws.String("test/s3_test.go")

	// Configure to use Minio Server
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("VmpKMFUxTnJOVmhXYkdoVFlXczFUVnBXYUU1alp5c3I+", "6642a81c991c42c989f7f42a637e6039", ""),
		Endpoint:         aws.String("http://10.95.162.158:50840"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)

	s3Client := s3.New(newSession)

	cparams := &s3.CreateBucketInput{
		Bucket: bucket, // Required
	}

	// Create a new bucket using the CreateBucket call.
	_, err := s3Client.CreateBucket(cparams)
	if err != nil {
		// Message from an error.
		fmt.Println(err.Error())
	}

	fp, err := os.Open("s3_test.go")

	defer fp.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	_, err = s3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String("hatlonely"),
		Key:    aws.String("test/s3_test.go"),
		Body:   fp,
	})

	// // // Upload a new object "testobject" with the string "Hello World!" to our "newbucket".
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("Hello from Minio!!"),
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n", *bucket, *key, err.Error())
		return
	}
	fmt.Printf("Successfully created bucket %s and uploaded data with key %s\n", *bucket, *key)

	out, err := s3Client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String("hatlonely"),
		Key:    aws.String("test/s3_test.go"),
	})
	defer out.Body.Close()

	file, err := os.Create("testobject_local")
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}
	defer file.Close()

	x := bufio.NewReader(out.Body)
	numBytes, err := x.WriteTo(file)
	if err != nil {
		fmt.Println("Write to File Err:", err.Error())
	}
	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")
}
