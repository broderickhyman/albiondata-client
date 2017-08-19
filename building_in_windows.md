# Building the Client in Windows
These instructions are specific to Windows 10 but should work in
earlier versions with a minor bit of tweaking.

## Installing Git
[Git Download Link](https://git-scm.com/download/win)

Download and install the Git software from the Download Link.
Choose to use Git in Git Bash only.

## Installing and Configuring Go
### Downloading and Installing
[Go Download Link](https://golang.org/doc/install#windows)

Grab the Microsoft Windows installer from the Download Link page.
Install via the MSI. After installation, open up Git Bash and run
`go version` to ensure a proper response.

### Testing Go
[Go Getting Started](https://golang.org/doc/install)

Follow the guide on this page for testing the install. In Windows
10, I had to create the `\go` directory in `%userprofile%\go`.
Other than that, the instructions should be followed completely.
Once that works, we're ready to move on.

## Installing Glide
Glide will help us pull in all of the dependencies for the project.

[Glide Download](https://github.com/Masterminds/glide/releases)

In the Download link above, find the most recent 64 bit release for
Windows. Extract the ZIP file and copy the glide.exe to your Go
executable path. Typically that is `C:\Go\bin`

Test Glide in Git Bash by typing `glide --version`. You should be
presented with a version string.

## Supporting Software

### MinGW-64
Used for building required libraries

#### Download and Install
[Download](https://sourceforge.net/projects/mingw-w64/files/latest/download)

**NOTE**: Ensure that when installing you select "Architecture: x86_64".
If you do not, you will get errors when trying to compile software during
the later steps.

#### Add to Windows PATH
* Right click Start and select System
* Click "Advanced System Settings" on the left side
* Click "Environment Variables" on the bottom right
* Under "System Variables" settings, find "Path" and click edit
* Click "New" and add the entry `C:\Program Files\mingw-w64\x86_64-7.1.0-posix-seh-rt_v5-rev1\mingw64\bin`
* Click "Ok"
* Open a new Git Bash window, and type `gcc --version`. You should greeted with a version string and description 

### WinPcap Developer
[Download](https://www.winpcap.org/devel.htm)

Download the WinPcap developer files. Extract the file contents to `C:\WpdPack`

**NOTE**: Ensure there is not a secondary WpdPack folder nested within
`C:\WpdPack` or your `go get` command will fail with a `pcap.h` warning.

## Building the Client

### Get the client from Github
For these instructions, I'll be using Git Bash.

Navigate to `~/go/`

Create the directory and subdirectories and navigate there:
```
mkdir -p src/github.com/regner
cd src/github.com/regner
```

Clone the repo

`git clone https://github.com/regner/albiondata-client.git`

### Fetch supporting Libraries

```
cd albiondata-client/
glide install
```

### Build the Client

```
go build
```

You should now be the proud owner of a new `albiondata-client.exe` executable!
