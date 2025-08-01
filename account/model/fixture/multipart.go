<<<<<<< HEAD
package fixture

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"runtime"
)

// MultipartImage used for instantiating a test fixture
// for creating multipart file uploads containing an image
type MultipartImage struct {
	imagePath     string
	ImageFile     *os.File
	MultipartBody *bytes.Buffer
	ContentType   string
}

// NewMultipartImage creates an image file for testing
// and creates a Multipart Form with this image file
// for testing
func NewMultipartImage(fileName string, contentType string) *MultipartImage {
	// create test file in same folder as this fixture
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)

	imagePath := filepath.Join(dir, fileName)

	// f, _ := os.Open(imagePath)
	f := createImage(imagePath)

	defer f.Close()

	// create a multipart write onto which we
	// will write the image file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// manually create form file as CreateFormFile will
	// force file's content type to "application/octet-stream"
	h := make(textproto.MIMEHeader)
	h.Set(
		"Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "imageFile", fileName),
	)
	h.Set("Content-Type", contentType)
	part, _ := writer.CreatePart(h)

	io.Copy(part, f)
	writer.Close()

	return &MultipartImage{
		imagePath:     imagePath,
		ImageFile:     f,
		MultipartBody: body,
		ContentType:   writer.FormDataContentType(),
	}
}

// GetFormFile extracts form file from multipart body
func (m *MultipartImage) GetFormFile() *multipart.FileHeader {
	_, params, _ := mime.ParseMediaType(m.ContentType)
	mr := multipart.NewReader(m.MultipartBody, params["boundary"])
	form, _ := mr.ReadForm(1024)
	files := form.File["imageFile"]

	return files[0]
}

// Close removes created file for test
func (m *MultipartImage) Close() {
	m.ImageFile.Close()
	os.Remove(m.imagePath)
}

// createImage used to create a quick example
// 1px x 1px image encoded as a PNG
func createImage(imagePath string) *os.File {
	rect := image.Rect(0, 0, 1, 1)
	img := image.NewRGBA(rect)

	f, _ := os.Create(imagePath)
	png.Encode(f, img)

	return f
}
=======
package fixture

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"runtime"
)

// MultipartImage used for instantiating a test fixture
// for creating multipart file uploads containing an image
type MultipartImage struct {
	imagePath     string
	ImageFile     *os.File
	MultipartBody *bytes.Buffer
	ContentType   string
}

// NewMultipartImage creates an image file for testing
// and creates a Multipart Form with this image file
// for testing
func NewMultipartImage(fileName string, contentType string) *MultipartImage {
	// create test file in same folder as this fixture
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)

	imagePath := filepath.Join(dir, fileName)

	// f, _ := os.Open(imagePath)
	f := createImage(imagePath)

	defer f.Close()

	// create a multipart write onto which we
	// will write the image file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// manually create form file as CreateFormFile will
	// force file's content type to "application/octet-stream"
	h := make(textproto.MIMEHeader)
	h.Set(
		"Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "imageFile", fileName),
	)
	h.Set("Content-Type", contentType)
	part, _ := writer.CreatePart(h)

	io.Copy(part, f)
	writer.Close()

	return &MultipartImage{
		imagePath:     imagePath,
		ImageFile:     f,
		MultipartBody: body,
		ContentType:   writer.FormDataContentType(),
	}
}

// GetFormFile extracts form file from multipart body
func (m *MultipartImage) GetFormFile() *multipart.FileHeader {
	_, params, _ := mime.ParseMediaType(m.ContentType)
	mr := multipart.NewReader(m.MultipartBody, params["boundary"])
	form, _ := mr.ReadForm(1024)
	files := form.File["imageFile"]

	return files[0]
}

// Close removes created file for test
func (m *MultipartImage) Close() {
	m.ImageFile.Close()
	os.Remove(m.imagePath)
}

// createImage used to create a quick example
// 1px x 1px image encoded as a PNG
func createImage(imagePath string) *os.File {
	rect := image.Rect(0, 0, 1, 1)
	img := image.NewRGBA(rect)

	f, _ := os.Create(imagePath)
	png.Encode(f, img)

	return f
}
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
