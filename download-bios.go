package main

import (
	"flag"
	"fmt"
	"github.com/NubeIO/nube-rubix-updater-cli/helpers/github"
	"github.com/NubeIO/nube-rubix-updater-cli/helpers/unzip"
)

func main() {
	o := flag.String("owner", "NubeIO", "repo owner")
	r := flag.String("repo", "rubix-bios", "repo")
	t := flag.String("tag", "latest", "version")
	tok := flag.String("token", "", "token")

	flag.Parse()
	owner := *o
	repo := *r
	tag := *t
	token := *tok
	deviceType := "armv7"
	dir := "/"
	if token == "" {
		fmt.Println("error: add token")
		fmt.Println("example:  ./download-bios --token=<YOUR TOKEN>")
		return
	}
	fmt.Println("owner", owner, "repo", repo, "tag", tag, "token", token)
	a := github.New()
	var version string
	if tag == "latest" {
		_repo := fmt.Sprintf("%s/%s", owner, repo)
		fmt.Println("repo", _repo)
		tags, err := a.GetLatestReleaseTag(_repo)
		if err != nil {
			fmt.Println("error:", _repo)
		}
		version = tags
	} else {
		version = tag
	}
	fmt.Println("version:", version)
	name, err := github.RetrieveAssets(owner, repo, version, token, deviceType)
	if err != nil {
		fmt.Println("error: main github.RetrieveAssets:", err)
		return
	}
	fmt.Println(dir)
	fmt.Println(name)

	uz := unzip.New()

	_, err = uz.Extract(name, "./builds")
	if err != nil {
		fmt.Println(err)
	}
}
