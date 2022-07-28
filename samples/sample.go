package samples

import "io"

func NoArgNoReturn() {
}

func NoArgOneReturn() int {
	return 0
}

func NoArgTwoReturn() (int, string) {
	return 0, ""
}

func NoArgTwoAndErrReturn() (int, string, error) {
	return 0, "", nil
}

func NoArgErrReturn() error {
	return nil
}

func OneArgNoReturn(x int) {
}

func ReturnWriter() io.Writer {
	return nil
}

type A struct{}

func (a A) NoArgNoReturn() {
}

func (a A) NoArgOneReturn() int {
	return 0
}

func (a A) NoArgTwoReturn() (int, string) {
	return 0, ""
}

func (a A) NoArgTwoAndErrReturn() (int, string, error) {
	return 0, "", nil
}

func (a A) NoArgErrReturn() error {
	return nil
}

func (a A) OneArgNoReturn(x int) {
}
