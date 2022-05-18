resource "oci_service_mesh_ingress_gateway" "test_ingress_gateway" {
  #Required
  compartment_id = var.compartment_ocid
  hosts {
    #Required
    listeners {
      #Required
      port     = local.ingress_gateway_hosts_listeners_port
      protocol = local.ingress_gateway_hosts_listeners_protocol

      #Optional
      tls {
        #Required
        mode = local.ingress_gateway_hosts_listeners_tls_mode

        #Optional
        client_validation {

          #Optional
          subject_alternate_names = local.ingress_gateway_hosts_listeners_tls_client_validation_subject_alternate_names
          trusted_ca_bundle {
            #Required
            type = local.ingress_gateway_hosts_listeners_tls_client_validation_trusted_ca_bundle_type

            #Optional
            ca_bundle_id = oci_certificates_management_ca_bundle.test_ca_bundle.id
          }
        }
        server_certificate {
          #Required
          type = local.ingress_gateway_hosts_listeners_tls_server_certificate_type

          #Optional
          certificate_id = var.certificate_id
        }
      }
    }
    name = local.ingress_gateway_hosts_name

    #Optional
    hostnames = local.ingress_gateway_hosts_hostnames
  }
  mesh_id = oci_service_mesh_mesh.test_mesh.id
  name    = local.ingress_gateway_name

  #Optional
  access_logging {

    #Optional
    is_enabled = local.ingress_gateway_access_logging_is_enabled
  }
  description   = local.ingress_gateway_description
  freeform_tags = { "bar-key" = "value" }
  mtls {

    #Optional
    maximum_validity = local.ingress_gateway_mtls_maximum_validity
  }
}

data "oci_service_mesh_ingress_gateway" "test_ingress_gateway" {
  #Required
  ingress_gateway_id = oci_service_mesh_ingress_gateway.test_ingress_gateway.id
}

locals {
  ingress_gateway_name                                                          = "test-ingress-gateway"
  ingress_gateway_description                                                   = "test ingress gateway description"
  ingress_gateway_hosts_name                                                    = "test-host"
  ingress_gateway_hosts_hostnames                                               = ["test.com"]
  ingress_gateway_hosts_listeners_port                                          = "8090"
  // allowed values for above are between 1024 and 65535
  ingress_gateway_hosts_listeners_protocol                                      = "HTTP"
  // allowed values for above are ("HTTP", "TLS_PASSTHROUGH", "TCP")
  ingress_gateway_hosts_listeners_tls_mode                                      = "DISABLED"
  // allowed values for above are ("DISABLED", "PERMISSIVE", "TLS", "MUTUAL_TLS")
  ingress_gateway_hosts_listeners_tls_client_validation_subject_alternate_names = ["test-alt-name"]
  ingress_gateway_hosts_listeners_tls_client_validation_trusted_ca_bundle_type  = "OCI_CERTIFICATES"
  // allowed values for above are ("OCI_CERTIFICATES", "LOCAL_FILE")
  ingress_gateway_hosts_listeners_tls_server_certificate_type                   = "OCI_CERTIFICATES"
  // allowed values for above are ("OCI_CERTIFICATES", "LOCAL_FILE")
  ingress_gateway_access_logging_is_enabled                                     = true
  ingress_gateway_mtls_maximum_validity                                         = 50
  // allowed values for above are between 45 and 90 days
}