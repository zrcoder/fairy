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

const imageCnt = 8

func main() {
	index := comp.Page().Body(
		comp.Tabs().TabsMode("sidebar").Swipeable(true).Tabs(
			comp.Tab().Title("Yaq 1").ID("1").Tab(carousel(1, 8)),
			comp.Tab().Title("Yaq 2").ID("2").Tab(carousel(2, 13)),
			comp.Tab().Title("Yaq 3").ID("3").Tab(carousel(3, 16)),
			comp.Tab().Title("Yaq 4").ID("4").Tab(carousel(4, 17)),
			comp.Tab().Title("Yaq 5").ID("5").Tab(carousel(5, 30)),
			comp.Tab().Title("Yaq 6").ID("6").Tab(carousel(6, 7)),
			comp.Tab().Title("Yaq 7").ID("7").Tab(carousel(7, 19)),
			comp.Tab().Title("Yaq 8").ID("8").Tab(carousel(8, 11)),
			comp.Tab().Title("Yaq 9").ID("9").Tab(carousel(9, 7)),
			comp.Tab().Title("Yaq 10").ID("10").Tab(carousel(10, 6)),
			comp.Tab().Title("Yaq 11").ID("11").Tab(carousel(11, 9)),
			comp.Tab().Title("Yaq 12").ID("12").Tab(carousel(12, 12, ".jpeg")),
			comp.Tab().Title("Yaq 13").ID("13").Tab(carousel(13, 9)),
			comp.Tab().Title("Yaq 14").ID("14").Tab(carousel(14, 10)),
			comp.Tab().Title("Yaq 15").ID("15").Tab(carousel(15, 12)),
			comp.Tab().Title("Yaq 16").ID("16").Tab(carousel(16, 14)),
			comp.Tab().Title("Yaq 17").ID("17").Tab(carousel(17, 11)),
			comp.Tab().Title("Yaq 18").ID("18").Tab(carousel(18, 12)),
			comp.Tab().Title("Yaq 19").ID("19").Tab(carousel(19, 8)),
			comp.Tab().Title("Yaq 20").ID("20").Tab(carousel(20, 11)),
			comp.Tab().Title("Yaq 21").ID("21").Tab(carousel(21, 7)),
		),
	)
	app := amisgo.New(
		conf.WithTheme(conf.ThemeDark),
		conf.WithIcon("/static/21/2.webp"),
	).
		StaticFS("/static/", http.FS(static.FS)).
		Mount("/", index)

	fmt.Println("Serving on http://localhost:8080")
	panic(app.Run(":8080"))
}

func carousel(id, cnt int, f ...string) any {
	opts := getOptions(id, cnt, f...)
	return comp.Carousel().
		Height("940").
		Auto(false).
		AlwaysShowArrow(true).
		Animation("slide").
		Options(
			opts...,
		)
}

func getOptions(id, cnt int, f ...string) []any {
	ext := ".webp"
	if len(f) > 0 {
		ext = f[0]
	}
	res := make([]any, cnt)
	prefix := "/static/" + strconv.Itoa(id) + "/"
	for i := range res {
		res[i] = comp.Schema{"image": prefix + strconv.Itoa(i+1) + ext}
	}
	return res
}
