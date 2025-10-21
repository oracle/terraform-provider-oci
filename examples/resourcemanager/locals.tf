locals {
  vcn_cidr_block           = "10.12.0.0/16"
  default_shape_name       = "VM.Standard.E3.Flex"
  operating_system         = "Oracle Linux"
  operating_system_version = "8"
  tcp_protocol             = 6
  private_endpoint_integ_test_vcn_cidr_block   = "10.12.0.0/16"
}