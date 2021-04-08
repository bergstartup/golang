package main

import "golang.org/x/tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left==nil{
		ch<-t.Value
		return
	}
	Walk(t.Left,ch)

	ch<-t.Value

	if t.Right==nil{
		return
	}

	Walk(t.Right,ch)
	return
}

func Same(t1, t2 *tree.Tree) bool{
	ch1:=make(chan int)
	go func(){
	 defer close(ch1)
	 Walk(t1,ch1)
	}()

	ch2:=make(chan int)
	go func(){
	 defer close(ch2)
	 Walk(t2,ch2)
	}()

	for a:=range ch1{
		fmt.Print(a," ")
		b:=<-ch2
		fmt.Println(b)
		if a!=b{
				return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	ch:=make(chan int)
	t:=tree.New(1)
	go func(){
	 defer close(ch)
	 Walk(t,ch)
	}()
	for a:= range ch{
		fmt.Println(a)
	}

	ch2:=make(chan int)
	t2:=tree.New(1)
	go func(){
	 defer close(ch2)
	 Walk(t2,ch2)
	}()
	for b:= range ch2{
		fmt.Println(b)
	}
}
