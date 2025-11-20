// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_dbmulticloud_oracle_db_aws_identity_connector" "test_oracle_db_aws_identity_connector" {
	aws_location="us-east1"
	compartment_id=var.compartment_ocid
	display_name="AWS_Tersi_Test"
	issuer_url= "https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com"
	oidc_scope="DBMC/aws"
	resource_id="ocid1.cloudvmcluster.test..awstersitest"
  
   service_role_details {
    role_arn                 = "arn:aws:iam::867344470629:role/OracleDatabaseKMS"
    service_private_endpoint = "https://kms.us-east-1.amazonaws.com"
    service_type             = "KMS"
  }
}   

data"oci_dbmulticloud_oracle_db_aws_identity_connector" "test_oracle_db_aws_identity_connector" {
  oracle_db_aws_identity_connector_id = oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector.id
}

output "aws_connector_id" {
  value = oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector.id
}
