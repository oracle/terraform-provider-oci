// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_generative_ai_dedicated_ai_cluster" "test_hosting_cluster" {
  #Required
  type                           = "HOSTING"
  compartment_id                 = var.compartment_ocid
  unit_count                     = var.hosting_cluster_unit_count
  unit_shape                     = var.hosting_cluster_shape

  #Optional
  display_name                  = var.hosting_cluster_display_name
  description                   = var.hosting_cluster_description
  #defined_tags not tested - cannot test in home region        
  freeform_tags                = var.test_freeform_tags
}

resource "oci_generative_ai_dedicated_ai_cluster" "test_fine_tuning_cluster" {
  #Required
  type                           = "FINE_TUNING"
  compartment_id                 = var.compartment_ocid
  unit_count                     = var.fine_tuning_cluster_unit_count
  unit_shape                     = var.fine_tuning_cluster_shape

  #Optional
  display_name                  = var.fine_tuning_cluster_display_name
  description                   = var.fine_tuning_cluster_description
  #defined_tags not tested - cannot test in home region   
  freeform_tags                = var.test_freeform_tags
}

data "oci_generative_ai_dedicated_ai_cluster" "test_hosting_cluster" {
  #Required
  dedicated_ai_cluster_id       = oci_generative_ai_dedicated_ai_cluster.test_hosting_cluster.id
}

data "oci_generative_ai_dedicated_ai_cluster" "test_fine_tuning_cluster" {
  #Required
  dedicated_ai_cluster_id       = oci_generative_ai_dedicated_ai_cluster.test_fine_tuning_cluster.id
}

data "oci_generative_ai_dedicated_ai_clusters" "test_clusters" {
  #Required
  compartment_id                = var.compartment_ocid
}

