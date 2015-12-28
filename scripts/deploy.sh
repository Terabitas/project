#!/bin/bash -e

if [[ $1 =~ ^v[0-9]+(\.[0-9]+)* ]]; then

echo "Deploy tag [$1]"

else

echo "Deploy branch [$1.$2]"

fi



