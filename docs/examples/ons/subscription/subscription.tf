// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_ons_subscription" "test_subscription" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  endpoint       = "${var.subscription_endpoint}"
  protocol       = "${var.subscription_protocol}"
  topic_id       = "${oci_ons_notification_topic.test_notification_topic.id}"

  #Optional
  defined_tags  = "${map("${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}", "${var.subscription_defined_tags_value}")}"
  freeform_tags = "${var.subscription_freeform_tags}"
}

data "oci_ons_subscriptions" "test_subscriptions" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  topic_id = "${oci_ons_subscription.test_subscription.topic_id}"
}
