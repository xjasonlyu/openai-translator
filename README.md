# OpenAI Translator

This repo was originally forked from: <https://github.com/zijiren233/openai-translator>.

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
