package utils

import (
	"bytes"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func Upload(fileDir string, data []byte) error {
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	bucket, err := client.Bucket("mydousheng")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = bucket.PutObject(fileDir, bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
func Delete(fileDir string) error {
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	bucket, err := client.Bucket("mydousheng")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = bucket.DeleteObject(fileDir)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
