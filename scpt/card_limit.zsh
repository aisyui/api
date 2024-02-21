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
#updated_at_n=`date --iso-8601=seconds`
data=`curl -sL "$host_users"|jq .`

n=`echo $data|jq length`
n=$((n - 1))

for ((i=0;i<=$n;i++))
do
	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	echo "{\"updated_at\":\"$updated_at_n\"} -s $host/users/$id"
	curl -X PATCH -H "Content-Type: application/json" -d "{\"raid_at\": \"$updated_at_n\",\"updated_at\":\"$updated_at_n\",\"token\":\"$token\",\"ten_at\":\"$updated_at_n\"}" -s $host/users/$id
done
