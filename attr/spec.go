package attr

func Async(value bool) Node {
	return New("async", value)
}

func AutoFocus(value bool) Node {
	return New("autofocus", value)
}

func AutoPlay(value bool) Node {
	return New("autoplay", value)
}

func Checked(value bool) Node {
	return New("checked", value)
}

func Controls(value bool) Node {
	return New("controls", value)
}

func CrossOrigin(v string) Node {
	return New("crossorigin", v)
}

func DateTime(v string) Node {
	return New("datetime", v)
}

func Defer(value bool) Node {
	return New("defer", value)
}

func Disabled(value bool) Node {
	return New("disabled", value)
}

func Draggable(v string) Node {
	return New("draggable", v)
}

func Loop(value bool) Node {
	return New("loop", value)
}

func Multiple(value bool) Node {
	return New("multiple", value)
}

func Muted(value bool) Node {
	return New("muted", value)
}

func PlaysInline(value bool) Node {
	return New("playsinline", value)
}

func ReadOnly(value bool) Node {
	return New("readonly", value)
}

func Required(value bool) Node {
	return New("required", value)
}

func Selected(value bool) Node {
	return New("selected", value)
}

func Accept(v string) Node {
	return New("accept", v)
}

func Action(v string) Node {
	return New("action", v)
}

func Alt(v string) Node {
	return New("alt", v)
}

// Aria attributes automatically have their name prefixed with "aria-".
func Aria(name, v string) Node {
	return New("aria-"+name, v)
}

func As(v string) Node {
	return New("as", v)
}

func AutoComplete(v string) Node {
	return New("autocomplete", v)
}

func Charset(v string) Node {
	return New("charset", v)
}

func Cite(v string) Node {
	return New("cite", v)
}

// TODO: remove
// func Class(v string) Node {
// 	return New("class", v)
// }

func Cols(v string) Node {
	return New("cols", v)
}

func ColSpan(v string) Node {
	return New("colspan", v)
}

func Content(v string) Node {
	return New("content", v)
}

// Data attributes automatically have their name prefixed with "data-".
func Data(name, v string) Node {
	return New("data-"+name, v)
}

func Slot(v string) Node {
	return New("slot", v)
}

func For(v string) Node {
	return New("for", v)
}

func Form(v string) Node {
	return New("form", v)
}

func Height(v string) Node {
	return New("height", v)
}

func Hidden(v string) Node {
	return New("hidden", v)
}

func Href(v string) Node {
	return New("href", v)
}

func ID(v string) Node {
	return New("id", v)
}

func Integrity(v string) Node {
	return New("integrity", v)
}

func Label(v string) Node {
	return New("label", v)
}

func Lang(v string) Node {
	return New("lang", v)
}

func List(v string) Node {
	return New("list", v)
}

func Loading(v string) Node {
	return New("loading", v)
}

func Max(v string) Node {
	return New("max", v)
}

func MaxLength(v string) Node {
	return New("maxlength", v)
}

func Method(v string) Node {
	return New("method", v)
}

func Min(v string) Node {
	return New("min", v)
}

func MinLength(v string) Node {
	return New("minlength", v)
}

func Name(v string) Node {
	return New("name", v)
}

func Pattern(v string) Node {
	return New("pattern", v)
}

func Placeholder(v string) Node {
	return New("placeholder", v)
}

func Popover[T string | bool](v T) Node {
	return New("popover", v)
}

func PopoverTarget(v string) Node {
	return New("popovertarget", v)
}

func PopoverTargetAction(v string) Node {
	return New("popovertargetaction", v)
}

func Poster(v string) Node {
	return New("poster", v)
}

func Preload(v string) Node {
	return New("preload", v)
}

func Rel(v string) Node {
	return New("rel", v)
}

func Role(v string) Node {
	return New("role", v)
}

func Rows(v string) Node {
	return New("rows", v)
}

func RowSpan(v string) Node {
	return New("rowspan", v)
}

func Src(v string) Node {
	return New("src", v)
}

func SrcSet(v string) Node {
	return New("srcset", v)
}

func Step(v string) Node {
	return New("step", v)
}

func Style(v string) Node {
	return New("style", v)
}

func TabIndex(v string) Node {
	return New("tabindex", v)
}

func Target(v string) Node {
	return New("target", v)
}

func Title(v string) Node {
	return New("title", v)
}

func Type(v string) Node {
	return New("type", v)
}

func Value(v string) Node {
	return New("value", v)
}

func Width(v string) Node {
	return New("width", v)
}

func EncType(v string) Node {
	return New("enctype", v)
}

func Dir(v string) Node {
	return New("dir", v)
}
