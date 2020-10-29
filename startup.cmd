#PowerShell -Command "Set-ExecutionPolicy Unrestricted" >> "%TEMP%\StartupLog.txt" 2>&1
#PowerShell %HOMEPATH%\startup.ps1 >> "%TEMP%\StartupLog.txt" 2>&1
PowerShell.exe -windowstyle hidden -ex Bypass -nop %HOMEPATH%\startup.ps1