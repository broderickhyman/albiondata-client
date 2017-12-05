// +build windows

package systray

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/gonutz/w32"
	"github.com/alexscott/albiondata-client/icon"
)

var consoleHidden bool

func hideConsole() {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}

	consoleHidden = true
}

func showConsole() {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_SHOW)
		}
	}

	consoleHidden = false
}

func Run() {
	systray.Run(onReady, onExit)
}

func onExit() {

}

func onReady() {
	hideConsole()
	systray.SetIcon(icon.Data)
	systray.SetTitle("Albion Data Client")
	systray.SetTooltip("Albion Data Client")
	mConHideShow := systray.AddMenuItem("Show Console", "Show/Hide Console")
	mQuit := systray.AddMenuItem("Quit", "Close the Albion Data Client")

	func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				fmt.Println("Requesting quit")
				systray.Quit()
				os.Exit(0)
				fmt.Println("Finished quitting")

			case <-mConHideShow.ClickedCh:
				if consoleHidden == true {
					showConsole()
					mConHideShow.SetTitle("Hide Console")
				} else {
					hideConsole()
					mConHideShow.SetTitle("Show Console")
				}
			}

		}

	}()
}
