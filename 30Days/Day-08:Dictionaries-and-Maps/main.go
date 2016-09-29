package main
import "fmt"
import "os"
import "bufio"


func main() {
    var n int
    fmt.Scan(&n)
    var phoneBook = make(map[string]string)
    for i := 0; i < n; i++ {
        var s1 string
        var s2 string
        fmt.Scanf("%s %s", &s1, &s2)
        phoneBook[s1] = s2 
    }
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        s := scanner.Text()
        if phoneBook[s] != "" {
            fmt.Printf("%s=%s\n",s,phoneBook[s])
        }else{
            fmt.Println("Not found")
        }
    }    
}
