# go-utils

some common util method collection

if you has a question, but don't has. welcome commit `issue` or `PR`. please commit `PR` to `dev` branch.

install: `go get github.com/hujingnb/go-utils`

you can see document on [godoc](https://pkg.go.dev/github.com/hujingnb/go-utils).

# package 

## `hstring`

dispose string [Example](./hstring/example_test.go)

import: `import "github.com/hujingnb/go-utils/hstring"`

function:

* `Reverse`: reverse string
* `CamelToSnake`
* `SnakeToCamel`
* `ToString`: change var to string
* `Pad`: Pad a string to a certain length with another string
* `Join`: Join array elements with a string

## `hhash`

hash function. [Example](./hhash/example_test.go)

import: `import "github.com/hujingnb/go-utils/hhash"`

function:

* `Md5By32`: get MD5 Hash Generator 32 Characters
* `Md5By16`: get MD5 Hash Generator 16 Characters
* `Sha1`: get sha1
* `Sha256`
* `Sha512`

## `harray`

dispose array. [Example](./harray/example_test.go)

import: `import "github.com/hujingnb/go-utils/harray"`

function:

* `Chunk`:  Split an array into chunks
* `Shuffle`: array shuffle
* `InArray`: check item is in array
* `IndexOf`: get item first index in array
* `Unique`: Removes duplicate values from an array
* `Equal`: check array is same
* `Intersect`: get intersection of arrays
* `Diff`: get difference of arrays
* `BinarySearch`: binary search
* `GetSureRandArr`: get sure rand array by same seed
* `Count`: count value of array
* `Reverse`: reverse array
* sort:
    * `SortBubble`
    * `SortInsert`: insertion sort
    * `SortSelect`: selection sort
    * `SortQuick`: quick sort
    * `SortMerge`: merge sort

## `hmap`

dispose map. [Example](./hmap/example_test.go)

import: `import "github.com/hujingnb/go-utils/hmap"`

function:

* `IterateByOrder`: iterate map by dict order
* `Equal`: check map is equal

## `hstruct`

dispose struct. [Example](./hstruct/example_test.go)

import: `import "github.com/hujingnb/go-utils/hstruct"`

function:

* `ToMap`: change to map
* `Name`: get struct name

## `hnumber`

dispose number. [Example](./hnumber/example_test.go)

import: `import "github.com/hujingnb/go-utils/hnumber"`

function:

* `HexInput`: read a string number on specify hex.
* `HexOutput`: change a number to specify hex.
* `RandRange`: get a int by range

## `hsystem`

system function. [Example](./hsystem/example_test.go)

import: `import "github.com/hujingnb/go-utils/hsystem"`

function:

* `PrintOutAndErrToFile`: write stdout and stderr to file

## `hcommon`

common function. [Example](./hcommon/example_test.go)

import: `import "github.com/hujingnb/go-utils/hcommon"`

function:

* `Copy`: depth copy for all kind

# PR items

1. please run all test function after change file. command: `go test ./...`
2. all function should has test function
3. all function should has example function in file `example_test.go`
4. all public function should in `README` file