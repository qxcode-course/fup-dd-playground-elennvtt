 package main
import "fmt"
func main() {
    var M, A, B int
    fmt.Scan(&M, &A, &B)

    C := M - (A + B)
 

    if C > A && C > B{
        fmt.Println(C)
    }else if A > B  && A > C {
        fmt.Println(A)
    }else{
        fmt.Println(B)
    }
}
