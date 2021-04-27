

set -eux

PROJECT_ROOT="${GITHUB_REPOSITORY}"

mkdir -p $PROJECT_ROOT
rmdir $PROJECT_ROOT
ln -s $GITHUB_WORKSPACE $PROJECT_ROOT
cd $PROJECT_ROOT

export GO111MODULE="on" 
export CGO_ENABLED=0
go mod download
go mod verify

EXT=''

if [ $GOOS == 'windows' ]; then
EXT='.exe'
fi

if [ -x "./build.sh" ]; then
  OUTPUT=`./build.sh "${CMD_PATH}"`
else
  go build "${CMD_PATH}"
  OUTPUT="${PROJECT_NAME}${EXT}"
fi

echo ${OUTPUT}
