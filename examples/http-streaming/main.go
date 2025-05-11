package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	a "github.com/teenjuna/goco/attr"
	e "github.com/teenjuna/goco/elem"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8000, "server port")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/streaming", handleStreaming)

	addr := fmt.Sprintf("localhost:%d", port)
	log.Printf("Running server on http://%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	page := page(r, "Index",
		e.P(nil, e.Text("This is just index page. Click on links above for examples!")),
	)

	if err := page.Render(w); err != nil {
		log.Printf("Index page: %v\n", err)
	}
}

func handleStreaming(w http.ResponseWriter, r *http.Request) {
	iter := func(yield func(string) bool) {
		timer := time.NewTimer(time.Second * 10)
		tick := time.Tick(time.Second)
		for {
			select {
			case <-timer.C:
				yield("Done!")
				return
			case t := <-tick:
				if !yield(t.Format(time.TimeOnly)) {
					return
				}
			}
		}
	}

	flush := e.Effect(func() error {
		w.(http.Flusher).Flush()
		return nil
	})

	page := page(r, "Streaming example",
		e.P(nil, e.Text("Streaming example (10 secs)")),
		flush,
		e.Ul(a.Class("list-disc list-inside"),
			e.Map(iter, func(item string) e.Renderer {
				return e.Slice{
					e.Li(nil, e.Text(item)),
					flush,
				}
			}),
		),
	)

	if err := page.Render(w); errors.Is(err, io.ErrClosedPipe) {
		// Ignore it. User has started stream and then navigated to another page.
	} else if err != nil {
		log.Printf("Streaming page: %v\n", err)
	}
}

func page(r *http.Request, title string, content ...e.Renderer) e.Renderer {
	link := func(text string, path string) e.Renderer {
		active := r.URL.Path == path
		return e.A(
			a.Slice{
				a.Href(path),
				a.Class().With("text-neutral-600", !active),
			},
			e.Text(text),
		)
	}

	slash := e.Span(nil, e.Text("/"))

	return e.Slice{
		e.DoctypeHTML(),
		e.HTML(a.Class("h-full"),
			e.Head(a.Class("h-full"),
				e.Meta(a.Charset("UTF-8")),
				e.Meta(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0")),
				e.Script(a.Src("https://unpkg.com/@tailwindcss/browser@4")),
				// e.Script(a.Slice{
				// 	a.Src("https://unpkg.com/htmx.org@2.0.4"),
				// 	a.Integrity("sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+"),
				// 	a.CrossOrigin("anonymous"),
				// }),
				e.Title(e.Text(title)),
			),
			e.Body(a.Class("h-full bg-neutral-300 text-neutral-900 font-mono"),
				e.Header(a.Class("py-2 px-4 gap-4 flex"),
					link("Index", "/"),
					slash,
					link("Streaming", "/streaming"),
				),
				e.Main(a.Class("py-2 px-4"),
					content...,
				),
			),
		),
	}
}
