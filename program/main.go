package main

import "log"

func main(){
	res := make([][]int, 0)
	for i := 0; i < 3; i++ {
		data := []int{i}

		datas := [][]int{data}
		res = append(datas, res...)
	}
	log.Printf("%+v\n", res)
}
