package mlotd

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"github.com/pkg/errors"
	"os"
)

type MetalLinkOfTheDay struct {
	URL        string
	VideoID    string
	VideoTitle string
}

func NewMetalLinkOfTheDay(url string) *MetalLinkOfTheDay {
	return &MetalLinkOfTheDay{
		URL: url,
	}
}

func (m *MetalLinkOfTheDay) GetDetails() error {
	client := youtube.Client{}
	video, err := client.GetVideo(m.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting video: %v\n", err)
		return errors.Wrap(err, "error getting video")
	}

	m.VideoID = video.ID
	m.VideoTitle = video.Title
	return nil
}
