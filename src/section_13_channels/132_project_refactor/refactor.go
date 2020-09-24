package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	// fmt.Println(photos)

	dir := "myDownloadImage"
	if _, err = os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModeDir)
	}

	for _, v := range photos {
		img, err := downloadImage(v.ThumbnailURL)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(img)

		format, err := decodeImage(img)
		// fmt.Println("Format image : ", format)
		if err != nil {
			log.Fatal(err)
		}

		fileName := fmt.Sprintf("%d.%s", v.ID, format)
		// "abc" + "." + format
		err = saveImage(filepath.Join(dir, fileName), img)
		if err != nil {
			log.Println(err)
		}
	}

}

func saveImage(fileName string, img []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(img))
	if err != nil {
		return err
	}
	return nil
}

func decodeImage(img []byte) (string, error) {
	_, format, err := image.Decode(bytes.NewReader(img))
	return format, err
}

func downloadImage(url string) ([]byte, error) {

	errMsg := func(err error) error {
		return fmt.Errorf("downloadImage : %v", err)
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, errMsg(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errMsg(err)
	}
	return body, nil
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
