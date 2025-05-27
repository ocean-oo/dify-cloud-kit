package main

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/langgenius/dify-cloud-kit/oss"
	"github.com/langgenius/dify-cloud-kit/oss/factory"
	"github.com/stretchr/testify/assert"
)

var ossArgs = oss.OSSArgs{
	Local: &oss.Local{
		Path: "/tmp/dify-oss-tests",
	},
	S3: &oss.S3{
		UseAws: true,
	},
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func TestAll(t *testing.T) {
	vendors := []string{
		"local",
	}
	key := randomString(10)

	size := 1 * 1024 * 1024
	data := make([]byte, size)

	for _, vendor := range vendors {
		storage, err := factory.Load(vendor, ossArgs)
		if err != nil {
			log.Fatal(err)
			continue
		}

		ossPaths, err := storage.List("/")
		assert.Equal(t, 0, len(ossPaths))
		assert.Nil(t, err)

		exist, err := storage.Exists(key)
		assert.Equal(t, false, exist)
		assert.Nil(t, err)

		err = storage.Save(key, data)
		assert.Nil(t, err)

		rdata, err := storage.Load(key)
		assert.Equal(t, data, rdata)
		assert.Nil(t, err)

		ossState, err := storage.State(key)
		assert.Equal(t, int64(size), ossState.Size)
		assert.Nil(t, err)

		exist, err = storage.Exists(key)
		assert.Equal(t, true, exist)
		assert.Nil(t, err)

		ossPaths, err = storage.List("/")
		assert.Equal(t, 1, len(ossPaths))
		assert.Nil(t, err)

		err = storage.Delete(key)
		assert.Nil(t, err)

		exist, err = storage.Exists(key)
		assert.Equal(t, false, exist)
		assert.Nil(t, err)

		ossPaths, err = storage.List("/")
		assert.Equal(t, 0, len(ossPaths))
		assert.Nil(t, err)
	}
}
