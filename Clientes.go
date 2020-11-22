package main

import (
	"fmt"
	"net"	
	"time"
	"encoding/gob"
)


func clientes() {

	var items [2]uint64
	c, err := net.Dial("tcp", ":9999")

	if err != nil {

		fmt.Println(err)
		return
	}
	
	err = gob.NewDecoder(c).Decode(&items)
	
	if err != nil {
	
		fmt.Println(err)
	
		} else {
		
		
		tunel := make(chan uint64)
		go proceso_cli(items[0], items[1], tunel)
		
		for {
		
			items[1] = <- tunel
			err := gob.NewEncoder(c).Encode(items[1])
		
			if err != nil {
		
				fmt.Println(err)
				return
		
			} 
		}
	}
}


func proceso_cli(id uint64, i uint64, tunel chan uint64) {
	
	for {
		
		fmt.Println(id, " -- ", i)
		i += 1
		
		tunel <- i
		time.Sleep(time.Millisecond * 500)
	}
}


func main() {

	go clientes()
	
	
	var input string
	fmt.Scanln(&input)
}