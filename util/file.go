package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsYamlFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if fileInfo.Mode().IsRegular() && filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml" {
		return true, nil
	}

	return false, nil
}

func IsYamlValid(path string) (bool, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}

	var data interface{}
	err = yaml.Unmarshal([]byte(fileContent), &data)
	if err != nil {
		return false, err
	}

	return true, nil
}
func ProcessYAMLFiles(folderPath string, processFunc func(string, []byte) error) error {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && (strings.HasSuffix(file.Name(), ".yaml") || strings.HasSuffix(file.Name(), ".yml")) {
			filePath := filepath.Join(folderPath, file.Name())
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}
			err = processFunc(filePath, data)
			return err
		}
	}

	return nil
}

func GetFileNameWithoutExtension(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func CreateOutputFileName(id string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s.yml", id, timestamp)
}

func CreateFileWithData(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// 任务类型
type TT int

const (
	Create  TT = iota + 1 //创建k8s任务
	Convert               //转换成k8s任务
)

func CreateNavalYAMLFile(outputPath string, id string, ty TT, waitSeconds int, content string) error {
	if ty == Create {
		templateMeta := `id: %s
	available: true
	type: %d
	wait_seconds: %d
	items:`

		templateContent := `- k8s_yaml_content: |
	%s`
		data := strings.TrimSpace(fmt.Sprintf(templateMeta, id, ty, waitSeconds)) + "\n"
		for _, block := range strings.Split(content, "---") {
			if strings.TrimSpace(block) == "" {
				continue
			}
			data += ("	" + fmt.Sprintf(templateContent, AddPrefixToEveryLine(strings.TrimSpace(block), "		")) + "\n")
		}
		return CreateFileWithData(outputPath, []byte(data))
	} else if ty == Convert {
		templateMeta := `id: %s
		available: true
		type: %d
		wait_seconds: %d
		items:`
		templateContent := `- docker_compose_content: |
	%s`
		data := strings.TrimSpace(fmt.Sprintf(templateMeta, id, ty, waitSeconds)) + "\n"
		for _, block := range strings.Split(content, "---") {
			if strings.TrimSpace(block) == "" {
				continue
			}
			data += ("	" + fmt.Sprintf(templateContent, AddPrefixToEveryLine(strings.TrimSpace(block), "		")) + "\n")
		}
		return CreateFileWithData(outputPath, []byte(data))
	}
	return nil
}

func ReadFile(path string) string {
	data, _ := ioutil.ReadFile(path)
	return string(data)
}

func AddPrefixToEveryLine(content string, prefix string) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		lines[i] = prefix + line
	}

	return strings.Join(lines, "\n")
}
