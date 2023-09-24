package vegur

import (
	"sync"

	"gio.tools/fonts/vegur/vegurbold"
	"gio.tools/fonts/vegur/vegurlight"
	"gio.tools/fonts/vegur/vegurregular"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(vegurbold.OTF)
		register(vegurlight.OTF)
		register(vegurregular.OTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(src []byte) {
	faces, err := opentype.ParseCollection(src)
	if err != nil {
		panic("failed to parse font: " + err.Error())
	}
	collection = append(collection, faces[0])
}
