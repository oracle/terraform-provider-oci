---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_key"
sidebar_current: "docs-oci-resource-kms-key"
description: |-
  Provides the Key resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_key
This resource provides the Key resource in Oracle Cloud Infrastructure Kms service.

Creates a new master encryption key.

As a management operation, this call is subject to a Key Management limit that applies to the total
number of requests across all management write operations. Key Management might throttle this call
to reject an otherwise valid request when the total rate of management write operations exceeds 10
requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_key" "test_key" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.key_display_name
	key_shape {
		#Required
		algorithm = var.key_key_shape_algorithm
		length = var.key_key_shape_length

		#Optional
		curve_id = oci_kms_curve.test_curve.id
	}
	management_endpoint = var.key_management_endpoint

	#Optional
	auto_key_rotation_details {

		#Optional
		last_rotation_message = var.key_auto_key_rotation_details_last_rotation_message
		last_rotation_status = var.key_auto_key_rotation_details_last_rotation_status
		rotation_interval_in_days = var.key_auto_key_rotation_details_rotation_interval_in_days
		time_of_last_rotation = var.key_auto_key_rotation_details_time_of_last_rotation
		time_of_next_rotation = var.key_auto_key_rotation_details_time_of_next_rotation
		time_of_schedule_start = var.key_auto_key_rotation_details_time_of_schedule_start
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	external_key_reference {
		#Required
		external_key_id = oci_kms_key.test_key.id
	}
	freeform_tags = {"Department"= "Finance"}
	is_auto_rotation_enabled = var.key_is_auto_rotation_enabled
	protection_mode = var.key_protection_mode
}
```

## Argument Reference

The following arguments are supported:

* `auto_key_rotation_details` - (Optional) (Updatable) The details of auto rotation schedule for the Key being create updated or imported.
	* `last_rotation_message` - (Optional) (Updatable) The last execution status message of auto key rotation. 
	* `last_rotation_status` - (Optional) (Updatable) The status of last execution of auto key rotation.
	* `rotation_interval_in_days` - (Optional) (Updatable) The interval of auto key rotation. For auto key rotation the interval should between 60 day and 365 days (1 year). Note: User must specify this parameter when creating a new schedule.
	* `time_of_last_rotation` - (Optional) (Updatable) A property indicating Last rotation Date. Example: `2023-04-04T00:00:00Z`.
	* `time_of_next_rotation` - (Optional) (Updatable) A property indicating Next estimated scheduled Time, as per the interval, expressed as date YYYY-MM-DD String. Example: `2023-04-04T00:00:00Z`. The time has no significance when scheduling an auto key rotation as this can be done anytime approximately the scheduled day, KMS ignores the time and replaces it with 00:00, for example 2023-04-04T15:14:13Z will be used as 2023-04-04T00:00:00Z. 
	* `time_of_schedule_start` - (Optional) (Updatable) A property indicating  scheduled start date expressed as date YYYY-MM-DD String. Example: `2023-04-04T00:00:00Z. The time has no significance when scheduling an auto key rotation as this can be done anytime approximately the scheduled day, KMS ignores the time and replaces it with 00:00, for example 2023-04-04T15:14:13Z will be used as 2023-04-04T00:00:00Z . Note : Today’s date will be used if not specified by customer.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the master encryption key.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `desired_state` - (Optional) (Updatable) Desired state of the key. Possible values : `ENABLED` or `DISABLED`
