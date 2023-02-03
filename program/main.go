package main


func main(){
	ln := NewHeadList([]int{1,2,3,4,5})
	ln.Print()
	println()

	n := reverseList(ln)
	n.Print()
}
