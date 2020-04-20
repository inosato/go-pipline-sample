#!/bin/bash

echo build domainifinder
go build -o domainifinder

echo build synonyms
cd ../synonyms
go build -o ../domainifinder/lib/synonyms

echo build available
cd ../available
go build -o ../domainifinder/lib/available

echo build sprinkle
cd ../sprinkle
go build -o ../domainifinder/lib/sprinkle

echo build coolify
cd ../coolify
go build -o ../domainifinder/lib/coolify

echo build domainify
cd ../domainify
go build -o ../domainifinder/lib/domainify