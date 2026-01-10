# Creating the queue with required parameters and CONSUMER_GROUPS capability
resource "oci_queue_queue" "test_queue" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.queue_display_name

  #Optional
  capabilities {
    type                                                    = "CONSUMER_GROUPS"
    is_primary_consumer_group_enabled                       = var.is_primary_consumer_group_enabled
    primary_consumer_group_display_name                     = var.primary_consumer_group_display_name
    primary_consumer_group_dead_letter_queue_delivery_count = var.primary_consumer_group_dead_letter_queue_delivery_count
    primary_consumer_group_filter                           = var.primary_consumer_group_filter
  }
}

# create a consumer group
resource "oci_queue_consumer_group" "test_consumer_group" {
  #Required
  queue_id     = oci_queue_queue.test_queue.id
  display_name = var.cg_display_name
}

# create another consumer group
resource "oci_queue_consumer_group" "test_consumer_group2" {
  #Required
  queue_id     = oci_queue_queue.test_queue.id
  display_name = var.cg_display_name2
}

data "oci_queue_consumer_groups" "test_consumer_groups" {
  #Optional
  queue_id = oci_queue_queue.test_queue.id
}

# create a queue with required parameters and LARGE_MESSAGES capability
resource "oci_queue_queue" "test_queue2" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.queue_display_name2

  #Optional
  capabilities {
    type = "LARGE_MESSAGES"
  }
}

# create a queue with required parameters and both CONSUMER_GROUPS and LARGE_MESSAGES capability
resource "oci_queue_queue" "test_queue3" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.queue_display_name3

  #Optional
  capabilities {
    type                                                    = "CONSUMER_GROUPS"
    is_primary_consumer_group_enabled                       = var.is_primary_consumer_group_enabled
    primary_consumer_group_display_name                     = var.primary_consumer_group_display_name
    primary_consumer_group_dead_letter_queue_delivery_count = var.primary_consumer_group_dead_letter_queue_delivery_count
    primary_consumer_group_filter                           = var.primary_consumer_group_filter
  }
  capabilities {
    type = "LARGE_MESSAGES"
  }
}