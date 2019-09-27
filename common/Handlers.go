package handlers

import (
  "fmt"
  "net/http"
  helpers "../helpers"
  repos "../repos"
)

//LoginPage GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request){
  var body, _= helpers.LoadFile("static/index.html")
  fmt.Fprintf(response, body)
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
