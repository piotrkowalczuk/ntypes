curl -L https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip > protoc.zip
unzip protoc.zip -d ./tmp/protoc
sudo mv ./tmp/protoc/bin/protoc /usr/local/bin/protoc
sudo mv ./tmp/protoc/include/google /usr/include/google
protoc --version