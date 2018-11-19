: ${PROTOC:="/usr/local/bin/protoc"}
: ${SCALAPBC:="./tmp/scalapbc/scalapbc-0.8.1/bin/scalapbc"}
PROTO_INCLUDE="-I=/usr/include -I=${GOPATH}/src"
VERSION=$(git describe --tags --always --dirty)
SCALA_VERSION="2.12.7"
DIR_PYTHON="./"
DIR_SCALA="./tmp/scala"

case $1 in
    lint)
        ${PROTOC} ${PROTO_INCLUDE} --lint_out=. *.proto
        ;;
    python)
        python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=${DIR_PYTHON}ntypes *.proto
        ;;
    scala)
        rm -rf ${DIR_SCALA}
        mkdir -p ${DIR_SCALA}/src/main/scala ${DIR_SCALA}/project
        ${SCALAPBC} -v361 --scala_out=./tmp/scala/src/main/scala ${PROTO_INCLUDE} *.proto
        .circleci/templates/build.sbt.sh ${VERSION:1} ${SCALA_VERSION} > ./tmp/scala/build.sbt
        .circleci/templates/plugins.sbt.sh > ./tmp/scala/project/plugins.sbt
        ;;
    golang | go)
        ${PROTOC} ${PROTO_INCLUDE} --go_out=${GOPATH}/src ${GOPATH}/src/github.com/piotrkowalczuk/ntypes/*.proto
        ;;
	*)
	    echo "code generation failure due to unknown language: ${1}"
        exit 1
        ;;
esac
