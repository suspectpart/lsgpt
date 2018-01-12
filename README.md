# About
Dumps GPT (Guid Partition Table) information on EFI systems. This project has no other purpose that for me to learn `Go`.

See [Wikipedia/GPT](https://en.wikipedia.org/wiki/GUID_Partition_Table) for information about the GPT structure. 

## Install
```bash
$ go get github.com/suspectpart/lsgpt
```

## Run
```bash
$ sudo $GOPATH/bin/lsgpt /dev/sdx
```

Replace `/dev/sdx` with the drive that contains your GPT.