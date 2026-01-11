@echo off
chcp 65001 > nul
title AutoOnline
color 0A

:menu
cls
echo.
echo      AutoOnline
echo.
echo  1. Запустить программу
echo  2. Запустить в фоне
echo  3. Выход
echo.
set /p c="Выбор: "

if "%c%"=="1" goto run
if "%c%"=="2" goto bg
if "%c%"=="3" exit
goto menu

:run

:: clear screen
cls
echo.
echo #####################################################
echo             ЗАПУСК "AutoOnline"
echo #####################################################
echo.

set "EXE_NAME=AutoOnline.exe"
echo Ищу программу %EXE_NAME%...
echo.

if exist "%EXE_NAME%" (
    set "MY_PROGRAM=%~dp0%EXE_NAME%"
    goto found_run
) else (
    goto not_found
)

:found_run
echo Найдено: %MY_PROGRAM%
echo.
echo Запускаю программу...
echo.


::run
"%MY_PROGRAM%"

echo.
echo Программа завершена.
echo.
pause
goto menu

:bg
cls
echo.
echo #####################################################
echo         ЗАПУСК В ФОНОВОМ РЕЖИМЕ
echo #####################################################
echo.

set "EXE_NAME=AutoOnline.exe"
echo Ищу программу %EXE_NAME%...
echo.


if exist "%EXE_NAME%" (
    set "MY_PROGRAM=%~dp0%EXE_NAME%"
    goto found_bg
) else (
    goto not_found
)

:found_bg
echo Найдено: %MY_PROGRAM%
echo.
echo Запускаю в фоновом режиме...
echo.


:: /b -- background mode
start /B "" "%MY_PROGRAM%"

echo Программа запущена в фоне.
echo Для остановки откройте Диспетчер задач.
echo.
pause
goto menu

:not_found
echo ОШИБКА: Файл %EXE_NAME% не найден!
echo.
echo Убедитесь что:
echo 1. Программа скомпилирована
echo 2. Файл находится в папке: %cd%
echo.
pause
goto menu