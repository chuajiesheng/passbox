package entity

import (
	"html/template"
)

type BaseContent struct {
	HeadHTML    template.HTML
	NavbarHTML  template.HTML
	ContentHTML template.HTML
}

type HeadContent struct {
	Script template.HTMLAttr
}

type NavbarContent struct {
	User string
}

type AnswerContent struct {
	Username string
	Password string
}

type QueryContent struct {
	CountLeft int
}
