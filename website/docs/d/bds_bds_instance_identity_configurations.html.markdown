---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_identity_configurations"
sidebar_current: "docs-oci-datasource-bds-bds_instance_identity_configurations"
description: |-
  Provides the list of Bds Instance Identity Configurations in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_identity_configurations
This data source provides the list of Bds Instance Identity Configurations in Oracle Cloud Infrastructure Big Data Service service.

Returns a list of all identity configurations associated with this Big Data Service cluster.


## Example Usage

```hcl
data "oci_bds_bds_instance_identity_configurations" "test_bds_instance_identity_configurations" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	compartment_id = var.compartment_id

	#Optional
	display_name = var.bds_instance_identity_configuration_display_name
	state = var.bds_instance_identity_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) The state of the identity config


## Attributes Reference

The following attributes are exported:

* `identity_configurations` - The list of identity_configurations.

### BdsInstanceIdentityConfiguration Reference

The following attributes are exported:

* `confidential_application_id` - identity domain confidential application ID for the identity config
* `display_name` - the display name of the identity configuration
* `iam_user_sync_configuration` - Information about the IAM user sync configuration.
	* `is_posix_attributes_addition_required` - whether to append POSIX attributes to IAM users
	* `state` - Lifecycle state of the IAM user sync config
	* `time_created` - Time when this IAM user sync config was created, shown as an RFC 3339 formatted datetime string.
	* `time_updated` - Time when this IAM user sync config was updated, shown as an RFC 3339 formatted datetime string.
* `id` - The id of the identity config
* `identity_domain_id` - Identity domain to use for identity config
* `state` - Lifecycle state of the identity configuration
* `time_created` - Time when this identity configuration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - Time when this identity configuration config was updated, shown as an RFC 3339 formatted datetime string.
* `upst_configuration` - Information about the UPST configuration.
	* `keytab_content` - The kerberos keytab content used for creating identity propagation trust config, in base64 format
	* `master_encryption_key_id` - Master Encryption key used for encrypting token exchange keytab.
	* `secret_id` - Secret ID for token exchange keytab
	* `state` - Lifecycle state of the UPST config
	* `time_created` - Time when this UPST config was created, shown as an RFC 3339 formatted datetime string.
	* `time_token_exchange_keytab_last_refreshed` - Time when the keytab for token exchange principal is last refreshed, shown as an RFC 3339 formatted datetime string.
	* `time_updated` - Time when this UPST config was updated, shown as an RFC 3339 formatted datetime string.
	* `token_exchange_principal_name` - Token exchange kerberos Principal name in cluster
	* `vault_id` - The instance OCID of the node, which is the resource from which the node backup was acquired.

