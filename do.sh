#!/bin/bash
for i in {1..10000}
do
  echo $i
   curl --location --request POST 'http://127.0.0.1:8080/url' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "alias": "oneMore",
       "url": "https://abdcd/kkk$i"
   }'
done