package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var newFolder strings.Builder

	goPath := fmt.Sprintf("%s/", os.Getenv("GOPATH"))
	log.Println(goPath)

	fmt.Fprintf(&newFolder, "%s", goPath)

	fmt.Print("Name your repository: ")
	reader := bufio.NewReader(os.Stdin)
	fileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Remove newline character
	fileName = strings.TrimSpace(fileName)

	newFolder.WriteString(fileName)
	str := newFolder.String()
	log.Println(str)
	if err := os.MkdirAll(str, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	pathToNewFolder := fmt.Sprintf("%s/%s", str, "main.go")

	err = os.WriteFile(pathToNewFolder, []byte(""), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", "mod", "init", fileName)

	// set dir of executeable cmd
	cmd.Dir = str

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}	