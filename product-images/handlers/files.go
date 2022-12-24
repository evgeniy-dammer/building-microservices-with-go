package handlers

import (
	"github.com/evgeniy-dammer/building-microservices-with-go/product-images/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"path/filepath"
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

// ServeHTTP implements the http.Handler interface
func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	filename := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", filename)

	if id == "" || filename == "" {
		f.invalidURI(r.URL.String(), rw)
		return
	}

	f.saveFile(id, filename, rw, r)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id string, filename string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for product", "id", id, "path", filename)

	fp := filepath.Join(id, filename)
	err := f.store.Save(fp, r.Body)

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
