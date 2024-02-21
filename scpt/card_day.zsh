#!/bin/zsh

case $OSTYPE in
	darwin*)
		alias date="/opt/homebrew/bin/gdate"
		;;
esac
host=https://api.syui.ai
token=`cat ~/.config/atr/api_card.json|jq -r .token`
host_users="$host/users?itemsPerPage=2550"
data=`curl -sL "$host_users"|jq .`
nd=`date +"%Y%m%d"`

n=`echo $data|jq length`
n=$((n - 1))

for ((i=0;i<=$n;i++))
do
	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	echo "{\"next\":\"$nd\"} -s $host/users/$id"
	curl -X PATCH -H "Content-Type: application/json" -d "{\"next\":\"$nd\",\"token\":\"$token\"}" -s $host/users/$id
done
