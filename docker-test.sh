#!/bin/bash
#
# docker-test.sh
#
# Copyright (c) 2017 Junpei Kawamoto
#
# This file is part of cgss.
#
# cgss is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# cgss is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with cgss.  If not, see <http://www.gnu.org/licenses/>.
#

#
# Run docker based client/server tests.
#
docker run -d --name cgss-server -p 13009:13009 itslabq/cgss

cat << EOS > cgss.yml
groups:
  - name: Group1
    servers:
      - address: 127.0.0.1
        port: 13009
EOS

cd client
go build -o client
cd ../
./client/client remote put cgss.yml 1 1
./client/client remote get cgss.yml --output cgss2.yml

docker kill cgss-server
docker rm cgss-server

[[ -z $(diff cgss.yml cgss2.yml && rm cgss.yml cgss2.yml) ]]
