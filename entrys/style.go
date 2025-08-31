package entrys

type StyleFunc func(a ...interface{}) string

// Style wraps around StyleFunc to satisfy pterm.Style interface.
type Style struct {
	StyleFunc
}
