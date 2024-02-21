#!/bin/zsh

if [ -z "$1" ];then
	exit
fi
echo delete-id
read
id=$1
data=`curl -sL "https://api.syui.ai/users/$id/card?itemsPerPage=2550"`
token=`cat ~/.config/atr/api_card.json|jq -r .token`
pass=`cat ~/.config/atr/api_card.json|jq -r .password`

n=`echo $data|jq length`

n=$((n - 1))

for ((i=0;i<=$n;i++))
do
	card_id=`echo $data|jq -r ".[$i].id"`
	echo $card
	curl -X DELETE -H "Content-Type: application/json" -d "{\"owner\":$id,\"password\":\"$pass\"}" https://api.syui.ai/cards/$card_id
done
#curl -X DELETE -H "Content-Type: application/json" https://api.syui.ai/users/$id
