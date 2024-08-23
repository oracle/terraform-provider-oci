---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_database_dbm_features_management"
sidebar_current: "docs-oci-resource-database_management-database_dbm_features_management"
description: |-
  Provides the Database Dbm Features Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_database_dbm_features_management
This resource provides the Database Dbm Features Management resource in Oracle Cloud Infrastructure Database Management service.

Enables a Database Management feature for the specified cloud database.


## Example Usage

```hcl
resource "oci_database_management_database_dbm_features_management" "test_database_dbm_features_management" {
	#Required
	database_id = oci_database_database.test_database.id
	enable_database_dbm_feature = var.enable_database_dbm_feature

	#Optional
	feature_details {
		#Required
		feature = var.database_dbm_features_management_feature_details_feature
		enable_database_dbm_feature = var.enable_database_dbm_feature

		#Optional
		connector_details {

			#Optional
			connector_type = var.database_dbm_features_management_feature_details_connector_details_connector_type
			database_connector_id = oci_database_management_database_connector.test_database_connector.id
			management_agent_id = oci_management_agent_management_agent.test_management_agent.id
			private_end_point_id = oci_database_management_private_end_point.test_private_end_point.id
		}
		database_connection_details {

			#Optional
			connection_credentials {

				#Optional
				credential_name = var.database_dbm_features_management_feature_details_database_connection_details_connection_credentials_credential_name
				credential_type = var.database_dbm_features_management_feature_details_database_connection_details_connection_credentials_credential_type
				named_credential_id = oci_database_management_named_credential.test_named_credential.id
				password_secret_id = oci_vault_secret.test_secret.id
				role = var.database_dbm_features_management_feature_details_database_connection_details_connection_credentials_role
				ssl_secret_id = oci_vault_secret.test_secret.id
				user_name = oci_identity_user.test_user.name
			}
			connection_string {

				#Optional
				connection_type = var.database_dbm_features_management_feature_details_database_connection_details_connection_string_connection_type
				port = var.database_dbm_features_management_feature_details_database_connection_details_connection_string_port
				protocol = var.database_dbm_features_management_feature_details_database_connection_details_connection_string_protocol
				service = var.database_dbm_features_management_feature_details_database_connection_details_connection_string_service
			}
		}
		is_auto_enable_pluggable_database = var.database_dbm_features_management_feature_details_is_auto_enable_pluggable_database
		management_type = var.database_dbm_features_management_feature_details_management_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database.
* `feature_details` - (Optional) The details required to enable the specified Database Management feature.
	* `connector_details` - (Optional) The connector details required to connect to an Oracle cloud database.
		* `connector_type` - (Optional) The list of supported connection types:
			* PE: Private endpoint
			* MACS: Management agent
			* EXTERNAL: External database connector
			* DIRECT: Direct connection 
		* `database_connector_id` - (Applicable when connector_type=EXTERNAL) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database connector.
		* `management_agent_id` - (Applicable when connector_type=MACS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent.
		* `private_end_point_id` - (Applicable when connector_type=PE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	* `database_connection_details` - (Optional) The connection details required to connect to the database.
		* `connection_credentials` - (Optional) The credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
			* `credential_name` - (Applicable when credential_type=DETAILS | NAME_REFERENCE | SSL_DETAILS) The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

				For example: inventorydb.abc112233445566778899 
			* `credential_type` - (Optional) The type of credential used to connect to the database.
			* `named_credential_id` - (Applicable when credential_type=NAMED_CREDENTIAL) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential where the database password metadata is stored. 
			* `password_secret_id` - (Applicable when credential_type=DETAILS | SSL_DETAILS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
			* `role` - (Applicable when credential_type=DETAILS | SSL_DETAILS) The role of the user connecting to the database.
			* `ssl_secret_id` - (Applicable when credential_type=SSL_DETAILS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
			* `user_name` - (Applicable when credential_type=DETAILS | SSL_DETAILS) The user name used to connect to the database.
		* `connection_string` - (Optional) The details of the Oracle Database connection string. 
			* `connection_type` - (Optional) The list of supported connection types:
				* BASIC: Basic connection details 
			* `port` - (Optional) The port number used to connect to the database.
			* `protocol` - (Optional) The protocol used to connect to the database.
			* `service` - (Optional) The service name of the database.
	* `feature` - (Required) The name of the Database Management feature.
	* `is_auto_enable_pluggable_database` - (Applicable when feature=DIAGNOSTICS_AND_MANAGEMENT) Indicates whether the pluggable database can be enabled automatically.
	* `management_type` - (Applicable when feature=DIAGNOSTICS_AND_MANAGEMENT) The management type for the database.
* `enable_database_dbm_feature` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Dbm Features Management
	* `update` - (Defaults to 20 minutes), when updating the Database Dbm Features Management
	* `delete` - (Defaults to 20 minutes), when destroying the Database Dbm Features Management
