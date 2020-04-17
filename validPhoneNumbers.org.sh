cat validPhoneNumbers.txt | tr '\\n' '\n' > tmp.txt
while read line
do
        result=$(echo "${line}" | grep '(' | wc -l)
        # parentheses exist
        if [ ${result} = 1 ] ;then
                formerOfFormer=$(echo "${line}" | awk -F'-' '{print $1}' | awk -F')' '{print $1}')
                latterOfFormer=$(echo "${line}" | awk -F'-' '{print $1}' | awk -F')' '{print $2}'| cut -c 2-4)
                latter=$(echo "${line}" | awk -F'-' '{print $2}' )

                if [[ "${formerOfFormer}" =~ ^'('[0-9]{3}$ ]];then
                        if [[ "${latterOfFormer}" =~ ^[0-9]{3}$ ]];then
                                if [[ "${latter}" =~ ^[0-9]{4}$ ]];then
                                        echo "${line}"
                                fi
                        fi
                fi
        fi
        # no paretheses exist
        if [ ${result} = 0 ];then
                # each token should have at least one numeric character.
                token=$(echo "${line}" | awk -F'-' '{print $1}')
                if [[ ! "${token}" =~ ^[0-9]{3}$ ]];then
                        continue
                fi
                token=$(echo "${line}" | awk -F'-' '{print $2}')
                if [[ ! "${token}" =~ ^[0-9]{3}$ ]];then
                        continue
                fi
                token=$(echo "${line}" | awk -F'-' '{print $3}')
                if [[ ! "${token}" =~ ^[0-9]{4}$ ]];then
                        continue
                fi
                echo "${line}"
        fi
done < tmp.txt
rm tmp.txt
