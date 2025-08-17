package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var ImCmd = "magick"

var Templates embed.FS

var HeightMapNameSuffix = "_height"
var MerMapNameSuffix = "_mer"

func GetTextureSubpath(p string, key string) (string, error) {
	subpaths := strings.Split(
		p,
		key,
	)

	if len(subpaths) > 1 {
		return subpaths[1], nil
	}

	return "", errors.New("")
}

func RunCmd(cmd *exec.Cmd) error {
	fmt.Print("66")
	fmt.Println(cmd)

	go cmd.Start()
	return nil
}

func PathToJson[T any](p string) (T, error) {
	var config T
	jsonFile, err := os.Open(p)

	if err != nil {
		return config, err
	}

	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)

	if err != nil {
		return config, err
	}

	json.Unmarshal(data, &config)
	return config, nil
}

func URLToJson[T any](p string) (T, error) {
	var config T

	resp, err := http.Get(p)
	if err != nil {
		return config, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	return parse[T](resp.Body)
}

func parse[T any](r io.Reader) (T, error) {
	var d T
	data, err := io.ReadAll(r)

	fmt.Println(string(data))

	if err != nil {
		fmt.Println("err")
		return d, err
	}

	json.Unmarshal(data, &d)

	fmt.Println("data")
	fmt.Println(d)

	return d, nil
}
