package emotion

import (
	"github.com/meinside/ms-cognitive-services-go"
)

type Client struct {
	ApiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

// Emotion Recognition in Image
func (c *Client) RecognizeImage(
	image interface{},
	rects []cognitive.Rectangle,
) (emotions []cognitive.Emotion, err error) {
	return cognitive.EmotionRecognizeImage(
		c.ApiKey,
		image,
		rects,
	)
}
