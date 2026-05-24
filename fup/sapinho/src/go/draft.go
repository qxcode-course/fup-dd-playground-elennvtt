package main
import "fmt"
func main() {
    var P,S,E int
    fmt.Scan(&P)
    fmt.Scan(&S)
    fmt.Scan(&E)

    pos := 0 
    for {
        pos2:= pos + S 

        if pos2 >= P{
            fmt.Printf("%d saiu\n", pos)
            break
        }
        fmt.Printf("%d %d\n", pos,pos2)
        pos = pos2 - E
    }
}
