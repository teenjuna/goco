package elem

import "github.com/teenjuna/goco/attr"

func DoctypeHTML() Node {
	return RawText("<!doctype html>")
}

func A(attrs attr.Renderer, children ...Renderer) Node {
	return New("a", false, attrs, children...)
}

func Address(attrs attr.Renderer, children ...Renderer) Node {
	return New("address", false, attrs, children...)
}

func Area(attrs ...attr.Renderer) Node {
	return New("area", true, attr.Slice(attrs))
}

func Article(attrs attr.Renderer, children ...Renderer) Node {
	return New("article", false, attrs, children...)
}

func Aside(attrs attr.Renderer, children ...Renderer) Node {
	return New("aside", false, attrs, children...)
}

func Audio(attrs attr.Renderer, children ...Renderer) Node {
	return New("audio", false, attrs, children...)
}

func Base(attrs ...attr.Renderer) Node {
	return New("base", true, attr.Slice(attrs))
}

func BlockQuote(attrs attr.Renderer, children ...Renderer) Node {
	return New("blockquote", false, attrs, children...)
}

func Body(attrs attr.Renderer, children ...Renderer) Node {
	return New("body", false, attrs, children...)
}

func Br(attrs ...attr.Renderer) Node {
	return New("br", true, attr.Slice(attrs))
}

func Button(attrs attr.Renderer, children ...Renderer) Node {
	return New("button", false, attrs, children...)
}

func Canvas(attrs attr.Renderer, children ...Renderer) Node {
	return New("canvas", false, attrs, children...)
}

func Cite(attrs attr.Renderer, children ...Renderer) Node {
	return New("cite", false, attrs, children...)
}

func Code(attrs attr.Renderer, children ...Renderer) Node {
	return New("code", false, attrs, children...)
}

func Col(attrs ...attr.Renderer) Node {
	return New("col", true, attr.Slice(attrs))
}

func ColGroup(attrs attr.Renderer, children ...Renderer) Node {
	return New("colgroup", false, attrs, children...)
}

func Data(attrs attr.Renderer, children ...Renderer) Node {
	return New("data", false, attrs, children...)
}

func DataList(attrs attr.Renderer, children ...Renderer) Node {
	return New("datalist", false, attrs, children...)
}

func Details(attrs attr.Renderer, children ...Renderer) Node {
	return New("details", false, attrs, children...)
}

func Dialog(attrs attr.Renderer, children ...Renderer) Node {
	return New("dialog", false, attrs, children...)
}

func Div(attrs attr.Renderer, children ...Renderer) Node {
	return New("div", false, attrs, children...)
}

func Dl(attrs attr.Renderer, children ...Renderer) Node {
	return New("dl", false, attrs, children...)
}

func Embed(attrs ...attr.Renderer) Node {
	return New("embed", true, attr.Slice(attrs))
}

func Form(attrs attr.Renderer, children ...Renderer) Node {
	return New("form", false, attrs, children...)
}

func FieldSet(attrs attr.Renderer, children ...Renderer) Node {
	return New("fieldset", false, attrs, children...)
}

func Figure(attrs attr.Renderer, children ...Renderer) Node {
	return New("figure", false, attrs, children...)
}

func Footer(attrs attr.Renderer, children ...Renderer) Node {
	return New("footer", false, attrs, children...)
}

func Head(attrs attr.Renderer, children ...Renderer) Node {
	return New("head", false, attrs, children...)
}

func Header(attrs attr.Renderer, children ...Renderer) Node {
	return New("header", false, attrs, children...)
}

func HGroup(attrs attr.Renderer, children ...Renderer) Node {
	return New("hgroup", false, attrs, children...)
}

func Hr(attrs ...attr.Renderer) Node {
	return New("hr", true, attr.Slice(attrs))
}

func HTML(attrs attr.Renderer, children ...Renderer) Node {
	return New("html", false, attrs, children...)
}

func IFrame(attrs attr.Renderer, children ...Renderer) Node {
	return New("iframe", false, attrs, children...)
}

