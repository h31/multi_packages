package package_3

type Stack struct {
	items   []int
	pointer int
}

func InitializeStack() *Stack {
	return &Stack{
		items:   make([]int, 20),
		pointer: -1,
	}
}

func (st *Stack) Push(el int) {
	if st.pointer+1 >= len(st.items) {
		newStack := make([]int, len(st.items)*2)
		newStack = append(st.items, newStack...)
		st.items = newStack
	}
	st.pointer += 1
	st.items[st.pointer] = el
}

func (st *Stack) Pop() int {
	if st.pointer >= 0 {
		res := st.items[st.pointer]
		st.pointer -= 1
		return res
	}

	return -1
}
