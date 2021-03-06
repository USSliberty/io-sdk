package main

import (
	"fmt"
)

func ExampleConfigLoad() {
	run("rm -Rvf /tmp/iosdk-test ; mkdir /tmp/iosdk-test")
	fmt.Println(ConfigLoad())
	DryRunPush("123456")
	Configure("/tmp/iosdk-test/javascript")
	fmt.Println(ConfigLoad())
	fmt.Print(run("ls -a /tmp/iosdk-test/.io*"))
	fmt.Println(Config.IoAPIKey)
	fmt.Println(len(Config.WhiskAPIKey))
	// Output:
	// stat /tmp/iosdk-test/.iosdk.v3: no such file or directory
	// Wrote /tmp/iosdk-test/.iosdk.v3
	// <nil>
	// /tmp/iosdk-test/.iosdk.v3
	// 123456
	// 101
}
