@ECHO OFF
ECHO Sleep1
ping 127.0.0.1 -n 2 > nul
ECHO Sleep2
ping 127.0.0.1 -n 10 > nul
