/*
Copyright © 2023 Vishal Tewatia <tewatiavishal3@gmail.com>
*/
package main

import "github.com/vishu42/gha-docs/cmd"

func main() {
	cg := cmd.CommandGroup{}
	cg.All().Execute()
}
