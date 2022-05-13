resource "oci_service_mesh_ingress_gateway_route_table" "test_ingress_gateway_route_table" {
  #Required
  compartment_id     = var.compartment_ocid
  ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
  name               = local.ingress_gateway_route_table_name
  route_rules {
    #Required
    destinations {
      #Required
      virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id

      #Optional
      port   = local.ingress_gateway_route_table_route_rules_destinations_port
      weight = local.ingress_gateway_route_table_route_rules_destinations_weight
    }
    type = local.ingress_gateway_route_table_route_rules_type

    #Optional
    ingress_gateway_host {
      #Required
      name = local.ingress_gateway_route_table_route_rules_ingress_gateway_host_name

      #Optional
      port = local.ingress_gateway_route_table_route_rules_ingress_gateway_host_port
    }
    is_grpc                 = local.ingress_gateway_route_table_route_rules_is_grpc
    is_host_rewrite_enabled = local.ingress_gateway_route_table_route_rules_is_host_rewrite_enabled
    is_path_rewrite_enabled = local.ingress_gateway_route_table_route_rules_is_path_rewrite_enabled
    path                    = local.ingress_gateway_route_table_route_rules_path
    path_type               = local.ingress_gateway_route_table_route_rules_path_type
  }

  #Optional
  description   = local.ingress_gateway_route_table_description
  freeform_tags = { "bar-key" = "value" }
  priority      = local.ingress_gateway_route_table_priority
}

data "oci_service_mesh_ingress_gateway_route_table" "test_ingress_gateway_route_table" {
  #Required
  ingress_gateway_route_table_id = oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table.id
}

data "oci_service_mesh_ingress_gateway_route_tables" "test_ingress_gateway_route_tables" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
}

locals {
  ingress_gateway_route_table_name                                  = "test-ingress-route-table"
  ingress_gateway_route_table_description                           = "test ingress gateway route table description"
  ingress_gateway_route_table_priority                              = 10
  // allowed values for above are between 1 and 1000, with 1 being highest priority.
  ingress_gateway_route_table_route_rules_type                      = "HTTP"
  // allowed values for above are ("HTTP", "TLS_PASSTHROUGH", "TCP")
  ingress_gateway_route_table_route_rules_destinations_port         = "8070"
  ingress_gateway_route_table_route_rules_destinations_weight       = 100
  ingress_gateway_route_table_route_rules_is_grpc                   = false
  ingress_gateway_route_table_route_rules_is_host_rewrite_enabled   = true
  ingress_gateway_route_table_route_rules_is_path_rewrite_enabled   = true
  ingress_gateway_route_table_route_rules_path                      = "/path"
  ingress_gateway_route_table_route_rules_path_type                 = "PREFIX"
  ingress_gateway_route_table_route_rules_ingress_gateway_host_name = "test-host"
  ingress_gateway_route_table_route_rules_ingress_gateway_host_port = "2000"
  // allowed values for above are between 1024 and 65535
}