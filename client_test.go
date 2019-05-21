package flickr

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	Client := &client{
		apiKey: "47be9848642f1a968a1d173caa83ab38",
	}

	response, errors := Client.request(params{
		"user_id": "131112890@N03",
	})

	if errors != nil {
		fmt.Println(errors)
	}

	fmt.Printf("%v", string(response))
}
