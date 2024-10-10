// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


resource "oci_desktops_desktop_pool" "test_desktop_pool" {
  #Required
  compartment_id  = var.compartment_id
  display_name    = var.desktop_pool_display_name
  contact_details = var.desktop_pool_contact_details

  are_privileged_users = var.desktop_pool_are_privileged_users
  availability_domain  = data.oci_identity_availability_domain.ad.name

  network_configuration {
    #Required
    subnet_id = var.desktop_pool_subnet_id
    vcn_id    = var.desktop_pool_vcn_id
  }

  device_policy {
    #Required
    audio_mode          = var.desktop_pool_device_policy_audio_mode
    cdm_mode            = var.desktop_pool_device_policy_cdm_mode
    clipboard_mode      = var.desktop_pool_device_policy_clipboard_mode
    is_display_enabled  = var.desktop_pool_device_policy_is_display_enabled
    is_keyboard_enabled = var.desktop_pool_device_policy_is_keyboard_enabled
    is_pointer_enabled  = var.desktop_pool_device_policy_is_pointer_enabled
    is_printing_enabled = var.desktop_pool_device_policy_is_printing_enabled
  }

  image {
    #Required
    image_id   = var.desktop_pool_image_id
    image_name = var.desktop_pool_image_name
  }

  availability_policy {
    #Required
    start_schedule {
      #Required
      cron_expression = var.desktop_pool_start_schedule_cron_expr
      timezone        = var.desktop_pool_start_schedule_timezone
    }
    stop_schedule {
      #Required
      cron_expression = var.desktop_pool_stop_schedule_cron_expr
      timezone        = var.desktop_pool_stop_schedule_timezone
    }
  }

  is_storage_enabled       = var.desktop_pool_is_storage_enabled
  storage_backup_policy_id = var.desktop_pool_backup_policy_id
  storage_size_in_gbs      = var.desktop_pool_storage_size_in_gbs

  shape_name   = var.desktop_pool_shape_name
  standby_size = var.desktop_pool_standby_size
  maximum_size = var.desktop_pool_maximum_size

  #Optional
  #  defined_tags         = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.desktop_pool_defined_tags_value)
  #  description          = var.desktop_pool_description
  #  freeform_tags        = var.desktop_pool_freeform_tags
  #  nsg_ids              = var.desktop_pool_nsg_ids
  #  time_start_scheduled = var.desktop_pool_time_start_scheduled
  #  time_stop_scheduled  = var.desktop_pool_time_stop_scheduled
  #  private_access_details {
    #    #Required
    #    subnet_id = var.desktop_pool_private_access_details_subnet_id
    #
    #    #Optional
    #    nsg_ids    = var.desktop_pool_private_access_details_nsg_ids
    #    private_ip = var.desktop_pool_private_access_details_private_ip
    #  }
}

data "oci_desktops_desktop_pools" "test_desktop_pools_datasource" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #  availability_domain = data.oci_identity_availability_domain.ad.name
  #  display_name        = var.desktop_pool_display_name
  id                     = oci_desktops_desktop_pool.test_desktop_pool.id
  #  state               = var.desktop_pool_state
}

data "oci_desktops_desktop_pool" "test_desktop_pool_datasource" {
  #Required
  desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id
}

data "oci_desktops_desktop_pool_desktops" "test_desktop_pool_desktops_datasource" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_id
  desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id
}

data "oci_desktops_desktop_pool_volumes" "test_desktop_pool_volumes_datasource" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_id
  desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id
  #Optional
  #  display_name = var.desktop_pool_display_name
  #  state = var.desktop_pool_state
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_desktops_desktops" "test_desktops_datasource" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_id
  desktop_pool_id = oci_desktops_desktop_pool.test_desktop_pool.id
  state = "ACTIVE"
}

data "oci_desktops_desktop" "test_desktop_datasource" {
  desktop_id = data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops_datasource.desktop_pool_desktop_collection.0.items.0.desktop_id
}


