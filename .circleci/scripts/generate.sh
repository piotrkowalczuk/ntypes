: ${PROTOC:="/usr/local/bin/protoc"}
PROTO_INCLUDE="-I=/usr/include -I=."

(
case $1 in
    lint)
        ${PROTOC} ${PROTO_INCLUDE} --lint_out=. *.proto
        ;;
    python)
        python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=./ntypes *.proto
        ;;
    java)
        rm -rf ./tmp/java
        mkdir -p ./tmp/java
        ${PROTOC} ${PROTO_INCLUDE} --java_out=./tmp/java *.proto
        ;;
    golang | go)
        ${PROTOC} ${PROTO_INCLUDE} --go_out=${GOPATH}/src *.proto
        ;;
	*)
	    echo "code generation failure due to unknown language: ${1}"
        exit 1
        ;;
esac
) 2>&1