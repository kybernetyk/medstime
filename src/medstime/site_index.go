package main

import (
    "web"
    "mustache"
)

func indexGet(ctx *web.Context) string {
    m := map[string]string{}
    s := mustache.RenderFile("templ/index.mustache", &m)
    return s
}