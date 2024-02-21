#!/bin/zsh
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
if [ -z "$1" ];then
	exit
fi
echo username
read
id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$1\")"|jq -r .id`

echo $id
read

#curl -X PATCH -H "Content-Type: application/json" -d "{\"model_mode\":1, \"model_skill\":1,\"model_attack\":1,\"model_limit\":1,\"model_critical\":1,\"model_critical_d\":10, \"model\":true,\"token\":\"$token\"}" -s $host/users/$id


curl -X PATCH -H "Content-Type: application/json" -d "{\"model_critical\":1, \"token\":\"$token\"}" -s $host/users/$id
