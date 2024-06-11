#!/bin/bash
# file: makefile.database.sh
# url: https://github.com/conneroisu/go-semantic-router/blob/main/scripts/makefile.database.sh
# title: makefile.database.sh
# description: A script to generate the database schema and models for the cse-ncaa project.
#
# Usage: make database

dbs=(
	"cse"
	"logs"
)

# for each known database
for db in "${dbs[@]}"; do
	awk 'FNR==1{print ""}1' ./data/"$db"/schemas/*.sql > ./data/"$db"/combined/schema.sql
	awk '!/--/' ./data/"$db"/combined/schema.sql > ./temp && mv ./temp ./data/"$db"/combined/schema.sql
	awk 'FNR==1{print ""}1' ./data/"$db"/seeds/*.sql > ./data/"$db"/combined/seeds.sql
	awk 'FNR==1{print ""}1' ./data/"$db"/queries/*.sql > ./data/"$db"/combined/queries.sql
done

for db in "${dbs[@]}"; do
	echo "===== $db ====="
	cd ./data/"$db" || exit
	echo "generating models"
	sqlc generate
	echo "^^^^^ $db ^^^^^"
	cd ../..
	rm ./data/"$db"/db.go
done
