# Building from source

**It is recommended that you clone exactly this git repository and build from root directory (not from `build` directory).**

**Requirements:**

* Go compiler 1.16

**How to build:**

Assuming you have cloned this git repository...

If you are on Windows, choose either options to build:

1. Run cmd in root directory and run:
```powershell
powershell -executionpolicy bypass -File .\build\build.ps1
```
2. Launch Windows Powershell, `cd` to root directory, and run:
```powershell
.\build\build.ps1
```

If you are on a Unix-like OS:

1. Go to terminal and `cd` to root directory. Then, run:
```sh
chmod +x ./build/build.sh
```
(to give execute permission to the script), and run:
```sh
./build/build.sh
```