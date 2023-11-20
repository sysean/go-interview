package main

func eval() {
	var (
		op    int
		cycle int
	)

	for {
		cycle++

		op = vm.Text[PC]
		PC++

		switch op {
		case IMM:
			AX = vm.Text[PC]
			PC++
		case LC:
			AX = vm.Data[AX]
		case LI:
			AX = vm.Text[AX]
		case SI:
			vm.Stack[SP] = AX
			SP++
		case ENT:
			SP--
			vm.Stack[SP] = BP
			BP = SP
			SP -= vm.Text[PC]
		case LEA:
			AX = BP + vm.Text[PC]
		case PUSH:
			SP--
			vm.Stack[SP] = AX
		case ADD:
			AX += vm.Stack[SP]
			SP++
		case EXIT:
			return
		}
	}
}
