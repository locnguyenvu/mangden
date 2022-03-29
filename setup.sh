#!/usr/bin/env bash

function get_machine_os() {
    unameOut="$(uname -s)"
    case "${unameOut}" in
        Linux*)     machine=Linux;;
        Darwin*)    machine=Mac;;
        CYGWIN*)    machine=Cygwin;;
        MINGW*)     machine=MinGw;;
        *)          machine="UNKNOWN:${unameOut}"
    esac
    echo ${machine}
}

function replace_repository_name() {
    repositoryName=$1
    file=$2

    foundPattern=$(grep "locnguyenvu/mangden" -l $file | wc -l)
    if [[ $foundPattern -eq 0 ]];then
        return
    fi
    echo $file
    machineOs=$(get_machine_os)
    if [ "${machineOs}" == "Mac" ]; then
        sed -i '' "s#locnguyenvu/mangden#${repositoryName}#g" $file
    fi
    if [ "${machineOs}" == "Linux" ]; then
        sed -i  "s#locnguyenvu/mangden#${repositoryName}#g" $file
    fi
}

repositoryName=''
while [ ! "${repositoryName}" ]; do
    read -p 'Enter repository name: ' repositoryName

    if [ ! "${repositoryName}" ]; then
        echo -ne "name is missing... "
    fi
done

echo "Download source ..."
curl -s -k -L  https://github.com/locnguyenvu/mangden/tarball/master | tar -xz
sourceDir=$(ls | grep "locnguyenvu-mangden")
cp -R ${sourceDir}/* .
cp ${sourceDir}/.env.example ./.env
rm -rf ${sourceDir}

echo "Replacing repository name ..."
go_files=$(find . -iname "*.go")
for file in ${go_files}; do
    replace_repository_name $repositoryName $file
done
replace_repository_name $repositoryName ./mdn

rm -rf ./setup.sh
