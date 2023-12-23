#!/bin/bash

parent_dir=$(dirname "$PWD")

cd "$parent_dir"

env_file_path="$parent_dir/.env"

docker-compose --env-file "$env_file_path" up -d

go_app_path="$parent_dir"

go run "$go_app_path"
