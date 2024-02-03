#!/bin/zsh
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`

username=ai
id=`curl -sL "$host/users?itemsPerPage=2000"|jq ".[]|select(.username == \"$username\")"|jq -r .id`

curl -sL $host/users/$id
read
echo $id
o=true
echo $o
curl -X PATCH -H "Content-Type: application/json" -d "{\"game\":true, \"game_test\":$o, \"game_end\":$o, \"token\":\"$token\"}" -s $host/users/$id
