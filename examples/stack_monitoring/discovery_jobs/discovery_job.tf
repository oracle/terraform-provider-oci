// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_management_agent_id_discovery" {}
variable "discovery_job_should_propagate_tags_to_discovered_resources" {
	default = false
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_stack_monitoring_discovery_job" "test_discovery_job" {
	#Required
	compartment_id = var.compartment_ocid
	discovery_client = "LA_SERVICE"
	discovery_type = "ADD"
	should_propagate_tags_to_discovered_resources = var.discovery_job_should_propagate_tags_to_discovered_resources
	discovery_details {
		#Required
		agent_id = var.stack_mon_management_agent_id_discovery
		properties {

			#Optional
			properties_map = {
			    "admin_server_host" = "somehost.us.oracle.com",
			    "admin_server_port" = "7001",
			    "admin_server_protocol" = "t3"
			}
		}
		resource_name = "terraformExample"
		resource_type = "WEBLOGIC_DOMAIN"
		license = "STANDARD_EDITION"

		#Optional
		credentials {
			#Required
			items {
				#Required
				credential_name = "Sk1YQ3JlZHM="
				credential_type = "Sk1YQ3JlZHM="
				properties {

					#Optional
					properties_map = {
					    "Username" = "d2VibG9naWM=",
					    "Password" = "d2VibG9naWM="
					}
				}
			}
		}
	}
	lifecycle {
		ignore_changes = [
			discovery_details, system_tags, defined_tags]
	}
}
data "oci_stack_monitoring_discovery_jobs" "test_discovery_jobs" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_stack_monitoring_discovery_job" "test_discovery_job" {
  #Required
  discovery_job_id = oci_stack_monitoring_discovery_job.test_discovery_job.id
}
