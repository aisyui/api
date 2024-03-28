#!/bin/zsh

d=${0:a:h}
cd $d
su=5000

go1.21.8 generate ./...
#go generate ./...
cp -rf $d/ent/openapi.json $d/tmp/

case $OSTYPE in
	darwin*)
		sed -i '' "s/255/$su/g" $d/ent/ogent/oas_parameters_gen.go
		sed -i '' "s/255/$su/g" $d/ent/openapi.json
		;;
	linux*)
		sed -i "s/255/$su/g" $d/ent/ogent/oas_parameters_gen.go
		sed -i "s/255/$su/g" $d/ent/openapi.json
		;;
esac

cp -rf $d/tmp/ogent ent/
#PASS=`cat token.json|jq -r .password` TOKEN=`cat token.json|jq -r .token` go build
#PASS=`cat token.json|jq -r .password` TOKEN=`cat token.json|jq -r .token` go run -mod=mod main.go
