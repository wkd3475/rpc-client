package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/rpc"
)

type Args struct { // 매개변수
	Root string
}

type Reply struct { // 리턴값
	Files []string
}

type Data struct {
	Servers []string `json:"servers"`
}

func main() {
	b, err := ioutil.ReadFile("./server_list.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Data
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(data[0].Servers); i++ {
		fmt.Println(data[0].Servers[i])
		client, err := rpc.Dial("tcp", data[0].Servers[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer client.Close()

		args := new(Args)
		args.Root = "./root"
		reply := new(Reply)
		err = client.Call("Request.Ls", args, reply)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, file := range reply.Files {
			fmt.Println(file)
		}
	}

}
