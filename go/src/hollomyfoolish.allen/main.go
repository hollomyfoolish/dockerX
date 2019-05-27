package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"log"
	"reflect"

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
}

const SAP_SME_NETWORK = "sap-sme"

func formatDir(srcDir string) string {
	srcDir = strings.Replace(srcDir, "\\", "/", -1)
	if strings.HasSuffix(srcDir, "/") != true{
		srcDir = srcDir + "/"
	}
	return srcDir
}

func updateComponentDockerNetwork(dir string) error {
	dir = formatDir(dir)
	file, err := os.OpenFile(dir + "docker-compose.yml", os.O_RDONLY, 0644)
	if err != nil{
		fmt.Printf("read compose file error %v\n", err)
		return err
	}
	if contents, e := ioutil.ReadAll(file); e == nil {
		var dcFile DockerComposeFile
        e = yaml.Unmarshal(contents, &dcFile)
        if e != nil {
			fmt.Printf("parse docker compose file error in dir: %s\n", dir)
			return e
		}
		for key, element := range dcFile.Services {
			fmt.Printf("check service %s \n", key)
			fmt.Printf("check service %v \n", reflect.TypeOf(element))
			if ss, ok := element.(map[interface{}]interface{}); ok{
				// update net works to sap-sme no matter there is one
				ss["networks"] = []string{SAP_SME_NETWORK}
			} else {
				fmt.Printf("%s is not a valid service\n", key)
			}
		}
		dcFile.Networks = make(map[string]interface{})
		sapsmeNw := make(map[string]bool)
		dcFile.Networks[SAP_SME_NETWORK] = sapsmeNw
		sapsmeNw["external"] = true

		d, err := yaml.Marshal(&dcFile)
        if err != nil {
			fmt.Printf("marshal compose error %v\n", dcFile)
			return err
        }
		err = ioutil.WriteFile(dir + "docker-compose.yml", d, 0644)
		if err == nil{
			fmt.Printf("update compose file %s network successfully\n", dir)
			return nil
		} else {
			fmt.Printf("update compose file %s network error\n", dir)
			return err
		}
	} else {
		fmt.Printf("read docker compose file error in dir: %s\n", dir)
		return e
	}
}

func main(){
	fmt.Println("Hello Go!")
	
	file, err := os.OpenFile("./files/docker-compose.yml", os.O_RDONLY,0644)
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

		if ioutil.WriteFile("./files/docker-compose.bak.yml", d, 0644) == nil {
			fmt.Println("写入文件成功:")
		}

        fmt.Printf("--- t dump:\n%s\n\n", string(d))

		// _, service := range t.Services{
		// 	fmt.Printf("--- t:\n%v\n\n", service)
		// }
	}

	updateComponentDockerNetwork("C:\\Users\\i311688\\Desktop\\MyTemp\\csm\\sld")

}