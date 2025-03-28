---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_identity_configuration"
sidebar_current: "docs-oci-resource-bds-bds_instance_identity_configuration"
description: |-
  Provides the Bds Instance Identity Configuration resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_identity_configuration
This resource provides the Bds Instance Identity Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Create an identity configuration for the cluster

## Example Usage

```hcl
resource "oci_bds_bds_instance_identity_configuration" "test_bds_instance_identity_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_identity_configuration_cluster_admin_password
	confidential_application_id = oci_dataflow_application.test_application.id
	display_name = var.bds_instance_identity_configuration_display_name
	identity_domain_id = oci_identity_domain.test_domain.id

	#Optional
	iam_user_sync_configuration_details {

		#Optional
		is_posix_attributes_addition_required = var.bds_instance_identity_configuration_iam_user_sync_configuration_details_is_posix_attributes_addition_required
	}
	upst_configuration_details {

		#Optional
		master_encryption_key_id = oci_kms_key.test_key.id
		vault_id = oci_kms_vault.test_vault.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) (Updatable) Base-64 encoded password for the cluster admin user.
* `confidential_application_id` - (Required) Identity domain confidential application ID for the identity config, required for creating identity configuration
* `display_name` - (Required) Display name of the identity configuration, required for creating identity configuration.
* `iam_user_sync_configuration_details` - (Optional) (Updatable) Details for activating/updating an IAM user sync configuration
	* `is_posix_attributes_addition_required` - (Optional) (Updatable) whether posix attribute needs to be appended to users, required for updating IAM user sync configuration
* `identity_domain_id` - (Required) Identity domain OCID to use for identity config, required for creating identity configuration
* `upst_configuration_details` - (Optional) (Updatable) Details for activating/updating UPST config on the cluster
	* `master_encryption_key_id` - (Optional) (Updatable) OCID of the master encryption key in vault for encrypting token exchange service principal keytab, required for activating UPST config
	* `vault_id` - (Optional) (Updatable) OCID of the vault to store token exchange service principal keyta, required for activating UPST config
* `activate_iam_user_sync_configuration_trigger` - (Optional) (Updatable) An optional property when set to "true" triggers Activate Iam User Sync Configuration and when set to "false" triggers Deactivate Iam User Sync Configuration.
* `activate_upst_configuration_trigger` - (Optional) (Updatable) An optional property when set to "true" triggers Activate Upst Configuration and when set to "false" triggers Deactivate Upst Configuration.
* `refresh_confidential_application_trigger` - (Optional) (Updatable) An optional property when set to "true" triggers Refresh Confidential Application.
* `refresh_upst_token_exchange_keytab_trigger` - (Optional) (Updatable) An optional property when set to "true"  triggers Refresh Upst Token Exchange Keytab.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Identity Configuration
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Identity Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Identity Configuration


## Import

BdsInstanceIdentityConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_identity_configuration.test_bds_instance_identity_configuration "bdsInstances/{bdsInstanceId}/identityConfigurations/{identityConfigurationId}" 
```

