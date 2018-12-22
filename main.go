package main

import (
	"github.com/gopherjs/vecty"
)

func main() {
	page := NewPage()

	vecty.SetTitle("Whisper Simulation")
	vecty.AddStylesheet("https://cdn.jsdelivr.net/gh/divan/whispervis@c395bc39dd76ba9d9118693da3d7bc90f234e5f9/css/bulma.css")
	vecty.AddStylesheet("https://cdn.jsdelivr.net/gh/divan/whispervis@c395bc39dd76ba9d9118693da3d7bc90f234e5f9/css/bulma-extensions.min.css")
	vecty.AddStylesheet("https://cdn.jsdelivr.net/gh/divan/whispervis@c395bc39dd76ba9d9118693da3d7bc90f234e5f9/css/custom.css")
	vecty.RenderBody(page)
}
