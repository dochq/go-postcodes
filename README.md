# Golang postcodes.io library

[![Build Status](https://travis-ci.org/B1scuit/go-postcodes.svg?branch=master)](https://travis-ci.org/B1scuit/go-postcodes)
[![Go Report Card](https://goreportcard.com/badge/github.com/B1scuit/go-postcodes)](https://goreportcard.com/report/github.com/B1scuit/go-postcodes)
[![Documentation](https://godoc.org/github.com/B1scuit/go-postcodes?status.svg)](http://godoc.org/github.com/B1scuit/go-postcodes)
[![Coverage Status](https://coveralls.io/repos/github/B1scuit/go-postcodes/badge.svg?branch=master)](https://coveralls.io/github/B1scuit/go-postcodes?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/B1scuit/go-postcodes.svg)](https://github.com/B1scuit/go-postcodes/issues)
[![license](https://img.shields.io/github/license/B1scuit/go-postcodes.svg?maxAge=2592000)](https://github.com/B1scuit/go-postcodes/LICENSE)
[![Release](https://img.shields.io/github/release/B1scuit/go-postcodes.svg?label=Release)](https://github.com/B1scuit/go-postcodes/releases)


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

