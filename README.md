# html2text
An html to text converter written in Go.  

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