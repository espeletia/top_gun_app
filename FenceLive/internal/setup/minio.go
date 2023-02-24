package setup

import (
	"FenceLive/internal/config"
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SetupMinio(config *config.Config) (*minio.Client, error) {
	ctx := context.Background()
	minioClient, err := minio.New(config.MinioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioConfig.AccessKeyId, config.MinioConfig.SecretAccessKey, ""),
		Secure: config.MinioConfig.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	bucketName := config.MinioConfig.Bucket
	location := config.MinioConfig.Location

	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil && !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Fatalln(err)
		}
	}
	return minioClient, nil
}
