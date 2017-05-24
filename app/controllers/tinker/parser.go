package tinker

type JSONTree struct {
	Value    string
	Name     string
	Children []JSONTree
}

func toTree(postfix string) JSONTree {
	var stack = NewStack()
	var rightOperand JSONTree
	var leftOperand JSONTree
	var trees []JSONTree
	var t JSONTree
	element := ""
	for _, ele := range postfix {
		element = string(ele)
		if IsOperator(element) == true {
			rightOperand = stack.Pop().(JSONTree)
			rightOperand.Name = "Right"
			leftOperand = stack.Pop().(JSONTree)
			leftOperand.Name = "Left"
			trees = append(trees, rightOperand)
			trees = append(trees, leftOperand)
			stack.Push(JSONTree{Value: element, Name: "Head", Children: []JSONTree{trees[len(trees)-2], trees[len(trees)-1]}})
		} else {
			t = JSONTree{element, "", nil}
			stack.Push(t)
		}
	}
	t = stack.Pop().(JSONTree)
	return t
}

/*
//Convert reverse polish to expression tree
func toTree(postfix string) Tree {
	var stack = NewStack()
	var rightOperand Tree
	var leftOperand Tree
	var trees []Tree
	var t Tree
	element := ""
	for _, ele := range postfix {
		element = string(ele)
		if IsOperator(element) == true {
			rightOperand = stack.Pop().(Tree)
			leftOperand = stack.Pop().(Tree)
			trees = append(trees, rightOperand)
			trees = append(trees, leftOperand)
			stack.Push(Tree{&trees[len(trees)-2], element, &trees[len(trees)-1]})
		} else {
			t = Tree{nil, element, nil}
			stack.Push(t)
		}
	}
	t = stack.Pop().(Tree)
	return t
}
*/

//reverse polish notation -> 2 + 3 -> 2 3 +
func toPostfix(expression string) string {
	//for final result
	pfix := ""
	//to pop into
	tptoken := ""
	//for non finalised tokens
	s := NewStack()
	//get rid of white space
	//expression = ReplaceSpace(expression)
	var str string

	//check if num -> add to final
	//in brackets add to stack, loop through, add contents to final result
	//if operator, go through stack and sort by operation -> output biggest first
	//loop through stack to pop any remaining tokens
	for _, expr := range expression {
		str = string(expr)
		if isalnum(str) {

			pfix += str
		} else {
			if str == "(" {
				s.Push(str)
			} else {
				if str == ")" {
					tptoken = s.Pop().(string)
					for tptoken != "(" {
						pfix += tptoken
						tptoken = s.Pop().(string)
					}
				} else {
					if s.Size == 0 {
						s.Push(str)
					} else {
						head := s.head.data.(string)
						for priority(str) <= priority(head) && head != "" {
							tptoken = s.Pop().(string)
							pfix += tptoken
							if s.Size == 0 {
								head = ""
							} else {
								head = s.head.data.(string)
							}
						}
						s.Push(str)
					}
				}
			}
		}
	}

	for s.Size != 0 {
		tptoken = s.Pop().(string)
		pfix += tptoken
	}
	return pfix
}

//Use unicode match to see if any whitespace in string, if so get rid of it
func ReplaceSpace(s string) string {
	var result []rune
	const badSpace = '\u0020'
	for _, r := range s {
		if r == badSpace {
			result = append(result, '\u00A0')
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func IsOperator(s string) bool {
	switch s {
	case "*", "/", "^", "+", "-":
		return true
	default:
		return false
	}
}

//BIMDAS
func priority(x string) int {
	if x == "(" {
		return 0
	}
	if x == "+" || x == "-" {
		return 1
	}
	if x == "*" || x == "/" || x == "%" {
		return 2
	}
	return 3
}

//is alpha numeric -> similar to c
//ie. is operand
func isalnum(x string) bool {
	c := int([]rune(x)[0])
	if (c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}
