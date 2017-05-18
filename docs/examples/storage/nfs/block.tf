resource "baremetal_core_volume" "TFBlock0" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}" 
  compartment_id = "${var.compartment_ocid}"
  display_name = "2TB NFS"
  size_in_mbs = "${var.2TB}"
}

resource "baremetal_core_volume_attachment" "TFBlock0Attach" {
    attachment_type = "iscsi"
    compartment_id = "${var.compartment_ocid}"
    instance_id = "${baremetal_core_instance.TFInstance.id}" 
    volume_id = "${baremetal_core_volume.TFBlock0.id}"
}

