package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"os"
	"strings"
)

// To hold the number retrieved files
var numberOfRetrievedFiles = 0

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : ", os.Args[0],
			" s3bucket destDirPath")
		os.Exit(1)
	}
	fmt.Println("Getting all files from the s3 bucket :", os.Args[1])
	fmt.Println("And will download them to :", os.Args[2])
	sess := makeSession()
	getBucketObjects(sess)
	// Print number of retrieved files
	fmt.Printf("We got %d files from our s3 bucket\n",
		numberOfRetrievedFiles)
}

func makeSession() *session.Session {
	// Specify profile to load for the session's config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})
	if err != nil {
		fmt.Println("failed to create session,", err)
		fmt.Println(err)
		os.Exit(2)
	}

	return sess
}

func getBucketObjects(sess *session.Session) {
	query := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Args[2]),
	}
	svc := s3.New(sess)

	// Flag used to check if we need to go further
	truncatedListing := true

	for truncatedListing {
		resp, err := svc.ListObjectsV2(query)

		if err != nil {
			// Print the error.
			fmt.Println(err.Error())
			return
		}
		// Get all files
		getObjectsAll(resp, svc)
		// Set continuation token
		query.ContinuationToken = resp.NextContinuationToken
		truncatedListing = *resp.IsTruncated
	}
}

func getObjectsAll(bucketObjectsList *s3.ListObjectsV2Output, s3Client *s3.S3) {
	//fmt.Println("One ring to rule them all")
	// Iterate through the files inside the bucket
	for _, key := range bucketObjectsList.Contents {
		fmt.Println(*key.Key)
		destFilename := *key.Key
		if strings.HasSuffix(*key.Key, "/") {
			fmt.Println("Got a directory")
			continue
		}
		numberOfRetrievedFiles++
		if strings.Contains(*key.Key, "/") {
			var dirTree string
			// split
			s3FileFullPathList := strings.Split(*key.Key, "/")
			fmt.Println(s3FileFullPathList)
			fmt.Println("destFilename " + destFilename)
			for _, dir := range s3FileFullPathList[:len(s3FileFullPathList)-1] {
				dirTree += "/" + dir
			}
			err := os.MkdirAll(os.Args[3]+"/"+dirTree, 0775)
			if err != nil {
				log.Fatal(err)
			}
		}
		out, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(os.Args[2]),
			Key:    key.Key,
		})
		if err != nil {
			log.Fatal(err)
		}
		destFilePath := os.Args[3] + destFilename
		destFile, err := os.Create(destFilePath)
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := io.Copy(destFile, out.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("File %s contanin %d bytes\n", destFilePath, bytes)
		_ := out.Body.Close()
		_ := destFile.Close()
	}
}
