package main

type ShellError struct {
	message string
	err     error
}

func (p *ShellError) Error() string {
	return p.message
}
