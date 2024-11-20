variable "idcs_open_id" {
}

variable "compartment_id" {
}

variable "vb_instance_network_endpoint_details_network_endpoint_type" {
  default = "PRIVATE"
}

variable "vb_instance_network_endpoint_details_network_security_group_ids" {
  default = []
}

variable "vb_instance_network_endpoint_details_private_endpoint_ip" {
  default = ""
}

variable "vb_instance_network_endpoint_details_subnet_id" {
  default = "subnetId"
}

resource "oci_visual_builder_vb_instance" "test_vb_instance_pe" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "displayNamePe"
  is_visual_builder_enabled = "true"
  idcs_open_id              = var.idcs_open_id
  node_count                = "1"

  network_endpoint_details {
    #Required
    network_endpoint_type = var.vb_instance_network_endpoint_details_network_endpoint_type
    subnet_id             = var.vb_instance_network_endpoint_details_subnet_id

    #Optional
    network_security_group_ids = var.vb_instance_network_endpoint_details_network_security_group_ids
    private_endpoint_ip        = var.vb_instance_network_endpoint_details_private_endpoint_ip
  }

}

data "oci_visual_builder_vb_instances" "test_vb_instance_pe" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = "displayNamePe"
  state        = "Active"
  filter {
    name = "id"
    values = [oci_visual_builder_vb_instance.test_vb_instance_pe.id]
  }
}

data "oci_visual_builder_vb_instance" "test_vb_instance_pe" {
  #Required
  vb_instance_id = oci_visual_builder_vb_instance.test_vb_instance_pe.id
}