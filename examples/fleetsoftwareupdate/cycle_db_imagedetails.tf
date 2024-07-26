variable "tenancy_ocid" {
}

variable "ssh_public_key" {
}

variable "region" {
}

variable "compartment_id" {
}

variable "db_target_1" {
}

variable "db_target_2" {
}

variable "db_software_image_1" {
}

variable "InstanceImageOCID" {
	type = "map"
	default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
	}
}
// Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
data "oci_core_images" "supported_shape_images" {
	compartment_id   = "${var.tenancy_ocid}"
	shape            = "VM.Standard2.1"
	operating_system = "Oracle Linux"
}

	
resource "oci_fleet_software_update_fsu_collection" "test_fsu_collection" {
	compartment_id = "${var.compartment_id}"
	fleet_discovery {
		strategy = "TARGET_LIST"
		targets = ["${var.db_target_1}"]
	}
	lifecycle {
		ignore_changes = ["defined_tags", "system_tags", "freeform_tags"]
	}
	service_type = "EXACS"
	source_major_version = "DB_19"
	type = "DB"
}

variable defined_tag_namespace_name {
	default = ""
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = "${var.tenancy_ocid}"
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

		is_retired = false
}

resource "oci_fleet_software_update_fsu_cycle" "test_fsu_cycle" {
	compartment_id = "${var.compartment_id}"
	fsu_collection_id = "${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}"
	goal_version_details {
		software_image_id = "${var.db_software_image_1}"
		type = "IMAGE_ID"
	}
	lifecycle {
		ignore_changes = ["defined_tags", "system_tags", "freeform_tags"]
	}
	type = "PATCH"
}
