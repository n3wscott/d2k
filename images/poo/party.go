package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	partyTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
    <svg
       version="1.1"
       id="Layer_1"
       x="0px"
       y="0px"
       viewBox="-0 -30 600 600"
       xml:space="preserve"
       width="600"
       height="600"
       xmlns="http://www.w3.org/2000/svg"
       xmlns:svg="http://www.w3.org/2000/svg">
 
  <defs>
    <radialGradient id="rg1" cx="0.474" cy="0.472" r="50%">
      <stop offset="5%" style="stop-color:{{ .color0 }};"/>
      <stop offset="15%" style="stop-color:{{ .color1 }};"/>
      <stop offset="30%" style="stop-color:{{ .color2 }};"/>
      <stop offset="40%" style="stop-color:{{ .color3 }};"/>
      <stop offset="75%" style="stop-color:{{ .color4 }};"/>
      <stop offset="90%" style="stop-color:{{ .color5 }};"/>
      <stop offset="100%" style="stop-color:{{ .color6 }};"/>
    </radialGradient>

    
    <mask id="poo-mask">
  <g
   id="poo"
   transform="matrix(4.1786926,0,0,4.380562,-1866.3929,-1999.1813)">
   <path
   id="path6295-5-8-0"
   style="display:inline;fill:#ffffff;fill-opacity:1;stroke:#945f01;stroke-width:4.50723;stroke-opacity:1"
   d="m 542.05579,496.99107 c 1.36857,4.09434 -0.57813,10.65416 -5.4254,14.83084 -14.83104,7.80679 -37.37243,14.53311 -43.88999,15.65657 -7.52599,-0.16384 -10.99026,-2.74497 -12.51133,-7.29551 -1.68717,-5.04751 0.61441,-10.49481 3.79938,-13.03279 4.28164,-3.09477 15.04361,-6.26498 18.70304,-11.029 4.72827,-6.15548 1.45054,-12.94358 -0.44456,-17.40553 8.35612,-2.94562 15.92252,8.50557 25.86763,12.75143 0,0 12.4141,1.07492 13.90123,5.52399 z" />
     <path
   id="path6295-4"
   style="display:inline;fill:#ffffff;fill-opacity:1;stroke:#945f01;stroke-width:4.50723;stroke-opacity:1"
   d="m 576.81793,554.06907 c 0.77225,6.17788 -4.83865,14.24211 -14.37978,17.77844 -27.38089,4.05345 -67.12396,3.22099 -78.31589,1.86255 -12.39118,-3.5569 -17.1819,-8.60396 -18.04019,-15.47017 -0.95203,-7.61608 4.84056,-14.0092 11.03452,-16.05195 10.67256,-0.57398 30.1607,4.05507 45.35116,3.74236 14.89757,-0.3067 25.5255,-5.56095 33.36501,-5.53642 13.13183,-0.82713 20.14603,6.9621 20.98517,13.67519 z" />
     <path
   id="path6295-5"
   style="display:inline;fill:#ffffff;fill-opacity:1;stroke:#945f01;stroke-width:3.63935;stroke-opacity:1"
   d="m 557.25679,512.64244 c 1.89857,4.64259 -0.80204,12.08078 -7.52646,16.81672 -20.57452,8.85213 -51.84531,16.47913 -60.88687,17.75302 -10.44051,-0.18577 -15.24635,-3.11253 -17.35645,-8.2724 -2.34057,-5.72338 0.85232,-11.90008 5.27071,-14.7779 8.23033,-2.66823 24.44838,-3.12639 36.26841,-6.53209 11.59213,-3.34003 18.80696,-9.63793 24.94601,-11.25101 10.10152,-3.37706 17.2216,1.21885 19.28465,6.26366 z" />
     <path
   id="path6295"
   style="fill:#ffffff;fill-opacity:1;stroke:#945f01;stroke-width:4.42086;stroke-opacity:1"
   d="m 568.58317,532.39137 c 2.34392,5.54896 -0.99017,14.43928 -9.29192,20.09981 -25.40065,10.58029 -64.00656,19.69629 -75.16899,21.21888 -12.88952,-0.22204 -18.82266,-3.72018 -21.42772,-9.88739 -2.88959,-6.84075 1.05225,-14.22331 6.50706,-17.66296 10.1609,-3.18914 30.18318,-3.73674 44.77582,-7.80732 14.31126,-3.99209 23.21847,-11.51952 30.79754,-13.4475 12.47101,-4.03636 21.26124,1.4568 23.80821,7.48648 z" />
     </g>
    </mask>
      </defs>
 
    
    <g mask="url(#poo-mask)" transform="rotate({{ .rotate }}, 300, 300)">
       <g id="spirals" transform="scale(5) translate(300,250) rotate({{ .spin }})" style="fill:none;stroke-width:80px;stroke-linecap:round;">
    <path id="spiral1" class="sprial" style="stroke:url(#rg1);" d="M0,0C3.36,0.06,6.6,-2.82,7.07,-7.07C7.66,-11.26,5.28,-16.67,0,-20C-5.18,-23.4,-13.27,-24.43,-21.21,-21.21C-29.13,-18.12,-36.69,-10.51,-40,0C-43.42,10.46,-42.3,23.68,-35.36,35.36C-28.53,47,-15.79,56.71,0,60C15.69,63.43,34.14,60.2,49.5,49.5C64.88,38.98,76.71,21.03,80,0C83.43,-20.92,78.08,-44.6,63.64,-63.64C49.44,-82.74,26.27,-96.74,0,-100C-26.16,-103.45,-55.05,-95.94,-77.78,-77.78C-100.6,-59.88,-116.76,-31.53,-120,0C-123.48,31.43,-113.8,65.49,-91.92,91.92C-70.34,118.47,-36.76,136.79,0,140C36.66,143.5,75.94,131.67,106.07,106.07C136.34,80.79,156.81,42,160,0C163.53,-41.9,149.54,-86.4,120.21,-120.21C91.24,-154.2,47.24,-176.83,0,-180C-47.13,-183.56,-96.85,-167.4,-134.35,-134.35C-172.06,-101.7,-196.86,-52.47,-200,0C-203.55,52.4,-185.28,107.29,-148.49,148.49C-112.14,189.9,-57.74,216.9,0,220C57.63,223.58,117.75,203.14,162.63,162.63C207.77,122.6,236.92,62.97,240,0C243.62,-62.86,220.98,-128.21,176.78,-176.78C133.05,-225.63,68.21,-256.95,0,-260C-68.13,-263.64,-138.65,-238.85,-190.92,-190.92C-243.49,-143.49,-276.98,-73.47,-280,0C-283.66,73.36,-256.71,149.1,-205.06,205.06C-153.94,261.36,-78.71,297,0,300C78.6,303.69,159.56,274.58,219.2,219.2C279.23,164.4,317.02,83.94,320,0C323.72,-83.84,292.44,-170.01,233.35,-233.35C174.84,-297.12,89.21,-337.04,0,-340C-89.07,-343.71,-180.46,-310.32,-247.49,-247.49C-315,-185.31,-357.03,-94.41,-360,0C-363.74,94.33,-328.18,190.91,-261.63,261.63C-195.75,332.86,-99.67,377.06,0,380C99.57,383.76,201.36,346.05,275.77,275.77C350.73,206.2,397.08,104.91,400,0"/>
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
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"5",
	"4",
	"3",
	"2",
	"1",
	"0",
	"-1",
	"-2",
	"-3",
	"-4",
	"-5",
	"-6",
	"-5",
	"-4",
	"-3",
	"-2",
	"-1",
}

func main() {
	party, err := template.New("party").Parse(partyTemplate)
	if err != nil {
		fmt.Printf("can't party :( %s\n\n", err)
		return
	}

	spin := 5 * 360 / len(rotations)
	scale := 1
	translate := 0

	fmt.Printf("\n\nmkdir out\n")

	for i, r := range rotations {
		file := fmt.Sprintf("poo-%02d.svg", i)
		png := fmt.Sprintf("poo-%02d.png", i)

		f, err := os.Create(file)
		if err != nil {
			fmt.Printf("can't party :( %s\n\n", err)
			continue
		}
		//fmt.Printf("writing file %s\n", f.Name())

		if i%4 == 0 {
			scale = scale * -1
			if scale < 0 {
				translate = -600
			} else {
				translate = 0
			}
		}

		props := map[string]string{
			"scale":     fmt.Sprintf("%d", scale),
			"translate": fmt.Sprintf("%d", translate),
			"rotate":    r,
			"spin":      fmt.Sprintf("%d", spin*i),
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

		fmt.Printf("inkscape -w 600 -h 630 %s -o out/%s\n", file, png)
	}

	fmt.Printf("\n\ncd out; convert -dispose previous -delay 5 -loop 0 *.png  party-poo.gif\n\n")
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