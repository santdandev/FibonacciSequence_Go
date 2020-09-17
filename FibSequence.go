package main

import "fmt"

func main(){
var n int
fmt.Print("Enter the number of terms to produce : ")
fmt.Scan(&n)
fmt.Print("Fibonacci :: ")
n1:=0
n2:=1
next:=0
for i:=1;i<=n;i++ {
if(i==1){
fmt.Print(n1, " - ")
continue
}
if(i==2){
fmt.Print(n2, " - ")
continue
}
next = n1 + n2
n1=n2
n2=next
fmt.Print(next, " - ")
}
}