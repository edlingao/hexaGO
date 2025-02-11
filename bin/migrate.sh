#!/bin/sh

if [ ! -f models/main.db ];
  then
    sqlite3 db/main.db < db/schema.sql;
fi

