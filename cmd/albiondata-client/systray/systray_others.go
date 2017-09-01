// +build linux darwin

package systray

func Run(runner func()) {
	runner()
}
