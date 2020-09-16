// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_id" {
}

variable "ssh_public_key" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_exadata_infrastructure" "test_exadata_infrastructure" {
  #Required
  admin_network_cidr          = "192.168.0.0/16"
  cloud_control_plane_server1 = "192.168.19.1"
  cloud_control_plane_server2 = "192.168.19.2"
  compartment_id              = var.compartment_id
  display_name                = "tstExaInfra"
  dns_server                  = ["192.168.10.10"]
  gateway                     = "192.168.20.1"
  infini_band_network_cidr    = "10.172.0.0/19"
  netmask                     = "255.255.0.0"
  ntp_server                  = ["192.168.10.20"]
  shape                       = "ExadataCC.Quarter3.100"
  time_zone                   = "US/Pacific"
  activation_file             = "activation.zip"

  #Optional
  corporate_proxy = "http://192.168.19.1:80"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

data "oci_database_exadata_infrastructure_download_config_file" "test_exadata_infrastructure_download_config_file" {
  #Required
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  #Optional
  base64_encode_content = true
}

data "oci_database_exadata_infrastructures" "test_exadata_infrastructures" {
  #Required
  compartment_id = var.compartment_id
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "tagNamespace1"
  name           = "testexamples-tag-namespace1"
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "local_file" "test_exadata_infrastructure_downloaded_config_file" {
  content  = data.oci_database_exadata_infrastructure_download_config_file.test_exadata_infrastructure_download_config_file.content
  filename = "${path.module}/exadata_infrastructure_config.zip"
}

