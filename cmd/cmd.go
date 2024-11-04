package cmd

type ICmd interface {
	Perform() error
}
