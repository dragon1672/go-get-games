#!/bin/sh

DIR=$(dirname "$0")
FILE=bundled.go
BIN=$(go env GOPATH)/bin

cd "$DIR" || exit
rm $FILE

# Sourced from https://www.svgrepo.com/collection/army-badges-5/ & https://www.svgrepo.com/collection/carbon-design-line-icons/20
# Install fyne with `$ go get fyne.io/fyne/v2/cmd/fyne`
{
  $BIN/fyne bundle -package assets -name EmptyIcon empty.svg;
  $BIN/fyne bundle -package assets -name BombIcon -append bomb.svg;
  $BIN/fyne bundle -package assets -name SignalIcon -append signal.svg;
  $BIN/fyne bundle -package assets -name TankIcon -append tank.svg;
  $BIN/fyne bundle -package assets -name TargetIcon -append target.svg;
  $BIN/fyne bundle -package assets -name N0Icon -append n0.svg;
  $BIN/fyne bundle -package assets -name N1Icon -append n1.svg;
  $BIN/fyne bundle -package assets -name N2Icon -append n2.svg;
  $BIN/fyne bundle -package assets -name N3Icon -append n3.svg;
  $BIN/fyne bundle -package assets -name N4Icon -append n4.svg;
  $BIN/fyne bundle -package assets -name N5Icon -append n5.svg;
  $BIN/fyne bundle -package assets -name N6Icon -append n6.svg;
  $BIN/fyne bundle -package assets -name N7Icon -append n7.svg;
  $BIN/fyne bundle -package assets -name N8Icon -append n8.svg;
  $BIN/fyne bundle -package assets -name N9Icon -append n9.svg;
} > $FILE
