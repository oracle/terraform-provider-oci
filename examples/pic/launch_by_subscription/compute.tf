// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_instance" "pic_instance" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "pic_instance"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.pic_subnet.id}"
    display_name     = "picprimaryvnic"
    assign_public_ip = true
    hostname_label   = "PICInstance"
  }

  source_details {
    source_type = "image"
    source_id   = "${lookup(data.oci_core_app_catalog_subscriptions.test_app_catalog_subscriptions.app_catalog_subscriptions[0], "listing_resource_id")}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  timeouts {
    create = "60m"
  }
}
