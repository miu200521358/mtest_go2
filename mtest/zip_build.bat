@echo off
set date00=%date: =0%
set yyyy=%date00:~0,4%
set mm=%date00:~5,2%
set dd=%date00:~8,2%

set time00=%time: =0%
set hh=%time00:~0,2%
set mn=%time00:~3,2%
set ss=%time00:~6,2%

@echo on
set filename=mtest_1.00.00_beta_%yyyy%%mm%%dd%_%hh%%mn%%ss%.exe
@REM wails build -trimpath -clean -o %filename%
wails build -trimpath -o %filename%
