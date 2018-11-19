#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/philbrookes/finalizer-operator/finalizer-operator/pkg/generated \
github.com/philbrookes/finalizer-operator/finalizer-operator/pkg/apis \
philbrookes:finalizer \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
