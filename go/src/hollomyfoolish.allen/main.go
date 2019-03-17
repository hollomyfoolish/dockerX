package main

import (
	"fmt"
	"io/ioutil"
	"os"
	// "strings"
	"log"

	"gopkg.in/yaml.v2"
)

// type Service struct	{
// 	Build string
// 	networks
// }

type DockerComposeFile struct {
	Version string
	Services map[string]interface{}
	Volumes map[string]interface{}
	Networks map[string]interface{}
	Configs map[string]interface{}
	Secrets map[string]interface{}
}

func main(){
	fmt.Println("Hello Go!")
	
	file, err := os.OpenFile("F:\\git_repos\\dockerX\\go\\src\\hollomyfoolish.allen\\files\\docker-compose.yml", os.O_RDONLY,0644)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	if contents, e := ioutil.ReadAll(file); e == nil {
		// result := strings.Replace(string(contents),"\n","",1)
		// fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file contents:",result)
		t := DockerComposeFile{}
    
        e = yaml.Unmarshal(contents, &t)
        if e != nil {
                log.Fatalf("error: %v", err)
		}

		d, err := yaml.Marshal(&t)
        if err != nil {
                log.Fatalf("error: %v", err)
        }

		if ioutil.WriteFile("F:\\git_repos\\dockerX\\go\\src\\hollomyfoolish.allen\\files\\docker-compose.bak.yml", d, 0644) == nil {
			fmt.Println("写入文件成功:")
		}

        fmt.Printf("--- t dump:\n%s\n\n", string(d))

		// _, service := range t.Services{
		// 	fmt.Printf("--- t:\n%v\n\n", service)
		// }
	}

}