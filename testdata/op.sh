#!/bin/sh

cd "${0%/*}" || exit

cmd=$(awk 'NR==1' < op_response)
status=$(awk 'NR==2' < op_response)
response=$(awk 'NR>=3' < op_response)


if [ "$cmd" != "$*" ]; then
    echo "unexpected cmd: $*" >&2
    exit 42
fi

if [ "$status" -eq 0 ]; then
    echo "$response"
else
    echo "$response" >&2
fi

exit "$status"
