#!/bin/zsh

host=https://api.syui.ai
pass=`cat ~/.config/atr/api_card.json|jq -r .password`
if [ -z "$1" ];then
	exit
fi

echo username
read
id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`

card=0
model=0
sword=0
curl -X POST -H "Content-Type: application/json" -d "{\"owner\":$id,\"card\":$card, \"model\":$model,\"sword\":$sword,\"password\":\"$pass\"}" -sL $host/ues

curl -sL api.syui.ai/users/$id/ue
