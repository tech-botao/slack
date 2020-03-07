# slack

> 简单的接口实现， slack通知功能


## install

```bash
go get -t github.com/tech-botao/slack
```


## configure

```bash
# 利用环境变量的时候
SLACK_TOKEN=XXXXXXXXXXXX

# token在程序中写入
slack.SetToken() 
```


## usage

```golang

import github.com/tech-botao/slack

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
```

