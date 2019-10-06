package handlers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "time"
  "math/rand"
  "github.com/thedevsaddam/renderer"
  linuxproc "github.com/c9s/goprocinfo/linux"
  helpers "../helpers"
  repos "../repos"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./static/*.html",
	}

	rnd = renderer.New(opts)
}


func CpuProcessHandler(response http.ResponseWriter, request *http.Request){
  rnd.HTML(response, http.StatusOK, "cpu_graph", nil)
}

func RamProcessHandler(response http.ResponseWriter, request *http.Request){
  rnd.HTML(response, http.StatusOK, "ram_graph", nil)
}


func AdminPageHandler(response http.ResponseWriter, request *http.Request){
  rnd.HTML(response, http.StatusOK, "admin", nil)
  //var body, _= helpers.LoadFile("static/chartjs_example3.html")
  //fmt.Fprintf(response, body)
}

//LoginPage GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request){
  var body, _= helpers.LoadFile("static/index.html")
  fmt.Fprintf(response, body)
}

type Values struct{
  X string`json:"x"`
  Y int `json:"y"`
}

//======= RAM DATA

type RamStruct struct{
  X int `json:"x"`
  Y int `json:"y"`
  Z int `json:"z"`
}

func RamData(response http.ResponseWriter, request *http.Request){
  memory, _ := linuxproc.ReadMemInfo("/proc/meminfo")
  total:=memory.MemTotal/1000
  consumida:= (memory.MemTotal - memory.MemFree)/1000
  per_consumo := (float64(memory.MemTotal - memory.MemFree)/float64(memory.MemTotal))*100
  ram_data := RamStruct{X:int(total),Y:int(consumida),Z:int(per_consumo)}
  byteArray, _ := json.Marshal(ram_data)
  fmt.Fprintf(response,string(byteArray))
}

//======= CPU DATA

type CpuStruct struct{
  X int `json:"x"`
  Y int `json:"y"`
}

func CpuData(response http.ResponseWriter, request *http.Request){
  cpu_data := RamStruct{X:rand.Intn(100),Y:rand.Intn(100)}
  byteArray, _ := json.Marshal(cpu_data)
  fmt.Fprintf(response,string(byteArray))
}



func AdminHandler(response http.ResponseWriter, request *http.Request){

  values := Values{X:time.Now().Format("2006-01-02T15:04:05Z"),Y:rand.Intn(100)}

  byteArray, _ := json.Marshal(values)

  fmt.Fprintf(response,string(byteArray))
}

//LoginPage POST
func LoginHandler(response http.ResponseWriter, request *http.Request){
  username:= request.FormValue("username")
  pass:= request.FormValue("pass")
  redirectTarget:= "/adminPage"

  if !helpers.IsEmpty(username) && !helpers.IsEmpty(pass){
    _userIsValid := repos.UserIsValid(username, pass)
    if _userIsValid {
        redirectTarget = "/adminPage"
    } else {
        redirectTarget = "/"
    }
  }
  http.Redirect(response, request, redirectTarget, 302)
}
