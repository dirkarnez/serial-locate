serial-locate
=============
A simple tool for searching the COM port

### Usage
- Direct query
```batch
serial-locate.exe --vid=2341 --pid=0043 --usb=true
```
- In microcontroller's downloader (e.g. `avrdude.exe`) script
```batch
@echo off
REM say searching for Arduino UNO COM port and pass to avrdude.exe

for /f "tokens=*" %%a in (
    'serial-locate.exe --vid=2341 --pid=0043 --usb=true'
) do (
    set comport=%%a
)

REM echo %comport% show "COM4" for example

C:\arduino-1.8.19\hardware\tools\avr\bin\avrdude.exe -C"C:\arduino-1.8.19\hardware\tools\avr\etc\avrdude.conf" -v -v -v -patmega328p -carduino -P\\.\%comport% -b115200 -D -Uflash:w:"application.hex":i

pause
```

### TODOs
- [ ] cheatsheet in `--help`
    - https://github.com/BrewPi/brewpi-script/blob/develop/autoSerial.py 
