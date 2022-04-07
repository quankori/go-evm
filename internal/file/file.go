package file

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func ReadFile() int64 {
	data, err := ioutil.ReadFile("last_block_sync.txt")
	if err != nil {
		fmt.Println(err)
	}
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func WriteFile(data string) {
	mydata := []byte(data)

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("last_block_sync.txt", mydata, 0777)
	// handle this error
	if err != nil {
		fmt.Println(err)
	}
}
