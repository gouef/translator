<img align=right width="168" src="docs/gouef_logo.png">

# Translator
This package provides a flexible translation system in Go, supporting multiple translation formats including YAML and PO files.

[![Static Badge](https://img.shields.io/badge/Github-gouef%2Ftranslator-blue?style=for-the-badge&logo=github&link=github.com%2Fgouef%2Ftranslator)](https://github.com/gouef/translator)

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/translator.svg)](https://pkg.go.dev/github.com/gouef/translator)
[![GitHub stars](https://img.shields.io/github/stars/gouef/translator?style=social)](https://github.com/gouef/translator/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/translator)](https://goreportcard.com/report/github.com/gouef/translator)
[![codecov](https://codecov.io/github/gouef/translator/branch/main/graph/badge.svg?token=YUG8EMH6Q8)](https://codecov.io/github/gouef/translator)

## Versions
![Stable Version](https://img.shields.io/github/v/release/gouef/translator?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/translator?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/translator?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)


## Features
- Support for multiple translation formats (YAML, PO/MO).
- Thread-safe implementation with `sync.RWMutex`.
- Supports hierarchical and nested keys for translations.
- Easy registration of multiple translators.
- Dynamic language switching.

## Installation

```sh
go get -u github.com/gouef/translator
```

## Usage

### Initializing the Translator

```go
package main

import (
	"fmt"
	"log"
	"github.com/gouef/translator"
)

func main() {
	translator.Init()

	// Load YAML translations
	yamlTranslator, err := translator.NewYaml("locales")
	if err != nil {
		log.Fatal(err)
	}
	translator.Register(yamlTranslator)

	// Set active language
	err = translator.SetLanguage("cs_CZ")
	if err != nil {
		log.Fatal(err)
	}

	// Translate a key
	fmt.Println(translator.Translate("hello"))
}
```

### Using PO Translations

```go
package main

import (
	"fmt"
	"log"
	"github.com/gouef/translator"
)

func main() {
	translator.Init()

	poTranslator, err := translator.NewPO("locales")
	if err != nil {
		log.Fatal(err)
	}
	translator.Register(poTranslator)

	err = translator.SetLanguage("cs_CZ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(translator.Translate("errors.hello"))
}
```

## API Documentation

### `translator.Init()`
Initializes the translator system by clearing registered translators.

### `translator.Register(translator Translator)`
Registers a new translation provider (YAML, PO, etc.). The last registered provider becomes the active one.

### `translator.SetLanguage(lang string) error`
Sets the active language for translation.

### `translator.GetLanguage() string`
Returns the currently active language.

### `translator.Translate(key string, args ...interface{}) string`
Translates a given key using the active translation provider. Supports placeholders with formatting.

## YAML Format

Example `locales/cs_CZ/messages.yaml`:

```yaml
hello: "Ahoj"
errors:
  hello: "Chyba: Ahoj"
```

## PO/MO Format

Example `locales/cs_CZ/messages.po`:

```po
msgid "hello"
msgstr "Ahoj"

msgid "errors.hello"
msgstr "Chyba: Ahoj"
```


## Contributing

Read [Contributing](CONTRIBUTING.md)

## Contributors

<div>
<span>
  <a href="https://github.com/JanGalek"><img src="https://raw.githubusercontent.com/gouef/translator/refs/heads/contributors-svg/.github/contributors/JanGalek.svg" alt="JanGalek" /></a>
</span>
<span>
  <a href="https://github.com/actions-user"><img src="https://raw.githubusercontent.com/gouef/translator/refs/heads/contributors-svg/.github/contributors/actions-user.svg" alt="actions-user" /></a>
</span>
</div>

## Join our Discord Community! 🎉

[![Discord](https://img.shields.io/discord/1334331501462163509?style=for-the-badge&logo=discord&logoColor=white&logoSize=auto&label=Community%20discord&labelColor=blue&link=https%3A%2F%2Fdiscord.gg%2FwjGqeWFnqK
)](https://discord.gg/wjGqeWFnqK)

Click above to join our community on Discord!
