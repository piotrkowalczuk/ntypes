#!/bin/sh

#define parameters which are passed in.
VERSION=$1
SCALA_VERSION=$2

#define the template.
cat  << EOF
name := "ntypes"
organization := "com.github.piotrkowalczuk"
version := "$VERSION"
scalaVersion := "$SCALA_VERSION"
description := "Set of types that helps to build complex protobuf messages that contains optional properties."
homepage := Some(url("https://github.com/piotrkowalczuk/ntypes"))

scmInfo := Some(
  ScmInfo(
    browseUrl = url("https://github.com/piotrkowalczuk/ntypes"),
    connection = "git@github.com:piotrkowalczuk/ntypes.git")
)

developers := List(
  Developer(
    id = "piotrkowalczuk",
    name = "Piotr Kowalczuk",
    email = "p.kowalczuk.priv@gmail.com",
    url = url("https://github.com/piotrkowalczuk")
  ),
  Developer(
    id = "fpopic",
    name = "Filip Popic",
    email = "filip.popic@gmail.com",
    url = url("https://github.com/fpopic")
  )
)

PB.targets in Compile := Seq(
  scalapb.gen(asciiFormatToString = true) -> (sourceManaged in Compile).value
)

resolvers += Resolver.bintrayRepo("piotrkowalczuk", "maven")

licenses += ("MIT", url("https://github.com/piotrkowalczuk/ntypes/blob/master/LICENSE.txt"))

credentials += Credentials(Path.userHome / ".bintray" / ".credentials")

publishMavenStyle := true

bintrayVcsUrl := Some("git@github.com:piotrkowalczuk/ntypes.git")

EOF