func Img(attrs ...attr.Renderer) Node {
	return New("img", true, attr.Slice(attrs))
}

func Input(attrs ...attr.Renderer) Node {
	return New("input", true, attr.Slice(attrs))
}

func Label(attrs attr.Renderer, children ...Renderer) Node {
	return New("label", false, attrs, children...)
}

func Legend(attrs attr.Renderer, children ...Renderer) Node {
	return New("legend", false, attrs, children...)
}

func Li(attrs attr.Renderer, children ...Renderer) Node {
	return New("li", false, attrs, children...)
}

func Link(attrs ...attr.Renderer) Node {
	return New("link", true, attr.Slice(attrs))
}

func Main(attrs attr.Renderer, children ...Renderer) Node {
	return New("main", false, attrs, children...)
}

func Menu(attrs attr.Renderer, children ...Renderer) Node {
	return New("menu", false, attrs, children...)
}

func Meta(attrs ...attr.Renderer) Node {
	return New("meta", true, attr.Slice(attrs))
}

func Meter(attrs attr.Renderer, children ...Renderer) Node {
	return New("meter", false, attrs, children...)
}

func Nav(attrs attr.Renderer, children ...Renderer) Node {
	return New("nav", false, attrs, children...)
}

func NoScript(attrs attr.Renderer, children ...Renderer) Node {
	return New("noscript", false, attrs, children...)
}

func Object(attrs attr.Renderer, children ...Renderer) Node {
	return New("object", false, attrs, children...)
}

func Ol(attrs attr.Renderer, children ...Renderer) Node {
	return New("ol", false, attrs, children...)
}

func OptGroup(attrs attr.Renderer, children ...Renderer) Node {
	return New("optgroup", false, attrs, children...)
}

func Option(attrs attr.Renderer, children ...Renderer) Node {
	return New("option", false, attrs, children...)
}

func P(attrs attr.Renderer, children ...Renderer) Node {
	return New("p", false, attrs, children...)
}

func Param(attrs ...attr.Renderer) Node {
	return New("param", true, attr.Slice(attrs))
}

func Picture(attrs attr.Renderer, children ...Renderer) Node {
	return New("picture", false, attrs, children...)
}

func Pre(attrs attr.Renderer, children ...Renderer) Node {
	return New("pre", false, attrs, children...)
}

func Progress(attrs attr.Renderer, children ...Renderer) Node {
	return New("progress", false, attrs, children...)
}

func Script(attrs attr.Renderer, children ...Renderer) Node {
	return New("script", false, attrs, children...)
}

func Section(attrs attr.Renderer, children ...Renderer) Node {
	return New("section", false, attrs, children...)
}

func Select(attrs attr.Renderer, children ...Renderer) Node {
	return New("select", false, attrs, children...)
}

func Source(attrs ...attr.Renderer) Node {
	return New("source", true, attr.Slice(attrs))
}

func Span(attrs attr.Renderer, children ...Renderer) Node {
	return New("span", false, attrs, children...)
}

func Summary(attrs attr.Renderer, children ...Renderer) Node {
	return New("summary", false, attrs, children...)
}

func SVG(attrs attr.Renderer, children ...Renderer) Node {
	return New("svg", false, attrs, children...)
}

func Table(attrs attr.Renderer, children ...Renderer) Node {
	return New("table", false, attrs, children...)
}

func TBody(attrs attr.Renderer, children ...Renderer) Node {
	return New("tbody", false, attrs, children...)
}

func Td(attrs attr.Renderer, children ...Renderer) Node {
	return New("td", false, attrs, children...)
}

func Template(attrs attr.Renderer, children ...Renderer) Node {
	return New("template", false, attrs, children...)
}

func Textarea(attrs attr.Renderer, children ...Renderer) Node {
	return New("textarea", false, attrs, children...)
}

func TFoot(attrs attr.Renderer, children ...Renderer) Node {
	return New("tfoot", false, attrs, children...)
}

func Th(attrs attr.Renderer, children ...Renderer) Node {
	return New("th", false, attrs, children...)
}

