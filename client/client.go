package main

import (
	"fmt"
	"net/rpc"

	"github.com/chiranSachintha/dc-project-1/common"
)

var (
	client *rpc.Client
)

func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000")

	/*--------------*/

	// create veg variable of type `common.Vegetable
	var veg common.Vegetable

	// get vegetable by name 'Carrot'
	if err := client.Call("VegStore.Get", "Carrot", &veg); err != nil {
		fmt.Println("Error:1 VegStore.Get()", err)
	} else {
		fmt.Printf("Vegetable '%s' found \n", veg)
	}

	/*--------------*/

	// add vegetable by name 'Carrot'
	if err := client.Call("VegStore.Add", common.Vegetable{
		ID:     1,
		Name:   "Carrot",
		Price:  32,
		Amount: 23,
	}, &veg); err != nil {
		fmt.Println("Error:2 VegStore.Add()", err)
	} else {
		fmt.Printf("Vegetable '%s' created \n", veg)
	}

	/*--------------*/

	// update vegetable by name 'Carrot'
	if err := client.Call("VegStore.Update", common.Vegetable{
		ID:     1,
		Name:   "Carrot",
		Price:  10,
		Amount: 10,
	}, &veg); err != nil {
		fmt.Println("Error:2 VegStore.Update()", err)
	} else {
		fmt.Printf("Vegetable '%s' created \n", veg)
	}

	/*--------------*/

	// delete vegetable by name 'Carrot'
	var reply string
	if err := client.Call("VegStore.Delete", "Carrot", &reply); err != nil {
		fmt.Println("Error:1 VegStore.Delete()", err)
	} else {
		fmt.Printf("Vegetable '%s' found and deleted \n", veg.Name)
	}
}
