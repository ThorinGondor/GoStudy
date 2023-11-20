package dao

import "fmt"

func Select(id int) bool {
	fmt.Printf("select * from dual where id = %d\n", id)
	return true
}
