package daemon

import "fmt"

var registerOnStart = []func(){
	NewInitGlobalDaemon().Handle,
}

var registerOnStop = []func(){
	func() {
		fmt.Println("server stopped")
	},
}

func RunStartBeforeFn() {
	for _, fn := range registerOnStart {
		fn()
	}
}

func RunStoppedFn() {
	for _, fn := range registerOnStop {
		fn()
	}
}
