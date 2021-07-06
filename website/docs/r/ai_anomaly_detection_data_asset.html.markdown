---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_data_asset"
sidebar_current: "docs-oci-resource-ai_anomaly_detection-data_asset"
description: |-
  Provides the Data Asset resource in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# oci_ai_anomaly_detection_data_asset
This resource provides the Data Asset resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Creates a new DataAsset.


## Example Usage

```hcl
resource "oci_ai_anomaly_detection_data_asset" "test_data_asset" {
	#Required
	compartment_id = var.compartment_id
	data_source_details {
		#Required
		data_source_type = var.data_asset_data_source_details_data_source_type

		#Optional
		atp_password_secret_id = oci_vault_secret.test_secret.id
		atp_user_name = oci_identity_user.test_user.name
		bucket = var.data_asset_data_source_details_bucket
		cwallet_file_secret_id = oci_vault_secret.test_secret.id
		database_name = oci_database_database.test_database.name
		ewallet_file_secret_id = oci_vault_secret.test_secret.id
		key_store_file_secret_id = oci_vault_secret.test_secret.id
		measurement_name = var.data_asset_data_source_details_measurement_name
		namespace = var.data_asset_data_source_details_namespace
		object = var.data_asset_data_source_details_object
		ojdbc_file_secret_id = oci_vault_secret.test_secret.id
		password_secret_id = oci_vault_secret.test_secret.id
		table_name = oci_nosql_table.test_table.name
		tnsnames_file_secret_id = oci_vault_secret.test_secret.id
		truststore_file_secret_id = oci_vault_secret.test_secret.id
		url = var.data_asset_data_source_details_url
		user_name = oci_identity_user.test_user.name
		version_specific_details {
			#Required
			influx_version = var.data_asset_data_source_details_version_specific_details_influx_version

			#Optional
			bucket = var.data_asset_data_source_details_version_specific_details_bucket
			database_name = oci_database_database.test_database.name
			organization_name = var.data_asset_data_source_details_version_specific_details_organization_name
			retention_policy_name = oci_identity_policy.test_policy.name
		}
		wallet_password_secret_id = oci_vault_secret.test_secret.id
	}
	project_id = oci_ai_anomaly_detection_project.test_project.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.data_asset_description
	display_name = var.data_asset_display_name
	freeform_tags = {"bar-key"= "value"}
	private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID for the data asset's compartment.
* `data_source_details` - (Required) Possible data sources
	* `atp_password_secret_id` - (Applicable when data_source_type=ORACLE_ATP) atp db password Secret Id
	* `atp_user_name` - (Applicable when data_source_type=ORACLE_ATP) atp db user name
	* `bucket` - (Applicable when data_source_type=ORACLE_OBJECT_STORAGE) Object storage bucket name
	* `cwallet_file_secret_id` - (Applicable when data_source_type=ORACLE_ATP) OCID of the secret containing the containers certificates of ATP wallet
	* `data_source_type` - (Required) Data source type where actually data asset is being stored
	* `database_name` - (Applicable when data_source_type=ORACLE_ATP) atp database name
	* `ewallet_file_secret_id` - (Applicable when data_source_type=ORACLE_ATP) OCID of the secret containing the PDB'S certificates of ATP wallet
	* `key_store_file_secret_id` - (Applicable when data_source_type=ORACLE_ATP) OCID of the secret containing Keystore.jks file of the ATP wallet
	* `measurement_name` - (Required when data_source_type=INFLUX) Measurement name for influx
	* `namespace` - (Applicable when data_source_type=ORACLE_OBJECT_STORAGE) Object storage namespace
	* `object` - (Applicable when data_source_type=ORACLE_OBJECT_STORAGE) File name
	* `ojdbc_file_secret_id` - (Applicable when data_source_type=ORACLE_ATP) OCID of the secret that contains jdbc properties file of ATP wallet
	* `password_secret_id` - (Required when data_source_type=INFLUX) Password Secret Id for the influx connection
	* `table_name` - (Applicable when data_source_type=ORACLE_ATP) atp database table name
	* `tnsnames_file_secret_id` - (Applicable when data_source_type=ORACLE_ATP) OCID of the secret that contains the tnsnames file of ATP wallet
	* `truststore_file_secret_id` - (Applicable when data_source_type=ORACLE_ATP) OCID of the secret containing truststore.jks file of the ATP wallet
	* `url` - (Required when data_source_type=INFLUX) public IP address and port to influx DB
	* `user_name` - (Required when data_source_type=INFLUX) Username for connection to Influx
	* `version_specific_details` - (Required when data_source_type=INFLUX) Possible data sources
		* `bucket` - (Required when influx_version=V_2_0) Bucket Name for influx connection
		* `database_name` - (Required when influx_version=V_1_8) DB Name for influx connection
		* `influx_version` - (Required) Data source type where actually data asset is being stored
		* `organization_name` - (Required when influx_version=V_2_0) Org name for the influx db
		* `retention_policy_name` - (Applicable when influx_version=V_1_8) retention policy is how long the bucket would last
	* `wallet_password_secret_id` - (Applicable when data_source_type=ORACLE_ATP) wallet password Secret ID in String format
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the Ai data asset
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `private_endpoint_id` - (Optional) OCID of Private Endpoint.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DataAsset.
* `data_source_details` - Possible data sources
	* `atp_password_secret_id` - atp db password Secret Id
	* `atp_user_name` - atp db user name
	* `bucket` - Object storage bucket name
	* `cwallet_file_secret_id` - OCID of the secret containing the containers certificates of ATP wallet
	* `data_source_type` - Data source type where actually data asset is being stored
	* `database_name` - atp database name
	* `ewallet_file_secret_id` - OCID of the secret containing the PDB'S certificates of ATP wallet
	* `key_store_file_secret_id` - OCID of the secret containing Keystore.jks file of the ATP wallet
	* `measurement_name` - Measurement name for influx
	* `namespace` - Object storage namespace
	* `object` - File name
	* `ojdbc_file_secret_id` - OCID of the secret that contains jdbc properties file of ATP wallet
	* `password_secret_id` - Password Secret Id for the influx connection
	* `table_name` - atp database table name
	* `tnsnames_file_secret_id` - OCID of the secret that contains the tnsnames file of ATP wallet
	* `truststore_file_secret_id` - OCID of the secret containing truststore.jks file of the ATP wallet
	* `url` - public IP address and port to influx DB
	* `user_name` - Username for connection to Influx
	* `version_specific_details` - Possible data sources
		* `bucket` - Bucket Name for influx connection
		* `database_name` - DB Name for influx connection
		* `influx_version` - Data source type where actually data asset is being stored
		* `organization_name` - Org name for the influx db
		* `retention_policy_name` - retention policy is how long the bucket would last
	* `wallet_password_secret_id` - wallet password Secret ID in String format
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the data asset.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Unique Oracle ID (OCID) that is immutable on creation.
* `private_endpoint_id` - OCID of Private Endpoint.
* `project_id` - The Unique project id which is created at project creation that is immutable on creation.
* `state` - The lifecycle state of the Data Asset.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the DataAsset was created. An RFC3339 formatted datetime string
* `time_updated` - The time the the DataAsset was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Asset
	* `update` - (Defaults to 20 minutes), when updating the Data Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Data Asset


## Import

DataAssets can be imported using the `id`, e.g.

```
$ terraform import oci_ai_anomaly_detection_data_asset.test_data_asset "id"
```

