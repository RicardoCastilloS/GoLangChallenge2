package mapper

import (
	"unicode"
   "bytes"
   "strings"
		
) 
var globalCounter = 0
var B bytes.Buffer   //I'd remove those global variables

type Interface interface {
   TransformRune(pos int)
   GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
   for pos, _ := range i.GetValueAsRuneSlice() {
      i.TransformRune(pos)
   }
   
}

type NewSkipString struct{
	Skip int
	Strn string
 }


 func  (r *NewSkipString) TransformRune(i int){
	s := strings.Split(r.Strn,"")   
		if(IsDigit(s[i]) ||IsLetter(s[i])){
		if (i+1-globalCounter)%3==0&&i!=0{         
			B.WriteString(strings.ToUpper(s[i]))
		} else {
         B.WriteString(strings.ToLower(s[i]))
		}
	} else{
      globalCounter ++
		B.WriteString(s[i])
   }
}
func (r NewSkipString) GetValueAsRuneSlice() []rune{
   b := []rune(r.Strn)
   return b
}

func IsLetter(s string) bool {
   for _, r := range s {
       if !unicode.IsLetter(r) {
           return false
       }
   }
   return true
}
func IsDigit(s string) bool {
   for _, r := range s {
       if !unicode.IsDigit(r) {
           return false
       }
   }
   return true
}




