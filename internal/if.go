package internal

import "io"

type IfBuilder struct {
	result Renderer
}

func (i IfBuilder) Render(w io.Writer) error {
	if i.result == nil {
		return nil
	}
	return i.result.Render(w)
}

func If(cond bool, content Renderer) IfBuilder {
	if !cond {
		return IfBuilder{}
	}
	return IfBuilder{result: content}
}

func (i IfBuilder) ElseIf(cond bool, content Renderer) IfBuilder {
	if !cond || i.result != nil {
		return i
	}
	return IfBuilder{result: content}
}

func (i IfBuilder) Else(content Renderer) IfBuilder {
	if i.result != nil {
		return i
	}
	return IfBuilder{result: content}
}
