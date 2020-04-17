#o/bin/bash
#sed -i s/\\\\n/\\$'\n'/g ./wordFrequency.txt
sed s/\\\\n/\\$'\n'/g ./wordFrequency.txt > tmp.txt
words=($(cat ./tmp.txt))
innerArray=()
array=()

# count each words
for i in `seq 0 ${#words[@]}`
do
	if [ -z ${words[${i}]} ];then
		continue
	fi
	if [ ${#array[@]} = 0 ] ;then # or if count=1
		innerArray=("${words[${i}]}-0")
		array=("${array[@]}" ${innerArray})
		continue
	fi
	for k in `seq 0 $((${#array[@]}-1))`  #if $result[@] is empty, loop would not be work.
	do
		arrayOp=`echo "${array[${k}]}" | awk -F'-' '{print $1}'`
		if [ "${words[${i}]}" = "${arrayOp}" ] ;then
			innerArray=`echo ${array[$k]} | awk -F'-' '{print $1 "-" $2+1}'`
			array[${k}]=${innerArray}
			break
		fi
		if [ ${k} -eq $((${#array[@]}-1)) ] ;then
                	innerArray=("${words[${i}]}-0")
                	array=("${array[@]}" ${innerArray})
                	break
		fi
	done
done
# if array is empty, exit here.
# show items with decendent order of word frequency.
echo ${array[@]} | sort -t "-" -k 2 -n -r | awk -F'-' '{print $1 " " $2+1}'
#for i in ${array[@]} ;do echo $i ;done  | sort -t "-" -k 2 -n -r | awk -F'-' '{print $1 " " $2+1}'
rm tmp.txt
