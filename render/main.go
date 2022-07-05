package main

import (
	"fmt"
	"os"

	"github.com/unrolled/render"
)

func main() {
	// ...
	r := render.New(render.Options{
		Directory:  "",                          // Specify what path to load the templates from.
		FileSystem: nil,                         // Specify filesystem from where files are loaded.
		Layout:     "layout/default1",           // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
		Extensions: []string{".html"},           // Specify extensions to load for templates.
		Delims:     render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
		Asset: func(name string) ([]byte, error) { // Load from an Asset function instead of file.
			return []byte("template content"), nil
		},
		AssetNames: func() []string { // Return a list of asset names for the Asset function
			return []string{"templates/home.html", "templates/layout/default.html"}
		},
	})

	err := r.HTML(os.Stdout, 200, "home", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
