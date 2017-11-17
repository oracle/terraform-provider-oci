#!/bin/bash

OCI_OS_DISTRO="darwin.tar.gz"
#ENV OCI_OS_DISTRO="linux.tar.gz"

LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/oracle/terraform-provider-oci/releases/latest)
LATEST_VERSION=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
ARTIFACT_URL="https://github.com/oracle/terraform-provider-oci/releases/download/$LATEST_VERSION/$OCI_OS_DISTRO"

wget -q -O terraform-oci.tar.gz "${ARTIFACT_URL}"

tar xzf terraform-oci.tar.gz -C ~/.terraform.d/plugins/
