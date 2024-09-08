@echo off
:: Устанавливаем заголовок окна командной строки
title Go Auto-Reload Server

:: Включаем расширение команд и выключаем эхо
setlocal enabledelayedexpansion

:: Переменная для хранения PID server.exe
set "serverPID="

:: Директория для поиска файлов .go
set "rootDir=C:\Users\Admin\GolandProjects\BizMart"

:: Проверка на существование каталога
if not exist "%rootDir%" (
    echo ERROR: Directory %rootDir% does not exist!
    exit /b
)

:: Метка начала скрипта
:start
echo Building and starting the server...

:: Переходим в каталог проекта
cd /d "%rootDir%"

:: Компилируем Go файл в исполняемый файл server.exe
go build -o server.exe main.go

:: Запускаем server.exe в фоне и получаем его PID
start /B server.exe
for /f "tokens=2" %%i in ('tasklist /fi "imagename eq server.exe" /fo list ^| find "PID:"') do set serverPID=%%i

echo Server started with PID: %serverPID%

:: Считываем текущую дату последней модификации всех файлов с расширением .go во всех поддиректориях
for /f "delims=" %%i in ('forfiles /S /p "%rootDir%" /m *.go /c "cmd /c echo @relpath @fdate @ftime"') do set lastModifiedTime=%%i

:: Переход к метке watch
:watch
:: Задержка в 2 секунды для снижения нагрузки на процессор
timeout /t 2 > nul

:: Переменная для отслеживания изменений во всех .go файлах
set modified=

:: Сравниваем текущую дату последней модификации с предыдущей для всех файлов .go в поддиректориях
for /f "delims=" %%i in ('forfiles /S /p "%rootDir%" /m *.go /c "cmd /c echo @relpath @fdate @ftime"') do (
    if NOT "%%i"=="%lastModifiedTime%" (
        set modified=true
        set lastModifiedTime=%%i
    )
)

:: Если изменение было обнаружено, перезапускаем сервер
if defined modified (
    echo Change detected, restarting server...
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
