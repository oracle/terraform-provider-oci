// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "idcs_access_token" {
}


provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_analytics_analytics_instance" "test_oce_instance_public" {
  compartment_id     = var.compartment_ocid
  description        = "OAC instance"
//  email_notification = var.email_notification
  feature_set        = "ENTERPRISE_ANALYTICS"
  license_type       = "LICENSE_INCLUDED"

  capacity {
    capacity_type  = "OLPU_COUNT"
    capacity_value = 2
  }

  name = "testoacinstance1"
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }
  state             = "ACTIVE"
  idcs_access_token = "eyJ4NXQjUzI1NiI6IjJ6d1lzd3VFenh3ZHdJdnRZakxZOG93Ul9MNzNRZGFjTVpmdnY1RlBPX0kiLCJ4NXQiOiJUX1JHWjhJRC1CQVpRc0h5YTFhQ19UUHRyTmciLCJraWQiOiJTSUdOSU5HX0tFWSIsImFsZyI6IlJTMjU2In0.eyJ1c2VyX3R6IjoiQW1lcmljYVwvQ2hpY2FnbyIsInN1YiI6InBhcm5lZXQucC5rYXVyQG9yYWNsZS5jb20iLCJ1c2VyX2xvY2FsZSI6ImVuIiwiaWRwX25hbWUiOiJsb2NhbElEUCIsInVzZXIudGVuYW50Lm5hbWUiOiJpZGNzLTlmZWIzZWY5MWNhNDQ0MWJiZWRhZGJjNTY0MGQ1ZmI4Iiwib25CZWhhbGZPZlVzZXIiOnRydWUsImlkcF9ndWlkIjoibG9jYWxJRFAiLCJhbXIiOlsiVVNFUk5BTUVfUEFTU1dPUkQiXSwiaXNzIjoiaHR0cHM6XC9cL2lkZW50aXR5Lm9yYWNsZWNsb3VkLmNvbVwvIiwidXNlcl90ZW5hbnRuYW1lIjoiaWRjcy05ZmViM2VmOTFjYTQ0NDFiYmVkYWRiYzU2NDBkNWZiOCIsImNsaWVudF9pZCI6IjMxMjczZThiNWRiYjRkZTM4YzNiMDg2ODkxNTM5NDBhIiwidXNlcl9pc0FkbWluIjp0cnVlLCJzdWJfdHlwZSI6InVzZXIiLCJzY29wZSI6InVybjpvcGM6aWRtOmcuaWRlbnRpdHlzb3VyY2V0ZW1wbGF0ZV9yIHVybjpvcGM6aWRtOnQuZ3JvdXBzLm1lbWJlcnMgdXJuOm9wYzppZG06dC5hcHAgdXJuOm9wYzppZG06dC51c2VyLmxvY2tlZHN0YXRlY2hhbmdlciB1cm46b3BjOmlkbTp0LmlkYnJpZGdlLmFkbWluIHVybjpvcGM6aWRtOnQudGVybXNvZnVzZSB1cm46b3BjOmlkbTp0LmlkY3NycHRzIHVybjpvcGM6aWRtOnQucmVxdWVzdHMgdXJuOm9wYzppZG06dC51c2VyLm1hbmFnZXIgdXJuOm9wYzppZG06dC5oZWxwZGVzay5zZWN1cml0eSB1cm46b3BjOmlkbTp0LnNlY3VyaXR5LmNsaWVudCB1cm46b3BjOmlkbTpnLmFwcHRlbXBsYXRlX3IgdXJuOm9wYzppZG06dC5idWxrLnVzZXIgdXJuOm9wYzppZG06dC5kaWFnbm9zdGljc19yIHVybjpvcGM6aWRtOnQuaWRiX2NvbnRhaW5lcnMgdXJuOm9wYzppZG06dC5pZGJyaWRnZS51c2VyIHVybjpvcGM6aWRtOnQudXNlci5tZSB1cm46b3BjOmlkbTpnLmFsbF9yIHVybjpvcGM6aWRtOnQudXNlci5zZWN1cml0eSB1cm46b3BjOmlkbTp0Lmdyb3Vwc19yIHVybjpvcGM6aWRtOnQuYXVkaXRfciB1cm46b3BjOmlkbTp0LmpvYi5hcHAgdXJuOm9wYzppZG06dC5vYXV0aGNvbnNlbnRzIHVybjpvcGM6aWRtOnQudXNlcnNfciB1cm46b3BjOmlkbTp0LnNvbWkgdXJuOm9wYzppZG06Zy5zaGFyZWRmaWxlcyB1cm46b3BjOmlkbTp0LmhlbHBkZXNrLnVzZXIgdXJuOm9wYzppZG06dC5yZXMuaW1wb3J0ZXhwb3J0IHVybjpvcGM6aWRtOnQuam9iLmlkZW50aXR5IHVybjpvcGM6aWRtOnQuY3VzdG9tY2xhaW1zIHVybjpvcGM6aWRtOnQuc2FtbCB1cm46b3BjOmlkbTp0Lm1mYSB1cm46b3BjOmlkbTp0LmRiLmFkbWluIHVybjpvcGM6aWRtOnQuc2NoZW1hcyB1cm46b3BjOmlkbTp0Lm1mYS51c2VyYWRtaW4gdXJuOm9wYzppZG06dC51c2VyLm1hbmFnZXIuam9iIHVybjpvcGM6aWRtOnQub2F1dGggdXJuOm9wYzppZG06dC5ncm91cHMgdXJuOm9wYzppZG06dC5qb2IuaW1wb3J0ZXhwb3J0IHVybjpvcGM6aWRtOnQuaWRicmlkZ2UudW5tYXBwZWQuaWRjc2F0dHJpYnV0ZXMgdXJuOm9wYzppZG06dC5rcmIuYWRtaW4gdXJuOm9wYzppZG06dC5uYW1lZGFwcGFkbWluIHVybjpvcGM6aWRtOnQuYmxrcnB0cyB1cm46b3BjOmlkbTp0LnNlbGZyZWdpc3RyYXRpb25wcm9maWxlIHVybjpvcGM6aWRtOnQudXNlci5hdXRoZW50aWNhdGUgdXJuOm9wYzppZG06dC5ncmFudHMgdXJuOm9wYzppZG06dC5hdXRoZW50aWNhdGlvbiB1cm46b3BjOmlkbTp0LmNvbnRhaW5lciB1cm46b3BjOmlkbTp0LmltYWdlcyB1cm46b3BjOmlkbTp0LmNhLmFkbWluIHVybjpvcGM6aWRtOnQuYnVsayB1cm46b3BjOmlkbTp0LmRlbGVnYXRlZC5ncm91cC5tZW1iZXJzIHVybjpvcGM6aWRtOnQuam9iLnNlYXJjaCB1cm46b3BjOmlkbTp0LmlkYnJpZGdlIHVybjpvcGM6aWRtOnQuc2V0dGluZ3MgdXJuOm9wYzppZG06dC51c2VyLm1hbmFnZXIuc2VjdXJpdHkgdXJuOm9wYzppZG06dC5jbG91ZGdhdGUgdXJuOm9wYzppZG06dC5pZGJyaWRnZS5zb3VyY2VldmVudCB1cm46b3BjOmlkbTp0LnBvbGljeSB1cm46b3BjOmlkbTp0LnVzZXJzIHVybjpvcGM6aWRtOnQucmVwb3J0cyB1cm46b3BjOmlkbTpnLmlkY3NycHRzbWV0YV9yIiwiY2xpZW50X3RlbmFudG5hbWUiOiJpZGNzLTlmZWIzZWY5MWNhNDQ0MWJiZWRhZGJjNTY0MGQ1ZmI4IiwidXNlcl9sYW5nIjoiZW4iLCJ1c2VyQXBwUm9sZXMiOlsiQXV0aGVudGljYXRlZCIsIkdsb2JhbCBWaWV3ZXIiLCJJZGVudGl0eSBEb21haW4gQWRtaW5pc3RyYXRvciJdLCJleHAiOjE2MDA5OTIwMTgsImlhdCI6MTYwMDk4ODQxOCwiY2xpZW50X2d1aWQiOiJlMmM1Y2ZjZGY3ZWM0ZmVmOTNhYzRlZDIzZTk3OTBjZiIsImNsaWVudF9uYW1lIjoiVGVzdGluZ1NESyIsImlkcF90eXBlIjoiTE9DQUwiLCJ0ZW5hbnQiOiJpZGNzLTlmZWIzZWY5MWNhNDQ0MWJiZWRhZGJjNTY0MGQ1ZmI4IiwianRpIjoiYjAyNzJiYWItM2Y1MC00YTMzLWJhZjYtN2RmMjg4MzdiNDRkIiwidXNlcl9kaXNwbGF5bmFtZSI6IlBhcm5lZXQgS2F1ciIsImFsc29UZW5hbnRzIjpbImlkY3MtY2EiXSwic3ViX21hcHBpbmdhdHRyIjoidXNlck5hbWUiLCJwcmltVGVuYW50Ijp0cnVlLCJ0b2tfdHlwZSI6IkFUIiwiY2FfZ3VpZCI6ImNhY2N0LTA2Y2JlODcwODExMzQwY2E4ZGJlMWU4NTgwYWFkZmI2IiwiYXVkIjpbImh0dHBzOlwvXC9pZGNzLTlmZWIzZWY5MWNhNDQ0MWJiZWRhZGJjNTY0MGQ1ZmI4LmlkZW50aXR5Lm9yYWNsZWNsb3VkLmNvbTo0NDMiLCJ1cm46b3BjOmxiYWFzOmxvZ2ljYWxndWlkPWlkY3MtOWZlYjNlZjkxY2E0NDQxYmJlZGFkYmM1NjQwZDVmYjgiXSwidXNlcl9pZCI6IjJlNWIzZmE3MmI1YTQ2OTFiOGZiOTNkMDY0NzYwNmE3IiwiY2xpZW50QXBwUm9sZXMiOlsiQXV0aGVudGljYXRlZCBDbGllbnQiLCJJZGVudGl0eSBEb21haW4gQWRtaW5pc3RyYXRvciJdLCJ0ZW5hbnRfaXNzIjoiaHR0cHM6XC9cL2lkY3MtOWZlYjNlZjkxY2E0NDQxYmJlZGFkYmM1NjQwZDVmYjguaWRlbnRpdHkub3JhY2xlY2xvdWQuY29tOjQ0MyJ9.pS85ZctqAADIgBv2u8yjq8zZKjimsNyosontgJqUkSzx6H1Rd7mQqSRMb395_cJ7AHz4M7KRDNdOSZWmP9OwCFQzkYoeQNWGCJM0gKJxxWAvbqLp048Nl2Ch2MOq8oA3h2BfDxJ-NKH0iAeQrEMeh8G3Ki3g4fACbPbFsKDv5MdII4ey-0xxskkbU3zCfgN7iRvSPGDyb65zxbkGsYZan_NZg_yCdJPz7EME2p8o5jsNxjCANKxPnJQ9U88JFvkQq6RPTvsihgOI4CtBu12Faa71VuFWltza0dwtbI8z8QImTKYx84WQgWUtszQ_DKxYJvfCeEzAhsulUSeBvaE-Mg"

  # Optional
  network_endpoint_details {
    #Required
    network_endpoint_type = "PUBLIC"

    #Optional
    whitelisted_ips = [oci_core_vcn.test_vcn.cidr_block]

    whitelisted_vcns {
      #Optional
      id              = oci_core_vcn.test_vcn.id
      whitelisted_ips = [oci_core_vcn.test_vcn.cidr_block]
    }
  }
}

# Create a private access channel for the instance
resource "oci_analytics_analytics_instance_private_access_channel" "test_private_access_channel" {
#Required
  analytics_instance_id = oci_analytics_analytics_instance.test_oce_instance_public.id
  display_name = "Example Private Access Channel"
  subnet_id = oci_core_subnet.test_subnet.id
  vcn_id = oci_core_vcn.test_vcn.id
  private_source_dns_zones {
    dns_zone = "examplecorp.com"
    description = "Example dns zone"
  }
}

# Create a vanity url for the instance
resource "oci_analytics_analytics_instances_vanity_url" "test_analytics_instances_vanity_url" {
  #Required
  analytics_instance_id = oci_analytics_analytics_instance.test_oce_instance_public.id
  ca_certificate        = file("/Users/papakaur/Downloads/TestVanityUrls/RootCA.crt")
  hosts                 = ["analyticsdev"]
  private_key           = file("/Users/papakaur/Downloads/TestVanityUrls/analyticsdevCA.key")
  public_certificate    = file("/Users/papakaur/Downloads/TestVanityUrls/analyticsdevCA.crt")

  #Optional
  description = "description"
  # passphrase is required if the private key was created with a passphrase
  passphrase  = "passphrase"
}

