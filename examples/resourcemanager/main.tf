// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// This example shows how to use Resource Manager data sources with the (local) `terraform_remote_state` data source
// so that output values from a different Resource Manager Stack can be referenced in this config.

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

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_resourcemanager_stacks" "stack_list" {
  compartment_id = var.compartment_id
  state          = "ACTIVE"
}

output "print-stacks" {
  value = [data.oci_resourcemanager_stacks.stack_list.stacks]
}

data "oci_resourcemanager_stack" "stack1" {
  stack_id = data.oci_resourcemanager_stacks.stack_list.stacks[0].id
}

output "print-stack1" {
  value = {
    id             = data.oci_resourcemanager_stack.stack1.stack_id
    compartment_id = data.oci_resourcemanager_stack.stack1.compartment_id
    display_name   = data.oci_resourcemanager_stack.stack1.display_name
    state          = data.oci_resourcemanager_stack.stack1.state
  }
}

// pull the statefile of a Resource Manager stack into this context
data "oci_resourcemanager_stack_tf_state" "stack1_tf_state" {
  stack_id   = data.oci_resourcemanager_stack.stack1.stack_id
  local_path = "stack1.tfstate"
}

output "print-tf-state" {
  value = {
    stack_id   = data.oci_resourcemanager_stack_tf_state.stack1_tf_state.stack_id
    local_path = data.oci_resourcemanager_stack_tf_state.stack1_tf_state.local_path
  }
}

// load that statefile into a remote state data source
data "terraform_remote_state" "external_stack_remote_state" {
  backend = "local"

  config = {
    path = data.oci_resourcemanager_stack_tf_state.stack1_tf_state.local_path
  }
}

// example of referencing an output value `subnet_id` from the remote state data source
//data "oci_core_subnet" "subnet1" {
//  subnet_id = data.terraform_remote_state.external_stack_remote_state.outputs.subnet_id
//}
//output print-subnet {
//  value = [data.oci_core_subnet.subnet1]
//}
