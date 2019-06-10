#!/bin/bash

SDK_PATH='github.com/aws/aws-sdk-go'
VERSION=`cat go.mod | grep "${SDK_PATH}" | cut -d' ' -f 2`
SRC_PATH="${GOPATH}/pkg/mod/${SDK_PATH}@${VERSION}/service"

NAMES=()
for file in `ls ${SRC_PATH}`; do
    if [ -f "${SRC_PATH}/${file}/service.go" ]; then
        # Find the proper casing of each service name. This will
        # always be the name of the struct matching the package name.
        NAMES+=( `cat ${SRC_PATH}/${file}/service.go | grep -i "type ${file} struct" | cut -d' ' -f 2` )
    fi
done

IMPORTS=(
    "${SDK_PATH}/aws/session"
    "github.com/go-nacelle/nacelle"
)

for name in ${NAMES[@]}; do
    IMPORTS+=( "${SDK_PATH}/service/${name,,}" )
done

echo "package awsutil" > services.go
echo "import (`printf '"%s"\n' ${IMPORTS[*]}`)" >> services.go

for name in ${NAMES[@]}; do
    cat <<-EOF >> services.go
    func New${name}ServiceInitializer() nacelle.Initializer {
        return NewServiceInitializer("${name,,}", func(sess *session.Session) interface{} {
            return ${name,,}.New(sess)
        })
    }

EOF
done

go fmt services.go
