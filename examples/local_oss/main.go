package main

import (
	"log"

	"github.com/langgenius/dify-cloud-kit/oss"
	"github.com/langgenius/dify-cloud-kit/oss/factory"
)

func main() {
	s, err := factory.Load("local", oss.OSSArgs{
		Local: &oss.Local{
			Path: "/sdad1/13ad13",
		},
	})
	if err != nil {
		log.Panic(err)
	}
	s.List("/test1")
}
