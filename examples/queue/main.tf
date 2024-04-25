// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Creating tag namespace and tag for defined tag
resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.tag_namespace_description
  name           = var.tag_namespace_name
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

# Creating the queue with all the optional parameters
resource "oci_queue_queue" "test_queue1" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.queue_display_name

  #Optional
  custom_encryption_key_id         = var.queue_custom_encryption_key_id # We can have dependency on the oci_kms_key and get the key id from that
  dead_letter_queue_delivery_count = var.queue_dead_letter_queue_delivery_count
  purge_trigger                    = var.purge_trigger
  purge_type                       = var.purge_type
  freeform_tags                    = var.queue_freeform_tags
  retention_in_seconds             = var.queue_retention_in_seconds
  timeout_in_seconds               = var.queue_timeout_in_seconds
  visibility_in_seconds            = var.queue_visibility_in_seconds
  channel_consumption_limit        = var.queue_channel_consumption_limit
}

# Purging the queue immediately after create if required. Queue is purged if purge trigger is set to any integer value. We are using the purge trigger and purge type optional parameter.
# We are purging the queue immediately after create.
resource "oci_queue_queue" "test_queue2" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.queue_display_name

  #Optional
  purge_trigger = 1
  purge_type = "normal"

}

# Normal queue creation if purge trigger parameter is unset. This will not trigger the purge queue operation. In addition, presence of purge type if purge trigger is unset is a no-op.
resource "oci_queue_queue" "test_queue3" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.queue_display_name

  #Optional
  purge_type = "normal"
}

data "oci_queue_queues" "test_queues" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = oci_queue_queue.test_queue1.display_name
  id             = oci_queue_queue.test_queue1.id
  state          = var.queue_state
}