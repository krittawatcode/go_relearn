package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Photosx []struct {
	AlbumID       int    `json:"albumId"`
	ID            int    `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	ThumbnailURL  string `json:"thumbnailUrl"`
	ThumbnailURL2 string `json:"thumbnailUrl2"`
}

func main() {
	// https://jsonplaceholder.typicode.com/photos
	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	fmt.Println(photos)
}

func getJson(url string, structType interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	switch v := structType.(type) {
	case *Photos:
		decoder := json.NewDecoder(res.Body)
		photos := structType.(*Photos)
		decoder.Decode(photos)
		fmt.Println("in *photos")
		return nil
	default:
		return fmt.Errorf("getJson : not support type %v ", v)
	}
}
