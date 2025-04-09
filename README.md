# OpenAI Translator

A Golang translation library compatible with the [OpenAI API](https://platform.openai.com/docs/overview).

## Installation

Using the Go command, from inside your project:

```shell
go get -u github.com/xjasonlyu/openai-translator
```

## Usage

```go
package main

import (
	"fmt"

	translator "github.com/xjasonlyu/openai-translator"
)

func main() {
	q := "Text to translate"
	text, _ := translator.Translate(
		q, "fr", "<API Key>",
		translator.WithBaseURL("<Custom API Endpoint>"),
	)
	fmt.Println(text)
}
```

## Credits

- [zijiren233/openai-translator](https://github.com/zijiren233/openai-translator)

## License

This project is open-sourced under the MIT license. See the [LICENSE](LICENSE) file for more details.
