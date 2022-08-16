@ECHO OFF
SET Exe=test.exe
SET BIN_WINDOWS = .\build\%Exe%

ECHO run
IF EXIST %BIN_WINDOWS% (
@ECHO ON
%BIN_WINDOWS% --address 127.0.0.1 --port 8081
)