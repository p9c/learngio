package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/mmcdole/gofeed"
)

func main() {
	gofont.Register()
	th := material.NewTheme()

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)

	go func() {
		w := app.NewWindow()
		// START INIT OMIT
		list := &layout.List{
			Axis: layout.Vertical,
		}
		gtx := layout.NewContext(w.Queue())
		// END INIT OMIT
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				list.Layout(gtx, len(feed.Items), func(i int) {
					f := feed.Items[i]
					layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func() {
							th.H5(f.Title).Layout(gtx)
						}),
						layout.Rigid(func() {
							th.Body1(f.Content).Layout(gtx)
						}),
						layout.Rigid(func() {
							//th.Body2(f.Image).Layout(gtx)
						}),
					)
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
