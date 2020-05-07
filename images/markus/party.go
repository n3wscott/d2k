package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	partyTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" preserveAspectRatio="xMidYMid meet" viewBox="0 0 160 220" style="enable-background:new 0 0 160 220;">

<defs>

<path d="M479.51 326.27C500.36 304.89 514.72 301.04 522.58 314.72C530.44 328.39 528.48 348.93 516.69 376.32C514.83 386.38 511.96 396.61 508.09 407.01C502.29 422.6 503.59 415.53 493.46 438.68C486.71 454.12 475.94 458.58 461.15 452.06L158.96 476.68L145.46 473.18C135.63 467.08 126.63 456.43 118.46 441.25C112.46 430.09 114.96 432.8 108.46 410.51C101.96 388.22 97.58 391.42 96.46 384.08C92.84 360.22 95.12 354.02 103.96 345.8C109.86 340.32 119.16 341.45 131.85 349.18L479.51 326.27Z" id="ears"></path>

<path d="M458.95 146.34C503.24 212.51 477.72 265.47 477.72 327.55C477.72 443.69 414.63 564.61 363.27 578.06C345.3 582.77 336.46 584.77 331.33 585.48C321.78 586.8 315.19 585.3 305.98 585.16C296.76 585.02 275.06 595.01 244.2 581.8C213.34 568.59 151.36 486.28 145.94 431.59C145.85 430.65 138.69 388.91 124.46 306.36C120.97 239.21 127.23 190.58 143.23 160.47C175.69 99.37 249.75 62.88 305.98 62.88C362.2 62.88 414.65 80.18 458.95 146.34Z" id="face"></path>
<path d="M502.64 325.58L506.05 375.74C514.3 361.84 519.13 350.01 520.55 340.24C522.68 325.58 517.3 312.53 512.55 310.62C507.8 308.72 497.65 318.92 495.55 323.12C494.15 325.93 496.51 326.74 502.64 325.58Z" id="r-innerear"></path>

<path d="M487.16 434.13C487.16 440.62 483.24 445.88 478.41 445.88C473.58 445.88 469.66 440.62 469.66 434.13C469.66 427.65 473.58 422.38 478.41 422.38C483.24 422.38 487.16 427.65 487.16 434.13Z" id="r-plug"></path>

<path d="M163.66 454.31C163.66 459.69 159.74 464.06 154.91 464.06C150.08 464.06 146.16 459.69 146.16 454.31C146.16 448.93 150.08 444.56 154.91 444.56C159.74 444.56 163.66 448.93 163.66 454.31Z" id="l-plug"></path>

<path d="M110.66 377.38L116.16 409.88C109.58 398.8 105.41 388.47 103.66 378.88C101.04 364.51 102.04 358.25 103.66 354.38C105.29 350.52 117.41 353.17 118.16 361.42C118.66 366.92 116.16 372.24 110.66 377.38Z" id="l-inner-ear"></path>

<path d="M284.36 43.6C283.32 41.81 281.91 40.08 280.1 38.39C279.66 37.97 280.13 37.25 280.69 37.49C289.97 41.55 313.16 51.68 350.26 67.89C344.42 58.87 340.77 53.22 339.31 50.97C339.04 50.55 339.63 50.11 339.95 50.48C343.54 54.6 349.59 58.73 358.12 62.88C374.24 70.71 371.27 70.19 387.2 75.54C403.12 80.88 399.04 79.85 400.81 80.27C401.86 80.52 407.62 80.23 418.08 79.4C418.57 79.36 418.77 80.02 418.33 80.25C417.31 80.79 414.76 82.13 410.68 84.29L437.1 103.43C437.33 100.89 437.47 99.3 437.53 98.66C437.57 98.18 438.26 98.15 438.35 98.62C438.7 100.52 439.16 102.12 439.74 103.43C440.66 105.52 441.85 107.12 443.32 108.22C456.24 117.84 463.15 123.16 464.05 124.19C464.5 124.71 465.42 125.34 466.81 126.08C471.32 124.55 474.14 123.6 475.27 123.21C475.7 123.07 475.96 123.66 475.58 123.89C471.5 126.22 469.64 128.14 470 129.64C470.8 132.96 472.76 134.86 473.58 136.1C489.48 160.26 496.17 193.24 496.44 235.61C496.71 277.97 490.65 372.81 475.28 373.08C461.46 373.32 466.26 254.32 435.21 198.78C404.15 143.24 390.39 154.19 363.48 150.23C336.57 146.28 296.27 166.14 280.94 166.14C264.56 166.14 235.02 158.52 211.42 166.14C190.06 173.04 153.13 190.29 148.07 219.31C146.28 229.57 142.82 263.49 137.68 321.09C139.55 365.94 138.5 388.06 134.53 387.44C121.84 385.47 97.78 216.37 112.55 180.99C127.31 145.6 132.18 136.97 156.21 111.25C180.24 85.53 232.18 52.01 286.92 52.01C287.61 52.01 286.64 47.54 284.36 43.6Z" id="hair"></path>

