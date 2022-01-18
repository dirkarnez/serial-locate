serial-locate
=============
A simple tool for searching the COM port

### Usage
- In microcontroller's downloader script
```batch
@echo off
REM say searching for Arduino UNO COM port and pass to avrdude.exe

for /f "tokens=*" %%a in (
    '%~dp0tools\serial-locate.exe --vid=2341 --pid=0043 --usb=true'
) do (
    set comport=%%a
)

REM echo %comport% show "COM4" for example

C:\arduino-1.8.19\hardware\tools\avr\bin\avrdude.exe -C"C:\arduino-1.8.19\hardware\tools\avr\etc\avrdude.conf" -v -v -v -patmega328p -carduino -P\\.\%comport% -b115200 -D -Uflash:w:"application.hex":i

pause
```