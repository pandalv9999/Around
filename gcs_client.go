package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"os"
)

func saveToGCS(r io.Reader, objectName string) (string, error) {
	getEnvVars()
	bucketName := os.Getenv("GCS_BUCKET_NAME")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}

	object := client.Bucket(bucketName).Object(objectName)
	wc := object.NewWriter(ctx)
	if _, err := io.Copy(wc, r); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", err
	}

	fmt.Printf("Image is saved to GCS: %s\n", attrs.MediaLink)
	return attrs.MediaLink, nil

}