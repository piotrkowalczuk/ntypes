: ${PROTOC:="/usr/local/bin/protoc"}
: ${SCALAPBC:="./.tmp/scalapbc/scalapbc-0.8.1/bin/scalapbc"}
PROTO_INCLUDE="-I=/usr/include -I=."

DIR_JAVA="./.tmp/java/"
DIR_PYTHON="./"

case $1 in
    lint)
        ${PROTOC} ${PROTO_INCLUDE} --lint_out=. *.proto
        ;;
    python)
        python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=${DIR_PYTHON}ntypes *.proto
        ;;
    scala)
        rm -rf ./.tmp/scala
        mkdir -p ./.tmp/scala
        ${SCALAPBC} -v361 --scala_out=./.tmp/scala/ntypes.jar ${PROTO_INCLUDE} *.proto
        ;;
    golang | go)
        ${PROTOC} ${PROTO_INCLUDE} --go_out=${GOPATH}/src *.proto
        ;;
	*)
	    echo "code generation failure due to unknown language: ${1}"
        exit 1
        ;;
esac