func THead(attrs attr.Renderer, children ...Renderer) Node {
	return New("thead", false, attrs, children...)
}

func Tr(attrs attr.Renderer, children ...Renderer) Node {
	return New("tr", false, attrs, children...)
}

func Ul(attrs attr.Renderer, children ...Renderer) Node {
	return New("ul", false, attrs, children...)
}

func Wbr(attrs ...attr.Renderer) Node {
	return New("wbr", true, attr.Slice(attrs))
}

func Abbr(attrs attr.Renderer, children ...Renderer) Node {
	return New("abbr", false, attrs, children...)
}

func B(attrs attr.Renderer, children ...Renderer) Node {
	return New("b", false, attrs, children...)
}

func Caption(attrs attr.Renderer, children ...Renderer) Node {
	return New("caption", false, attrs, children...)
}

func Dd(attrs attr.Renderer, children ...Renderer) Node {
	return New("dd", false, attrs, children...)
}

func Del(attrs attr.Renderer, children ...Renderer) Node {
	return New("del", false, attrs, children...)
}

func Dfn(attrs attr.Renderer, children ...Renderer) Node {
	return New("dfn", false, attrs, children...)
}

func Dt(attrs attr.Renderer, children ...Renderer) Node {
	return New("dt", false, attrs, children...)
}

func Em(attrs attr.Renderer, children ...Renderer) Node {
	return New("em", false, attrs, children...)
}

func FigCaption(attrs attr.Renderer, children ...Renderer) Node {
	return New("figcaption", false, attrs, children...)
}

func H1(attrs attr.Renderer, children ...Renderer) Node {
	return New("h1", false, attrs, children...)
}

func H2(attrs attr.Renderer, children ...Renderer) Node {
	return New("h2", false, attrs, children...)
}

func H3(attrs attr.Renderer, children ...Renderer) Node {
	return New("h3", false, attrs, children...)
}

func H4(attrs attr.Renderer, children ...Renderer) Node {
	return New("h4", false, attrs, children...)
}

func H5(attrs attr.Renderer, children ...Renderer) Node {
	return New("h5", false, attrs, children...)
}

func H6(attrs attr.Renderer, children ...Renderer) Node {
	return New("h6", false, attrs, children...)
}

func I(attrs attr.Renderer, children ...Renderer) Node {
	return New("i", false, attrs, children...)
}

func Ins(attrs attr.Renderer, children ...Renderer) Node {
	return New("ins", false, attrs, children...)
}

func Kbd(attrs attr.Renderer, children ...Renderer) Node {
	return New("kbd", false, attrs, children...)
}

func Mark(attrs attr.Renderer, children ...Renderer) Node {
	return New("mark", false, attrs, children...)
}

func Q(attrs attr.Renderer, children ...Renderer) Node {
	return New("q", false, attrs, children...)
}

func S(attrs attr.Renderer, children ...Renderer) Node {
	return New("s", false, attrs, children...)
}

func Samp(attrs attr.Renderer, children ...Renderer) Node {
	return New("samp", false, attrs, children...)
}

func Small(attrs attr.Renderer, children ...Renderer) Node {
	return New("small", false, attrs, children...)
}

func Strong(attrs attr.Renderer, children ...Renderer) Node {
	return New("strong", false, attrs, children...)
}

func Sub(attrs attr.Renderer, children ...Renderer) Node {
	return New("sub", false, attrs, children...)
}

func Sup(attrs attr.Renderer, children ...Renderer) Node {
	return New("sup", false, attrs, children...)
}

func Time(attrs attr.Renderer, children ...Renderer) Node {
	return New("time", false, attrs, children...)
}

func Title(children ...Renderer) Node {
	return New("title", false, nil, children...)
}

func U(attrs attr.Renderer, children ...Renderer) Node {
	return New("u", false, attrs, children...)
}

func Var(attrs attr.Renderer, children ...Renderer) Node {
	return New("var", false, attrs, children...)
}

func Video(attrs attr.Renderer, children ...Renderer) Node {
	return New("video", false, attrs, children...)
}
