package main

import ("fmt"
	"../Mapper/mapper"	
) 

func main() {
	s := mapper.NewSkipString{ 2, "Aspiration.com"}
	mapper.MapString(&s)
	s.Strn=mapper.B.String()
	fmt.Println("\n",s)
 }
 

 

