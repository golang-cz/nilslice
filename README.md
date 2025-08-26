# nilslice.Initialize(obj) - fix `null` arrays rendered by encoding/json

Golang's `json.Marshal()` function has a very unfortunate behavior when encoding `nil` slices. It renders a `null` value in JSON instead of an empty array `[]`.

This issue has a couple of open issues (ie. https://github.com/golang/go/issues/27589) but none of the proposed solutions were accepted into the standard library as of Aug 2023.

This lightweight Go package will help you mitigate the issue by recursively initializing all `nil` slices in a given object using a `reflect` package.

```go
type Payload struct {
	Items []Item `json:"items"`
}

payload := &Payload{}

b, _ := json.Marshal(payload)
fmt.Println(string(b))
// {"items": null}

b, _ = json.Marshal(nilslice.Initialize(payload))
fmt.Println(string(b))
// {"items": []}
```

## Fixed in encoding/json/v2 or bytedance/sonic

The original issue https://github.com/golang/go/issues/27589 has been addressed in both `encoding/json/v2` and `github.com/bytedance/sonic` packages.

## Install
```shell
$ go get github.com/golang-cz/nilslice
```

```go
import "github.com/golang-cz/nilslice"
```

## Authors
- [golang.cz](https://golang.cz)

## License
[MIT license](./LICENSE)
