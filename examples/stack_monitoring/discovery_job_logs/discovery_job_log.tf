// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_management_agent_id_discovery" {}


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
}


data "oci_stack_monitoring_discovery_job_logs" "test_discovery_job_logs" {
  #Required
  discovery_job_id = oci_stack_monitoring_discovery_job.test_discovery_job.id
}