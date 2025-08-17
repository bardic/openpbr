#! /bin/sh

cd ..

mkdir test_out

go run cmd/openvibe/openvibe.go transform \
    --target-assets "blocks" \
    --in "./pkg/cli/test_data/sample/" \
    --out "./test_out" \
    --modifiers "%IN/%FILENAME.%FILEEXT,%OUT/%FILENAME.jpg"