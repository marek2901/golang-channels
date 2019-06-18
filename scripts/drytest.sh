#!/usr/bin/env bash

echo SELECT EM ALL
sqlite3 ../lol.db "SELECT * FROM CSV_RECORDS;"

echo COUNT
sqlite3 ../lol.db "SELECT COUNT(*) FROM CSV_RECORDS;"

echo YAYY
