package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	partyTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg version="1.1" id="my-icon" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 viewBox="0 0 640 640" style="enable-background:new 0 0 640 640;" xml:space="preserve">

<defs>
<path d="M147.49 137.22C166.47 137.22 182.16 152 183.3 170.95C186.95 231.78 196.07 383.88 210.68 627.23C190.8 627.23 178.38 627.23 173.41 627.23C163.48 627.23 155.18 619.69 154.22 609.81C145.07 515.29 120.19 258.39 111.26 166.17C109.76 150.66 121.96 137.22 137.54 137.22C146.29 137.22 139.68 137.22 147.49 137.22Z" id="eoRlRX90x">
</path>
<path d="M499.14 137.22C517.06 137.22 530.9 152.96 528.62 170.73C516.88 262.03 484.34 515.01 472.2 609.45C470.89 619.62 462.24 627.23 451.99 627.23C447.08 627.23 434.79 627.23 415.14 627.23C437.71 381.34 451.81 227.66 457.45 166.19C458.96 149.78 472.72 137.22 489.2 137.22C497.94 137.22 491.33 137.22 499.14 137.22Z" id="f28f1I7Mej">
</path>
<path d="M218.4 137.22C237.1 137.22 252.31 152.3 252.47 171.01C252.99 231.84 254.3 383.91 256.38 627.23L205.81 627.23C192.55 381.38 184.25 227.73 180.94 166.27C180.09 150.49 192.65 137.22 208.46 137.22C217.2 137.22 210.59 137.22 218.4 137.22Z" id="b4XTskvW5e">
</path>
<path d="M289.31 137.22C307.86 137.22 322.79 152.46 322.42 171C321.2 231.83 318.16 383.91 313.29 627.23L257.12 627.23C253.69 381.4 251.54 227.76 250.69 166.31C250.46 150.31 263.37 137.22 279.37 137.22C288.11 137.22 281.5 137.22 289.31 137.22Z" id="fd2DJcBom">
</path>
<path d="M360.22 137.22C378.54 137.22 393.08 152.66 391.98 170.95C388.33 231.78 379.2 383.88 364.6 627.23L312.63 627.23C316.92 381.4 319.61 227.76 320.68 166.3C320.96 150.16 334.13 137.22 350.28 137.22C359.02 137.22 352.41 137.22 360.22 137.22Z" id="aa6yd3GEA">
</path>
<path d="M431.48 137.22C449.58 137.22 463.74 152.83 461.99 170.85C456.08 231.7 441.3 383.83 417.65 627.23L362.88 627.23C377.71 381.38 386.99 227.72 390.69 166.26C391.68 149.95 405.19 137.22 421.53 137.22C430.28 137.22 423.67 137.22 431.48 137.22Z" id="b1WMzuu2q8">
</path>
</defs>


