#!/usr/bin/env bash
# Set resource_type (permanent/transient) and stack_id to run this script
#export resource_type=transient
#export stack_id=123

cd ../../../oci
go test -v -run TestResourceDiscoveryUpdateStack