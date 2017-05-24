// Go's concurrency primitives make it easy to
// express concurrent concepts, such as
// this binary tree comparison.
//
// Trees may be of different shapes,
// but have the same contents. For example:
//
//        4               6
//      2   6          4     7
//     1 3 5 7       2   5
//                  1 3
//
// This program compares a pair of trees by
// walking each in its own goroutine,
// sending their contents through a channel
// to a third goroutine that compares them.

package tinker

import (
	//"fmt"
	"math/rand"
	//"reflect"
	//"strconv"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan result, lvl int) {
	if t == nil {
		return
	}
	lvl++
	Walk(t.Left, ch, lvl)
	ch <- result{t.Value, lvl}
	Walk(t.Right, ch, lvl)
}

type result struct {
	val string
	lvl int
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan result {
	ch := make(chan result)
	go func() {
		Walk(t, ch, 0)
		close(ch)
	}()
	return ch
}

// Compare reads values from two Walkers
// that run simultaneously, and returns true
// if t1 and t2 have the same contents.
/*func Compare(t1, t2 *Tree) bool {
	c1, c1lvl, c2, c2lvl := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}*/

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for g, _ := range rand.Perm(n) { //
		t = t.insert(toChar(g)) // (v+1)*k)
	}
	return t
}

func toChar(i int) string {
	return string('A' + i)
}

func toNum(i string) int {
	return int([]rune(i)[0])
}

func (t *Tree) insert(v string) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	//fmt.Print(v)
	//fmt.Print(t.Value)
	//fmt.Println()
	if toNum(v) < toNum(t.Value) {
		//fmt.Print("left")
		t.Left = t.Left.insert(v)
		return t
	}
	t.Right = t.Right.insert(v)
	//fmt.Print("right")
	return t
}

//Function to get width of a given level
func (t *Tree) width(lvl int) int {
	if t == nil {
		return 0
	} else if lvl == 1 {
		return 1
	} else if lvl > 1 {
		return t.Left.width(lvl-1) + t.Right.width(lvl-1)
	}
	return lvl
}

//Gets height of tree
func (t *Tree) height() int {
	if t == nil {
		return 0
	} else {
		left := t.Left.height()
		right := t.Right.height()

		if left < right {
			return right + 1
		} else {
			return left + 1
		}
	}
}

func (t *Tree) pprint() string {

	tree_height := t.height()
	var result string
	ch := Walker(t)
	for c := range ch {
		result = result + ppspaces(&c.lvl, &tree_height)
		//fmt.Print(c.lvl)
		result = result + c.val
	}
	return result
}


func ppspaces(lvl *int, tree_height *int) string {
	symbol := "-"
	space := ""
	for i := *tree_height; i > *tree_height-*lvl; i-- {
		space += symbol
	}
	return space
}



