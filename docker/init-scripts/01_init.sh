#! /bin/bash

sed "s/mytable/$TABLE_NAME/g" /tmp/init_template.sql > /tmp/init.sql
chmod 644 /tmp/init.sql

psql -U $POSTGRES_USER -d $POSTGRES_DB -a -f /tmp/init.sql
