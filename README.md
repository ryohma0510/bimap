# bimap

BiMap is a Go package that provides a bidirectional map. It allows you to map keys to values and also values back to keys.

This package is primarily designed to manage the correspondence tables for enum conversions, such as mapping between protobuf enum integer values and their corresponding string values used in databases.

# 🚀 install

```sh
$ go get github.com/ryohma0510/bimap
```

```golang
import "github.com/ryohma0510/bimap"
```

# 💡 Example Usage

```golang
import (
    "github.com/samber/lo"
    "github.com/ryohma0510/bimap"
)

type GenderStr string

const (
	GenderStrMale   GenderStr = "male"
	GenderStrFemale GenderStr = "female"
)

type GenderInt int

const (
	GenderIntMale GenderInt = iota + 1
	GenderIntFemale
)

var genderMap = lo.Must(bimap.NewFromMap(map[GenderInt]GenderStr{
  GenderIntMale: GenderStrMale,
  GenderIntFemale: GenderStrFemale,
}))


func doSomething(genderInt GenderInt) GenderStr {
  // convert key to value
  strGender, err := genderMap.Get(genderInt)
  // ...
  // do something
  // ...
  // convert value to key
  genderStr, err := genderMap.InverseGet(valueFromDB)

  return genderStr
}

```
