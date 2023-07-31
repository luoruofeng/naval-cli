package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GeneratePath(host string, port int, path string) string {
	return fmt.Sprintf("http://%s:%d/%s", host, port, path)
}

func FileDo(host string, port int, data []byte, httpMehod string, path string) {
	req, err := http.NewRequest(httpMehod, GeneratePath(host, port, path), bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("创建请求失败", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-yaml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送消息失败", err)
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败", err)
		return
	}
	fmt.Println(string(respData))
}

func AddHttpRequest(host string, port int, filePath string, method string, path string) {
	if r, err := IsPathExists(filePath); err != nil || !r {
		fmt.Println("文件不存在")
		return
	}

	if IsDirectory(filePath) {
		ProcessYAMLFiles(filePath, func(inputPath string, data []byte) error {
			fmt.Println("发送文件:", inputPath)
			FileDo(host, port, []byte(ReadFile(inputPath)), method, path)
			return nil
		})
	} else {
		fmt.Println("发送文件:", filePath)
		FileDo(host, port, []byte(ReadFile(filePath)), method, path)
	}
}

func DeleteHttpRequest(host string, port int, id string, method string, path string) {
	url := GeneratePath(host, port, path)
	req, err := http.NewRequest(http.MethodDelete, url+"/"+id, strings.NewReader(""))
	if err != nil {
		fmt.Println("创建请求失败", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送消息失败", err)
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败", err)
		return
	}
	fmt.Println(string(respData))
}
