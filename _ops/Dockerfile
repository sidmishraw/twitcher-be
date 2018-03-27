#
#  BSD 3-Clause License
#
# Copyright (c) 2018, Sidharth Mishra
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are met:
#
# * Redistributions of source code must retain the above copyright notice, this
#  list of conditions and the following disclaimer.
#
# * Redistributions in binary form must reproduce the above copyright notice,
#  this list of conditions and the following disclaimer in the documentation
#  and/or other materials provided with the distribution.
#
# * Neither the name of the copyright holder nor the names of its
#  contributors may be used to endorse or promote products derived from
#  this software without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
#
# Dockerfile
# @author Sidharth Mishra
# @created Tue Aug 29 2017 17:44:08 GMT-0700 (PDT)
# @last-modified Mon Mar 26 2018 20:42:46 GMT-0700 (PDT)
#

# Following is the Dockerfile for Twitcher's Golang backend REST service.

# Base image is Go v1.10
FROM golang:1.10

LABEL author="Sidharth Mishra <sidmishraw@gmail.com>"
LABEL twitcher.version="0.0.1"

# exposing port 8080 of the container
EXPOSE 8080 

# make the project directory in the container
RUN mkdir -p /go/src/github.com/sidmishraw

# copy the source files into the container
# Let the build context for this image be at the project root
# then copy the project root into the container
COPY . /go/src/github.com/sidmishraw/twitcher

# Set the working directory for twitcher
WORKDIR /go/src/github.com/sidmishraw/twitcher

# Clean the go directory
RUN go clean

# Build twitcher
RUN go build -o twitcher -i

# Set the twitcher app built to be the entry point for the container?
ENTRYPOINT [ "/go/src/github.com/sidmishraw/twitcher/twitcher" ]
