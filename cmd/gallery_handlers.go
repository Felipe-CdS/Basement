package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	gallery_view "nugu.dev/basement/views/gallery"
)

func (app *application) Gallery(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		putGallery(w, r, app.RWBucketURL)
		return
	}

	namesList, err := GetBucketList(app.RWBucketURL)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var component templ.Component
	switch viewType := r.URL.Query().Get("v"); viewType {
	case "grid":
		component = gallery_view.Gallery(getGridGallery(namesList, app.ReadBucketURL))
	case "list":
		component = gallery_view.Gallery(getListGallery(namesList, app.ReadBucketURL))
	default:
		component = gallery_view.Gallery(templ.NopComponent)
	}
	component.Render(r.Context(), w)
}

func getGridGallery(namesList []string, readBucketURL string) templ.Component {
	var component_objs []gallery_view.BucketBodyView

	for i, x := range namesList {
		holder := gallery_view.BucketBodyView{Name: x}
		component_objs = append(component_objs, holder)

		// limit requests to bucket on testing
		if i == 0 {
			// break
		}
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

	if len(r.MultipartForm.File) == 0 {
		fmt.Print("asdjklsdjk")
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
