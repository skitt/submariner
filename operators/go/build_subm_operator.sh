#!/bin/bash
set -ex

version=${1:-dev}
push_image=${2:-false}
op_gen_dir=$(pwd)
op_out_dir=$op_gen_dir/submariner-operator

function build_subm_operator() {

  pushd $op_out_dir
  go mod vendor

  operator-sdk build quay.io/submariner/submariner-operator:$version --verbose
  if [[ $push_image = true ]]; then
    docker push quay.io/submariner/submariner-operator:$version
  else
    echo "Skipping pushing SubM Operator image to Quay"
  fi

  popd
}

build_subm_operator

