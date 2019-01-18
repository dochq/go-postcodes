# Golang postcodes.io library

[![Build Status](https://travis-ci.org/DocHQ/go-postcodes.svg?branch=master)](https://travis-ci.org/DocHQ/go-postcodes)
[![Go Report Card](https://goreportcard.com/badge/github.com/DocHQ/go-postcodes)](https://goreportcard.com/report/github.com/DocHQ/go-postcodes)
[![Documentation](https://godoc.org/github.com/DocHQ/go-postcodes?status.svg)](http://godoc.org/github.com/DocHQ/go-postcodes)
[![Coverage Status](https://coveralls.io/repos/github/DocHQ/go-postcodes/badge.svg?branch=master)](https://coveralls.io/github/DocHQ/go-postcodes?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/DocHQ/go-postcodes.svg)](https://github.com/DocHQ/go-postcodes/issues)
[![license](https://img.shields.io/github/license/DocHQ/go-postcodes.svg?maxAge=2592000)](https://github.com/DocHQ/go-postcodes/LICENSE)
[![Release](https://img.shields.io/github/release/DocHQ/go-postcodes.svg?label=Release)](https://github.com/DocHQ/go-postcodes/releases)

## Usage

```
import (
  "github.com/DocHQ/go-postcodes"
  "fmt"
)

func main(){
  post, err := postcodesio.Lookup("AA1 1AA")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post.Result)
}
```

