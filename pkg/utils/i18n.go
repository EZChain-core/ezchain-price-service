package utils

import (
	"text/template"

	i "github.com/kataras/i18n"
	"github.com/gertd/go-pluralize"
)

// alias type for use based on package
type I18n = i.I18n

var TranI18n *I18n

var pluralizeClient *pluralize.Client = pluralize.NewClient()

func getFuncs(current *i.Locale) template.FuncMap {
	return template.FuncMap{
		"plural": func(word string, count int) string {
			return pluralizeClient.Pluralize(word, count, true)
		},
	}
}

func InitI18n() *I18n {
	var err error
	if TranI18n == nil {
		// default glob 'll load from current directory run program
		TranI18n, err = i.New(i.Glob("./locales/*/*", i.LoaderConfig{
			// Set custom functions per locale!
			Funcs: getFuncs,
		}), "en-US", "vi-VN")
		if err != nil {
			panic(err)
		}
	}
	return TranI18n
}