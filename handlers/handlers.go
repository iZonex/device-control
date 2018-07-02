package handlers

import (
	"device-control/status"
	"device-control/util"
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
}

// func main() {
// 	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
// 		var user User
// 		json.NewDecoder(r.Body).Decode(&user)

// 		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
// 	})

// 	http.ListenAndServe(":8080", nil)
// }

// PageData information about main page

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/status.html"))
	data := PageData{
		PageTitle: "IOT",
	}
	tmpl.Execute(w, data)
}

func WifiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/wifi.html"))
	data := PageData{
		PageTitle: "IOT",
	}
	tmpl.Execute(w, data)
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/server.html"))
	data := PageData{
		PageTitle: "IOT",
	}
	tmpl.Execute(w, data)
}

func DeviceInformationHandler(w http.ResponseWriter, r *http.Request) {

	statusInfo := status.StatusInfo()
	util.ResponseOk(w, statusInfo)

}
