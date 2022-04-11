package main

import (
	"LogoGenerator/emojipedia"
	"LogoGenerator/image"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("LogoGenerator - Create simple logos using emoji\n\n")

	if len(os.Args) != 4 {
		fmt.Printf(`Usage: %s "emoji search term" "color" <size>`, os.Args[0])
		os.Exit(1)
	}

	if len(os.Args[2]) > 7 || len(os.Args[2]) < 6 {
		fmt.Println("Color must be in hex format (leading # is optional) e.g. #ad5ff2")
		os.Exit(1)
	}

	size, err := strconv.Atoi(os.Args[3])

	if err != nil {
		fmt.Println("Size must be an integer less than or equal to 2048")
		os.Exit(1)
	} else if size > 2048 {
		fmt.Println("Size must be an integer less than or equal to 2048")
		os.Exit(1)
	}

	fmt.Println("Searching for emoji...")
	emoji, err := emojipedia.Search(os.Args[1])

	if err != nil {
		if err == emojipedia.ErrNoEmoji {
			fmt.Println("Couldn't find an emoji with the search term provided")
		} else if err == emojipedia.ErrNoUrl {
			fmt.Println("Couldn't fetch image URL for emoji")
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
	}

	fmt.Println("Please wait, generating image...")
	err = image.Generate(emoji, os.Args[2], size)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Logo saved to output.png!")
}
