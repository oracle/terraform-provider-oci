resource "oci_service_mesh_virtual_deployment" "test_virtual_deployment" {
  #Required
  compartment_id = var.compartment_ocid
  listeners {
    #Required
    port     = local.virtual_deployment_listeners_port
    protocol = local.virtual_deployment_listeners_protocol
  }
  name = local.virtual_deployment_name
  service_discovery {
    #Required
    hostname = local.virtual_deployment_service_discovery_hostname
    type     = local.virtual_deployment_service_discovery_type
  }
  virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id

  #Optional
  access_logging {

    #Optional
    is_enabled = local.virtual_deployment_access_logging_is_enabled
  }
  description   = local.virtual_deployment_description
  freeform_tags = { "bar-key" = "value" }
}

data "oci_service_mesh_virtual_deployment" "test_virtual_deployment" {
  #Required
  virtual_deployment_id = oci_service_mesh_virtual_deployment.test_virtual_deployment.id
}

locals {
  virtual_deployment_name                       = "test-virtual-deployment"
  virtual_deployment_description                = "test virtual deployment description"
  virtual_deployment_listeners_port             = "8080"
  virtual_deployment_listeners_protocol         = "HTTP"
  // allowed values for above for above are ("HTTP", "TLS_PASSTHROUGH", "TCP", "HTTP2", "GRPC")
  virtual_deployment_service_discovery_hostname = "test.com"
  virtual_deployment_service_discovery_type     = "DNS"
  virtual_deployment_access_logging_is_enabled  = true
}
