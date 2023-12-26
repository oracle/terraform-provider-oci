// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "knowledge_base_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "knowledge_base_defined_tags" {
  default  = { "example-tag-namespace-all.example-tag" = "value" }
}

variable "knowledge_base_id" {
  default = "id"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_adm_knowledge_base" "example_knowledge_base" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name  = "Example_Knowledge_Base"
}

resource "oci_adm_vulnerability_audit" "example_vulnerability_audit" {
  #Required
  knowledge_base_id = oci_adm_knowledge_base.example_knowledge_base.id

  #Optional
  build_type = "MAVEN"
  compartment_id = var.compartment_ocid
  application_dependencies {
    gav = "com.google.guava:guava:29.0-jre"
    node_id = "node_id"
    #Optional
    application_dependency_node_ids = []
  }
  #Optional
  usage_data {
    namespace   = "id758to84zun"
    bucket      = "usaga-data-TERSI-2463"
    object      = "usage-data-with-vulnerable-class-method.json.gz"
    source_type = "objectStorageTuple"
  }

  source {
    type = "OCI_RESOURCE"
    oci_resource_id = "ocid1.example.ocid"
  }
  display_name = "Example_Vulnerability_Audit"
}

resource "oci_adm_vulnerability_audit" "example_vulnerability_audit_polyglot" {
  #Required
  knowledge_base_id = oci_adm_knowledge_base.example_knowledge_base.id

  #Optional
  build_type = "UNSET"

  application_dependencies {
    gav = ""
    purl = "pkg:deb/ubuntu/openjdk-6@6b30?distro=14.04"
    node_id = "node_id"
    #Optional
    application_dependency_node_ids = []
  }

  source {
    type = "OCI_RESOURCE"
    oci_resource_id = "ocid1.example.ocid"
  }
  display_name = "Example_Polyglot_Vulnerability_Audit"
}

data "oci_adm_knowledge_base" "example_knowledge_base" {
  knowledge_base_id = oci_adm_knowledge_base.example_knowledge_base.id
}

data "oci_adm_knowledge_bases" "example_knowledge_bases" {
  compartment_id = var.compartment_ocid
}

data "oci_adm_vulnerability_audit" "example_vulnerability_audit" {
  vulnerability_audit_id = oci_adm_vulnerability_audit.example_vulnerability_audit.id
}

data "oci_adm_vulnerability_audit_application_dependency_vulnerabilities" "example_application_dependency_vulnerabilities" {
  vulnerability_audit_id = oci_adm_vulnerability_audit.example_vulnerability_audit.id
  sort_by = "dfs"
  root_node_id = "node_id"
  depth = 0
}

data "oci_adm_vulnerability_audit_application_dependency_vulnerabilities" "example_application_dependency_vulnerabilities_with_depth" {
  vulnerability_audit_id = oci_adm_vulnerability_audit.example_vulnerability_audit.id
  gav = "com.google.guava:guava:29.0-jre"
  cvss_v2greater_than_or_equal = "1.5"
}

data "oci_adm_vulnerability_audit_application_dependency_vulnerabilities" "example_application_dependency_vulnerabilities_polyglot" {
  vulnerability_audit_id = oci_adm_vulnerability_audit.example_vulnerability_audit_polyglot.id
  purl = "pkg:deb/ubuntu/openjdk-6@6b30?distro=14.04"
  severity_greater_than_or_equal = "LOW"
}
