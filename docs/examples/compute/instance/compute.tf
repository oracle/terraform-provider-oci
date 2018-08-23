resource "oci_core_instance" "TFInstance" {
  count               = "${var.NumInstances}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "TFInstance${count.index}"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.ExampleSubnet.id}"
    display_name     = "primaryvnic"
    assign_public_ip = true
    hostname_label   = "tfexampleinstance${count.index}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"

    # Apply this to set the size of the boot volume that's created for this instance.
    # Otherwise, the default boot volume size of the image is used.
    # This should only be specified when source_type is set to "image".
    #boot_volume_size_in_gbs = "60"
  }

  # Apply the following flag only if you wish to preserve the attached boot volume upon destroying this instance
  # Setting this and destroying the instance will result in a boot volume that should be managed outside of this config.
  # When changing this value, make sure to run 'terraform apply' so that it takes effect before the resource is destroyed.
  #preserve_boot_volume = true

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data           = "${base64encode(file(var.BootStrapFile))}"
  }
  defined_tags = "${
    map(
      "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag2.name}", "awesome-app-server"
    )
  }"
  freeform_tags = "${map("freeformkey${count.index}", "freeformvalue${count.index}")}"
  timeouts {
    create = "60m"
  }
}
