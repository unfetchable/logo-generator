# Logo Generator

Create simple logos using emoji

## Building

Ensure you have [go installed](https://golang.org/doc/install) before continuing.

You can build the application using the following command:

```
$ go build
```

This will create an executable (`LogoGenerator`) you can use to run the application.

## Usage

The application takes an emoji name/search term, a hex color and a size. The search term provided is passed to Emojipedia and the first result is used.

```
$ ./LogoGenerator "emoji search term" "color" size
```

### Example

```
$ ./LogoGenerator "fire" "#ad5ff2" 256
```