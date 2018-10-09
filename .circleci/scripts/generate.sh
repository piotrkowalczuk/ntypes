: ${PROTOC:="/usr/local/bin/protoc"}
: ${SCALAPBC:="./tmp/scalapbc/scalapbc-0.8.1/bin/scalapbc"}
PROTO_INCLUDE="-I=/usr/include -I=."
VERSION=$(git describe --tags --always --dirty)
POM_FILE="<?xml version=\"1.0\" encoding=\"UTF-8\"?> \
<project xmlns=\"http://maven.apache.org/POM/4.0.0\" \
         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" \
         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd\"> \
    <modelVersion>4.0.0</modelVersion> \
    <groupId>com.github.piotrkowalczuk</groupId> \
    <artifactId>ntypes</artifactId> \
    <version>${VERSION:1}</version> \
    <dependencies> \
        <dependency> \
            <groupId>com.thesamet.scalapb</groupId> \
            <artifactId>compilerplugin_2.12</artifactId> \
            <version>0.8.1</version> \
        </dependency> \
    </dependencies> \
</project>"

DIR_JAVA="./tmp/java/"
DIR_PYTHON="./"

case $1 in
    lint)
        ${PROTOC} ${PROTO_INCLUDE} --lint_out=. *.proto
        ;;
    python)
        python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=${DIR_PYTHON}ntypes *.proto
        ;;
    scala)
        rm -rf ./tmp/scala
        mkdir -p ./tmp/scala ./tmp/protobuf
        ${SCALAPBC} -v361 --scala_out=./tmp/scala/ntypes.jar ${PROTO_INCLUDE} *.proto
        zip -ur ./tmp/scala/ntypes.jar ./README.md ./LICENSE.txt
        echo "${POM_FILE}" > ./tmp/pom.xml
        cp ntypes.proto ./tmp/protobuf/ntypes.proto
        cd ./tmp/ && zip -ur ./scala/ntypes.jar ./pom.xml ./protobuf/ntypes.proto
        ;;
    golang | go)
        ${PROTOC} ${PROTO_INCLUDE} --go_out=${GOPATH}/src *.proto
        ;;
	*)
	    echo "code generation failure due to unknown language: ${1}"
        exit 1
        ;;
esac