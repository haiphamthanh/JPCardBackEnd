#!/bin/bash

# Setup your workspace
# Link: https://sourabhbajaj.com/mac-setup/Go/
export GOPATH=$HOME/go/JPCardBackEnd # Put your current directory path
export PATH=$PATH:$(go env GOPATH)/bin # Embed your path to global environments 
echo $GOPATH

# Create folders
# https://hackernoon.com/how-to-create-golang-rest-api-project-layout-configuration-part-1-am733yi7?ref=hackernoon.com
mkdir -p \
$GOPATH/src/github.com/boilerplate/pkg \
$GOPATH/src/github.com/boilerplate/cmd \
$GOPATH/src/github.com/boilerplate/db/scripts \
$GOPATH/src/github.com/boilerplate/scripts

# Make docker file
cd $GOPATH/src/github.com/boilerplate && go mod init github.com/boilerplate