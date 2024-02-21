#!/bin/zsh

case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
host_users="$host/users?itemsPerPage=2550"
updated_at_n=`date --iso-8601=seconds -d '1 days ago'`
now_at=`date --iso-8601=seconds`
raid_at_n=`date --iso-8601=seconds -d '1 days ago'`
data=`curl -sL "$host_users"|jq .`
nd=`date +"%Y%m%d"`

n=`echo $data|jq length`
n=$((n - 1))

for ((i=0;i<=$n;i++))
do
	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	model=`echo $data|jq ".[$i]"|jq -r .model`
	#echo $model
	if [ "$model" = false ];then
		card=`curl -sL "$host/users/$id/card?itemsPerPage=2550"|jq -r ".[]|select(.skill == \"model\")"`
		if [ -n "$card" ];then
			echo "\n[${id}] $name"
			curl -X PATCH -H "Content-Type: application/json" -d "{\"model\":true,\"model_limit\": 1,\"token\":\"$token\"}" -s $host/users/$id
		fi
	fi
done
