package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

var HOME_FOLDER = os.Getenv("HOME") + "/.http-runner"

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func createConfig() error {
	fmt.Println("createConfig...")
	relativePath, _ := filepath.Abs("./")
	return Copy(relativePath + "/config.example.yaml", HOME_FOLDER + "/config.yml")
}

func getConfigYaml() []byte {
	dat, err := ioutil.ReadFile(HOME_FOLDER + "/config.yml")
	if err != nil {
		fmt.Println("%s", err)
		fmt.Println("Create config.yaml...")
		err2 := createConfig()
		if err2 != nil {
			fmt.Printf("Error during create config.yml : %s", err2)
		} else {
			return getConfigYaml()
		}
	}

	return dat
}

type Config struct {
	port string
	host string
	security struct {
		auth_type string
		basic_auth struct {
			login string
			password string
		}
	}
}

func Config() {
	t := nil

	data := getConfigYaml()
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}


func server() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
		//runCmd()
	})

	fmt.Print("Runnig on http://0.0.0.0/")
	http.ListenAndServe(":80", nil)

}


func runCmd() {
	out, err := exec.Command("/bin/sh", "/Users/stephane/.http-runner/scripts/ls.sh").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}

	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

func main() {
	fmt.Printf("hello, world\n")
	//runCmd()
	config()

}