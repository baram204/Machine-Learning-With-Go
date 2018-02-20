#!/bin/bash

cd docker/jupyter
docker-compose up -d
docker exec jupyter_jupyter_1 jupyter notebook list