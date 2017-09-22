@echo off
:: Copyright 2017 Google Inc. All rights reserved.
:: Use of this source code is governed by the Apache 2.0
:: license that can be found in the LICENSE file.
setlocal
set APPENGINE_DEV_APPSERVER=%~dp0\dev_appserver.py
set GOARCH=
set GOBIN=
set GOOS=


:: Note that due to the nature of BAT files, if the optional
:: "--dev_appserver Z:\path\to\dev_appserver.py" arguments are
:: provided, they must appear exactly as the first and second
:: arguments after the command name for this to work properly.
if "%1"=="--dev_appserver" (
    set APPENGINE_DEV_APPSERVER=%~2
)

:: Set a GOPATH if one is not set.
if not "%GOPATH%"=="" goto havepath
set GOPATH=%~dp0\gopath
:havepath

:: Call goapp.py to get the correct executable to call, but call it from here
:: because Windows fork-and-exec doesn't work like linux: it spawns the process
:: in the background instead of replacing the current process.
FOR /F "delims=" %%i IN ('python %~dpn0 --print-command %*') DO set actualcmd=%%i

:: actualcmd looks like "C:\path\to\sdk\goroot-1.x\bin\goapp". GOROOT must be set
:: correctly based on that path.
setlocal ENABLEDELAYEDEXPANSION
set search=\bin\%~n0
FOR /F "tokens=1 delims= " %%i IN ("%actualcmd%") DO set goapproot=%%i
set GOROOT=!goapproot:%search%=!

%actualcmd%
