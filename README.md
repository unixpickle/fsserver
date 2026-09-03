# fsserver

**fsserver** (pronounced "FS server") is an HTTP server which serves contents from the filesystem. It supports directory listings, index files, and various MIME types.

# Installation

This project is written in [Go](https://go.dev/doc/install). You will need Go 1.25 or newer in order to compile it from source.

Install it like this:

    go install github.com/unixpickle/fsserver@latest

You should now have an `fsserver` command.

# Usage

The `fsserver` command can take the following options:

    -index="index.html": the index filename
    -path=".": the directory to serve
    -port=80: server port number
    -silent=false: disable logging

For example, you can serve the directory "/Users/alex/Desktop" on port 8080 using this command:

    fsserver -path=/Users/alex/Desktop -port=8080

# Updating web assets

Web-related files are stored in the [lib/assets](lib/assets) directory and embedded in the binary with [`go:embed`](https://pkg.go.dev/embed). Changes to these files are included automatically the next time the project is built; no generation step is needed.

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
