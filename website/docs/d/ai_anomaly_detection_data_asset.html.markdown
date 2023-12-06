---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_data_asset"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-data_asset"
description: |-
  Provides details about a specific Data Asset in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_data_asset
This data source provides details about a specific Data Asset resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Gets a DataAsset by identifier

## Example Usage

```hcl
data "oci_ai_anomaly_detection_data_asset" "test_data_asset" {
	#Required
	data_asset_id = oci_ai_anomaly_detection_data_asset.test_data_asset.id
}
```

## Argument Reference

The following arguments are supported:

* `data_asset_id` - (Required) The OCID of the Data Asset.


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

