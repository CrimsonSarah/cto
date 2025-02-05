package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	checkDir()

	// Is this governed by any TOS??
	// TODO: Get card list from https://documenter.getpostman.com/view/14059948/TzecB4fH
	resp, err := http.Get("https://images.digimoncard.io/images/cards/BT5-103.jpg")

	if err != nil {
		panic(err)
	}

	image := make([]byte, resp.ContentLength)
	io.ReadFull(resp.Body, image)
	os.WriteFile("data/textures/cards/BT5-103.jpg", image, 0644)
}

func checkDir() {
	fileBytes, err := os.ReadFile("go.mod")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Expected to run in `cto/client`, but could not find the expected `go.mod`. Aborting.")
		panic(err)
	}

	fileStr := string(fileBytes)

	if !strings.Contains(fileStr, "module github.com/CrimsonSarah/cto/client") {
		panic("Expected to run in `cto/client`, but `go.mod` does not seem to include the correct `module` directive. Aborting.")
	}
}
