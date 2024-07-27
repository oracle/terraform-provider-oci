
	# Need to have this block even though it's empty; for import testing
	provider "oci" {
	}

	variable "tenancy_ocid" {
	}

	variable "compartment_id" {
    }

    variable "instance_id" {
    }

    variable "instance_maintenance_event_id" {
    }

	variable "ssh_public_key" {
	}

	variable "region" {
		default = "us-ashburn-1"
	}


data "oci_core_instance_maintenance_events" "test_instance_maintenance_event" {
  compartment_id = var.compartment_id
  filter {
    name = "id"
    values = [oci_core_instance_maintenance_event.test_instance_maintenance_event.id]
  }
  instance_action = "REBOOT_MIGRATION"
  instance_id = var.instance_id
  state = "SCHEDULED"
}

		// Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
		data "oci_core_images" "supported_shape_images" {
			compartment_id   = var.tenancy_ocid
			shape            = "VM.Standard2.1"
			operating_system = "Oracle Linux"
		}


resource "oci_core_instance_maintenance_event" "test_instance_maintenance_event" {
  can_delete_local_storage = "true"
  display_name = "displayName2"
  freeform_tags = {
    "Department" = "Accounting"
  }
  instance_maintenance_event_id = var.instance_maintenance_event_id
  time_window_start = "2025-01-12T15:04:05Z"
}
