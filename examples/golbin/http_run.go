package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	golbin "github.com/abhishekkr/gol/golbin"
)

func main() {
	dashr_ip := flag.String("fqdn", "127.0.0.1", "IP/FQDN to run HTTP listener at")
	dashr_port := flag.String("http", "8001", "port to run HTTP listener at")
	www_data := flag.String("www", "www-data", "path to media dashr static site content")
	media_content := flag.String("media", "dummy-media-files", "path to media setup root of Playbooks, Roles Dir")
	dashr_config := flag.String("config", "config", "path to fetch/save Config used by Static Site Content")
	flag.Parse()

	connection_string := fmt.Sprintf("%s:%s", *dashr_ip, *dashr_port)
	www_data_uri := fmt.Sprintf("/%s/", *www_data)
	media_content_uri := fmt.Sprintf("/%s/", *media_content)
	dashr_config_uri := fmt.Sprintf("/%s/", *dashr_config)

	dashr_fs := http.FileServer(http.Dir(*www_data))
	http.Handle(www_data_uri, http.StripPrefix(www_data_uri, dashr_fs))

	media_fs := http.FileServer(http.Dir(*media_content))
	http.Handle(media_content_uri, http.StripPrefix(media_content_uri, media_fs))

	config_fs := http.FileServer(http.Dir(*dashr_config))
	http.Handle(dashr_config_uri, http.StripPrefix(dashr_config_uri, config_fs))

	http.HandleFunc("/", handleSystem)

	log.Println("media Dashr @", connection_string)
	if err := http.ListenAndServe(connection_string, nil); err != nil {
		fmt.Println("ERROR: Failed to start server.", err.Error())
	}
}

/****/
func exec(command string) string {
	kon := golbin.Console{Command: command}
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	kon.Run()
	fmt.Println(kon)

	return kon.StdOutput
}

func handleSystem(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	_command := req.FormValue("command")
	_output := exec(_command)

	data, _ := json.Marshal(fmt.Sprintf("{\"command\":	\"%s\"}", _output))
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}
