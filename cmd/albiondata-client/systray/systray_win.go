// +build windows

package systray

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/regner/albiondata-client/client"
	"github.com/regner/albiondata-client/icon"
)

var myRUNNER func()

func Run(runner func()) {
	myRUNNER = runner

	systray.Run(onReady, onExit)
}

func onExit() {

}

func onReady() {
	client.HideConsole()
	systray.SetIcon(icon.Data)
	systray.SetTitle("Albion Data Client")
	systray.SetTooltip("Albion Data Client")
	mConHideShow := systray.AddMenuItem("Show Console", "Show/Hide Console")
	if !client.CanHideConsole {
		mConHideShow.Disable()
	}
	mQuit := systray.AddMenuItem("Quit", "Close the Albion Data Client")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				fmt.Println("Requesting quit")
				systray.Quit()
				fmt.Println("Finished quitting")

			case <-mConHideShow.ClickedCh:
				if client.ConsoleHidden == true {
					client.ShowConsole()
					client.ConsoleHidden = false
					mConHideShow.SetTitle("Hide Console")
				} else {
					client.HideConsole()
					client.ConsoleHidden = true
					mConHideShow.SetTitle("Show Console")
				}
			}

		}

	}()

	myRUNNER()
}
