package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Controller struct {
	s3Client *S3Client
}

func (c *Controller) handlePublic(w http.ResponseWriter, req *http.Request) {

	println(req.Method + " " + req.URL.Path)

	filepath := "public/index.html"

	switch req.URL.Path {
	case "/style.css":
		w.Header().Set("Content-Type", "text/css")
		filepath = "public/style.css"
	case "/PhotoViewer.js":
		w.Header().Set("Content-Type", "text/javascript")
		filepath = "public/PhotoViewer.js"
	}

	data, err := Asset(filepath)

	if err != nil {
		http.Error(w, "Asset "+req.URL.Path+" not found", http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}

func (c *Controller) handleObjectsEndpoint(w http.ResponseWriter, req *http.Request) {

	println(req.Method + " " + req.URL.Path)

	switch req.Method {
	case http.MethodGet:
		c.listObjects(w, req)
	case http.MethodDelete:
		c.bulkDeleteObjects(w, req)
	default:
		http.Error(w, "Method not implemented: "+req.Method+" "+req.URL.Path, http.StatusNotImplemented)
	}
}

func (c *Controller) listObjects(w http.ResponseWriter, req *http.Request) {

	path := strings.TrimPrefix(req.URL.Path, "/api/objects/")
	response := c.s3Client.listObjects(path)

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(response)

	if err != nil {
		println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func (c *Controller) bulkDeleteObjects(w http.ResponseWriter, req *http.Request) {

	var keys []string

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(req.Body).Decode(&keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.s3Client.bulkDeleteKeys(keys)

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(keys)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		print(err.Error())
		return
	}
	w.Write(js)

}
