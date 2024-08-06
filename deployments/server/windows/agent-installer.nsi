!include nsDialogs.nsh
!include LogicLib.nsh

Name "Screen Capture Agent"
Outfile ScreenCaptureAgentInstaller.exe
XPStyle on

Var Dialog
Var Label
Var Text

Page custom nsDialogsPage nsDialogsPageLeave
Page instfiles

Function nsDialogsPage

  nsDialogs::Create 1018
  Pop $Dialog

  ${If} $Dialog == error
          Abort
  ${EndIf}

  ${NSD_CreateLabel} 0 0 100% 12u "Enter your device name provided by the admin here."
  Pop $Label

  ${NSD_CreateText} 0 13u 100% 12u $Text
  Pop $Text

  nsDialogs::Show

FunctionEnd

Function nsDialogsPageLeave

  ${NSD_GetText} $Text $0

  ; Generate the batch file with the device name
  FileOpen $9 "$APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\screen-capture-agent.bat" w
  FileWrite $9 "@echo off$\r$\n"
  FileWrite $9 'START "Screen Capture Agent" "$PROGRAMFILES\ScreenCaptureAgent\screen-server.exe" -vhost $0$\r$\n'
  FileWrite $9 "exit$\r$\n"
  FileClose $9

FunctionEnd

Section
  SetOutPath "$PROGRAMFILES\ScreenCaptureAgent"
  File "screen-server.exe"
SectionEnd
