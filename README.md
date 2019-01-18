# Golang postcodes.io library

## Usage

```
import (
  "github.com/B1scuit/go-postcodes"
  "fmt"
)

func main(){
  post, err := postcodesio.Lookup("AA1 1AA")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post)
}
```

