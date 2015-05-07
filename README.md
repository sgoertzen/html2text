# html2text
An html to text converter written in Go.  This library will strip the html tags from the source and perform clean up on the text.  This includes things like adding new lines correctly and appending on urls from links.

[![Build Status](https://travis-ci.org/sgoertzen/html2text.svg)](https://travis-ci.org/sgoertzen/html2text)

## Install
go get github.com/sgoertzen/html2text

## Usage
```sh
import "github.com/sgoertzen/html2text"

func main() {
	t := html2text.Textify("<div>hello</div>")
	log.Println(t)  // Outputs "hello"
}
```

## More Examples
See the included html2text_test.go file for more usage examples.

## Enhancements
If you encounter html that doesn't work properly please open an issue with the specific html and desired text.
