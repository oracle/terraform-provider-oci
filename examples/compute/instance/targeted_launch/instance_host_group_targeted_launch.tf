variable "compute_host_group_id" {
  default = ""
}

resource "oci_core_instance" "test_customer_bare_metal_host_targeted_launch" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestHostGroup"
  shape               = var.shape

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "tlhostgroup"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid
  }

  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }

  timeouts {
    create = "60m"
  }

  placement_constraint_details {
    type = "HOST_GROUP"
    compute_host_group_id = var.compute_host_group_id
  }

}

