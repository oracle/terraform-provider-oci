// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "wlp_agent_agent_version" {
  default = "1.0.147"
}

variable "wlp_agent_certificate_signed_request" {
  default = "-----BEGIN CERTIFICATE REQUEST-----\nMIID1DCCArwCAQAwggGNMQswCQYDVQQGEwJVUzEPMA0GA1UEChMGT3JhY2xlMYIB\nKDCCASQGA1UECxOCARtvcGMtaW5zdGFuY2U6b2NpZDEuaW5zdGFuY2Uub2MxLnBo\neC5hbnlocWxqdDd4bTQ1Y2ljeWkya2E1cHVtcWd5dWhkZXVhaHJjYXh6NHN3bGtv\nZmo2dXhjdmtubnhkaGEsb3BjLWNvbXBhcnRtZW50Om9jaWQxLnRlbmFuY3kub2Mx\nLi5hYWFhYWFhYXFvZ2d6c2p1dDJ1NjR3cWxpeWQ0ZXlkM2RsNGlwc3UyNmxncXg0\nYmlob2ZudmU1bGk1aHEsb3BjLXRlbmFudDpvY2lkMS50ZW5hbmN5Lm9jMS4uYWFh\nYWFhYWFxb2dnenNqdXQydTY0d3FsaXlkNGV5ZDNkbDRpcHN1MjZsZ3F4NGJpaG9m\nbnZlNWxpNWhxMQwwCgYDVQQDEwN3bHAxMzAxBgkqhkiG9w0BCQEWJHdvcmtsb2Fk\ncHJvdGVjdGlvbl91c19ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMDLbqoECIIh02HvkusRyGGI/cqK9Wrg7xDn/Wwg1C9noOo+\nbHmU5sBervLUHKXuC3IUwM0GgytjLsOjMWI9ex0ZunQONwwAe/MDD+YQcnqbOnmb\naUrdp0gB231SRqCUST1xf9y8shlK3zXrav+qgtF1bDihsGh6O4DMLPYIsOZAXo6M\nrGPokj1nViLdvFaBBG4Q1sgximufh/eqFCaUawIUOeQ7XcDqeWM+G8IA3vIuWqbr\nSoI61/COgq6eDsUMu/ZcMNF0UYRV4bWwVM18Cx8Tlp0kH/mbnlHxBMxz1x/cbHmQ\nEwPrSKWo8Gn2B1HeXWhVGNPa4Xs0xn/kaW1QaS8CAwEAAaAAMA0GCSqGSIb3DQEB\nCwUAA4IBAQABiABQPOngTCA24KzY6GcyVi/4H6nhOu6smAgnPM2PoJEoog5yvnLR\nTvoyec0TTIIiRZtDIYejRMUyGZxR1o1Hgrkq80OmqfRZW57e2WPRgpHcp87Yfp0B\nRmkobQMRSAypZDGCdco2cuQ4F7GG0KFMb1Tf+b/XQnf6L3cd9PCHPECOVe1LFJV3\nqxhNkkxd+REI8iihLjzslqJFufYTkfmL2xamhS2nzGbG5XcfURdqx6S2ZDVoCkNy\nikohM9PlBrWAXWYALRqgcy1KFH9lQ9+tIqpnGbOHOyIqFPmoMKX2ugisTWMpgTp9\nxICh2HMz77KABXXf/t58HDODI4Wx8yJA\n-----END CERTIFICATE REQUEST-----\n"
}

variable "wlp_agent_os_info" {
  default = "Oracle Linux Server_8.5_amd64"
}



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  //version             = "5.39.0"
  /*tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region*/
}

resource "oci_cloud_guard_wlp_agent" "test_wlp_agent" {
  #Required
  agent_version              = var.wlp_agent_agent_version
  certificate_signed_request = var.wlp_agent_certificate_signed_request
  compartment_id             = var.compartment_id
  os_info                    = var.wlp_agent_os_info
}

data "oci_cloud_guard_wlp_agents" "test_wlp_agents" {
  #Required
  compartment_id = var.compartment_id
}
