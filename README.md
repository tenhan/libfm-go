# libFFM-gp
Pure Golang implemented library for FM (factorization machines)

## Usage
```go
package main

import (
	"fmt"
	"github.com/tenhan/libfm-go/models"
)

func main()  {
	fm := models.FM{}
	err := fm.LoadModelFromJsonFile("models/fm_model.json")
	if err != nil {
		panic(err)
	}
	var input = []float64{1,0,1,1,0,1,0,0,0,0,0}
	ret,err := fm.Predict(input)
	fmt.Printf("predict: %f,err: %v\n",ret,err)
}
```

## Command line
```shell
> git clone github.com/tenhan/libfm-go
> cd libfm-go
> go build
> ./libfm-go -fm="./models/fm_model.json" -input="1,0,1,0,1,0,1,0,1,0,1"
0.9169972048780838
```

