package controllers

import (
    "fmt"
    "net/http"
)

type IdeaController struct{}

func (c *IdeaController) GetIdeas(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Here is a list of ideas")
}