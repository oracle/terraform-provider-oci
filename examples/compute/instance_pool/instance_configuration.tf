variable "windows_instance_shape" {
  default = "VM.Standard2.1"
}

resource "oci_core_instance_configuration" "test_instance_config" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    /*
      Attach multiple block volumes
    */
    block_volumes {
      create_details {
        compartment_id       = var.compartment_ocid
        display_name         = "TestCreateVolumeDetails-1"
        availability_domain  = data.oci_identity_availability_domain.ad.name
        size_in_gbs          = 50
        vpus_per_gb          = 20 // min vpus
        is_auto_tune_enabled = false
        block_volume_replicas {
          display_name        = "TestCreateVolumeDetails-1"
          availability_domain = data.oci_identity_availability_domain.ad.name
        }
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails-1"
        is_read_only                        = true
        is_shareable                        = true
      }
    }

    block_volumes {
      create_details {
        compartment_id      = var.compartment_ocid
        display_name        = "TestCreateVolumeDetails-2"
        availability_domain = data.oci_identity_availability_domain.ad.name
        size_in_gbs         = 50
        vpus_per_gb         = 20 // min vpus
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails-2"
        is_read_only                        = true
        is_shareable                        = true
      }
    }

    launch_details {
      compartment_id                      = var.compartment_ocid
      ipxe_script                         = "ipxeScript"
      shape                               = var.windows_instance_shape
      display_name                        = "TestInstanceConfigurationLaunchDetails"
      is_pv_encryption_in_transit_enabled = false
      preferred_maintenance_action        = "LIVE_MIGRATE"
      launch_mode                         = "NATIVE"

      agent_config {
        is_management_disabled = false
        is_monitoring_disabled = false
      }

      availability_config {
        recovery_action             = "RESTORE_INSTANCE"
        is_live_migration_preferred = false
      }

      launch_options {
        network_type = "PARAVIRTUALIZED"
      }

      instance_options {
        are_legacy_imds_endpoints_disabled = false
      }

      create_vnic_details {
        assign_private_dns_record = true
        assign_public_ip       = true
        display_name           = "TestInstanceConfigurationVNIC"
        skip_source_dest_check = false
      }

      licensing_configs {
        type = "WINDOWS"
        license_type = "OCI_PROVIDED"
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = data.oci_core_images.supported_windows_shape_images.images[0]["id"]
      }
    }
  }
}

output "supported_windows_shape_images" {
  value = data.oci_core_images.supported_windows_shape_images.images[0]["id"]
}

# Gets a list of all images that support a given Instance shape
data "oci_core_images" "supported_windows_shape_images" {
  compartment_id   = var.tenancy_ocid
  shape            = var.windows_instance_shape
  operating_system = "Windows"
}