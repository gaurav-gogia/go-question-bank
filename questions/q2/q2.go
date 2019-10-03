/*
	Complete a function named Fact, the function returns factorial of a number

	Go Documentation: https://godoc.org
*/

package q2
import "fmt"
 
var factVal int = 1 
var i int = 1
var n int
 
// Fact function, please do NOT change it's name :)
func Fact(num int) int {
	// Your code goes here
	if(n >0){        
        for i:=1; i<=n; i++ {
            factVal *= int(i)  
        }
         
    }    
    return factVal  
}

func main(){    
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",Fact(n))
}

