package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    fmt.Print("[ ")
    first := true

    for i := 0; i <= 10; i++{
        if i == N{
            continue
        }
        if !first {
            fmt.Print(" ")
        }
        if i == 10 {
            fmt.Print("ceu")
        }else{
            fmt.Print(i)
        }
        first=false
            
        }
        fmt.Print(" ]")
    }
    
