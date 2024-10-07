package _go

import (
	"fmt"
	"os"
)

// main 生成500个函数
func main() {
	file, err := os.Create("generated_functions.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("package main\n\n")
	file.WriteString("import \"fmt\"\n\n")

	for i := 1; i <= 500; i++ {
		funcName := fmt.Sprintf("Function%d", i)
		funcBody := fmt.Sprintf("func %s() {\n\tfmt.Println(\"This is %s\")\n}\n\n", funcName, funcName)
		file.WriteString(funcBody)
	}

	fmt.Println("Generated 500 functions in generated_functions.go")
}
