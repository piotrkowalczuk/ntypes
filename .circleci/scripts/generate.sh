: ${PROTOC:="/usr/local/bin/protoc"}
: ${SCALAPBC:="/usr/local/bin/scalapbc"}
PROTO_INCLUDE="-I=/usr/include -I=."

DIR_JAVA="./tmp/java/"
DIR_PYTHON="./"

case $1 in
    lint)
        ${PROTOC} ${PROTO_INCLUDE} --lint_out=. *.proto
        ;;
    python)
        python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=${DIR_PYTHON}ntypes *.proto
        ;;
    java)
        rm -rf ./tmp/java
        mkdir -p ./tmp/java
        ${PROTOC} ${PROTO_INCLUDE} --java_out=./tmp/java *.proto
        ;;
    scala)
        rm -rf ./tmp/scala
        mkdir -p ./tmp/scala
        ${SCALAPBC} --scala_out=ntypes.jar ${PROTO_INCLUDE} *.proto
        ;;
    golang | go)
        ${PROTOC} ${PROTO_INCLUDE} --go_out=${GOPATH}/src *.proto
        ;;
	*)
	    echo "code generation failure due to unknown language: ${1}"
        exit 1
        ;;
esac