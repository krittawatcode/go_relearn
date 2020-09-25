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
	"runtime/trace"
	"sync"
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

const MAX_DOWNLOAD = 1000

func main() {
	defer func() {
		fmt.Println("Main exit!")
	}()

	log.SetFlags(log.Ltime)
	dir := "myDownloadImage_" + time.Now().Format("09_04_05")
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, 0700)
	}
	f, err := os.Create(dir + ".trace.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	photos := Photos{}
	err = getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	fmt.Println(len(photos[0:MAX_DOWNLOAD]))

	chImg := make(chan Image, len(photos[0:MAX_DOWNLOAD]))
	counter := sync.WaitGroup{}
	token := make(chan struct{}, 2000)

	for _, v := range photos[0:MAX_DOWNLOAD] {
		v := v
		counter.Add(1)
		go func() {
			defer counter.Done()
			if v.ID > 4500 {
				v.ThumbnailURL = "http://abc.jpg"
			}
			token <- struct{}{}
			img, err := downloadImage(v.ThumbnailURL)
			<-token
			if err != nil {
				log.Println(err)
				return
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

	go func() {
		counter.Wait()
		close(chImg)
	}()

	for v := range chImg {
		err := saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}

}

func saveImage(fileName string, img []byte) error {

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
