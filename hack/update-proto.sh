#!/usr/bin/env bash

# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

script_dir=$(dirname $0 | xargs realpath)


cd pkg/apis/topology

go-to-protobuf \
  --output-base=. \
  --go-header-file "$script_dir/boilerplate.go.txt" \
  --proto-import vendor \
  --packages +./v1alpha1=v1alpha1 \
  --keep-gogoproto=false \
  --apimachinery-packages "-k8s.io/apimachinery/pkg/util/intstr"

# Hack to get the go_package option right
sed s',go_package =.*,go_package = "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1";,' \
  -i v1alpha1/generated.proto
