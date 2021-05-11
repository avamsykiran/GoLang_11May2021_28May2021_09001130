package main

import "fmt"

type pen struct {
	nib float32
	inkColor string
}

func (p pen) display(){
	fmt.Println("Nib: ",p.nib,"\tInkColor: ",p.inkColor)
}

func main(){
	p1 := pen{0.4,"blue"}

	p1.display();
}
