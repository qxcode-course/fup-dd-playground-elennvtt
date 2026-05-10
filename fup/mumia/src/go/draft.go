package main
import "fmt"
func main() {
    var nome string
    var num int

    fmt.Scan(&nome)
    fmt.Scan(&num)

    if num < 12 {
        fmt.Println(nome + " " + "eh"+ " " + "crianca")
    }else if num <18 {
        fmt.Println(nome + " " + "eh"+ " " +"jovem")
    }else if num < 65 {
        fmt.Println(nome + " " + "eh"+ " " +"adulto")
    }else if num < 1000 {
        fmt.Println(nome + " " + "eh"+ " " +"idoso")
    }else{
        fmt.Println(nome + " " + "eh"+ " " +"mumia")
    }
}
