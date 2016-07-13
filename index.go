package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type helloWorld struct {
	Hello string `json:"hello"`
	World int    `json:"world"`
}

type helloWorlds []helloWorld

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hellos := helloWorlds{
		helloWorld{
			Hello: "Hello World",
			World: 1,
		},
		helloWorld{
			Hello: "Goodbye Cruel World",
			World: 69,
		},
	}

	if err := json.NewEncoder(w).Encode(hellos); err != nil {
		fmt.Println(err)
	}

	return
}
