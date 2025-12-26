.MODEL SMALL
.STACK 100H
.DATA
MSG1 DB &#39;Enter how many numbers: $&#39;
MSG2 DB &#39;Enter the numbers: $&#39;
MSG3 DB &#39;Array elements: $&#39;
NEWLINE DB 13,10,&#39;$&#39;
N DB ? ; user input count
ARR DB 50 DUP(?) ; maximum 50 numbers
.CODE
MAIN PROC
MOV AX, @DATA
MOV DS, AX
;--------------------------------------
; Ask how many inputs
;--------------------------------------
MOV AH, 9
LEA DX, MSG1
INT 21H
MOV AH, 9
LEA DX, NEWLINE
INT 21H
; read single digit number (1–9)
MOV AH, 1
INT 21H
SUB AL, &#39;0&#39;
MOV N, AL ; store count in N
MOV AH, 9
LEA DX, NEWLINE
INT 21H
;--------------------------------------
; Ask user to enter numbers
;--------------------------------------
MOV AH, 9
LEA DX, MSG2
INT 21H
MOV AH, 9
LEA DX, NEWLINE
INT 21H

;--------------------------------------
; INPUT LOOP
;--------------------------------------
MOV CL, N ; number of inputs
MOV CH, 0
MOV SI, 0
INPUT_LOOP:
MOV AH, 1
INT 21H
SUB AL, 30H
MOV ARR[SI], AL
INC SI
LOOP INPUT_LOOP
;--------------------------------------
; DISPLAY ARRAY CONTENTS
;--------------------------------------
MOV AH, 9
LEA DX, NEWLINE
INT 21H
MOV AH, 9
LEA DX, MSG3
INT 21H
MOV AH, 9
LEA DX, NEWLINE
INT 21H
MOV CL, N
MOV CH, 0
MOV SI, 0
DISPLAY_LOOP:
MOV AL, ARR[SI]
ADD AL, 30H
MOV DL, AL
MOV AH, 2
INT 21H
INC SI
LOOP DISPLAY_LOOP
;--------------------------------------
; EXIT
;--------------------------------------
MOV AH, 4CH

INT 21H
MAIN ENDP
END MAIN