#!/bin/zsh

host=https://api.syui.ai
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
echo username card cp
read
if [ -z "$1" ];then
	exit
fi
if [ -z "$2" ];then
	exit
fi
if [ -z "$3" ];then
	exit
fi


id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`
card=$2
cp=$3
s=normal
skill=normal
count=1
author=ai
curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card,\"status\":\"$s\",\"cp\":$cp,\"password\":\"$pass\",\"skill\":\"$skill\",\"count\":$count,\"author\":\"$author\"}" -sL $host/cards
