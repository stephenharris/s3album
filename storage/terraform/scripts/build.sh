#!/usr/bin/env bash

python3 -m venv /tmp/thumbnailgenerator/venv
source /tmp/thumbnailgenerator/venv/bin/activate

pip install Pillow

curdir=$(pwd)
cd /tmp/thumbnailgenerator/venv/lib/python3.9/site-packages
zip -r $curdir/thumbnail_generator.zip .
cd -
zip -g thumbnail_generator.zip thumbnail_generator.py