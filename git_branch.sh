#!/bin/sh
git  rev-parse --abbrev-ref HEAD | grep -v HEAD || \
git  describe --tags HEAD || \
git  rev-parse HEAD