// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_desktops_desktop_pool" "test_desktop_pool_datasource" {
  #Required
  desktop_pool_id = var.desktop_pool_id
}

data "oci_desktops_desktop_pool_desktops" "test_desktop_pool_desktops_datasource" {
  #Required
  compartment_id = var.compartment_id
  desktop_pool_id = var.desktop_pool_id

  #Optional
  #availability_domain = data.oci_identity_availability_domain.ad.name
}

data "oci_desktops_desktops" "test_desktops_datasource" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #availability_domain = data.oci_identity_availability_domain.ad.name
  #state = "ACTIVE"
  #desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool_datasource.id
  #desktop_pool_id = var.desktop_pool_id
}

data "oci_desktops_desktop" "test_desktop_datasource" {
  #Required
  desktop_id = var.desktop_id
}
