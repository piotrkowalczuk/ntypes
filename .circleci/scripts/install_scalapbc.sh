curl -L https://github.com/scalapb/ScalaPB/releases/download/v0.7.4/scalapbc-0.7.4.zip > scalapbc.zip

rm -rf ./.tmp/scalapbc
mkdir -p ./.tmp/scalapbc
unzip scalapbc.zip -d ./.tmp/scalapbc

# scalapbc has no version flag