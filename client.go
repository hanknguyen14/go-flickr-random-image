package flickr

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	host   string = "api.flickr.com/services/rest/"
	method string = "flickr.photos.search"
)

type params map[string]string

type client struct {
	apiKey   string
	apiToken string
}

func (Client *client) request(Params params) ([]byte, error) {
	url := fmt.Sprintf("https://%s?method=%s&api_key=%s&format=json&nojsoncallback=1", host, method, Client.apiKey)

	if len(Client.apiToken) > 0 {
		url = fmt.Sprintf("%s&auth_token=%s", url, Client.apiToken)
	}

	for key, value := range Params {
		url = fmt.Sprintf("%s&%s=%s", url, key, value)
	}

	response, errors := http.Get(url)

	if errors != nil {
		return nil, errors
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
