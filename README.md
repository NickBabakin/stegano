# stegano

## Introduction

***stegano*** is a simple tool to hide information in bmp files.

***Warning***: info, hidden by ***stegano***, can not be visible by human eyes, but can be extracted via special tools. ***stegano*** does not encrypt information.

***stegano*** puts given information in a /bmp file in a way, that human eye is highly unlikely to see any changes. ***stegano*** changes last bit to 0 or 1 in first bytes of .bmp data section (thus header of bmp file is not corrupted). The amount of bytes changed depends on the size of the given information.

Size of hidden message is also hidden, it is stored before the message itself, in first 16 bytes of .bmp data section. So the maximum size of given message can be 65536 bytes. ***stegano*** also checks size of given info and .bmp file, so there is no undefined behaviour.

## Usage

*stegano* has two commands: *hide* and *extract*


**hide** puts message in file
```
./stegano hide --help

Usage of hide:
  -f string
        REQUIRED: source file used as a container
  -m string
        REQUIRED: message you want to hide
  -r string
        name of file to be created (default "result.bmp")
  -s    shows first 160 bytes of source container and container with hidden data
```
**extract** extracts message from file
```
niickson@niickson-laptop:~/goProjects/stegano$ ./
stegano extract --help
Usage of extract:
  -f string
        REQUIRED: file to extract info from
```

## Example

Hiding
```
./stegano hide -f bmp_example.bmp -m "Hello, World!" -s

Container before hiding
42 4D F6 D4 01 00 00 00 00 00 36 00 00 00 28 00 
00 00 C8 00 00 00 C8 00 00 00 01 00 18 00 00 00 
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
00 00 00 00 00 00 00 00 FF 00 00 FF 00 00 FF 00 
00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 
FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 
00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 
00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 
FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 
00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 00 FF 00 
Container after hiding
42 4D F6 D4 01 00 00 00 00 00 36 00 00 00 28 00 
00 00 C8 00 00 00 C8 00 00 00 01 00 18 00 00 00 
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
00 00 00 00 00 00 00 00 FE 00 00 FE 00 00 FE 00 
00 FE 01 01 FE 01 00 FF 00 00 FF 00 00 FE 00 01 
FF 00 00 FF 00 01 FE 01 01 FE 01 01 FE 00 00 FF 
01 00 FF 01 00 FE 00 01 FF 00 01 FF 01 01 FE 00 
01 FE 01 01 FE 00 00 FE 01 00 FE 00 00 FE 00 01 
FE 01 00 FF 01 01 FE 01 01 FE 01 01 FF 01 00 FF 
01 01 FE 00 01 FE 00 01 FF 00 01 FF 00 00 FE 01 
Message 'Hello, World!' is hidden.
New file name is result.bmp
```

Extracting

```
./stegano extract -f result.bmp 

Extraction ended successfully.
The message is 'Hello, World!'
```