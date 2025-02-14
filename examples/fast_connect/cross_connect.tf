// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "cross_connect_secret_version_cak" {
  default = null
}

variable "cross_connect_secret_version_ckn" {
  default = null
}

resource "oci_core_cross_connect" "cross_connect" {
  #Required
  compartment_id        = var.compartment_ocid
  location_name         = data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations[0].name
  port_speed_shape_name = data.oci_core_cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.cross_connect_port_speed_shapes[1].name

  #Optional
  cross_connect_group_id = oci_core_cross_connect_group.cross_connect_group.id
  display_name           = var.cross_connect_display_name

  #far_cross_connect_or_cross_connect_group_id = oci_core_far_cross_connect_or_cross_connect_group.far_cross_connect_or_cross_connect_group.id
  #near_cross_connect_or_cross_connect_group_id = oci_core_near_cross_connect_or_cross_connect_group.near_cross_connect_or_cross_connect_group.id

  #Set Cross Connect to Active to provision (required to provision virtual circuits).
  #You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up
  is_active = true
}

data "oci_core_cross_connects" "cross_connects" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  cross_connect_group_id = oci_core_cross_connect_group.cross_connect_group.id
  display_name           = var.cross_connect_display_name
  #state = var.cross_connect_state
}

output "cross_connects" {
  value = data.oci_core_cross_connects.cross_connects.cross_connects
}

resource "oci_core_cross_connect" "test_cross_connect_for_macsec" {
  compartment_id        = var.compartment_ocid
  location_name         = data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations[0].name
  port_speed_shape_name = data.oci_core_cross_connect_port_speed_shapes.cross_connect_port_speed_shapes.cross_connect_port_speed_shapes[0].name
  display_name = "MacSecTestForCrossConnect"
  is_active = true
  macsec_properties {
        #Required
        state = "ENABLED"
        #Optional
        encryption_cipher = "AES256_GCM"
        primary_key {
            #Required
            connectivity_association_key_secret_id = var.secret_ocid_cak
            connectivity_association_name_secret_id = var.secret_ocid_ckn
            #Optional, api will always create with current version, but can use to update
            connectivity_association_key_secret_version = var.cross_connect_secret_version_cak
            connectivity_association_name_secret_version = var.cross_connect_secret_version_ckn
        }
        is_unprotected_traffic_allowed = false

    }
}


