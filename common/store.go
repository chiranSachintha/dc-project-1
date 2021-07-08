package common

import "fmt"

// Vegetable struct represents a vegetable.
type Vegetable struct {
	ID     int
	Name   string
	Price  float32
	Amount float32
}

type VegStore struct {
	database map[string]Vegetable
}

var (
	vegStore *VegStore
)

func (c *VegStore) Add(vegetable Vegetable) error {

	// check if vegetable already exists in the database
	if _, ok := c.database[vegetable.Name]; ok {
		return fmt.Errorf("vegetable with name '%s' already exists", vegetable.Name)
	}

	// add vegetable `p` in the database
	c.database[vegetable.Name] = vegetable
	// c.Commit()

	// return `nil` error
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

	// return `nil` error
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

func GetStore() *VegStore {
	return vegStore
}
