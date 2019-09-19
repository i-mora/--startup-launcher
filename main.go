package main

import rootCMD "./cmd"

func main() {
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
