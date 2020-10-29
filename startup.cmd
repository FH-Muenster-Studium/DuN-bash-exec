#PowerShell -Command "Set-ExecutionPolicy Unrestricted" >> "%TEMP%\StartupLog.txt" 2>&1
#PowerShell %HOMEPATH%\startup.ps1 >> "%TEMP%\StartupLog.txt" 2>&1
PowerShell.exe -windowstyle hidden %HOMEPATH%\startup.ps1 && exit