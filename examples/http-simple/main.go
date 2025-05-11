package main

import (
	"log"
	"net/http"
	"slices"
	"strings"

	a "github.com/teenjuna/goco/attr"
	e "github.com/teenjuna/goco/elem"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var items []string
		if q := r.URL.Query().Get("items"); q != "" {
			items = strings.Split(q, ",")
		}

		if err := Page(TodoList(items)).Render(w); err != nil {
			w.Write([]byte("failed to render: " + err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}

	})
	if err := http.ListenAndServe("localhost:9000", nil); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}

func Page(content ...e.Renderer) e.Renderer {
	return e.Slice{
		e.DoctypeHTML(),
		e.HTML(nil,
			e.Head(nil,
				e.Meta(a.Charset("UTF-8")),
				e.Meta(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0")),
				e.Title(e.Text("http-simple example")),
			),
			e.Body(nil, content...),
		),
	}

}

func TodoList(items []string) e.Renderer {
	return e.
		If(len(items) != 0, e.Lazy(func() e.Renderer {
			return e.Ul(a.ID("todo-list"),
				e.Map(slices.Values(items), TodoItem),
			)
		})).
		Else(e.P(nil, e.Text("No items")))
}

func TodoItem(item string) e.Renderer {
	return e.Li(a.Class("todo-item"), e.Text(item))
}
