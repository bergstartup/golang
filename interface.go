package main
import "fmt"

type vector []int
type Int int

type math interface{
  Square() Int
}

func (a Int) Square() Int{
  return a*a
}

func (v vector) Square() Int{
  a:=0
  for _,j:= range v{
    a+=j*j
  }
  return Int(a)
}

func main(){
  var i_int,j_vector math
  i_int=Int(5)
  j_vector=vector{1,2,3} 
  fmt.Println(i_int.Square(),j_vector.Square())
}
