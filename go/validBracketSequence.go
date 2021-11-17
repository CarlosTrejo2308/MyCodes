package main

func validBracketSaquence(sequence string) bool {
	stack := []rune{}
	for _, char := range sequence {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == '(' && char != ')' || top == '[' && char != ']' || top == '{' && char != '}' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {

}
