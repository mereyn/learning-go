package main

func main() {
	// initialize option 1
	// var colors map[string]string

	// initialize option 2
	// colors := make(map[string]string)
	// colors["res"] = "#FF0000"

	// sample
	// colors := make(map[int]string)
	// colors[10] = "#FF0000"

	// delete value on existing map
	// delete(colors, 10)

	// map of key is string and value is tring
	colors := colors{
		"red":    "#FF0000",
		"maroon": "#800000",
		"white":  "#ffffff",
	}
	colors.printMap()
	// fmt.Println(colors)
}
