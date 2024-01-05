// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "random_string" "projectname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_ocid
  name           = join("", ["T",  random_string.topicname.result])
}

resource "oci_devops_project" "test_project" {
  #Required
  compartment_id = var.compartment_ocid
  name           = join("", ["A", random_string.projectname.result])
  notification_config {
    #Required
    topic_id = oci_ons_notification_topic.test_notification_topic.id
  }
}

resource "oci_devops_deploy_pipeline" "test_deploy_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  description   = "description"
  display_name  = "displayName"
}

resource "oci_devops_deploy_artifact" "test_deploy_command_spec_artifact" {
  project_id              = oci_devops_project.test_project.id
  display_name = "Display_name"
  deploy_artifact_type = "COMMAND_SPEC"
  argument_substitution_mode = "NONE"
  deploy_artifact_source {
    deploy_artifact_source_type = "INLINE"
    base64encoded_content       = "YXBpVmVyc2lvbjogYmF0Y2gvdjEKa2luZDogSm9iCm1ldGFkYXRhOgogIGdlbmVyYXRlTmFtZTogaGVsbG93b3JsZAogIGxhYmVsczoKICAgIGFwcDogaGVsbG93b3JsZApzcGVjOgogIHR0bFNlY29uZHNBZnRlckZpbmlzaGVkOiAxMjAKICB0ZW1wbGF0ZToKICAgIHNwZWM6CiAgICAgIGNvbnRhaW5lcnM6CiAgICAgICAgLSBuYW1lOiBoZWxsb3dvcmxkCiAgICAgICAgICBpbWFnZTogcGh4Lm9jaXIuaW8vYXgwMjJ3dmdtanBxL2hlbGxvd29ybGQtb2tlLXZlcmlmaWVyOmxhdGVzdAogICAgICAgICAgY29tbWFuZDoKICAgICAgICAgICAgLSAiL2Jpbi9iYXNoIgogICAgICAgICAgICAtICItYyIKICAgICAgICAgICAgLSAic2xlZXAgMjsgZWNobyBIZWxsbyBXb3JsZDsiCiAgICAgIHJlc3RhcnRQb2xpY3k6IE5ldmVy"
  }
}

data "oci_identity_availability_domain" "devops_ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = lower(data.oci_identity_availability_domain.devops_ad.name)
  cidr_block = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  //defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "MySubnet"
  dns_label = "dnslabel"
  freeform_tags = {
    "Department" = "Finance"
  }
  lifecycle {
    ignore_changes = [
      defined_tags]
  }
  prohibit_internet_ingress = "false"
  prohibit_public_ip_on_vnic = "false"
  route_table_id = oci_core_vcn.test_vcn.default_route_table_id
  security_list_ids = [
    oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  //defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
  display_name = "displayName"
  dns_label = "dnslabel"
  freeform_tags = {
    "Department" = "Finance"
  }
  lifecycle {
    ignore_changes = [
      defined_tags]
  }
}

resource "oci_devops_deploy_stage" "test_shell_deploy_stage" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
    }
  }
  deploy_stage_type = "SHELL"
  command_spec_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_command_spec_artifact.id
  timeout_in_seconds = 36000
  container_config {
    container_config_type = "CONTAINER_INSTANCE_CONFIG"
    shape_name = "CI.Standard.E4.Flex"
    compartment_id = var.compartment_ocid
    availability_domain = data.oci_identity_availability_domain.devops_ad.name
    shape_config {
      ocpus = 1.0
      memory_in_gbs = 2.0
    }
    network_channel {
      network_channel_type = "SERVICE_VNIC_CHANNEL"
      subnet_id = oci_core_subnet.test_subnet.id
      nsg_ids = []
    }
  }
}