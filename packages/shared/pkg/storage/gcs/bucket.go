package gcs

import (
	"cloud.google.com/go/storage"

	"github.com/khulnasoft/infra/packages/shared/pkg/utils"
)

type BucketHandle = storage.BucketHandle

func newBucket(bucket string) *BucketHandle {
	return client.Bucket(bucket)
}

var (
	templateBucketName = utils.RequiredEnv("TEMPLATE_BUCKET_NAME", "bucket for storing template files")

	TemplateBucket = newBucket(templateBucketName)
)
