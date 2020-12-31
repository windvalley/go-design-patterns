package bridge

func ExampleNewCommonMessageBox() {
	var implementer SendImplementer

	implementer = new(Mail)
	messageBox := NewCommonMessageBox(implementer)
	messageBox.SendMessage("hello bridge")

	implementer = new(Phone)
	messageBox = NewCommonMessageBox(implementer)
	messageBox.SendMessage("hello bridge")

	// Output:
	// send content(hello bridge) by mail
	// send content(hello bridge) by phone
}

func ExampleNewImportantMessageBox() {
	var implementer SendImplementer

	implementer = new(Mail)
	messageBox := NewImportantMessageBox(implementer)
	messageBox.SendMessage("hello bridge")

	implementer = new(Phone)
	messageBox = NewImportantMessageBox(implementer)
	messageBox.SendMessage("hello bridge")

	// Output:
	// send content([Important]hello bridge) by mail
	// send content([Important]hello bridge) by phone
}
