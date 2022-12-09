resource "oci_service_mesh_access_policy" "test_access_policy" {
  #Required
  compartment_id = var.compartment_ocid
  mesh_id        = oci_service_mesh_mesh.test_mesh.id
  name           = local.access_policy_name
  rules {
    #Required
    action = local.access_policy_rules_action
    destination {
      #Required
      type = local.access_policy_rules_destination_type

      #Optional
      hostnames          = local.access_policy_rules_destination_hostnames
      ip_addresses       = local.access_policy_rules_destination_ip_addresses
      ports              = local.access_policy_rules_destination_ports
      protocol           = local.access_policy_rules_destination_protocol
    }
    source {
      #Required
      type = local.access_policy_rules_source_type

      #Optional
      virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id
    }
  }

  #Optional
  description   = local.access_policy_description
  freeform_tags = { "bar-key" = "value" }
}

data "oci_service_mesh_access_policy" "test_access_policy" {
  #Required
  access_policy_id = oci_service_mesh_access_policy.test_access_policy.id
}

data "oci_service_mesh_access_policies" "test_access_policies" {
  #Required
  compartment_id = var.compartment_ocid
  mesh_id        = oci_service_mesh_mesh.test_mesh.id
}

locals {
  access_policy_name                           = "test-access-policy"
  access_policy_description                    = "test access policy description"
  access_policy_rules_action                   = "ALLOW"
  access_policy_rules_source_type              = "VIRTUAL_SERVICE"
  // allowed values for above are ("ALL_VIRTUAL_SERVICES", "VIRTUAL_SERVICE", "EXTERNAL_SERVICE", "INGRESS_GATEWAY")
  access_policy_rules_destination_type         = "EXTERNAL_SERVICE"
  access_policy_rules_destination_hostnames    = ["test.com"]
  access_policy_rules_destination_ip_addresses = ["0.0.0.0/0"]
  access_policy_rules_destination_ports        = ["8080"]
  access_policy_rules_destination_protocol     = "HTTP" // allowed values are ("HTTP", "HTTPS", "TCP")
}