package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Passphrase1  string
	Passphrase2 string
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		var details ContactDetails
		if len(r.FormValue("passphrase1")) != 0 {
			details = ContactDetails{
				Passphrase1:   r.FormValue("passphrase1"),
				//	Passphrase2:   r.FormValue("passphrase2"),
			}
		} else {
			details = ContactDetails{
				//Passphrase1:   r.FormValue("passphrase1"),
					Passphrase2:   r.FormValue("passphrase2"),
			}
		}



		// do something with details
		_ = details
		fmt.Println(details)
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}