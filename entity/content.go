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
	User         string
	UserLoginURL template.HTMLAttr
}

type AnswerContent struct {
	Username string
	Password string
}

type MessageContent struct {
	Header    string
	Message   string
	Link      template.HTMLAttr
	LinkTitle string
}

type QueryContent struct {
	CountLeft int
}
