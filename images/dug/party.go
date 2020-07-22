package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	partyTemplate = `<?xml version="1.0" encoding="utf-8"?>
<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 viewBox="0 0 500 600" style="enable-background:new 0 0 500 600;" xml:space="preserve">

<style type="text/css">
	.hair{
   opacity:1;
   fill-opacity:0;
   stroke:#000000;
   stroke-width:10;
   stroke-opacity:1;
  }
  .h0 { stroke:{{ .color0 }}; }
  .h1 { stroke:{{ .color1 }}; }
  .h2 { stroke:{{ .color2 }}; }
  .h3 { stroke:{{ .color3 }}; }
  .h4 { stroke:{{ .color4 }}; }
  .h5 { stroke:{{ .color5 }}; }
  .h6 { stroke:{{ .color6 }}; }
  .h7 { stroke:{{ .color7 }}; }
  .h8 { stroke:{{ .color8 }}; }
  .h9 { stroke:{{ .color9 }}; }
	.head{
   opacity:1;
   fill-opacity:0;
   stroke:#000000;
   stroke-width:10;
   stroke-opacity:1;
  }
	.face{
   opacity:1;
   fill-opacity:0;
   stroke:#000000;
   stroke-width:10;
   stroke-opacity:1;
  }
	.bg{
   opacity:1;
   fill-opacity:1;
   fill:#FFFFFF};
  }
</style>

<g transform="rotate({{ .rotate }} 250 300)">
   <path d="M140.0 74.0L145.0 273.0L114.0 273.0L114.0 284.0L95.0 284.0L95.0 293.0L85.0 293.0L85.0 314.0L71.0 314.0L71.0 344.0L83.0 344.0L83.0 371.0L105.0 371.0L105.0 384.0L136.0 384.0L136.0 394.0L145.0 394.0L156.0 394.0L145.0 394.0L145.0 424.0L133.0 424.0L133.0 442.0L125.0 442.0L126.0 472.0L114.0 472.0L114.0 512.0L104.0 512.0L104.0 548.0L114.0 548.0L114.0 562.0L176.0 562.0L176.0 552.0L282.0 552.0L282.0 540.0L358.0 540.0L358.0 530.0L372.0 530.0L372.0 523.0L385.0 523.0L385.0 509.0L393.0 509.0L393.0 194.0L393.0 188.0L384.0 100.0L324.0 80.0L284.0 75.0L200.0 70.0L140.0 70.0" class="bg" />
   <path d="M140.0 74.0L145.0 273.0L114.0 273.0L114.0 284.0L95.0 284.0L95.0 293.0L85.0 293.0L85.0 314.0L71.0 314.0L71.0 344.0L83.0 344.0L83.0 371.0L105.0 371.0L105.0 384.0L136.0 384.0L136.0 394.0L145.0 394.0L156.0 394.0L145.0 394.0L145.0 424.0L133.0 424.0L133.0 442.0L125.0 442.0L126.0 472.0L114.0 472.0L114.0 512.0L104.0 512.0L104.0 548.0L114.0 548.0L114.0 562.0L176.0 562.0L176.0 552.0L282.0 552.0L282.0 540.0L358.0 540.0L358.0 530.0L372.0 530.0L372.0 523.0L385.0 523.0L385.0 509.0L393.0 509.0L393.0 194.0L393.0 188.0" class="head" id="b3ga0ukOiC" />
   <path d="M184 42.0L172.0 42.0L172.0 55.0L164.0 55.0L164.0 64.0L156.0 64.0L156.0 89.0L149.0 89.0L149.0 74.0L140.0 74.0" class="hair h0" id="b3ga0ukOiC" />
   <path d="M234.0 32.0C234.0 34.0 234.0 38.0 234.0 44.0L224.0 44.0L224.0 56.0L224.0 44.0L217.0 44.0L203.0 44.0L203.0 55.0L194.0 55.0L194.0 65.0L184.0 65.0L184.0 74.0L175.0 74.0L175.0 87.0" class="hair h1" id="lyp0VpGtl" />
   <path d="M263.0 32.0L263.0 45.0L254.0 45.0L254.0 54.0L244.0 54.0L244.0 66.0L233.0 66.0L233.0 76.0L225.0 76.0L225.0 89.0L225.0 76.0L215.0 76.0L215.0 62.0L215.0 76.0L205.0 76.0L205.0 89.0" class="hair h2" id="b3yzX8E1pX" />
   <path d="M293.0 31.0C293.0 33.0 293.0 37.0 293.0 44.0L283.0 44.0L283.0 55.0L272.0 55.0L272.0 65.0L263.0 65.0L263.0 76.0L254.0 76.0L255.0 95.0L265.0 95.0L255.0 95.0L255.0 105.0" class="hair h3" id="a2fM1pj2z4" />
   <path d="M323.0 32.0L323.0 44.0L314.0 44.0L314.0 54.0L304.0 54.0L304.0 65.0L295.0 65.0L295.0 76.0L284.0 76.0L285.0 105.0L275.0 105.0L275.0 116.0L265.0 116.0L265.0 125.0L254.0 125.0L254.0 136.0" class="hair h4" id="gKX4UXdoY" />
   <path d="M364.0 32.0L364.0 56.0L364.0 44.0L375.0 44.0L365.0 44.0L344.0 44.0L344.0 54.0L333.0 54.0L333.0 64.0L324.0 64.0L324.0 74.0L313.0 74.0L313.0 95.0L304.0 95.0L304.0 107.0L304.0 95.0L284.0 95.0" class="hair h4" id="b15yb52IZg" />
   <path d="M304.0 157.0L304.0 144.0L318.0 144.0L304.0 144.0L304.0 124.0" class="hair h5" id="erCUCNiMh" />
   <path d="M404.0 32.0L404.0 43.0L394.0 43.0L394.0 54.0L384.0 54.0L384.0 65.0L374.0 65.0L374.0 75.0L364.0 75.0L354.0 75.0L354.0 63.0L354.0 75.0L344.0 75.0L344.0 93.0L354.0 93.0L354.0 105.0L354.0 93.0L363.0 93.0L363.0 74.0L363.0 93.0C352.0 93.0 345.0 93.0 343.0 93.0C338.0 93.0 334.0 93.0 333.0 93.0L333.0 104.0L324.0 104.0L324.0 115.0L314.0 115.0L314.0 124.0C304.0 124.0 297.0 124.0 294.0 124.0C293.0 124.0 293.0 120.0 294.0 112.0L294.0 124.0L284.0 124.0L284.0 136.0" class="hair h5" id="c24ibJw9hm" />
   <path d="M404.0 61.0L404.0 74.0L394.0 74.0L394.0 93.0L407.0 93.0L383.0 93.0L383.0 103.0L374.0 103.0L374.0 115.0L363.0 115.0L363.0 123.0L334.0 124.0L333.0 155.0L323.0 155.0L323.0 164.0L314.0 164.0L314.0 175.0L304.0 175.0L304.0 185.0" class="hair h6" id="c208FZHKIB" />
   <path d="M344.0 124.0L355.0 124.0L355.0 144.0L367.0 144.0L355.0 144.0L355.0 155.0L355.0 144.0L334.0 144.0L333.0 124.0L344.0 124.0L344.0 113.0" class="hair h6" id="bF61TP1sU" />
   <path d="M407.0 113.0L393.0 113.0L393.0 124.0L384.0 124.0L384.0 145.0L398.0 145.0L384.0 145.0L384.0 154.0L374.0 154.0L374.0 164.0L363.0 164.0L363.0 174.0L353.0 174.0L354.0 184.0L353.0 174.0L344.0 174.0L344.0 163.0L344.0 174.0L334.0 174.0L334.0 205.0L334.0 194.0L344.0 194.0" class="hair h7" id="d27oyKNn3U" />
   <path d="M374.0 193.0C372.0 193.0 369.0 193.0 363.0 193.0L363.0 202.0L353.0 202.0L353.0 213.0L343.0 213.0L343.0 223.0L334.0 223.0L334.0 235.0" class="hair h8" id="a1GXgwtNJy" />
   <path d="M393.0 162.0L393.0 174.0L383.0 174.0L383.0 185.0" class="hair h8" id="dtrpDwaQm" />   
   <path d="M374.0 213.0L374.0 224.0L363.0 223.0L363.0 234.0" class="hair h9" id="d1Aha6GJ0u" />
   
   <path d="M344.0 314.0L353.0 314.0L353.0 303.0L363.0 303.0L363.0 283.0L374.0 283.0L374.0 264.0L364.0 264.0L364.0 253.0L344.0 253.0L344.0 264.0" class="face" id="c8rOyqmRz" />
   <path d="M170.0 275.0L169.0 254.0" class="face" id="dmNKltYyA" />
   <path d="M205.0 203.0L194.0 203.0L194.0 213.0L184.0 213.0L184.0 223.0L174.0 223.0L174.0 233.0L162.0 233.0" class="face" id="d1lC99SerX" />
</g>

</svg>
`
)

