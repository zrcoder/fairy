package main

import (
	"fmt"
	"net/http"
	"strconv"

	"gitee.com/rdor/fairy/static"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
)

func main() {
	app := amisgo.New(
		conf.WithTheme(conf.ThemeDark),
		conf.WithIcon("/static/21/2.webp"),
	)
	index := app.Page().Body(
		app.Tabs().TabsMode("sidebar").Swipeable(true).Tabs(
			app.Tab().Title("Yaq").ID("0").Tab(carousel(app, 0, 1, ".gif")),
			app.Tab().Title("Yaq 1").ID("1").Tab(carousel(app, 1, 8)),
			app.Tab().Title("Yaq 2").ID("2").Tab(carousel(app, 2, 13)),
			app.Tab().Title("Yaq 3").ID("3").Tab(carousel(app, 3, 16)),
			app.Tab().Title("Yaq 4").ID("4").Tab(carousel(app, 4, 17)),
			app.Tab().Title("Yaq 5").ID("5").Tab(carousel(app, 5, 30)),
			app.Tab().Title("Yaq 6").ID("6").Tab(carousel(app, 6, 7)),
			app.Tab().Title("Yaq 7").ID("7").Tab(carousel(app, 7, 19)),
			app.Tab().Title("Yaq 8").ID("8").Tab(carousel(app, 8, 11)),
			app.Tab().Title("Yaq 9").ID("9").Tab(carousel(app, 9, 7)),
			app.Tab().Title("Yaq 10").ID("10").Tab(carousel(app, 10, 6)),
			app.Tab().Title("Yaq 11").ID("11").Tab(carousel(app, 11, 9)),
			app.Tab().Title("Yaq 12").ID("12").Tab(carousel(app, 12, 12, ".jpeg")),
			app.Tab().Title("Yaq 13").ID("13").Tab(carousel(app, 13, 9)),
			app.Tab().Title("Yaq 14").ID("14").Tab(carousel(app, 14, 10)),
			app.Tab().Title("Yaq 15").ID("15").Tab(carousel(app, 15, 12)),
			app.Tab().Title("Yaq 16").ID("16").Tab(carousel(app, 16, 14)),
			app.Tab().Title("Yaq 17").ID("17").Tab(carousel(app, 17, 11)),
			app.Tab().Title("Yaq 18").ID("18").Tab(carousel(app, 18, 12)),
			app.Tab().Title("Yaq 19").ID("19").Tab(carousel(app, 19, 8)),
			app.Tab().Title("Yaq 20").ID("20").Tab(carousel(app, 20, 11)),
			app.Tab().Title("Yaq 21").ID("21").Tab(carousel(app, 21, 7)),
			app.Tab().Title("Yaq 22").ID("22").Tab(carousel(app, 22, 5)),
			app.Tab().Title("Yaq 23").ID("23").Tab(carousel(app, 23, 6, ".jpeg")),
			app.Tab().Title("Yaq 24").ID("24").Tab(carousel(app, 24, 17)),
		),
	)

	app.StaticFS("/static", http.FS(static.FS))
	app.Mount("/", index)

	fmt.Println("Serving on http://localhost:8080")
	panic(app.Run(":8080"))
}

func carousel(app *amisgo.App, id, cnt int, f ...string) any {
	opts := getOptions(app, id, cnt, f...)
	return app.Carousel().
		Height("940").
		Auto(false).
		AlwaysShowArrow(true).
		Animation("slide").
		Options(
			opts...,
		)
}

func getOptions(app *amisgo.App, id, cnt int, f ...string) []comp.CarouselOption {
	ext := ".webp"
	if len(f) > 0 {
		ext = f[0]
	}
	res := make([]comp.CarouselOption, cnt)
	prefix := "/static/" + strconv.Itoa(id) + "/"
	for i := range res {
		res[i] = app.CarouselOption().Image(prefix + strconv.Itoa(i+1) + ext)
	}
	return res
}
