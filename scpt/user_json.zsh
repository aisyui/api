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
n=`echo $data|jq length`
n=$((n - 1))

f=~/ai/card/public/json/user.json
if [ -f $f ];then
	rm $f
fi

echo "{" >! $f

for ((i=0;i<=$n;i++))
do
	name=`echo $data|jq ".[$i]"|jq -r .username`
	id=`echo $data|jq ".[$i]"|jq -r .id`
	j="\"$name\":$id"
	j="\"$name\":\"$id\""
	if [ $n -ne $i ];then
		j="${j},"
	fi
	echo $j
	echo $j >> $f
done

	echo "}" >> $f

	cat $f|jq .
