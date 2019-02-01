package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {

	hello := &HelloWorld{Message: "Hello World"}
	write(hello)

}

func write(hello *HelloWorld) error {

	fmt.Println(proto.MarshalTextString(hello))

	b, err := proto.Marshal(hello)
	if err != nil {
		return fmt.Errorf("could not encode hello %v", err)
	}

	f, err := os.OpenFile("db.blob", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open db %v", err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write to %v", err)
	}

	b, err = ioutil.ReadFile("db.blob")
	if err != nil {
		return fmt.Errorf("could not read file %v", err)
	}

	newWorld := &HelloWorld{}
	proto.Unmarshal(b, newWorld)

	fmt.Println(newWorld.Message)

	return nil
}
