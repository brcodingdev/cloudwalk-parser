package app

// Report contract to print report
type Report interface {
	PrintMatch() (string, error)
}
