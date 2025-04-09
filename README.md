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
	"log"
	"os"

	openai "github.com/xjasonlyu/openai-translator"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	apiURL := os.Getenv("OPENAI_API_URL")

	translator := openai.NewTranslator(
		apiKey,
		openai.WithBaseURL(apiURL),
	)

	text, err := translator.TranslateText("Hello, world!", "ZH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text) // "你好，世界！"
}

```

## Credits

- [zijiren233/openai-translator](https://github.com/zijiren233/openai-translator)

## License

This project is open-sourced under the MIT license. See the [LICENSE](LICENSE) file for more details.
