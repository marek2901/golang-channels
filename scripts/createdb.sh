#!/usr/bin/env bash

if !(which sqlite3 &> /dev/null); then
    echo missing sqlite binary
    exit 1
fi

rm ../lol.db &> /dev/null
sqlite3 ../lol.db "CREATE TABLE CSV_RECORDS(ID INTEGER PRIMARY KEY AUTOINCREMENT, YEAR TEXT NOT NULL, REGION TEXT NOT NULL, CONSUMPTION_TYPE TEXT NOT NULL, CONSUMPTION TEXT NOT NULL);"