var colors = []string{
	"#FECE7A",
	"#7EFF7A",
	"#7EFFFF",
	"#7BA3FE",
	"#CB70FE",
	"#FC49F5",
	"#FC4EA7",
	"#FD5258",
}

var rotations = []string{
	"0",
	"2",
	"4",
	"6",
	"8",
	"6",
	"4",
	"2",
	"0",
	"-2",
	"-4",
	"-6",
	"-8",
	"-6",
	"-4",
	"-2",
}

func main() {
	party, err := template.New("party").Parse(partyTemplate)
	if err != nil {
		fmt.Printf("can't party :( %s\n\n", err)
		return
	}

	for i, r := range rotations {
		f, err := os.Create(fmt.Sprintf("dug-%02d.svg", i))
		if err != nil {
			fmt.Printf("can't party :( %s\n\n", err)
			continue
		}
		fmt.Printf("writing file %s\n", f.Name())

		props := map[string]string{
			"rotate": r,
		}

		for c := 0; c <= 9; c++ {
			key := fmt.Sprintf("color%d", c)
			props[key] = colors[(c+i)%len(colors)]
		}

		if err := party.Execute(f, props); err != nil {
			fmt.Printf("really can't party :( :( %s\n\n", err)
			return
		}
		_ = f.Close()
	}
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

#F0C070
#FFCF7F
f7e6c8

*/
