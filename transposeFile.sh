#! /bin/bash
#
# Input as below.
# a b c d
# 1 2 3 4
# 1 2 3 4
# 
# output as below.
# a 1 1
# b 2 2
# c 3 3
# d 4 4
#

columns=`cat transposeFile.txt | head -n 1 | awk -F ' ' '{print NF}'`
for i in `seq 1 ${columns}` ;do
	line=`cat transposeFile.txt | awk -v i="${i}" '{print $i}' | tr '\n' ' ' | sed -e 's/ $//g'`
	#echo -n ${line[0]}
	for i in `seq 0 "${#line[@]}"` ;do echo -n " ${line[i]}" ;done | sed -e 's/ $//g'
	#echo "" #necessary for LeetCode environment.
done
