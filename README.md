## go-flickr-random-image

A minimalistic Flickr random image API client for Go

## Install

```
$ go get github.com/hanknguyen14/go-flickr-random-image
```

### GetRandomImage
Get a random image of a user.

```go
import (
  "github.com/hanknguyen14/go-flickr-random-image"
)

Client := &flickr.client{
  apiKey: "key",
  apiToken: "token", // optional
}

userId := "123123123"
numberOfImages := 50; //Random in 50 first images

response, err := Client.getRandomImage(userId, numberOfImages)
```
