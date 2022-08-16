package go_debug_tool

func ExamplePrintTypeAndValue() {
  stringVar := "test string"
  PrintTypeAndValue("stringVar", stringVar)

  // Output:
  // [stringVar] 型: string, 値: test string
}