<path d="M256.52 292.23C256.52 297.33 227.72 291.41 210.05 292.23C192.39 293.04 161.96 307.11 161.96 302.01C161.96 296.91 193.16 272.12 218.21 272.12C243.25 272.12 256.52 287.13 256.52 292.23Z" id="l-brow"></path>

<path d="M429.35 283.77C427.17 289.78 401.91 278.33 371.47 278.33C341.02 278.33 315.08 292.73 311.41 285.94C307.74 279.15 332.87 257.68 363.32 257.68C393.76 257.68 431.52 277.75 429.35 283.77Z" id="r-brow"></path>

<path d="M324.03 419.11C343.75 415.46 369.87 418.28 378.81 422.56C400.44 432.93 392.97 452.89 392.93 456.53C392.88 460.65 386.13 445.65 363.06 442.76C352.45 441.42 334.57 443.06 309.43 447.66C302.28 444.43 297.18 442.79 294.13 442.76C290.96 442.72 286.78 443.81 281.59 446.03C267.08 444.9 254.52 445.45 243.93 447.66C228.7 450.85 217.13 456.31 209.24 464.06C206.08 453.87 207.17 445.63 212.5 439.34C220.51 429.9 222.66 427.39 241.25 424.13C253.64 421.95 270.29 422.68 291.2 426.3C299.93 423.94 310.87 421.55 324.03 419.11Z" id="stash"></path>

<path d="M342.93 497.39C344.64 505.98 339.81 514.05 328.45 521.62C317.08 529.18 306.42 534.78 296.47 538.42C284.98 534.62 274.59 529.93 265.31 524.36C251.38 516.01 244.02 508.38 244.02 505C244.02 501.42 261.46 501.36 296.33 504.82C297.5 504.94 298.69 504.9 299.85 504.7C305.6 503.73 319.96 501.29 342.93 497.39Z" id="soul"></path>

<path d="M321.75 364.64C334.86 374.59 341.41 384.1 341.41 393.18C341.41 407.33 331.14 397.15 317.26 405.52C303.39 413.9 300.54 424.18 285.91 424.18C271.29 424.18 270.5 407.04 258.76 405.52C246.29 403.91 238.91 415.45 238.91 397.68C238.91 389.29 247.5 376.57 264.69 359.51C271.33 318.11 279.4 297.83 288.91 298.68C298.42 299.54 309.37 321.52 321.75 364.64Z" id="nose"></path>

<path d="M333.91 551.6L363.95 539.56C375.01 529.19 383.09 515.53 388.18 498.56C393.09 482.17 394.62 469.78 392.76 461.37C392.64 460.82 393.19 460.35 393.71 460.57C406.44 465.74 418.87 460.88 431.02 445.98C442.07 432.43 455.61 406.71 471.66 368.82C472.73 366.31 476.46 367.04 476.5 369.77C477.39 430.44 465.85 477.91 441.89 512.17C425.76 535.25 409.07 552.82 388.18 565.72C358.74 583.9 334.47 591.3 326.82 593.23C313.74 596.53 307.25 594.73 297.58 594.73C290.94 594.73 282.13 602.81 264.94 594.73C259.14 592 229.22 578.12 206.91 557.49C184.6 536.85 186.38 529.14 175.69 512.17C156.39 481.55 141.84 438.27 132.04 382.32C131.58 379.69 135.27 378.57 136.35 381.02C157.48 428.94 173.93 459.89 185.71 473.87C196.35 486.5 204.39 485.28 209.83 470.21C210.25 469.06 211.96 469.37 211.94 470.6C211.4 502.73 212.94 522.35 216.55 529.46C222.56 541.32 231.34 551.6 249.6 551.6C267.86 551.6 286.08 543.46 289.58 543.46C292.04 543.46 306.81 546.18 333.91 551.6Z" id="beard"></path>

<path d="M131.3 335.3C137.54 335.52 141.07 336.79 141.9 339.1C143.15 342.58 151.69 372.08 157.08 378C162.48 383.93 174.47 384.64 195.48 385.2C216.49 385.76 225.27 383.87 241.12 380.24C256.97 376.62 254.79 374.87 258.88 370.72C260.78 368.79 264.48 366.48 268.53 354.31C272.58 342.14 276.34 329.3 278.48 325.97C282.48 319.74 283.03 319.6 288.08 319.6C293.13 319.6 293.88 319.79 298.68 325.97C299.14 326.55 305.95 341.93 312.1 354.31C318.25 366.7 328.94 369.84 331.08 370.72C333.23 371.6 379.88 374.56 401.48 369.24C423.08 363.93 425.74 362.24 428.84 357.24C430.9 353.91 432.85 346.98 434.68 336.44C436.38 326.74 437.98 319.74 439.48 315.44C441.73 308.99 443.18 312.14 444.28 311.44C445.01 310.98 455.01 311.38 474.28 312.64L475.28 308.44L441.72 289.26C410.39 285.38 382.99 284.61 359.51 286.95C336.03 289.29 316.48 294.75 300.84 303.32L274.12 305.42C256.59 300.41 237.54 298.29 216.97 299.06C196.4 299.83 174.32 303.48 150.72 310.02L132.46 316.62L131.3 335.3Z" id="glasses"></path>

