package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	partyTemplate = `<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 viewBox="0 0 700 700" style="enable-background:new 0 0 700 700;" xml:space="preserve">
 
<style type="text/css">
	.box{
   opacity:1;
   fill-opacity:1;
   fill:#000000;
   stroke:#000000;
   stroke-width:1;
   stroke-opacity:1;
  }
  .b0 { fill:{{ .color0 }}; }
  .b1 { fill:{{ .color1 }}; }
  .b2 { fill:{{ .color2 }}; }
  .b3 { fill:{{ .color3 }}; }
  .b4 { fill:{{ .color4 }}; }
  .b5 { fill:{{ .color5 }}; }
  .b6 { fill:{{ .color6 }}; }
</style>


<g transform="rotate({{ .rotate }} 350 350)">
<g transform="translate(100 100)">

<rect class="box b6" x="200" y="0" rx="1" ry="1" width="100" height="100" />

<rect class="box b4" x="100" y="100" rx="1" ry="1" width="100" height="100" />
<rect class="box b5" x="200" y="100" rx="1" ry="1" width="100" height="100" />
<rect class="box b4" x="300" y="100" rx="1" ry="1" width="100" height="100" />

<rect class="box b3" x="0" y="200" rx="1" ry="1" width="100" height="100" />
<rect class="box b0" x="200" y="200" rx="1" ry="1" width="100" height="100" />
<rect class="box b3" x="400" y="200" rx="1" ry="1" width="100" height="100" />


<rect class="box b2" x="0" y="400" rx="1" ry="1" width="100" height="100" />
<rect class="box b1" x="100" y="400" rx="1" ry="1" width="100" height="100" />
<rect class="box b1" x="300" y="400" rx="1" ry="1" width="100" height="100" />
<rect class="box b2" x="400" y="400" rx="1" ry="1" width="100" height="100" />

</g>
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
	"5",
	"10",
	"15",
	"20",
	"15",
	"10",
	"5",
	"0",
	"-5",
	"-10",
	"-15",
	"-20",
	"-15",
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
		f, err := os.Create(fmt.Sprintf("agababov-%02d.svg", i))
		if err != nil {
			fmt.Printf("can't party :( %s\n\n", err)
			continue
		}
		fmt.Printf("writing file %s\n", f.Name())

		props := map[string]string{
			"rotate": r,
		}

		for c := 0; c <= 6; c++ {
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
