provider "oci" {
  tenancy_ocid     = var.compartment_ocid
  region           = var.region
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  retry_duration_seconds = "1800"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_compute_capacity_report" "test_compute_capacity_report_flex_shape" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  shape_availabilities {
    #Required
    instance_shape = var.compute_capacity_report_shape_availabilities_instance_shape_flex

    #Optional
    fault_domain = var.compute_capacity_report_shape_availabilities_fault_domain
    instance_shape_config {

      #Optional
      memory_in_gbs = var.compute_capacity_report_shape_availabilities_instance_shape_config_memory_in_gbs
      nvmes         = var.compute_capacity_report_shape_availabilities_instance_shape_config_nvmes
      ocpus         = var.compute_capacity_report_shape_availabilities_instance_shape_config_ocpus
    }
  }
}

resource "oci_core_compute_capacity_report" "test_compute_capacity_report_fix_shape" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  shape_availabilities {
    #Required
    instance_shape = var.compute_capacity_report_shape_availabilities_instance_shape_fix
  }
}