// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable config_file_profile {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "tenancy_ocid" {
}

variable "project_description" {
  default = "description"
}

variable "vault_id" { 
    
}

variable "key_id" { 

}

variable "subnet" {

}

provider "oci" {
  auth = "SecurityToken"
  config_file_profile = var.config_file_profile
  region = var.region
}

resource "oci_managed_kafka_kafka_cluster_config" "test_kafka_cluster_config"{
    compartment_id = var.compartment_ocid
    latest_config {
        properties = {
        "properties" = "properties"
        }
    }
    lifecycle {
        ignore_changes = [defined_tags]
    }
}

data "oci_managed_kafka_kafka_cluster_config_versions" "test_kafka_cluster_config_versions" {
    kafka_cluster_config_id = "${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}"
}

	
data "oci_managed_kafka_kafka_cluster_config_version" "test_kafka_cluster_config_version" {
      kafka_cluster_config_id = "${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}"
      version_number = "1"
}

//DEPENDENCIES
variable defined_tag_namespace_name {
  default = ""
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}

resource "oci_managed_kafka_kafka_cluster_config" "test_kafka_cluster_config_2" {
      compartment_id = var.compartment_ocid
      defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}"
      display_name = "displayName2"
      
      freeform_tags = {
        "Department" = "Accounting"
      }

      latest_config {
        properties = {
            "properties2" = "properties2"
        }
        version_number = "11"
      }

      lifecycle {
          ignore_changes = [defined_tags]
      }
}


resource "oci_managed_kafka_kafka_cluster" "test_kafka_cluster" {
    access_subnets {
        subnets = ["${var.subnet}"]
    }
    broker_shape {
        node_count = "3"
        ocpu_count = "2"
    }
    cluster_config_id = "${oci_managed_kafka_kafka_cluster_config.test_kafka_cluster_config.id}"
    cluster_config_version = "1"
    cluster_type = "DEVELOPMENT"
    compartment_id = "${var.compartment_ocid}"
    coordination_type = "ZOOKEEPER"
    kafka_version = "3.7.0"
    lifecycle {
        ignore_changes = ["defined_tags"]
    }
  
    timeouts {
      create = "60m"
      delete = "60m"
    }
}

resource "oci_vault_secret" "test_secrets" {
    compartment_id = "${var.compartment_ocid}"
    key_id = "${var.key_id}"
    secret_content {
        #Required
        content_type = "BASE64"
        #Optional
        content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    }
    # Secret names should be unique within a vault.if you see an error with 400 name already exists, change name and try again Avoid entering confidential information.
    secret_name = "testsecretname10"
    vault_id = "${var.vault_id}"
}

resource "oci_managed_kafka_kafka_cluster_superusers_management" "test_kafka_cluster_superusers_management" {
    compartment_id = "${var.compartment_ocid}"
    enable_superuser = "true"
    kafka_cluster_id = "${oci_managed_kafka_kafka_cluster.test_kafka_cluster.id}"
    secret_id = "${oci_vault_secret.test_secrets.id}"
}
