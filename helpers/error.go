package helpers

func PanicIfErr(err any) {
	if err != nil {
		panic(err)
	}
}
