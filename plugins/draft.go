package main

import (
    "github.com/gohugoio/hugo/identity"
    "github.com/gohugoio/hugo/resources/page"
)

func init() {
    identity.Add("draft", func(p page.Page) bool {
        return p.IsDraft()
    })
}