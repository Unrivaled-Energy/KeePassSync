package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	endpoint  string
	accessID  string
	accessKey string
	useSSL    bool

	bucketName string
	ctx        context.Context
}

func (m Minio) initMinio() *minio.Client {

	// m.ctx =

	// Initialize minio client object.
	minioClient, err := minio.New(m.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.accessID, m.accessKey, ""),
		Secure: m.useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = minioClient.MakeBucket(m.ctx, m.bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(m.ctx, m.bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", m.bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", m.bucketName)
	}
	return minioClient
}

func (m Minio) uploadfile(minioClient *minio.Client, objectName string, filePath string, contentType string) {

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(m.ctx, m.bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

}
