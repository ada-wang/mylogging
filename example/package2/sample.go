/*
Copyright github.com/ada-wang    wanggang-info@ruc.edu.cn
All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package package2

import "github.com/ada-wang/mylogging"

////////////////////////////////////////
// for mylogging
////////////////////////////////////////
const (
	pkgLogID = "package2"
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
