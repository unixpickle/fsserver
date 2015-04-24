# fsserver

**fsserver** (pronounced "FS server") is an HTTP server which serves contents from the filesystem. It supports directory listings, index files, and various MIME types.

# Installation

This project is written in [Go](https://golang.org/doc/install). You will need Go in order to compile it from source.

You can download the repository like this:

    go get -u github.com/unixpickle/fsserver

And install it like this:

    go install github.com/unixpickle/fsserver

You should now have an `fsserver` command.

# Usage

The `fsserver` command can take the following options:

    -index="index.html": the index filename
    -path=".": the directory to serve
    -port=80: server port number
    -silent=false: disable logging

For example, you can serve the directory "/Users/alex/Desktop" on port 8080 using this command:

    fsserver -path=/Users/alex/Desktop -port=8080

# Updating [lib/bindata.go](lib/bindata.go)

Notice that web-related files are stored in the [lib/assets](lib/assets) directory. Changing files in this directory will not update them in the Go code; the Go code uses [go-bindata](https://github.com/jteeuwen/go-bindata) to compile these assets statically.

So, if you don't have go-bindata, install it like this:

    go get -u github.com/jteeuwen/go-bindata
    go install github.com/jteeuwen/go-bindata/go-bindata

Now you can generate bindata.go from the assets directory:

    cd lib
    go-bindata -pkg=fsserver assets/...
    cd -

# TODO

 * Send the "Content-Length" header

# License

**fsserver** is licensed under the BSD 2-clause license. See [LICENSE](LICENSE).

```
Copyright (c) 2015, Alexander Nichol.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer. 
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```