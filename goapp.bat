@echo off
:: Copyright 2017 Google Inc. All rights reserved.
:: Use of this source code is governed by the Apache 2.0
:: license that can be found in the LICENSE file.

:: get the path to the current file and strip the file extension, which is the name
:: of the python file which wraps the go tools
python %~dpn0 %*
