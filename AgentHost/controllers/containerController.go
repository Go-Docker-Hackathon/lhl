package controllers

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/fsouza/go-dockerclient"
	"github.com/Go-Docker-Hackathon/team-LHL/AgentHost/resource"
)

func getPostData(r *http.Request) map[string]string {
	postData := make(map[string]string)
	if "POST" == r.Method {
		r.ParseForm()
		for k, v := range r.PostForm {
			postData[string(k)] = strings.Join(v,",")
		}
	}
	return postData
}

func CreateContainer(w http.ResponseWriter, r *http.Request) string {
	
	data := getPostData(r)
	resources := strings.Split(data["resources"], ",")
	client, _ := docker.NewClient("unix:///var/run/docker.sock")
	image := resource.GetImage(resources)
	config := &docker.Config{
            Image: image,
    }
    containerOptions := docker.CreateContainerOptions{
        Config: config,
        HostConfig: nil,
    }
	fmt.Println("--------creating---image is--%v", image, "--------")
    exec, err := client.CreateContainer(containerOptions)
    s, _ := json.Marshal(exec)
	if err != nil {
		fmt.Println(err)
	}
	return string(s)
}