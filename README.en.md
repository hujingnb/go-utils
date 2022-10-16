# go-utils

some common util method collection

if you has a question, but don't has. welcome commit `issue` or `PR`. please commit `PR` to `dev` branch.

install: `go get github.com/hujingnb/go-utils`

you can see document on [godoc](https://pkg.go.dev/github.com/hujingnb/go-utils).

# package 

* `hstring`: [dispose string](./hstring/README.en.md)
* `hhash`: [hash](./hhash/README.en.md)
* `harray`: [dispose array](./harray/README.en.md)
* `hmap`: [dispose map](./hmap/README.en.md)
* `hstruct`: [dispose struct](./hstruct/README.en.md)
* `hnumber`: [dispose number](./hnumber/README.en.md)
* `hsystem`: [system](./hsystem/README.md)
* `hcommon`: [common function (eg: depth copy)](./hcommon/README.en.md)

# PR items

1. please run all test function after change file. command: `go test ./...`
2. all function should has test function
3. all function should has example function in file `example_test.go`
4. all public function should in `README` file