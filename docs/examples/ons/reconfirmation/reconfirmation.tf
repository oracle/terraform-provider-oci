data "oci_ons_reconfirmation" "test_reconfirmation" {
  #Required
  subscription_id = "${oci_ons_subscription.test_subscription.id}"
}

output confirmation_url {
  value = "${data.oci_ons_reconfirmation.test_reconfirmation.url}"
}
