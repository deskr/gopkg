package image

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	simage "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"

	"github.com/nfnt/resize"
)

// Type for supported profile picture types
type Type string

// Size in pixels
type Size struct {
	Width  int
	Height int
}

// Image content and type
type Image struct {
	Content []byte
	Type    Type
}

const (
	// Invalid is an invalid image type
	Invalid Type = ""
	// GIF image/gif
	GIF Type = "image/gif"
	// JPEG image/jpeg
	JPEG Type = "image/jpeg"
	// PNG image/png
	PNG Type = "image/png"
)

var validImageTypes = [...]Type{
	GIF,
	JPEG,
	PNG,
}

// ErrInvalidData when the image data is invalid
var ErrInvalidData = errors.New("Invalid image data")

// ErrInvalidType when an image type is invalid
var ErrInvalidType = errors.New("Invalid image type")

// ProcessError image processing failed (resizing)
type ProcessError struct {
	err error
}

func (v ProcessError) Error() string {
	return fmt.Sprintf("Failed to process image: %+v", v.err)
}

func (v Type) String() string {
	return string(v)
}

// IsValid returns true if the image type is either gif/png/jpeg
func (v Type) IsValid() bool {
	switch v {
	case GIF:
		return true
	case PNG:
		return true
	case JPEG:
		return true
	default:
		return false
	}
}

// NewImage creates a new profile picture
func NewImage(r io.Reader, typ Type, maxSize Size) (pic Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	img, err := createProfileImage(typ, data)
	if err != nil {
		return
	}
	// Resize the image
	img = resizeProfileImage(img, maxSize)

	bytes, err := func(img simage.Image, typ Type) (b []byte, err error) {
		var buf bytes.Buffer
		switch typ {
		case GIF:
			err = gif.Encode(&buf, img, nil)
			break
		case PNG:
			err = png.Encode(&buf, img)
			break
		case JPEG:
			err = jpeg.Encode(&buf, img, nil)
			break
		}
		b = buf.Bytes()
		return
	}(img, typ)

	if err != nil {
		err = ProcessError{err}
		return
	}

	pic = Image{bytes, typ}
	return
}

// Base64 returns base64 encoded data
func (v Image) Base64() string {
	return base64.StdEncoding.EncodeToString(v.Content)
}

// Image returns an image from the Content bytes
func (v Image) Image() (img simage.Image, err error) {
	img, err = createProfileImage(v.Type, v.Content)
	return
}

// IsValid checks if the profile pic is valid
func (v Image) IsValid() bool {
	if len(v.Content) < 10 {
		return false
	}

	if v.Type == Invalid {
		return false
	}
	return true
}

func createProfileImage(typ Type, data []byte) (img simage.Image, err error) {
	switch typ {
	case GIF:
		img, err = gif.Decode(bytes.NewReader(data))
		break
	case PNG:
		img, err = png.Decode(bytes.NewReader(data))
		break
	case JPEG:
		img, err = jpeg.Decode(bytes.NewReader(data))
		break
	default:
		err = ErrInvalidType
		break
	}
	if err != nil {
		err = ProcessError{err}
		return
	}
	return
}

func resizeProfileImage(img simage.Image, maxSize Size) simage.Image {
	s := img.Bounds().Size()
	if s.X > maxSize.Width || s.Y > maxSize.Height {
		return resize.Thumbnail(uint(maxSize.Width), uint(maxSize.Height), img, resize.Lanczos3)
	}
	return img
}
