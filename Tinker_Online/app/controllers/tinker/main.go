// Project_Tinker project main.go
package tinker

import (
	//"bufio"
	"fmt"
	//"os"
	//"regexp"
	//"reflect"
	//"strings"
	"encoding/json"
)

/*
	ta := Tree{nil, "A", nil}
	tb := Tree{&ta, "B", nil}
	tj := Tree{nil, "J", nil}
	t1 := Tree{&tb, "G", &tj}
	pprint(&t1)
*/

func main() {
	fmt.Println("= Project Tinker =")
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		expression, _ := reader.ReadString('\n')
		fmt.Println(expression)
	*/
	//postfix := toPostfix("(3+4*5/6)")
	//postfix := toPostfix("1+1")
	//postfix := toPostfix(expression)
	//fmt.Println(postfix)
	//trees := toTree(postfix)
	//fmt.Println(trees.Value)
	//fmt.Println(trees.Left.Value)
	//fmt.Println(trees.Right.Value)
	//fmt.Println(trees.Right.Left.Value)
	//fmt.Println(trees.Right.Right.Value)
	//fmt.Println(trees.Right.Right.Left.Value)
	//fmt.Println(trees.Right.Right.Right.Value)
	//fmt.Println(trees.Right.Value)
	//fmt.Print(trees.Left.Left.Value)
	//fmt.Print(trees.Left.Left.Left.Value)
	//fmt.Print(trees.Right.Value)

	/*
		for g := 0; g < trees.Size; g++ {
			fmt.Println(trees.Size)
			tree := trees.Pop().(Tree)

			if tree.Left != nil {
				//fmt.Println("a" + tree.Left.Value)
			}
			fmt.Println(tree.Value)
			if tree.Right != nil {
				fmt.Println(tree.Right.Value)
			}
		}
	*/
	//fmt.Print("a")
	//t1 := New(26, 1)
	//pprint(t1)
}

//func init() {
//	go g()
//}
type Export struct {

}

func (c Export) Postfix(expr string) string {
	postfix:= toPostfix(expr)
	tree:= toTree(postfix)
	tojson, err := json.Marshal(tree)
	if err != nil {
		return ""
	}
	return string(tojson)
}

func g() {
	fmt.Println("= Project Tinker =")
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		expression, _ := reader.ReadString('\n')
		fmt.Println(expression)
	*/
	//postfix := toPostfix("(3+4*5/6)")
	postfix := toPostfix("1+1")
	//postfix := toPostfix(expression)
	fmt.Println(postfix)
}

