package main

import (
	appcofig "github.com/takumi-kawaguchi/database_manager/appconfig"
)

func main() {
	appcofig.Init()
	// err := appcofig.Init()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
