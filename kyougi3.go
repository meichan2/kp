package main
import "fmt"
func main(){
    var a,b,A[],c,d,e,f int
    fmt.Scan(&a)
    for i:=0;i<a;i++{
        fmt.Scan(&b)
        A=append(A,b)
    }
    fmt.Scan(&c)
    fmt.Scan(&d)
    for i:=0;i<d;i++{
        fmt.Scan(&e)
        fmt.Scan(&f)
        e=e-1
        c=c-A[e]*f
        if c<0{
            c=c+A[e]*f
        }
    }
    fmt.Println(c)
}
