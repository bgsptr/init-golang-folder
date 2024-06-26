package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	child := []string{
	"cmd", 
	"config", 
	"handler", 
	"middleware", 
	"migrations", 
	"model", 
	"repository", 
	"service",
}

	var newFolder strings.Builder

	// replace with your src directory
	goPath := fmt.Sprintf("%s/src/", os.Getenv("GOPATH"))

	fmt.Fprintf(&newFolder, "%s", goPath)

	fmt.Print("Name your repository: ")
	reader := bufio.NewReader(os.Stdin)
	fileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Remove newline character
	fileName = strings.TrimSpace(fileName)

	// regex 1 or more whitespace
	reWhtSp := regexp.MustCompile(`[\s\p{Zs}]{1,}`)
	fileName = reWhtSp.ReplaceAllString(fileName, "-")

	newFolder.WriteString(fileName)
	str := newFolder.String()
	log.Println(str)
	if err := os.MkdirAll(str, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	i := len(child) - 1

	for len(child) > 0 {
		err := os.Mkdir(fmt.Sprintf("%s/%s", str, child[i]), os.ModePerm)
		if err != nil {
			log.Fatal(err)
			continue
		}
		child = child[:len(child)-1]
		i--
	}

	pathToNewFolder := fmt.Sprintf("%s/cmd/%s", str, "main.go")

	err = os.WriteFile(pathToNewFolder, []byte("package main"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("golang.org/x/%s", fileName))

	// set dir of executeable cmd
	cmd.Dir = str

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}	