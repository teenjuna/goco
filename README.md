# goco

**Go** package for HTML **co**mponents.

The goal is to provide a strict yet powerful API for rendering HTML on the server.
Think of something like [JSX](https://en.wikipedia.org/wiki/JSX_(JavaScript)), but in Go.

## Features
- Ready to use
[spec elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements)
and ability to create custom ones
- Ready to use
[spec attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes)
and ability to create custom ones
- `Text` and `Textf` for `HTML-escaped` text
- `Raw` and `Rawf` for raw text (please don't put user input here)
- `If` and `Switch` for conditional rendering
- `Seq` and `Map` for working with Go iterators
- `Lazy` for avoiding doing work when it's not needed
- `Effect` for rendering side effects

## Quick example

Here is a simple HTTP server that renders todo list with items from URL query parameter.

```
go get github.com/teenjuna/goco
```

```go
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
```

Check out [examples](./examples) directory for more examples.

## How it works

Both elements and attributes are built around a `Renderer` interface. This interface represents
something that can render to `io.Writer`. Each component you create is just a tree of `Renderer`
objects. When you call `.Render(w)` on your component, it will call `.Render(w)` on each object
in the tree.

One important note regarding `Renderer` interface: there is an `internal.Renderer`, but both
`attr` and `elem` packages have their own `Renderer` interface that embeds `internal.Renderer`
along with a private "marker" method. These private methods ensure that one can't implement
`attr.Renderer` outside of `attr` package or `elem.Renderer` outside of `elem` package. In short,
it doesn't let you fuck something up. You're not limited, though: each package provides everything
you need to make something custom.

## Is it fast?

Probably. I didn't measure. But, as I described in the section above, each render is just a series of
writes to `io.Writer`. I also tried to make as little allocations as possible by:
- Preventing string concatenations
- Returning structs instead of pointers to structs or interfaces where possible

So it should be fast.

## Acknowledgements

This package is just a slightly different take on something that already exists, namely the popular
[gomponents](https://github.com/maragudk/gomponents) package. It's a cool package and I'd love to
use it, but I don't quite like some of it's design decisions. And, since it's a simple thing, I
decided to just write my own.