<style type="text/css">
	.st0{fill:#FFFFFF;}
	.st1{fill:#4A6475;}
	.st2{fill:#DDB7B1;}
	.st3{fill:#134C8E;}  //shirt
	.st4{fill:#444444;}
	.st5{fill:#4F4F4F;}
	.st6{fill:#026FE8;}  // k8s logo
	.st7{fill:#F7F9FD;}
	.st8{fill:{{ .color }};} // hair
	.st9{fill:#756042;}
	.st10{fill:#CE9691;}
	.st11{fill:#1F3342;}
</style>


<defs>
<g id="matt">
<g transform="scale(0.5 0.5)">
<g transform="rotate({{ .rotate }} 80 110)">
<g transform="translate(-885 -75)">
// face
<path class="st2" d="M907.8,115.5c-3.9,14.2,2.6,77.6,5.2,86.6c2.6,9,29.7,80.1,68.5,58.2c38.8-22,42.7-45.2,42.7-45.2
	s15.5-80.1,0-100.8S907.8,115.5,907.8,115.5z"/>
// hair
<path class="st8" d="M1033.1,126.5c-2.6-8.2-12.9-32.7-32.3-35.4c-9.3-1.3-40.1-5.5-50.4-2.7c-10.3,2.7-39.5,12.7-47.8,25.9
	c-15.5,24.5-3.9,158.8,65.9,161.6c35,1.4,63.3-28.2,63.3-68.9c0-5.1,1.3-15.3,1.2-21.8C1034,171.6,1035.4,133.6,1033.1,126.5z
	 M991.8,236.9c-3.9-9.5-11.7-6.6-23.3-8.2c-15.5-2.1-18.1,4.4-20.7,9.6c-0.9,1.7-8.3-4.2-9-5.2c-12.5-16.3-22-12.9-28-58.5
	c0-10.8,1.2-23.8,6-30.4c4-5.4,10.3-9.5,14.2-8.2c3.9,1.4,2.6,16.4,3.9,13.6c1.3-2.7,7.8-19.1,14.2-25.9c6.5-6.8,6.5,16.4,9,17.7
	c2.6,1.4,14.2-24.5,14.2-24.5s14.2,12.3,14.2,17.7c0,5.5,7.8-8.2,10.3-12.3c0,0,12.9,23.2,12.9,20.4s7.8-12.3,7.8-9.5
	c0,2.4,12.9,36.5,2.5,64.7C1018.9,216.4,994.9,244.5,991.8,236.9z"/>
// Mushtash
<path class="st9" d="M929.7,230.1c0,0,12.9-16.4,18.1-15c5.2,1.4,41.4-1.4,44-1.4c2.6,0,16.8,17.7,16.8,17.7s2.6-32.9-25.9-31.5
	c-7.7,0.4-28.4,0-31,0C949.1,199.9,929.7,205.5,929.7,230.1z"/>
// nose
<path class="st10" d="M958.2,178.9c0,28.6-2.6,30.6,12.9,30.6c11.9,0,9-30.6,9-30.6H958.2z"/>
</g>
</g>
</g>
</g>

</defs>

<g>


<g transform="translate(100 60)"><use xlink:href="#matt"></use></g> <!-- left -->
<g transform="translate(450 70)"><use xlink:href="#matt"></use></g> <!-- right -->
<g transform="translate(275 0)"><use xlink:href="#matt"></use></g> <!-- top middle -->

<g transform="translate(260 65)"><use xlink:href="#matt"></use></g>

<g transform="translate(200 15)"><use xlink:href="#matt"></use></g>
<g transform="translate(240 35)"><use xlink:href="#matt"></use></g>
<g transform="translate(310 20)"><use xlink:href="#matt"></use></g>
<g transform="translate(200 50)"><use xlink:href="#matt"></use></g>
<g transform="translate(140 50)"><use xlink:href="#matt"></use></g>
<g transform="translate(150 70)"><use xlink:href="#matt"></use></g>

<g transform="translate(390 80)"><use xlink:href="#matt"></use></g>
<g transform="translate(340 65)"><use xlink:href="#matt"></use></g>
<g transform="translate(310 90)"><use xlink:href="#matt"></use></g>
<g transform="translate(370 50)"><use xlink:href="#matt"></use></g>



<use xlink:href="#eoRlRX90x" opacity="1" fill="#e73b3b" fill-opacity="1"></use>
<use xlink:href="#eoRlRX90x" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="2" stroke-opacity="1"></use>
<use xlink:href="#f28f1I7Mej" opacity="1" fill="#fcebeb" fill-opacity="1"></use>
<use xlink:href="#f28f1I7Mej" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="2" stroke-opacity="1"></use>
<use xlink:href="#b4XTskvW5e" opacity="1" fill="#fcebeb" fill-opacity="1"></use>
<use xlink:href="#b4XTskvW5e" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="2" stroke-opacity="1"></use>
<use xlink:href="#fd2DJcBom" opacity="1" fill="#e73b3b" fill-opacity="1"></use>
<use xlink:href="#fd2DJcBom" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="2" stroke-opacity="1"></use>
<use xlink:href="#aa6yd3GEA" opacity="1" fill="#fcebeb" fill-opacity="1"></use>
<use xlink:href="#aa6yd3GEA" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="2" stroke-opacity="1"></use>
<use xlink:href="#b1WMzuu2q8" opacity="1" fill="#e73b3b" fill-opacity="1"></use>
<use xlink:href="#b1WMzuu2q8" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="2" stroke-opacity="1"></use>

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

var rotations = []int{
	0,
	10,
	20,
	10,
	0,
	-10,
	-20,
	-10,
}

func main() {
	party, err := template.New("party").Parse(partyTemplate)
	if err != nil {
		fmt.Printf("can't party :( %s\n\n", err)
		return
	}

	for i, r := range rotations {
		f, err := os.Create(fmt.Sprintf("kernel-matt-%02d.svg", i))
		if err != nil {
			fmt.Printf("can't party :( %s\n\n", err)
			continue
		}
		fmt.Printf("writing file %s\n", f.Name())

		if err := party.Execute(f, map[string]string{
			"rotate": fmt.Sprintf("%d", r),
			"color":  colors[i%len(colors)],
		}); err != nil {
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

*/
