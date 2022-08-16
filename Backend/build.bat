@ECHO OFF
SET Exe=test.exe
SET BIN_WINDOWS=.\build\%Exe%
SET CMD_SRC=.\main.go

ECHO Check folder
IF NOT EXIST .\build (
    ECHO Create folder
	MKDIR .\build
)
ECHO Check file
IF EXIST %BIN_WINDOWS% (
    ECHO Delete %Exe%...
    DEL %BIN_WINDOWS%
)

echo Building %Exe%...
@ECHO ON
go build -v -o %BIN_WINDOWS%
@ECHO OFF