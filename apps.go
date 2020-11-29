package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Project []struct {
	CreateTime     time.Time `json:"createTime"`
	LifecycleState string    `json:"lifecycleState"`
	Name           string    `json:"name"`
	Parent         struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"parent"`
	ProjectID     string `json:"projectId"`
	ProjectNumber string `json:"projectNumber"`
}

var (
	email = flag.String("email", "", "scan email address")
)

func main() {
	flag.Parse()
	if *email == "" {
		os.Exit(1)
	}
	out, err := exec.Command("gcloud", "projects", "list", "--format=json").Output()
	if err != nil {
		log.Fatal(err)
	}

	var projects Project
	json.Unmarshal(out, &projects)
	fmt.Println("[+] Total Project ", len(projects))
	for _, p := range projects {
		// fmt.Println(p.ProjectID)
		GetIamPolicy(p.ProjectID, *email)
	}
}

func GetIamPolicy(projects, filteremail string) {
	filter := "--filter=" + filteremail
	out, err := exec.Command("gcloud", "projects", "get-iam-policy", projects, filter, "--format=json").Output()
	if err != nil {
		log.Fatal(err)
	}

	if len(out) > 3 {
		fmt.Println("[+] Found  in ", projects, "with email", filteremail)
	}

}
