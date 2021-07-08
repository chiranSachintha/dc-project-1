package common

import (
	"fmt"
)

// Vegetable struct represents a vegetable.
type Vegetable struct {
	ID     int
	Name   string
	Price  float32
	Amount float32
}

func (veg Vegetable) GetPrice() float32 {
	return veg.Price
}

var (
	store *VegStore
)

type VegStore struct {
	database map[string]Vegetable
}

func (c *VegStore) Add(vegetable Vegetable, reply *Vegetable) error {

	// check if vegetable already exists in the database
	if _, ok := c.database[vegetable.Name]; ok {
		return fmt.Errorf("vegetable with name '%s' already exists", vegetable.Name)
	}

	// add vegetable `p` in the database
	c.database[vegetable.Name] = vegetable
	*reply = vegetable

	return nil
}

func (c *VegStore) Get(name string, reply *Vegetable) error {

	// get vegetable with name `p` from the database
	result, ok := c.database[name]

	// check if vegetable exists in the database
	if !ok {
		return fmt.Errorf("vegetable with name '%s' does not exist", name)
	}

	*reply = result

	return nil
}

func (c *VegStore) GetAllVeg(args string, reply *[]Vegetable) error {

	var arr []Vegetable

	for _, element := range c.database {
		arr = append(arr, element)
	}

	*reply = arr

	return nil
}

func (c *VegStore) Update(vegetable Vegetable, reply *Vegetable) error {

	_, ok := c.database[vegetable.Name]

	if !ok {
		return fmt.Errorf("vegetable with id '%s' does not exist", vegetable.Name)
	}

	c.database[vegetable.Name] = vegetable
	// c.Write()
	*reply = vegetable

	return nil
}

func (c *VegStore) Delete(name string, reply *string) error {

	// get vegetable with id `p` from the database
	_, ok := c.database[name]

	// check if vegetable exists in the database
	if !ok {
		return fmt.Errorf("vegetable with name '%s' does not exist", name)
	} else {
		delete(c.database, name)
		// c.Write()
		*reply = "deleted successfully."
	}

	return nil
}

func GetStore() *VegStore {
	return &VegStore{
		database: make(map[string]Vegetable),
	}
}

// to read from the file
// func (c *VegStore) Read() {
// 	file, _ := ioutil.ReadFile("veg-db.json")
// 	json.Unmarshal(file, &vegStore.database)
// }

// // write to the file
// func (c *VegStore) Write() {
// 	jsonString, _ := json.Marshal(vegStore.database)
// 	ioutil.WriteFile("veg-db.json", jsonString, os.ModePerm)
// }
