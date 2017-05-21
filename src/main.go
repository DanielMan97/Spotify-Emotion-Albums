package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/meinside/ms-cognitive-services-go"
)

const (
	EmotionApiKey = "Enter Emotion API Key"
)

type MapSF struct {
	key   string
	value float64
}

type MapSFList []MapSF

func (p MapSFList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p MapSFList) Len() int      { return len(p) }
func (p MapSFList) Less(i, j int) bool {
	return p[i].value < p[j].value
}
func sortMapByValue(m map[string]float64) MapSFList {
	p := make(MapSFList, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapSF{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	if imgBytes, err := ioutil.ReadFile("File Directory to Image!"); err == nil {
		if emotions, err := cognitive.EmotionRecognizeImage(
			EmotionApiKey,
			imgBytes,
			nil,
		); err == nil {
			var m map[string]float64
			m = make(map[string]float64)
			m["surprise"] = emotions[0].Scores["surprise"]
			m["fear"] = emotions[0].Scores["fear"]
			m["happiness"] = emotions[0].Scores["happiness"]
			m["neutral"] = emotions[0].Scores["neutral"]
			m["sadness"] = emotions[0].Scores["sadness"]
			m["anger"] = emotions[0].Scores["anger"]
			m["contempt"] = emotions[0].Scores["contempt"]
			m["disgust"] = emotions[0].Scores["disgust"]

			searchSpotify(sortMapByValue(m)[7].key)
		} else {
			fmt.Printf("Could not recognize image! %s\n", err)
		}
	} else {
		fmt.Printf("Error file Read.\n")
	}
}
