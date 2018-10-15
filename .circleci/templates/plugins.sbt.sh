#!/bin/sh

#define parameters which are passed in.
VERSION=$1
SCALA_VERSION=$2

#define the template.
cat  << EOF
// Bintray
addSbtPlugin("org.foundweekends" % "sbt-bintray" % "0.5.4")

// Protoc
addSbtPlugin("com.thesamet" % "sbt-protoc" % "0.99.18")
libraryDependencies += "com.thesamet.scalapb" %% "compilerplugin" % "0.8.1"
EOF