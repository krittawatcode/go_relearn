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
	"time"
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

type Image struct {
	filePath string
	img      []byte
}

func main() {
	defer func() {
		fmt.Println("Main exit!")
	}()
	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	// fmt.Println(photos)

	dir := "myDownloadImage_" + time.Now().Format("09_04_05")
	chImg := make(chan Image, len(photos))

	for _, v := range photos {
		v := v
		go func() {
			img, err := downloadImage(v.ThumbnailURL)
			if err != nil {
				log.Fatal(err)
			}
			format, err := decodeImage(img)
			if err != nil {
				log.Fatal(err)
			}
			// log.Printf("Downloaded : %v\n", v.ID)
			absolutefileName := filepath.Join(dir, fmt.Sprintf("%d.%s", v.ID, format))
			chImg <- Image{filePath: absolutefileName, img: img}
		}()
	}

	for range photos {
		v := <-chImg
		saveImage(dir, v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}
	time.Sleep(10 * time.Second)
}

func saveImage(dir string, fileName string, img []byte) error {

	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, 0700)
	}

	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("saveImage - cannot create file: %v", err)
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(img))
	if err != nil {
		return fmt.Errorf("saveImage - cannot save file: %v", err)
	}
	log.Printf("\tSaved : %v\n", fileName)
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
