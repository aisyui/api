#!/bin/zsh
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
if [ -z "$1" ];then
	exit
fi
if [ -z "$2" ];then
	did=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.did == \"$1\")"|jq -r .id`
	echo $did
	exit
fi

echo old-id new-id
read

echo account delete $1 token
curl -X PATCH -H "Content-Type: application/json" -d "{\"delete\":true,\"token\":\"$token\"}" -s $host/users/$1

echo account refresh $2 token
curl -X PATCH -H "Content-Type: application/json" -d "{\"delete\":false,\"token\":\"$token\"}" -s $host/users/$2
