# Persian Funny Name Generator

The Persian Funny Name Generator is a Go package that generates humorous Persian names in both Finglish (Persian written with the Latin alphabet) and Persian script.

## Installation

To install the package, use the following command:

```sh
go get github.com/Ahmadrezadl/funny-name-generator
```

## Usage

Below is a basic example of how to use the Persian Funny Name Generator in your Go project.

```go
package main

import NameGenerator "github.com/Ahmadrezadl/funny-name-generator"

func main() {
	NameGenerator.Randomize()
	println("Finglish Name: " + NameGenerator.FinglishName())
	println("Persian Name: " + NameGenerator.PersianName())
}
```

## Output

When you run the code, it will output a random funny name in both Finglish and Persian script. For example:

```
Finglish Name: Ali Kalle Siah
Persian Name: کمال شیکم طلایی
```

## Functions

- `Randomize()`: Generates a new set of names.
- `FinglishName()`: Returns the generated name in Finglish.
- `PersianName()`: Returns the generated name in Persian script.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements.ails.
