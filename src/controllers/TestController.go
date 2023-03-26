package controllers

import (
    "fmt"
    "net/http"
)

type TestController struct{}

func (c *TestController) GetHomePageText(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Fern sux")
}