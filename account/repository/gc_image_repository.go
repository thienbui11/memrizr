<<<<<<< HEAD
package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/thienbui11/MindStack/account/model"
	"github.com/thienbui11/MindStack/account/model/apperrors"
)

type gcImageRepository struct {
	Storage    *storage.Client
	BucketName string
}

// NewImageRepository is a factory for initializing User Repositories
func NewImageRepository(gcClient *storage.Client, bucketName string) model.ImageRepository {
	return &gcImageRepository{
		Storage:    gcClient,
		BucketName: bucketName,
	}
}

func (r *gcImageRepository) DeleteProfile(ctx context.Context, objName string) error {
	bckt := r.Storage.Bucket(r.BucketName)

	object := bckt.Object(objName)

	if err := object.Delete(ctx); err != nil {
		log.Printf("Failed to delete image object with ID: %s from GC Storage\n", objName)
		return apperrors.NewInternal()
	}

	return nil
}

func (r *gcImageRepository) UpdateProfile(
	ctx context.Context,
	objName string,
	imageFile multipart.File,
) (string, error) {
	bckt := r.Storage.Bucket(r.BucketName)

	object := bckt.Object(objName)
	wc := object.NewWriter(ctx)

	// set cache control so profile image will be served fresh by browsers
	// To do this with object handle, you'd first have to upload, then update
	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"

	// multipart.File has a reader!
	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("Unable to write file to Google Cloud Storage: %v\n", err)
		return "", apperrors.NewInternal()
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	imageURL := fmt.Sprintf(
		"https://storage.googleapis.com/%s/%s",
		r.BucketName,
		objName,
	)

	return imageURL, nil
}
=======
package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/jacobsngoodwin/memrizr/account/model"
	"github.com/jacobsngoodwin/memrizr/account/model/apperrors"
)

type gcImageRepository struct {
	Storage    *storage.Client
	BucketName string
}

// NewImageRepository is a factory for initializing User Repositories
func NewImageRepository(gcClient *storage.Client, bucketName string) model.ImageRepository {
	return &gcImageRepository{
		Storage:    gcClient,
		BucketName: bucketName,
	}
}

func (r *gcImageRepository) DeleteProfile(ctx context.Context, objName string) error {
	bckt := r.Storage.Bucket(r.BucketName)

	object := bckt.Object(objName)

	if err := object.Delete(ctx); err != nil {
		log.Printf("Failed to delete image object with ID: %s from GC Storage\n", objName)
		return apperrors.NewInternal()
	}

	return nil
}

func (r *gcImageRepository) UpdateProfile(
	ctx context.Context,
	objName string,
	imageFile multipart.File,
) (string, error) {
	bckt := r.Storage.Bucket(r.BucketName)

	object := bckt.Object(objName)
	wc := object.NewWriter(ctx)

	// set cache control so profile image will be served fresh by browsers
	// To do this with object handle, you'd first have to upload, then update
	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"

	// multipart.File has a reader!
	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("Unable to write file to Google Cloud Storage: %v\n", err)
		return "", apperrors.NewInternal()
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	imageURL := fmt.Sprintf(
		"https://storage.googleapis.com/%s/%s",
		r.BucketName,
		objName,
	)

	return imageURL, nil
}
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
