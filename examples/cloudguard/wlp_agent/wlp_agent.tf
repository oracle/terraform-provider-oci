// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_id" {}

variable "wlp_agent_agent_version" {
  default = "1.0.147"
}

variable "wlp_agent_certificate_signed_request" {
  default = "-----BEGIN CERTIFICATE REQUEST-----\nMIID2DCCAsACAQAwggGRMQwwCgYDVQQDDAN3bHAxCzAJBgNVBAYTAlVTMQ8wDQYD\nVQQKDAZPcmFjbGUxggEsMIIBKAYDVQQLDIIBH29wYy1pbnN0YW5jZTpvY2lkMS5p\nbnN0YW5jZS5vYzEucGh4LmFueWhxbGpyanNuZ2xucWM1am1wcjdwajR0dG5uZXgz\ndW91cWMzbGJveXZsdnpzeHFnZm43ZmNjdWd1YSxvcGMtY29tcGFydG1lbnQ6b2Np\nZDEuY29tcGFydG1lbnQub2MxLi5hYWFhYWFhYTJrNXJicDN1bDJ1Y2loMno2bWlx\na3lxaTV6Z3Z6Y3lxaWlmcWN2cnpnbTJ4Znh3eTJpd3Esb3BjLXRlbmFudDpvY2lk\nMS50ZW5hbmN5Lm9jMS4uYWFhYWFhYWEzNnkzeWlkdGUybzdqenlhNXBkdXRlY2Rx\nYjdiZDdxZjJyc3h1c3c0NW5mazNvbmg1dTRxMTMwMQYJKoZIhvcNAQkBFiR3b3Jr\nbG9hZHByb3RlY3Rpb25fdXNfZ3JwQG9yYWNsZS5jb20wggEiMA0GCSqGSIb3DQEB\nAQUAA4IBDwAwggEKAoIBAQCsEFP2pZr0jX7gMF6hh8CYkM4VZsjguKhQeYaRwORD\n0F+5iyIv6XhDLtOSoBlYDgcwxj9CXjT9KWCg3mJFe0viMocVZj2PUlX38MzZVvgC\nCx+5x4DOcyZo/C9QtAGUhUshYHxY6YTUV6ybVQFavf0R91UaObaTA7ZsOXYNTQZ5\nVXZJ+uPStv9m4YcLj/8C4BLiTaHEgsdIz9tcXyoB3LDMDLS7lGeuuR8f2lx+M1kF\n4PcDFH22nSlumiLS24HNfV6cHqN6yDQBDUoqaP4GLKd3CH8yMOFdJlKiHR44Az7G\niYeUT8l03xYB59cl0VTn4aho2evVUCKs076IteAIgerZAgMBAAGgADANBgkqhkiG\n9w0BAQsFAAOCAQEALpm6WwvEqowU8nXG2Pkb7luW/T/CA1+alFxdpRYssez3tTGq\n2HxfQcxpv9VCZT2RZE/34pUbPKfPvtW5Ve9QKHfz3uNxOtZQh5o/02+nHlKqX5wT\nWOYHDdGA+QmYXC+OeY1qCT697q7dwZOnd8MefZ8n+x4Pl7sU2+NqkuIOei0Tl6T+\ncDvYbWewAuZgTxfvTrEDntrsRpgTEagFfWEXftpLrSgXwZlUF+5bZVPHNv+aeXx5\n4R4gCaAvfRk4xBhVrXMkXL1xq5lzikfVq4C0g+cXiZqplq+1PcIdRZhJiap/vxhM\ntCxCsRac/0XmuPiNO01KwcWPDl14MPLpDTf2qQ==\n-----END CERTIFICATE REQUEST-----\n"
}

variable "wlp_agent_os_info" {
  default = "Oracle Linux Server_8.5_amd64"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  version             = "7.19.0"
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
