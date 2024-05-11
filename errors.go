package main

type ShellError struct {
	message string
	err     error
}

func (s *ShellError) Error() string {
	return s.message
}
