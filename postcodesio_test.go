package postcodesio

// Example usage:
import (
	"fmt"
	"github.com/DocHQ/go-postcodes"
)

func ExampleLookup() {
	post, err := postcodesio.Lookup("AA1 1AA")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post.Result)
}
