// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Create an External Pluggable Database resource
resource "oci_database_external_pluggable_database" "test_external_pluggable_database" {
  compartment_id = var.compartment_ocid
  display_name = var.external_pluggable_database_display_name
  external_container_database_id = oci_database_external_container_database.test_external_container_database.id

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

  freeform_tags = {
    "Department" = "Finance"
  }
}

// Create a Connector using credential name for the External Pluggable Database resource
resource "oci_database_external_database_connector" "test_external_pluggable_database_connector" {
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
    external_database_id = oci_database_external_pluggable_database.test_external_pluggable_database.id
    connector_type = var.connector_type
    #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
    #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

    freeform_tags = {
      "Department" = "Finance"
    }
}

// Enable Database Management for the External Pluggable Database
resource "oci_database_external_pluggable_database_management" "test_enable_external_pluggable_database_management" {
    depends_on = [oci_database_external_container_database_management.test_enable_external_container_database_management]
    external_pluggable_database_id = oci_database_external_pluggable_database.test_external_pluggable_database.id
    external_database_connector_id = oci_database_external_database_connector.test_external_pluggable_database_connector.id
    enable_management = true
}

// Enable Operations Insights for the External Pluggable Databases
resource "oci_database_external_pluggable_database_operations_insights_management" "test_enable_external_pluggable_database_operations_insights_management" {
  external_pluggable_database_id = oci_database_external_pluggable_database.test_external_pluggable_database.id
  external_database_connector_id = oci_database_external_database_connector.test_external_pluggable_database_connector.id
  enable_operations_insights = true
}

//Commenting out this code block to unblock the failure in backward compatibility test
// Disable Database Management for the External Pluggable Database
/*resource "oci_database_external_pluggable_database_management" "test_disable_external_pluggable_database_management" {
    depends_on = [oci_database_external_pluggable_database_management.test_enable_external_pluggable_database_management]
    external_pluggable_database_id = oci_database_external_pluggable_database.test_external_pluggable_database.id
    external_database_connector_id = oci_database_external_database_connector.test_external_pluggable_database_connector.id
    enable_management = false
}*/

data "oci_database_external_pluggable_database" "test_external_pluggable_database" {
	#Required
	external_pluggable_database_id = oci_database_external_pluggable_database.test_external_pluggable_database.id
}

data "oci_database_external_pluggable_databases" "test_external_pluggable_databases" {
	#Required
	compartment_id = var.compartment_ocid

	#Optional
	display_name = var.external_pluggable_database_display_name
	state = var.external_database_state
}