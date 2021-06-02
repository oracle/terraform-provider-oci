// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Create an External Container Database resource
resource "oci_database_external_container_database" "test_external_container_database" {
  compartment_id = var.compartment_ocid
  display_name = var.external_container_database_display_name

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

  freeform_tags = {
    "Department" = "Finance"
  }
}

// Create a Connector using credential name for the External Container Database resource
resource "oci_database_external_database_connector" "test_external_container_database_connector" {
    connection_credentials {
        credential_type = var.credential_type
        credential_name = var.credential_name
    }
    connection_string {
        hostname = var.hostname
        port = var.port
        protocol = var.protocol
        service = var.service
    }
    connector_agent_id = var.connector_agent_id
    display_name = var.external_database_connector_display_name
    external_database_id = oci_database_external_container_database.test_external_container_database.id
    connector_type = var.connector_type
    #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
    #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

    freeform_tags = {
      "Department" = "Finance"
    }
}

// Enable Database Management

resource "oci_database_external_container_database_management" "test_enable_external_container_database_management" {
    external_container_database_id = oci_database_external_container_database.test_external_container_database.id
    external_database_connector_id = oci_database_external_database_connector.test_external_container_database_connector.id
    license_model = var.license_model
    enable_management = true
}

// Disable Database Management

/*resource "oci_database_external_container_database_management" "test_disable_external_container_database_management" {
    depends_on = [oci_database_external_container_database_management.test_enable_external_container_database_management,
    oci_database_external_pluggable_database_management.test_disable_external_pluggable_database_management]
    external_container_database_id = oci_database_external_container_database.test_external_container_database.id
    external_database_connector_id = oci_database_external_database_connector.test_external_container_database_connector.id
    enable_management = false
}*/

data "oci_database_external_container_database" "test_external_container_database" {
	#Required
	external_container_database_id = oci_database_external_container_database.test_external_container_database.id
}

data "oci_database_external_container_databases" "test_external_container_databases" {
	#Required
	compartment_id = var.compartment_ocid

	#Optional
	display_name = var.external_container_database_display_name
	state = var.external_database_state
}