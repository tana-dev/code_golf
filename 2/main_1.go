package main
import("fmt"
"os"
"strconv")
func main() {
a, _ := strconv.Atoi(os.Args[1])
for i := 2; i <= a; i++ {
for j := 2; j < i; j++ {
if 0 == i%j {
break
} else if j == (i - 1) {
// k := strconv.Itoa(i)
if i != a { fmt.Print(i)
     fmt.Print(",")} else { fmt.Print(i) } } } } }
