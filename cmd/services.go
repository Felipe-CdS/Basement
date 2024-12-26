package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func GetBucketList(bucketURL string) ([]string, error) {

	res, err := http.Get(bucketURL)
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	type BucketBody struct {
		Name string `json:"name"`
	}

	type P struct {
		Objects []BucketBody `json:"objects"`
	}

	var data P

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Println(err)
	}

	var namesList []string

	for _, o := range data.Objects {
		namesList = append(namesList, o.Name)
	}

	return namesList, nil
}

func PutInBucket(bucketURL string, fr multipart.File, fn string, fs int) error {

	//setup the req
	url := fmt.Sprintf("%s%s", bucketURL, fn)
	req, err := http.NewRequest(http.MethodPut, url, fr)
	req.ContentLength = int64(fs)
	req.Header.Add("Cache-Control", "public, max-age=31536000, immutable")

	if err != nil {
		return fmt.Errorf("Error creating request")
	}

	//exec the req and get the resp
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("Error sending request to bucket...")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Resp status code: %d", resp.StatusCode)
	}

	return nil
}
