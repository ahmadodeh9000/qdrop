# qdrop

Send a file from your computer to your phone by scanning a QR code. No apps, no cables, no cloud.

```
% qdrop report.pdf
FILE :   report.pdf (245.3 KB)
LINK :   http://192.168.1.12:53211/a3f9c2d41b7e8a05

[QR code]
⬛⬛⬛⬛⬛⬛⬛⬜⬛⬛⬜⬜⬜⬜⬜⬜⬛⬜⬛⬛⬛⬛⬛⬛⬛
⬛⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬜⬜⬛
⬛⬜⬛⬛⬛⬜⬛⬜⬛⬜⬜⬛⬛⬛⬜⬛⬜⬜⬛⬜⬛⬛⬛⬜⬛
⬛⬜⬛⬛⬛⬜⬛⬜⬛⬛⬜⬛⬛⬜⬛⬜⬜⬜⬛⬜⬛⬛⬛⬜⬛
⬛⬜⬛⬛⬛⬜⬛⬜⬛⬛⬛⬛⬜⬜⬜⬛⬜⬜⬛⬜⬛⬛⬛⬜⬛
⬛⬜⬜⬜⬜⬜⬛⬜⬜⬛⬜⬜⬛⬛⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬛
⬛⬛⬛⬛⬛⬛⬛⬜⬛⬜⬛⬜⬛⬜⬛⬜⬛⬜⬛⬛⬛⬛⬛⬛⬛
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜
⬛⬛⬛⬛⬜⬜⬛⬜⬛⬜⬜⬜⬛⬜⬛⬜⬛⬛⬜⬜⬛⬛⬛⬜⬛
⬜⬜⬜⬜⬛⬜⬜⬛⬜⬛⬜⬜⬜⬜⬛⬜⬛⬜⬜⬛⬜⬜⬜⬛⬜
⬜⬛⬛⬛⬛⬜⬛⬛⬜⬜⬜⬛⬛⬛⬜⬜⬜⬜⬛⬛⬛⬜⬜⬜⬜
⬜⬛⬛⬛⬛⬜⬜⬛⬜⬜⬜⬛⬜⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬜⬜
⬜⬜⬛⬜⬜⬛⬛⬜⬜⬛⬜⬜⬛⬛⬜⬛⬜⬛⬛⬛⬛⬜⬛⬛⬛
⬜⬜⬜⬜⬛⬛⬜⬜⬛⬛⬛⬜⬛⬜⬜⬛⬛⬜⬛⬛⬛⬜⬜⬜⬛
⬜⬛⬜⬛⬜⬜⬛⬛⬜⬜⬜⬜⬜⬛⬛⬜⬛⬛⬜⬜⬛⬜⬛⬛⬜
⬛⬜⬛⬛⬛⬜⬜⬜⬛⬛⬜⬜⬛⬜⬜⬜⬛⬛⬛⬛⬛⬜⬜⬜⬛
⬜⬜⬜⬛⬛⬛⬛⬜⬜⬛⬜⬜⬛⬛⬜⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛
⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜⬜⬛⬛⬛⬜⬛⬜⬜⬜⬛⬜⬛⬜⬛
⬛⬛⬛⬛⬛⬛⬛⬜⬜⬜⬜⬜⬜⬛⬛⬜⬛⬜⬛⬜⬛⬜⬛⬛⬛
⬛⬜⬜⬜⬜⬜⬛⬜⬜⬜⬛⬜⬛⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬛⬜
⬛⬜⬛⬛⬛⬜⬛⬜⬜⬛⬛⬜⬜⬜⬜⬜⬛⬛⬛⬛⬛⬛⬜⬛⬛
⬛⬜⬛⬛⬛⬜⬛⬜⬛⬛⬜⬛⬜⬛⬛⬜⬜⬛⬛⬜⬛⬛⬛⬛⬛
⬛⬜⬛⬛⬛⬜⬛⬜⬛⬜⬜⬛⬛⬜⬛⬛⬜⬜⬛⬜⬛⬜⬛⬛⬜
⬛⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬛⬛⬛⬛⬛⬜⬛⬜⬜⬛⬜⬛⬜⬜
⬛⬛⬛⬛⬛⬛⬛⬜⬛⬛⬛⬜⬜⬜⬛⬜⬜⬜⬛⬛⬛⬛⬛⬛⬛
```

Scan the code with your phone's camera and the file downloads. qdrop exits as soon as the transfer is done.

## Install

```
go install github.com/ahmadodeh9000/qdrop@latest
```

Or build from source:

```
git clone https://github.com/ahmadodeh9000/qdrop.git
cd qdrop
go build
```

## Usage

```
qdrop <file>
```

Your computer and phone must be on the same Wi-Fi network.

## How it works

- Starts a small HTTP server on a random port
- Generates a random token for the URL, so only someone who scans the code can download the file
- Shows the link as a QR code in your terminal
- Shuts down after the first download

## Note
I build it because I wanted to learn golang, so if you find any bug please feel free to reach me out

## License

MIT