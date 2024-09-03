@echo off
:: Устанавливаем заголовок окна командной строки
title Go Auto-Reload Server

:: Включаем расширение команд и выключаем эхо
setlocal enabledelayedexpansion

:: Переменная для хранения PID server.exe
set "serverPID="

:: Метка начала скрипта
:start
echo Building and starting the server...

:: Компилируем Go файл в исполняемый файл server.exe
go build -o server.exe main.go

:: Запускаем server.exe в фоне и получаем его PID
start /B server.exe
for /f "tokens=2" %%i in ('tasklist /fi "imagename eq server.exe" /fo list ^| find "PID:"') do set serverPID=%%i

echo Server started with PID: %serverPID%

:: Считываем текущую дату последней модификации всех файлов с расширением .go в переменную lastModifiedTime
for /f "delims=" %%i in ('forfiles /m *.go /c "cmd /c echo @fdate @ftime"') do set lastModifiedTime=%%i

:: Переход к метке watch
:watch
:: Задержка в 2 секунды для снижения нагрузки на процессор
timeout /t 2 > nul

:: Переменная для отслеживания изменений во всех .go файлах
set modified=

:: Сравниваем текущую дату последней модификации с предыдущей
for /f "delims=" %%i in ('forfiles /m *.go /c "cmd /c echo @fdate @ftime"') do (
    if NOT "%%i"=="%lastModifiedTime%" (
        set modified=true
        set lastModifiedTime=%%i
    )
)

:: Если изменение было обнаружено, перезапускаем сервер
if defined modified (
    taskkill /f /pid %serverPID% >nul 2>&1
    goto start
)

:: Возвращаемся к метке watch
goto watch

:cleanup
:: Убиваем процесс server.exe при выходе из BAT файла
if not "%serverPID%"=="" (
    taskkill /f /pid %serverPID% >nul 2>&1
)
exit /b
