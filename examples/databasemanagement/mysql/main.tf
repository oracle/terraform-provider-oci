// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
 
variable "tenancy_ocid" {}
variable "start_time" {}
variable "end_time" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "digest"{}
variable "test_managed_database_id"{}


provider "oci" {
 // version = "5.36.0"
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {  

 default =  "ocid1.test.oc1..<unique_ID>EXAMPLE-compartment-Value"
}

# List managed MySQL database resources in a compartment
data "oci_database_management_managed_my_sql_databases" "test_managed_my_sql_databases" {
  #Required
  compartment_id = var.compartment_id
}

# Get managed MySQL database resource
data "oci_database_management_managed_my_sql_database" "test_managed_my_sql_database" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
}

# Get configuration data for a managed MySQL database resource
data "oci_database_management_managed_my_sql_database_configuration_data" "test_managed_my_sql_database_configuration_data" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
}

# Get SQL data for a managed MySQL database resource
data "oci_database_management_managed_my_sql_database_sql_data" "test_managed_my_sql_database_sql_data" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
  filter_column = "COUNT_STAR"
  
  //  The start_time and end_time of the time range to retrieve the SQL data of a Managed Database
  //in UTC in ISO-8601 format, for example "2024-04-14T17:23:13.000Z". 
  start_time = var.start_time    
  end_time = var.end_time

 
}

data "oci_database_management_managed_my_sql_database_outbound_replications" "test_managed_my_sql_database_outbound_replications" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.3.id
}

data "oci_database_management_managed_my_sql_database_inbound_replications" "test_managed_my_sql_database_inbound_replications" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
}

data "oci_database_management_managed_my_sql_database_high_availability_members" "test_managed_my_sql_database_high_availability_members" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.3.id
}

data "oci_database_management_managed_my_sql_database_general_replication_information" "test_managed_my_sql_database_general_replication_information" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.3.id
}

data "oci_database_management_managed_my_sql_database_binary_log_information" "test_managed_my_sql_database_binary_log_information" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.3.id
}

data "oci_database_management_managed_my_sql_database_query_detail" "test_managed_my_sql_database_query_detail" {
  #Required
 // digest                     = var.digest
  digest  = "8e6a579423a4d2efea0a5884b7dbb72697023a99cc8e1922c3a6cbc355ab4487"
  managed_my_sql_database_id = var.test_managed_database_id
  //managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.3.id
}


data "oci_database_management_managed_my_sql_database_digest_errors" "test_managed_my_sql_database_digest_errors" {
  #Required
  //digest                     = var.digest
  digest  = "8e6a579423a4d2efea0a5884b7dbb72697023a99cc8e1922c3a6cbc355ab4487"
   managed_my_sql_database_id = var.test_managed_database_id
  //managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.3.id
}




