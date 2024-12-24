package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func PutInBucket() {
	fmt.Println("akldsjksd")
}
