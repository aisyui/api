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

while true
do


for ((id=1221;id<=1230;id++))
do
	t=`curl -sL $host/users/$id|jq .code`
	sleep 3
	if [ $t -ne 404 ];then
		echo ok
		curl -X PATCH -H "Content-Type: application/json" -d "{\"raid_at\": \"$updated_at_n\",\"updated_at\":\"$updated_at_n\",\"token\":\"$token\",\"ten_at\":\"$updated_at_n\"}" -s $host/users/$id
	else
		echo $t
	fi
done


done