* `display_name` - (Required) (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `external_key_reference` - (Optional) A reference to the key on external key manager.
	* `external_key_id` - (Required) ExternalKeyId refers to the globally unique key Id associated with the key created in external vault in CTM
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_auto_rotation_enabled` - (Optional) (Updatable) A parameter specifying whether the auto key rotation is enabled or not.
* `key_shape` - (Required) The cryptographic properties of a key.
	* `algorithm` - (Required) The algorithm used by a key's key versions to encrypt or decrypt. Only AES algorithm is supported for `External` keys.
	* `curve_id` - (Optional) Supported curve IDs for ECDSA keys.
	* `length` - (Required) The length of the key in bytes, expressed as an integer. Supported values include the following:
		* AES: 16, 24, or 32
		* RSA: 256, 384, or 512
		* ECDSA: 32, 48, or 66 
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
* `protection_mode` - (Optional) The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default, a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported. A protection mode of `EXTERNAL` mean that the key persists on the customer's external key manager which is hosted externally outside of oracle. Oracle only hold a reference to that key. All cryptographic operations that use a key with a protection mode of `EXTERNAL` are performed by external key manager.  
* `restore_from_file` - (Optional) (Updatable) Details where key was backed up.
    * `content_length` - (Required) (Updatable) content length of key's backup binary file
    * `content_md5` - (Optional) (Updatable) content md5 hashed value of key's backup file
    * `restore_key_from_file_details` - (Required) Key backup file content.
* `restore_from_object_store` - (Optional) (Updatable) Details where key was backed up
    * `bucket` - (Optional) (Updatable) Name of the bucket where key was backed up
    * `destination` - (Required) (Updatable) Type of backup to restore from. Values of "BUCKET", "PRE_AUTHENTICATED_REQUEST_URI" are supported
    * `namespace` - (Optional) (Updatable) Namespace of the bucket where key was backed up
    * `object` - (Optional) (Updatable) Object containing the backup
    * `uri` - (Optional) (Updatable) Pre-authenticated-request-uri of the backup
* `restore_trigger` - (Optional) (Updatable) An optional property when flipped triggers restore from restore option provided in config file. 
* `time_of_deletion` - (Optional) (Updatable) An optional property for the deletion time of the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `auto_key_rotation_details` - The details of auto rotation schedule for the Key being create updated or imported.
	* `last_rotation_message` - The last execution status message of auto key rotation. 
	* `last_rotation_status` - The status of last execution of auto key rotation.
	* `rotation_interval_in_days` - The interval of auto key rotation. For auto key rotation the interval should between 60 day and 365 days (1 year). Note: User must specify this parameter when creating a new schedule.
	* `time_of_last_rotation` - A property indicating Last rotation Date. Example: `2023-04-04T00:00:00Z`.
	* `time_of_next_rotation` - A property indicating Next estimated scheduled Time, as per the interval, expressed as date YYYY-MM-DD String. Example: `2023-04-04T00:00:00Z`. The time has no significance when scheduling an auto key rotation as this can be done anytime approximately the scheduled day, KMS ignores the time and replaces it with 00:00, for example 2023-04-04T15:14:13Z will be used as 2023-04-04T00:00:00Z. 
	* `time_of_schedule_start` - A property indicating  scheduled start date expressed as date YYYY-MM-DD String. Example: `2023-04-04T00:00:00Z. The time has no significance when scheduling an auto key rotation as this can be done anytime approximately the scheduled day, KMS ignores the time and replaces it with 00:00, for example 2023-04-04T15:14:13Z will be used as 2023-04-04T00:00:00Z . Note : Today’s date will be used if not specified by customer.
* `compartment_id` - The OCID of the compartment that contains this master encryption key.
* `current_key_version` - The OCID of the key version used in cryptographic operations. During key rotation, the service might be in a transitional state where this or a newer key version are used intermittently. The `currentKeyVersion` property is updated when the service is guaranteed to use the new key version for all subsequent encryption operations. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `external_key_reference_details` - Key reference data to be returned to the customer as a response.
	* `external_key_id` - ExternalKeyId refers to the globally unique key Id associated with the key created in external vault in CTM.
	* `external_key_version_id` - Key version ID associated with the external key.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the key.
* `is_auto_rotation_enabled` - A parameter specifying whether the auto key rotation is enabled or not.
* `is_primary` - A Boolean value that indicates whether the Key belongs to primary Vault or replica vault.
* `key_shape` - The cryptographic properties of a key.
	* `algorithm` - The algorithm used by a key's key versions to encrypt or decrypt. Only AES algorithm is supported for `External` keys.
	* `curve_id` - Supported curve IDs for ECDSA keys.
	* `length` - The length of the key in bytes, expressed as an integer. Supported values include the following:
		* AES: 16, 24, or 32
		* RSA: 256, 384, or 512
		* ECDSA: 32, 48, or 66 
* `protection_mode` - The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default, a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported. A protection mode of `EXTERNAL` mean that the key persists on the customer's external key manager which is hosted externally outside of oracle. Oracle only hold a reference to that key. All cryptographic operations that use a key with a protection mode of `EXTERNAL` are performed by external key manager. 
* `replica_details` - Key replica details 
	* `replication_id` - ReplicationId associated with a key operation 
* `restored_from_key_id` - The OCID of the key from which this key was restored.
* `state` - The key's current lifecycle state.  Example: `ENABLED` 
* `time_created` - The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Key
	* `update` - (Defaults to 20 minutes), when updating the Key
	* `delete` - (Defaults to 20 minutes), when destroying the Key


## Import

Keys can be imported using the `id`, e.g.

```
$ terraform import oci_kms_key.test_key "managementEndpoint/{managementEndpoint}/keys/{keyId}" 
```

