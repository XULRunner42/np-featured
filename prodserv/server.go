package main

import (
    "github.com/hoisie/web.go"
    "github.com/hoisie/mustache.go"
    "trans"
)

func getProducts() (ret []map[string]string) {
    /*
    ret = []map[string]string{ {
	"url":	"/product/1",
	"title":"Steaming Pile #1",
	"img":	"http://example.us/pile-shit.jpg",
    }, {
	"url": "/product/2",
	"title":"Steaming Pile #2",
	"img":	"http://example.us/pile-shit.jpg",
    }, {
	"url": "/product/4",
	"title":"Lookin' Good Template",
	"img":	"http://nerdland.info/template.png",
    }, {
	"url": "/product/3",
	"title":"Steaming Pile #3",
	"img":	"http://example.us/pile-shit.jpg",
	},
    }
    */

    return trans.NpToAmazon()
}

func getPages() (ret []map[string]string) {

    // there are so many pages, 

    ret = []map[string]string{
	{ "page": "1" },
	{ "page": "2" },
	{ "page": "3" },
	{ "page": "4" },
	{ "page": "5" },
	{ "page": "6" },
    }

    return ret
}

func renderPage(ctx *web.Context, val string) {

    products := getProducts()
    pages := getPages()

    in := map[string]interface{} {
	"products": products,
	"pages": pages,
    }

    out := mustache.Render("{{>productpage}}", in)

    ctx.WriteString(out)
}

func index(ctx *web.Context, val string) {
    renderPage(ctx, "1")
}

func renderProduct(ctx *web.Context, val string) {
    out := "not implemented product view<br/><a href='/'>back to front</a>"
    ctx.WriteString(out)
}

func main() {
    web.Get("/product/(.*)", renderProduct)
    web.Get("/page/(.*)", renderPage)
    web.Get("/(.*)", index)
    web.Run("0.0.0.0:8080")
}
