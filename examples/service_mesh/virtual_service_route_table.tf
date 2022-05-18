resource "oci_service_mesh_virtual_service_route_table" "test_virtual_service_route_table" {
  #Required
  compartment_id = var.compartment_ocid
  name           = local.virtual_service_route_table_name
  route_rules {
    #Required
    destinations {
      #Required
      virtual_deployment_id = oci_service_mesh_virtual_deployment.test_virtual_deployment.id
      weight                = local.virtual_service_route_table_route_rules_destinations_weight

      #Optional
      port = local.virtual_service_route_table_route_rules_destinations_port
    }
    type = local.virtual_service_route_table_route_rules_type

    #Optional
    is_grpc   = local.virtual_service_route_table_route_rules_is_grpc
    path      = local.virtual_service_route_table_route_rules_path
    path_type = local.virtual_service_route_table_route_rules_path_type
  }
  virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id

  #Optional
  description   = local.virtual_service_route_table_description
  freeform_tags = { "bar-key" = "value" }
  priority      = local.virtual_service_route_table_priority
}

data "oci_service_mesh_virtual_service_route_table" "test_virtual_service_route_table" {
  #Required
  virtual_service_route_table_id = oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table.id
}

locals {
  virtual_service_route_table_name                            = "test-virtual-service-route-table"
  virtual_service_route_table_route_rules_destinations_weight = "100"
  virtual_service_route_table_route_rules_type                = "HTTP"
  // allowed values for above are ("HTTP", "TLS_PASSTHROUGH", "TCP")
  virtual_service_route_table_description                     = "test virtual service route table description"
  virtual_service_route_table_priority                        = 1
  // allowed values for above are between 1 and 1000, with 1 being highest priority.
  virtual_service_route_table_route_rules_destinations_port   = "8080"
  virtual_service_route_table_route_rules_is_grpc             = true
  virtual_service_route_table_route_rules_path                = "/path"
  virtual_service_route_table_route_rules_path_type           = "PREFIX"
}