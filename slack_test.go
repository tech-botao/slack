package slack

import "fmt"

func Example_Send() {

	builder := NewBuilder().
		Pretext("hello").
		Text("Hello Text").
		Color("red").
		AddField("1", "one")

	err := builder.Send()
	if err != nil {
		fmt.Println(err)
	}

	// output:

}

func Example_SendError() {

	err := fmt.Errorf("this is err")
	SendError(err)

	// output:

}
