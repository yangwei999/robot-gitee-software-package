#!/bin/sh

set -euo pipefail

# repo_url is the url to push code, it contains username and token
repo_url=$1
# repo is the repo name of repo_url
repo=$2
user=$3
email=$4
# get code from the pr of ci repo, therefore it can guarantee that the code is CI checked
ci_repo_link=$5
ci_repo=$6

if [ ! -d $ci_repo ]; then
  mkdir $ci_repo
  git clone -q $ci_repo_link $ci_repo
fi

cd $ci_repo

git checkout master

git fetch origin

git checkout $repo

cd ..

if [ -d $repo ]; then
    rm -rf $repo
fi

git clone $repo_url

cp $ci_repo/* $repo

cd $repo

git config user.name $user
git config user.email $email

# get the files size more than 50MB but not in .git folder
large_files=$(find . -path '*/.git' -prune -o -type f -size +50M -print)
if [ -n "${large_files}" ]; then
    git lfs install
    git lfs track --filename ${large_files}
fi

git add .
git commit -m 'add spec and rpm'
git push

cd ..

rm -rf $repo