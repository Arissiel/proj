package main

import (
	"fmt"
	"net/http"
	"proj/server/handler"
)

func main() {
	handler.NewMap()
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/mutate", handler.HandlerMutation)
	fmt.Println("server started listening on port 13281")
	http.ListenAndServe(":13281", nil)

}

// m := make(map[string]string) // всегда будет создаваться при старте сервера

// m["что-то из ручки"] = "something from handler" // Передать ключ и значение в Query parameters на ручке ПОСТ

// delete(m, "key from DELETE handler") // TODO: delete element from map with DELETE handler

// resp := m["key from handler"] // TODO: return key from map for GET handler
