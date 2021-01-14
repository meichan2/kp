package main
import "fmt"
func main(){
    var a,b,c,d,e,f,g,j int
    s:= []int{}
    fmt.Scan(&a)
    fmt.Scan(&b)
    for i:=0;i<a;i++{
        fmt.Scan(&c)
        s=append(s,c)
    }
    fmt.Scan(&c)
    for i:=0;i<c;i++{
        fmt.Scan(&d)
        fmt.Scan(&e)
        f=0
        j=0
        for k:=d-1;k<e;k++{
            f=f+s[k]
            j++
        }

        f=f/j

        if f<b{
           g=b-f
        }else{
            g=0
        }
        for k:=d-1;k<e;k++{
            s[k]=s[k]+g
        }

    }

    for i:=0;i<a;i++{
        fmt.Print(s[i])
        if i+1==a{
            fmt.Println("")
        }else{
            fmt.Print(" ")
        }
    }
}
