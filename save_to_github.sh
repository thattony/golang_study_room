#!/bin/bash

# current path to all golang tutorials
GO_TUT_PATH=/home/thattony/Documents/golang__projects

# add && commit with message
git add * && git commit -m $1 

git push -u origin master
