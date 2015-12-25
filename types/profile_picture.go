package types

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/nfnt/resize"
)

// ImageType for supported profile picture types
type ImageType string

// ImageSize in pixels
type ImageSize struct {
	Width  int
	Height int
}

// ProfilePicture for pic
type ProfilePicture struct {
	Content   []byte
	ImageType ImageType
}

const (
	// ImageInvalid is an invalid image type
	ImageInvalid ImageType = ""
	// ImageGIF image/gif
	ImageGIF ImageType = "image/gif"
	// ImageJPEG image/jpeg
	ImageJPEG ImageType = "image/jpeg"
	// ImagePNG image/png
	ImagePNG ImageType = "image/png"
)

var validImageTypes = [...]ImageType{
	ImageGIF,
	ImageJPEG,
	ImagePNG,
}

// ErrInvalidImageData when the image data is invalid
var ErrInvalidImageData = errors.New("Invalid image data")

// ErrInvalidImageType when an image type is invalid
var ErrInvalidImageType = errors.New("Invalid image type")

// ErrProcessImage image processing failed (resizing)
type ErrProcessImage struct {
	err error
}

func (v ErrProcessImage) Error() string {
	return fmt.Sprintf("Failed to process image: %+v", v.err)
}

func (v ImageType) String() string {
	return string(v)
}

// NewProfilePicture creates a new profile picture
func NewProfilePicture(data []byte, maxSize ImageSize) (pic ProfilePicture, err error) {
	if len(data) < 10 {
		err = ErrInvalidImageData
		return
	}

	typ, err := detectImageType(data)
	if err != nil {
		return
	}

	img, err := createProfileImage(*typ, data)
	if err != nil {
		return
	}
	// Resize the image
	img = resizeProfileImage(img, maxSize)

	bytes, err := func(img image.Image, typ ImageType) (b []byte, err error) {
		var buf bytes.Buffer
		switch typ {
		case ImageGIF:
			err = gif.Encode(&buf, img, nil)
			break
		case ImagePNG:
			err = png.Encode(&buf, img)
			break
		case ImageJPEG:
			err = jpeg.Encode(&buf, img, nil)
			break
		}
		b = buf.Bytes()
		return
	}(img, *typ)

	if err != nil {
		err = ErrProcessImage{err}
		return
	}

	pic = ProfilePicture{bytes, *typ}
	return
}

// Base64 returns base64 encoded data
func (v ProfilePicture) Base64() string {
	return base64.StdEncoding.EncodeToString(v.Content)
}

// Image returns an image from the Content bytes
func (v ProfilePicture) Image() (img image.Image, err error) {
	img, err = createProfileImage(v.ImageType, v.Content)
	return
}

// IsValid checks if the profile pic is valid
func (v ProfilePicture) IsValid() bool {
	if len(v.Content) < 10 {
		return false
	}

	if v.ImageType == ImageInvalid {
		if _, err := detectImageType(v.Content); err != nil {
			return false
		}
	}
	return true
}

func createProfileImage(typ ImageType, data []byte) (img image.Image, err error) {
	switch typ {
	case ImageGIF:
		img, err = gif.Decode(bytes.NewReader(data))
		break
	case ImagePNG:
		img, err = png.Decode(bytes.NewReader(data))
		break
	case ImageJPEG:
		img, err = jpeg.Decode(bytes.NewReader(data))
		break
	default:
		err = ErrInvalidImageType
		break
	}
	if err != nil {
		err = ErrProcessImage{err}
		return
	}
	return
}

func resizeProfileImage(img image.Image, maxSize ImageSize) image.Image {
	s := img.Bounds().Size()
	if s.X > maxSize.Width || s.Y > maxSize.Height {
		return resize.Thumbnail(uint(maxSize.Width), uint(maxSize.Height), img, resize.Lanczos3)
	}
	return img
}

func detectImageType(data []byte) (*ImageType, error) {
	typ := http.DetectContentType(data)
	for _, t := range validImageTypes {
		if typ == string(t) {
			return &t, nil
		}
	}
	return nil, ErrInvalidImageType
}
