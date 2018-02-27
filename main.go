package main

func main() {

	SendTest()
}

func SendTest() {
	mycontent := " my dear令"

	email := NewEmail("fengshaomin@bjsasc.com",
		"test golang email", mycontent)

	println(email.to)

	if err := email.SendEmail(); err != nil {

		println(err)
	}
}
