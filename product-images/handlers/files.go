package handlers

import (
	"github.com/evgeniy-dammer/building-microservices-with-go/product-images/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// UploadREST implements the http.Handler interface
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// get parameters from path
	id := vars["id"]
	filename := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", filename)

	if id == "" || filename == "" {
		f.invalidURI(r.URL.String(), rw)
		return
	}

	// save the file
	f.saveFile(id, filename, rw, r.Body)
}

// UploadMultipart handles file uploading
func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	// parse multipart form
	err := r.ParseMultipartForm(128 * 1024)

	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	// get the product id from multipart form
	id, idErr := strconv.Atoi(r.FormValue("id"))
	f.log.Info("Process form for id", "id", id)

	if idErr != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected integer id", http.StatusBadRequest)
		return
	}

	// get the uploading file from multipart form
	file, header, err := r.FormFile("file")

	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected file	", http.StatusBadRequest)
		return
	}

	// save file onto disk
	f.saveFile(r.FormValue("id"), header.Filename, rw, file)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id string, filename string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", filename)

	// prepare the full path
	fullPath := filepath.Join(id, filename)

	// save the file
	err := f.store.Save(fullPath, r)

	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}

// invalidURI sends error to response
func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Info("Invalid path", "path", uri)
	http.Error(rw, "Invalid path should be in format: /[id]/[filepath]", http.StatusBadRequest)
}
