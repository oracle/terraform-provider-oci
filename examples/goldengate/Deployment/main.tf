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
  	default = "displayName"
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

variable "deployment_ogg_data_certificate" {
}

variable "deployment_ogg_data_key" {
}

variable "deployment_state" {
  	default = "ACTIVE"
}

#variable defined_tag_namespace_name { default = "" }

provider "oci" {
  	tenancy_ocid     = var.tenancy_ocid
  	user_ocid        = var.user_ocid
  	fingerprint      = var.fingerprint
  	private_key_path = var.private_key_path
  	region           = var.region
}


resource "oci_core_network_security_group" "test_network_security_group" {
	compartment_id = var.compartment_ocid
	vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
	cidr_block = "10.0.0.0/24"
  display_name = "TestSubnet"
	compartment_id = var.compartment_ocid
	vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
	cidr_block = "10.0.0.0/16"
	compartment_id = var.compartment_ocid
}

resource "oci_golden_gate_deployment_backup" "test_deployment_backup" {
  	#Required
	bucket         = oci_objectstorage_bucket.test_bucket.name
  	compartment_id = var.compartment_ocid
  	deployment_id  = oci_golden_gate_deployment.test_backup_deployment.id
  	display_name   = var.deployment_display_name
  	namespace 	   = data.oci_objectstorage_namespace.test_namespace.namespace
  	object = "object"
  	lifecycle {
		ignore_changes = ["defined_tags", "system_tags", "freeform_tags"]
  	}

}

resource "oci_objectstorage_bucket" "test_bucket" {
	compartment_id = var.compartment_ocid
	name = "tfTestBucketold"
	namespace = data.oci_objectstorage_namespace.test_namespace.namespace
}

resource "oci_golden_gate_deployment" "test_backup_deployment" {
	#Required
  	compartment_id          = var.compartment_ocid
  	cpu_core_count          = var.deployment_cpu_core_count
  	deployment_type         = var.deployment_deployment_type
  	display_name            = var.deployment_display_name
  	is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
  	license_model           = var.deployment_license_model
	subnet_id 				= oci_core_subnet.test_subnet.id
  	ogg_data {
		admin_password  = var.deployment_ogg_data_admin_password
    	admin_username  = var.deployment_ogg_data_admin_username
    	deployment_name = var.deployment_ogg_data_deployment_name
  	}
}

resource "oci_golden_gate_deployment" "test_deployment" {
  #Required
  compartment_id          = var.compartment_ocid
  cpu_core_count          = var.deployment_cpu_core_count
  deployment_type         = var.deployment_deployment_type
  display_name            = var.deployment_display_name
  is_auto_scaling_enabled = var.deployment_is_auto_scaling_enabled
  license_model           = var.deployment_license_model
  subnet_id               = oci_core_subnet.test_subnet.id

  #Optional
  #defined_tags         = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.deployment_defined_tags_value)
  deployment_backup_id = oci_golden_gate_deployment_backup.test_deployment_backup.id
  description          = var.deployment_description
  fqdn                 = var.deployment_fqdn
  freeform_tags        = var.deployment_freeform_tags
  is_public            = var.deployment_is_public
  nsg_ids              = [oci_core_network_security_group.test_network_security_group.id]
  ogg_data {
    #Required
    admin_password  = var.deployment_ogg_data_admin_password
    admin_username  = var.deployment_ogg_data_admin_username
    deployment_name = var.deployment_ogg_data_deployment_name

    #Optional
    #certificate = var.deployment_ogg_data_certificate
    #key         = var.deployment_ogg_data_key
  }
}

data "oci_objectstorage_namespace" "test_namespace" {
}

data "oci_golden_gate_deployments" "test_deployments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.deployment_display_name
  state        = var.deployment_state
}