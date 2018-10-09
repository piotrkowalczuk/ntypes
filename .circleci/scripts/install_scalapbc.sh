curl -L https://github.com/scalapb/ScalaPB/releases/download/v0.8.1/scalapbc-0.8.1.zip > scalapbc.zip

rm -rf ./.tmp/scalapbc
mkdir -p ./.tmp/scalapbc
unzip scalapbc.zip -d ./.tmp/scalapbc

# scalapbc has no version flag