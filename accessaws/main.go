package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	//cmd := exec.Command("aws", "s3", "ls", "--profile", "profile1")
	cmd1 := exec.Command("aws", "s3", "mb", "s3://aka-labs8", "--profile", "profile1")
	//aws s3 ls --profile example_1

	//err := cmd.Run()aws

	// if err != nil {
	// 	log.Fatal(err)
	// }
	out, err := cmd1.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
