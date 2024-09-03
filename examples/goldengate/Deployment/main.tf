variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}
variable "test_subnet_id" {}
variable "objectstorage_bucket_name" {}
variable "objectstorage_namespace" {}
variable password_secret_id {}
variable identity_domain_id {}
variable group_id {}

variable "deployment_cpu_core_count" {
  	default = 1
}

variable "deployment_deployment_type" {
  	default = "OGG"
}

variable "deployment_description" {
  	default = "description"
}

variable "deployment_display_name" {
  	default = "TERRAFORM_backward_compatibility_test"
}

variable "deployment_fqdn" {
}

variable "deployment_freeform_tags" {
  	default = { "bar-key" = "value" }
}

variable "deployment_is_auto_scaling_enabled" {
  	default = false
}

variable "deployment_is_public" {
  	default = false
}

variable "deployment_license_model" {
  	default = "LICENSE_INCLUDED"
}

variable "deployment_ogg_data_admin_password" {
  	default = "BEstrO0ng_#11"
}

variable "deployment_ogg_data_admin_username" {
  default = "adminUsername"
}

variable "deployment_ogg_data_deployment_name" {
  default = "deployment"
}

variable "deployment_ogg_data_credential_store" {
	default = "GOLDENGATE"
}

variable "deployment_ogg_data_certificate" {
}

variable "deployment_ogg_data_key" {
}

variable "deployment_state" {
  	default = "ACTIVE"
}

provider "oci" {
  	tenancy_ocid     = var.tenancy_ocid
  	user_ocid        = var.user_ocid
  	fingerprint      = var.fingerprint
  	private_key_path = var.private_key_path
  	region           = var.region
}

resource "oci_golden_gate_deployment_backup" "test_deployment_backup" {
  	#Required
  	bucket         = var.objectstorage_bucket_name
  	compartment_id = var.compartment_id
  	deployment_id  = oci_golden_gate_deployment.test_deployment.id
  	display_name   = var.deployment_display_name
  	namespace      = var.objectstorage_namespace
  	object = "object"
  	lifecycle {
  		ignore_changes = [defined_tags, system_tags, freeform_tags]
  	}
	locks {}
}

resource "oci_golden_gate_deployment" "test_deployment" {
	#Required
  	compartment_id          = var.compartment_id
  	cpu_core_count          = var.deployment_cpu_core_count
  	deployment_type         = var.deployment_deployment_type
  	display_name            = var.deployment_display_name
  	is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
  	license_model           = var.deployment_license_model
  	subnet_id               = var.test_subnet_id
  	ogg_data {
    	admin_password  = var.deployment_ogg_data_admin_password
    	admin_username  = var.deployment_ogg_data_admin_username
    	deployment_name = var.deployment_ogg_data_deployment_name
		group_to_roles_mapping {
			security_group_id = var.group_id
		}
  	}
	locks {}
}

resource "oci_golden_gate_deployment" "test_deployment_GOLDENGATE" {
	compartment_id          = var.compartment_id
	cpu_core_count          = var.deployment_cpu_core_count
	deployment_type         = var.deployment_deployment_type
	display_name            = var.deployment_display_name
	is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
	license_model           = var.deployment_license_model
	subnet_id               = var.test_subnet_id
	ogg_data {
		deployment_name 	= var.deployment_ogg_data_deployment_name
		credential_store 	= var.deployment_ogg_data_credential_store
		admin_username  	= var.deployment_ogg_data_admin_username
		password_secret_id 	= var.password_secret_id
	}
}

resource "oci_golden_gate_deployment" "test_deployment_IAM" {
	count					= var.identity_domain_id != "" ? 1 : 0
	compartment_id          = var.compartment_id
	cpu_core_count          = var.deployment_cpu_core_count
	deployment_type         = var.deployment_deployment_type
	display_name            = var.deployment_display_name
	is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
	license_model           = var.deployment_license_model
	subnet_id               = var.test_subnet_id
	ogg_data {
		deployment_name 	= var.deployment_ogg_data_deployment_name
		credential_store 	= "IAM"
		identity_domain_id  = var.identity_domain_id
	}
}

resource "oci_golden_gate_deployment" "test_deployment_from_backup" {
  	deployment_backup_id = oci_golden_gate_deployment_backup.test_deployment_backup.id

  	compartment_id          = var.compartment_id
  	cpu_core_count          = var.deployment_cpu_core_count
  	deployment_type         = var.deployment_deployment_type
  	display_name            = var.deployment_display_name
  	is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
  	license_model           = var.deployment_license_model
  	subnet_id               = var.test_subnet_id
  	ogg_data {
    	admin_password  = var.deployment_ogg_data_admin_password
    	admin_username  = var.deployment_ogg_data_admin_username
    	deployment_name = var.deployment_ogg_data_deployment_name
  	}
}

data "oci_golden_gate_deployments" "test_deployments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.deployment_display_name
  state        = var.deployment_state
}
