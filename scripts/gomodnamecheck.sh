#!/usr/bin/env bash

filename=go.mod

if [[ ! -f go.mod ]]; then
  echo "ERROR: go.mod not found in pwd."
  echo "Please run this from the root of the terraform provider repository"
  exit 1
fi

# read first line of go.mod which declares the module name
read -r firstline<$filename
if [ -z "$firstline" ]; then
  echo "no module name found in go.mod file"
  exit 1;
fi

# parse the module name
module_name=$(echo $firstline | sed "s/^module \(.*\)/\1/")

# check if the module name is fully qualified and tested module name that is discoverable by go get
# github.com/oracle/terraform-provider-oci
echo "discovered module name: $module_name"
pat=".*.com\/.*"
if [[ $module_name =~ $pat ]]; then
  if [ $module_name != "github.com/oracle/terraform-provider-oci" ]; then
      echo "detecting unknown module name for terraform-provider-oci"
      exit 1;
  fi
else
    echo "module name not in format of a qualified name that can be discovered via go get"
    exit 1;
fi