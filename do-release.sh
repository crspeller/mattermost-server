#!/bin/bash

set -o
MM_VERSION=4.1.2

LONG_RELEASE=$MM_VERSION-rc1
SHORT_RELEASE=${MM_VERSION:0:3}
MM_FILENAME='mattermost-team-'$LONG_RELEASE'-linux-amd64.tar.gz'
RELEASE_BRANCH=release-$SHORT_RELEASE
RELEASE_TAG=v$LONG_RELEASE

echo "Long Release: "$LONG_RELEASE
echo "Short Release: "$SHORT_RELEASE
echo "MM Filename: "$MM_FILENAME
echo "Release Branch: "$RELEASE_BRANCH
echo "Release Tag: "$RELEASE_TAG

# Do branch
git checkout -b $RELEASE_BRANCH

#
# Update docker files
#
cd docker
rm -r `find . -maxdepth 1 -type d ! -path . | sort | head -n1`
cp -r dev $SHORT_RELEASE

cd $SHORT_RELEASE
sed -i'.bak' 's|RUN wget --no-check-certificate https://s3.amazonaws.com/mattermost-travis-master/mattermost.tar.gz|ADD http://releases.mattermost.com/'$MM_FILENAME'/|g' ./Dockerfile
sed -i'.bak' 's|RUN tar -zxvf mattermost.tar.gz --strip-components=1 && rm mattermost.tar.gz|RUN tar -zxvf /'$MM_FILENAME' --strip-components=1 && rm /'$MM_FILENAME'|g' ./Dockerfile
rm Dockerfile.bak

rm Dockerrun.aws.zip
cd Dockerrun.aws
sed -i'.bak' 's|mattermost/platform:dev|mattermost/platform:'$SHORT_RELEASE'|g' ./Dokerrun.aws.json
rm Dokerrun.aws.json.bak
zip -r Dockerrun.aws.zip .
mv Dockerrun.aws.zip ../
cd ../../..

# Commit docker changes
git add .
git commit -m "Docker updates for "$LONG_RELEASE

#
# Bump version.go
#

sed -i.bak 's|var versions = \[\]string{|var versions = \[\]string{\n\t"'$MM_RELEASE'",|g' model/version.go
rm model/version.go.bak

git add .
git commit -m "Bumping version to: "$LONG_RELEASE

#
# Tag and push
#

git tag -a $RELEASE_TAG -m "Mattermost release "$RELEASE_TAG
git push origin $RELEASE_BRANCH
git push origin $RELEASE_TAG
