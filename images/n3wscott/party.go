package main

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"text/template"
)

const (
	partyTemplate = `<?xml version="1.0" encoding="utf-8"?>
<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
    viewBox="0 0 160 220" style="enable-background:new 0 0 160 220;" xml:space="preserve">
<style type="text/css"> 
	.st0{fill:{{ .color }};} <!-- face -->
	.st1{fill:#D68C80;} <!-- nose -->
	.st2{fill:#8D6553;} <!-- hair -->
</style>
<g transform="rotate({{ .rotate }} 80 110)">
  <g transform="translate(-775 -83)">
    <path class="st0" d="M920.9,177.9c0,47.1-30.7,112.5-68.5,112.5c-37.8,0-68.5-65.4-68.5-111.5s30.7-85.3,68.5-85.3
    	C890.2,92.6,920.9,130.8,920.9,177.9z"/>
    <path class="st1" d="M862.4,182.9c-1.9-0.6-15.7,0-15.7,0s-8.8,38.3-5.7,42.1c3.1,3.8,10,6.9,12.6,6.9s13.2-3.1,13.8-6.9
    	C868,221.2,862.4,182.9,862.4,182.9z"/>
    <path class="st2" d="M882.7,111.6c-1.6,0.4-56.2,7.3-69.4,19.5c-13.2,12.2-10.9-3.2-10.9-3.2s-9.1,0.5-12.9,34.2
    	c-0.4,3.6,0.2,37-1.6,34.7c-1.8-2.3-3.2-5.4-3.6-14.5c-0.6-11.4-8.2-63,24.9-79.8c31.3-15.9,77.9-6.8,82.6-5.4
    	c15.4,4.5,19.3,8.4,25.2,27.9s7.9,56.5,6.6,61.5c-1.4,5-1.6,13.6-3.2,10.4c-0.9-1.8-0.9-9.5-0.5-15c0.1-1.4-3.4-34.2-7.9-42.9
    	C907.4,130.4,892.7,109.4,882.7,111.6z"/>
  </g>
</g>
</svg>
`
)

var colors = []string{
	"#FECE7A",
	"#FECE7A",
	"#7EFF7A",
	"#7EFF7A",
	"#7EFFFF",
	"#7EFFFF",
	"#7BA3FE",
	"#7BA3FE",
	"#CB70FE",
	"#CB70FE",
	"#FC49F5",
	"#FC49F5",
	"#FC4EA7",
	"#FC4EA7",
	"#FD5258",
	"#FD5258",
}

var rotations = []string{
	"0",
	"5",
	"10",
	"12.5",
	"15",
	"12.5",
	"10",
	"5",
	"0",
	"-5",
	"-10",
	"-12.5",
	"-15",
	"-12.5",
	"-10",
	"-5",
}

func main() {
	party, err := template.New("party").Parse(partyTemplate)
	if err != nil {
		fmt.Printf("can't party :( %s\n\n", err)
		return
	}

	for i, r := range rotations {
		f, err := os.Create(fmt.Sprintf("n3wscott-%02d.svg", i))
		if err != nil {
			fmt.Printf("can't party :( %s\n\n", err)
			continue
		}
		fmt.Printf("writing file %s\n", f.Name())

		if err := party.Execute(f, map[string]string{
			"rotate": r,
			"color": colors[i%len(colors)],
		}); err != nil {
			fmt.Printf("really can't party :( :( %s\n\n", err)
			return
		}
		_ = f.Close()
	}
}

func doGif() {
	files := []string{"g1.gif", "g2.gif","g3.gif", "g2.gif"}

	// load static image and construct outGif
	outGif := &gif.GIF{}
	for _, name := range files {
		f, _ := os.Open(name)
		inGif, _ := gif.Decode(f)
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	// save to out.gif
	f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}

/*

Parrot colors:
#FECE7A
#7EFF7A
#7EFFFF
#7BA3FE
#CB70FE
#FC49F5
#FC4EA7
#FD5258

*/
