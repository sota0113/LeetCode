#! /bin/bash
columns=`cat transposeFile.txt | head -n 1 | awk -F ' ' '{print NF}'`
for i in `seq 1 ${columns}` ;do
	line=`cat transposeFile.txt | awk -v i="${i}" '{print $i}' | tr '\n' ' ' | sed -e 's/ $//g'`
	echo -n ${line[0]}
	for i in `seq 1 "${#line[@]}"` ;do echo -n " ${line[i]}" ;done | sed -e 's/ $//g'
	echo "" #necessary for LeetCode environment.
done
