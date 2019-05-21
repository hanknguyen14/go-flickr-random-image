package flickr

import (
	"testing"
)

func TestGetRandomImage(t *testing.T) {
	Client := &client{
		apiKey: "47be9848642f1a968a1d173caa83ab38",
	}

	photo, errors := Client.getRandomImage("131112890@N03", 50)

	if errors != nil {
		t.Error(errors)
	}

	if len(photo.Title) == 0 {
		t.Error("Invalid Image Title")
	}

	if len(photo.URL) == 0 {
		t.Error("Invalid Image URL")
	}

	if len(photo.Filetype) == 0 {
		t.Error("Invalid Image File Type")
	}
}
