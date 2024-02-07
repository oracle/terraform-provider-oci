---
subcategory: "Vault"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vault_secret"
sidebar_current: "docs-oci-resource-vault-secret"
description: |-
  Provides the Secret resource in Oracle Cloud Infrastructure Vault service
---

# oci_vault_secret
This resource provides the Secret resource in Oracle Cloud Infrastructure Vault service.

Creates a new secret according to the details of the request.


## Example Usage

```hcl
resource "oci_vault_secret" "test_secret" {
	#Required
	compartment_id = var.compartment_id
	secret_content {
		#Required
		content_type = var.secret_secret_content_content_type

		#Optional
		content = var.secret_secret_content_content
		name = var.secret_secret_content_name
		stage = var.secret_secret_content_stage
	}
	secret_name = oci_vault_secret.test_secret.name
	vault_id = oci_kms_vault.test_vault.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.secret_description
	freeform_tags = {"Department"= "Finance"}
	key_id = oci_kms_key.test_key.id
	metadata = var.secret_metadata
	rotation_config {
		#Required
		target_system_details {
			#Required
			target_system_type = var.secret_rotation_config_target_system_details_target_system_type

			#Optional
			adb_id = oci_vault_adb.test_adb.id
			function_id = oci_functions_function.test_function.id
		}

		#Optional
		is_scheduled_rotation_enabled = var.secret_rotation_config_is_scheduled_rotation_enabled
		rotation_interval = var.secret_rotation_config_rotation_interval
	}
	secret_content {
		#Required
		content_type = var.secret_secret_content_content_type

		#Optional
		content = var.secret_secret_content_content
		name = var.secret_secret_content_name
		stage = var.secret_secret_content_stage
	}
	secret_rules {
		#Required
		rule_type = var.secret_secret_rules_rule_type

		#Optional
		is_enforced_on_deleted_secret_versions = var.secret_secret_rules_is_enforced_on_deleted_secret_versions
		is_secret_content_retrieval_blocked_on_expiry = var.secret_secret_rules_is_secret_content_retrieval_blocked_on_expiry
		secret_version_expiry_interval = var.secret_secret_rules_secret_version_expiry_interval
		time_of_absolute_expiry = var.secret_secret_rules_time_of_absolute_expiry
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the secret.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A brief description of the secret. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `key_id` - (Optional) The OCID of the master encryption key that is used to encrypt the secret. You must specify a symmetric key to encrypt the secret during import to the vault. You cannot encrypt secrets with asymmetric keys. Furthermore, the key must exist in the vault that you specify. 
* `metadata` - (Optional) (Updatable) Additional metadata that you can use to provide context about how to use the secret during rotation or other administrative tasks. For example, for a secret that you use to connect to a database, the additional metadata might specify the connection endpoint and the connection string. Provide additional metadata as key-value pairs.
* `rotation_config` - (Optional) (Updatable) Defines the frequency of the rotation and the information about the target system
	* `is_scheduled_rotation_enabled` - (Optional) (Updatable) Enables auto rotation, when set to true rotationInterval must be set. 
	* `rotation_interval` - (Optional) (Updatable) The time interval that indicates the frequency for rotating secret data, as described in ISO 8601 format. The minimum value is 1 day and maximum value is 360 days. For example, if you want to set the time interval for rotating a secret data as 30 days, the duration is expressed as "P30D." 
	* `target_system_details` - (Required) (Updatable) The TargetSystemDetails provides the targetSystem type and type-specific connection metadata 
		* `adb_id` - (Required when target_system_type=ADB) (Updatable) The unique identifier (OCID) for the autonomous database that Vault Secret connects to. 
		* `function_id` - (Required when target_system_type=FUNCTION) (Updatable) The unique identifier (OCID) of the Oracle Cloud Infrastructure Functions that vault secret connects to. 
		* `target_system_type` - (Required) (Updatable) Unique identifier of the target system that Vault Secret connects to. 
* `secret_content` - (Optional) (Updatable) The content of the secret and metadata to help identify it.
	* `content` - (Optional) (Updatable) The base64-encoded content of the secret.
	* `content_type` - (Optional) (Updatable) The base64-encoded content of the secret.
	* `name` - (Optional) (Updatable) Names should be unique within a secret. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	* `stage` - (Optional) (Updatable) The rotation state of the secret content. The default is `CURRENT`, meaning that the secret is currently in use. A secret version that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example, you might create or update a secret and mark its rotation state as `PENDING` if you haven't yet updated the secret on the target system. When creating a secret, only the value `CURRENT` is applicable, although the value `LATEST` is also automatically applied. When updating a secret, you can specify a version's rotation state as either `CURRENT` or `PENDING`. 
* `secret_name` - (Required) A user-friendly name for the secret. Secret names should be unique within a vault. Avoid entering confidential information. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. 
* `secret_rules` - (Optional) (Updatable) A list of rules to control how the secret is used and managed.
	* `is_enforced_on_deleted_secret_versions` - (Applicable when rule_type=SECRET_REUSE_RULE) (Updatable) A property indicating whether the rule is applied even if the secret version with the content you are trying to reuse was deleted. 
	* `is_secret_content_retrieval_blocked_on_expiry` - (Applicable when rule_type=SECRET_EXPIRY_RULE) (Updatable) A property indicating whether to block retrieval of the secret content, on expiry. The default is false. If the secret has already expired and you would like to retrieve the secret contents, you need to edit the secret rule to disable this property, to allow reading the secret content. 
	* `rule_type` - (Required) (Updatable) The type of rule, which either controls when the secret contents expire or whether they can be reused.
	* `secret_version_expiry_interval` - (Applicable when rule_type=SECRET_EXPIRY_RULE) (Updatable) A property indicating how long the secret contents will be considered valid, expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. The secret needs to be updated when the secret content expires. The timer resets after you update the secret contents. The minimum value is 1 day and the maximum value is 90 days for this property. Currently, only intervals expressed in days are supported. For example, pass `P3D` to have the secret version expire every 3 days. 
	* `time_of_absolute_expiry` - (Applicable when rule_type=SECRET_EXPIRY_RULE) (Updatable) An optional property indicating the absolute time when this secret will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. The minimum number of days from current time is 1 day and the maximum number of days from current time is 365 days. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - (Required) The OCID of the vault where you want to create the secret.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where you want to create the secret.
* `current_version_number` - The version number of the secret version that's currently in use.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A brief description of the secret. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the secret.
* `key_id` - The OCID of the master encryption key that is used to encrypt the secret. You must specify a symmetric key to encrypt the secret during import to the vault. You cannot encrypt secrets with asymmetric keys. Furthermore, the key must exist in the vault that you specify. 
* `last_rotation_time` - A property indicating when the secret was last rotated successfully, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `lifecycle_details` - Additional information about the current lifecycle state of the secret.
* `metadata` - Additional metadata that you can use to provide context about how to use the secret or during rotation or other administrative tasks. For example, for a secret that you use to connect to a database, the additional metadata might specify the connection endpoint and the connection string. Provide additional metadata as key-value pairs. 
* `next_rotation_time` - A property indicating when the secret is scheduled to be rotated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `rotation_config` - Defines the frequency of the rotation and the information about the target system
	* `is_scheduled_rotation_enabled` - Enables auto rotation, when set to true rotationInterval must be set. 
	* `rotation_interval` - The time interval that indicates the frequency for rotating secret data, as described in ISO 8601 format. The minimum value is 1 day and maximum value is 360 days. For example, if you want to set the time interval for rotating a secret data as 30 days, the duration is expressed as "P30D." 
	* `target_system_details` - The TargetSystemDetails provides the targetSystem type and type-specific connection metadata 
		* `adb_id` - The unique identifier (OCID) for the autonomous database that Vault Secret connects to. 
		* `function_id` - The unique identifier (OCID) of the Oracle Cloud Infrastructure Functions that vault secret connects to. 
		* `target_system_type` - Unique identifier of the target system that Vault Secret connects to. 
* `rotation_status` - Additional information about the status of the secret rotation
* `secret_name` - The user-friendly name of the secret. Avoid entering confidential information.
* `secret_rules` - A list of rules that control how the secret is used and managed.
	* `is_enforced_on_deleted_secret_versions` - A property indicating whether the rule is applied even if the secret version with the content you are trying to reuse was deleted. 
	* `is_secret_content_retrieval_blocked_on_expiry` - A property indicating whether to block retrieval of the secret content, on expiry. The default is false. If the secret has already expired and you would like to retrieve the secret contents, you need to edit the secret rule to disable this property, to allow reading the secret content. 
	* `rule_type` - The type of rule, which either controls when the secret contents expire or whether they can be reused.
	* `secret_version_expiry_interval` - A property indicating how long the secret contents will be considered valid, expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. The secret needs to be updated when the secret content expires. The timer resets after you update the secret contents. The minimum value is 1 day and the maximum value is 90 days for this property. Currently, only intervals expressed in days are supported. For example, pass `P3D` to have the secret version expire every 3 days. 
	* `time_of_absolute_expiry` - An optional property indicating the absolute time when this secret will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. The minimum number of days from current time is 1 day and the maximum number of days from current time is 365 days. Example: `2019-04-03T21:10:29.600Z` 
* `state` - The current lifecycle state of the secret.
* `time_created` - A property indicating when the secret was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_current_version_expiry` - An optional property indicating when the current secret version will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the secret, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the Vault in which the secret exists

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Secret
	* `update` - (Defaults to 20 minutes), when updating the Secret
	* `delete` - (Defaults to 20 minutes), when destroying the Secret


## Import

Secrets can be imported using the `id`, e.g.

```
$ terraform import oci_vault_secret.test_secret "id"
```

