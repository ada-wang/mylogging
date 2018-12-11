/*
Copyright github.com/ada-wang    wanggang-info@ruc.edu.cn
All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"github.com/ada-wang/mylogging"
	"github.com/ada-wang/mylogging/example/package1"
	"github.com/ada-wang/mylogging/example/package2"
)

////////////////////////////////////////
// for mylogging
////////////////////////////////////////
const (
	pkgLogID = "main"
	level    = mylogging.DEBUG
)

var logger = mylogging.MustGetLogger(pkgLogID)

func init() {
	mylogging.SetModuleLevel(level, pkgLogID)
}

////////////////////////////////////////
// for mylogging - end
////////////////////////////////////////

// Print for logger example
func Print() {
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("error")
	logger.Critical("critical")
}

func main() {

	// for main
	Print()
	// for package1
	package1.Print()
	// for package2
	package2.Print()

	// you can override package-log-level in main-package
	mylogging.SetModuleLevel(mylogging.NOTICE, "package1")
	mylogging.SetModuleLevel(mylogging.ERROR, "package2")

	// for package1
	package1.Print()
	// for package2
	package2.Print()
}
