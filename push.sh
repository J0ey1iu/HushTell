#!/bin/bash

git push
docker build . -t joeyliu086/hushtell
docker push joeyliu086/hushtell