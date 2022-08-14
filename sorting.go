package main
import (
	"fmt"
	"sort"
)

type loom struct{
	name string
	age int

}

type Looms []loom

func (looms Looms) Len() int {
	return len(looms)
}
func (l Looms)Swap(i,j int){
	l[i],l[j] = l[j],l[i]
}

func (l Looms)Less (i,j int)bool{
	return l[i].age > l[j].age
}

func main(){
	origin := make([]loom,5)
	origin[0] = loom{"ilman",32}
	origin[1]= loom{"howie",40}
	origin[2]= loom{"maddie",56}
	origin[3]=loom{"gloomie",88}
	origin[4]=loom{"jhonie",21}

	fmt.Printf("before capacity:  %d and length %d\n",cap(origin),len(origin))

	origin = append(origin,loom{"maggie",30},loom{"stevie",37})
	
	fmt.Printf("after cap %d and length %d\n",cap(origin),len(origin))

	origin = append(origin,loom{"gerrie",30},loom{"dowwy",67})
	
	fmt.Printf("after 2 cap %d and length %d\n",cap(origin),len(origin))
	fmt.Printf("origin: %v \n",origin)

	sort.Sort(Looms(origin))
	fmt.Printf("origin: %v \n",origin)

}
