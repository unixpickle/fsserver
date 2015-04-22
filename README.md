# fsserver

**fsserver** (pronounced "FS server") is an HTTP server which serves contents from the filesystem. It supports directory listings, index files, and various MIME types.

# Dependencies & Setup

You can get the dependencies like this:

    go get github.com/hoisie/mustache

Install go-bindata (if you don't already have it) like this:

    go get -u github.com/jteeuwen/go-bindata
    go install github.com/jteeuwen/go-bindata/go-bindata

Generate the bindata:

    cd lib
    go-bindata -debug -pkg=fsserver assets/...
    cd -

Finally, you can install fsserver:

    go install .

# Usage

You can run fsserver from the command line. Here is the usage:

    -index="index.html": the index filename
    -path=".": the directory to serve
    -port=80: server port number
    -silent=false: disable logging

For example, you can serve the directory "/Users/alex/Desktop" on port 8080 using this command:

    fsserver -path=/Users/alex/Desktop -port=8080

# TODO

 * Human-readable file sizes
 * Style the directory listing page
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