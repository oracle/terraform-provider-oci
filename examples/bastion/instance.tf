variable instance_image_ocid {
  type = map(string)
  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.9-2021.04.09-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaanj7qmui2ux5hbiwtbtkzajuvvhuzo2y7755stim22ue6msqwv2ja"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaw2wavtqrd3ynbrzabcnrs77pinccp55j2gqitjrrj2vf65sqj5kq"
  }
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.bastion_ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard1.1"

  agent_config {
    are_all_plugins_disabled  = false
    is_management_disabled    = false
    is_monitoring_disabled    = false
    plugins_config  {
      desired_state = "ENABLED"
      name          = "Bastion"
    }
  }

  create_vnic_details {
    subnet_id        = oci_core_subnet.test_subnet.id
    display_name     = "Primaryvnic"
    assign_public_ip = true
    hostname_label   = "testinstance"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }

  metadata = {
    ssh_authorized_keys = var.session_key_details_public_key_content
  }

  timeouts {
    create = "60m"
  }
}