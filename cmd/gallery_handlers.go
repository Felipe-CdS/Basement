package main

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	gallery_view "nugu.dev/basement/views/gallery"
)

func (app *application) Gallery(w http.ResponseWriter, r *http.Request) {

	if err := app.AuthMiddleware(w, r); err != nil {
		return
	}

	htmxReq, err := strconv.ParseBool(r.Header.Get("HX-Request"))

	if err != nil {
		htmxReq = false
	}

	if !htmxReq {
		component := gallery_view.Gallery()
		component.Render(r.Context(), w)
		return
	}

	if r.Method == "PUT" {
		putGallery(w, r, app.RWBucketURL)
		return
	}

	namesList, err := GetBucketList(app.RWBucketURL)

	var component templ.Component

	switch viewType := r.URL.Query().Get("v"); viewType {
	case "grid":
		component = getGridGallery(namesList, app.ReadBucketURL)
	case "list":
		component = getListGallery(namesList, app.ReadBucketURL)
	default:
		component = gallery_view.Gallery()
	}

	component.Render(r.Context(), w)
}

func getGridGallery(namesList []string, readBucketURL string) templ.Component {
	var component_objs []gallery_view.BucketBodyView

	for _, x := range namesList {
		holder := gallery_view.BucketBodyView{Name: x}
		component_objs = append(component_objs, holder)
	}

	return gallery_view.GridView(component_objs, readBucketURL)
}

func getListGallery(namesList []string, readBucketURL string) templ.Component {
	return gallery_view.ListView(namesList, readBucketURL)
}

func putGallery(w http.ResponseWriter, r *http.Request, bucketURL string) {
	if parseErr := r.ParseMultipartForm(32 << 20); parseErr != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	for _, fheaders := range r.MultipartForm.File {
		for _, header := range fheaders {
			f, openErr := header.Open()

			if openErr != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			defer f.Close()

			PutInBucket(bucketURL, f, header.Filename, int(header.Size))
		}
	}

	component := gallery_view.SuccessDialog()
	component.Render(r.Context(), w)
}
