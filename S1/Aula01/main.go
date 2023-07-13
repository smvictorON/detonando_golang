package main

import (
	"fmt"
)

type Client struct {
	Name     string
	LastName string
}

func (c Client) FullName() string {
	return fmt.Sprintf("%s %s", c.Name, c.LastName)
}

func main() {
	client := Client{"Victor", "Mussio"}
	fmt.Println(client.FullName())
}