<path d="M153.48 353.4C154.33 363.14 157.15 370.07 161.96 374.2C169.39 378.93 178.43 381.38 189.08 381.57C205.06 381.85 247.48 376.47 253.08 370.52C263.71 359.22 268.13 336.13 268.48 327.4C268.83 318.68 271.78 318.3 263.28 310C259.03 305.85 231.36 303.07 206.76 304.32C182.16 305.57 159.11 312.78 156.88 315.8C152.43 321.85 152.21 338.8 153.48 353.4Z" id="l-lenz"></path>

<path d="M312.66 339.78C317.01 352.48 324.26 360.93 329.26 363.78C334.26 366.63 349.91 370.03 373.46 367.98C397.01 365.93 419.36 362.88 423.46 355.58C427.56 348.28 429.41 342.53 428.26 327.58C427.11 312.63 427.36 302.78 418.86 295.78C410.36 288.78 389.21 290.73 371.86 291.58C354.51 292.43 322.66 298.83 316.66 301.58C310.66 304.33 308.46 307.03 307.06 313.78C305.66 320.53 308.31 327.08 312.66 339.78Z" id="r-lenz"></path>

</defs>

<g transform="rotate({{.rotate}} 80 190)">
<g transform="translate(-12 15)">
<g transform="scale(0.30 0.30)">

<use xlink:href="#ears" opacity="1" fill="#d4b8a9" fill-opacity="1"></use>
<use xlink:href="#ears" opacity="1" fill-opacity="0" stroke="#d4b8a9" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#r-innerear" opacity="1" fill="#CE9691" fill-opacity="1"></use>
<use xlink:href="#r-innerear" opacity="1" fill-opacity="0" stroke="#CE9691" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#l-inner-ear" opacity="1" fill="#CE9691" fill-opacity="1"></use>
<use xlink:href="#l-inner-ear" opacity="1" fill-opacity="0" stroke="#CE9691" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#r-plug" opacity="1" fill="{{.color1}}" fill-opacity="1"></use>
<use xlink:href="#r-plug" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#l-plug" opacity="1" fill="{{.color1}}" fill-opacity="1"></use>
<use xlink:href="#l-plug" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#face" opacity="1" fill="#d4b8a9" fill-opacity="1"></use>
<use xlink:href="#hair" opacity="1" fill="{{.color2}}" fill-opacity="1"></use>
<use xlink:href="#l-brow" opacity="1" fill="{{.color1}}" fill-opacity="1"></use>
<use xlink:href="#r-brow" opacity="1" fill="{{.color1}}" fill-opacity="1"></use>
<use xlink:href="#stash" opacity="1" fill="{{.color1}}" fill-opacity="1"></use>
<use xlink:href="#soul" opacity="1" fill="{{.color1}}" fill-opacity="1"></use>
<use xlink:href="#nose" opacity="1" fill="#CE9691" fill-opacity="1"></use>
<use xlink:href="#beard" opacity="1" fill="{{.color2}}" fill-opacity="1"></use>
<use xlink:href="#glasses" opacity="1" fill="#000000" fill-opacity="1"></use>
<use xlink:href="#glasses" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#l-lenz" opacity="1" fill="#d6cec9" fill-opacity="1"></use>
<use xlink:href="#l-lenz" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="1" stroke-opacity="1"></use>
<use xlink:href="#r-lenz" opacity="1" fill="#d6cec9" fill-opacity="1"></use>
<use xlink:href="#r-lenz" opacity="1" fill-opacity="0" stroke="#000000" stroke-width="1" stroke-opacity="1"></use>

</g></g></g>

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
	"2.5",
	"5",
	"7.5",
	"10",
	"7.5",
	"5",
	"2.5",
	"0",
	"-2.5",
	"-5",
	"-7.5",
	"-10",
	"-7.5",
	"-5",
	"-2.5",
}

func main() {
	party, err := template.New("party").Parse(partyTemplate)
	if err != nil {
		fmt.Printf("can't party :( %s\n\n", err)
		return
	}

	for i, r := range rotations {
		f, err := os.Create(fmt.Sprintf("markus-%02d.svg", i))
		if err != nil {
			fmt.Printf("can't party :( %s\n\n", err)
			continue
		}
		fmt.Printf("writing file %s\n", f.Name())

		if err := party.Execute(f, map[string]string{
			"rotate": r,
			"color1": colors[i%len(colors)],
			"color2": colors[(i+len(colors)/2)%len(colors)],
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
