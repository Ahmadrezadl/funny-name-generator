## Persian Funny Name Generator
Import:
```
go get github.com/Ahmadrezadl/funny-name-generator
```
Code:
```go
package main

import NameGenerator "github.com/Ahmadrezadl/funny-name-generator"

func main() {
	NameGenerator.Randomize()
	println("Finglish Name: " + NameGenerator.FinglishName())
	println("Persian Name: " + NameGenerator.PersianName())
}
```
Output:
```
Finglish Name: Ali Koon Siah
Persian Name: کمال مغز طلایی
```
