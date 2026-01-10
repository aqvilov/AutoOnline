@echo off
chcp 65001 > nul
:: means UTF-8

title AutoOnline
color 0A

echo.
echo #####################################################
echo             ЗАПУСКАЕМ "AutoOnline"
echo #####################################################
echo.

set "EXE_NAME=AutoOnline.exe"

echo Ищу программу %EXE_NAME%...
echo.


if exist "%EXE_NAME%" (
    set "MY_PROGRAM=%~dp0%EXE_NAME%"
    goto found
) else (
    goto not_found
)

:not_found
echo ОШИБКА: Файл %EXE_NAME% не найден!
echo.
pause
exit /b 1

:found
echo Найдено: %MY_PROGRAM%
echo.
echo Запускаю программу...
echo.

::run the program
"%MY_PROGRAM%"

echo.
echo Программа завершена.
echo.
echo Нажмите любую клавишу для выхода...
pause > nul