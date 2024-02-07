---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_cdn_config"
sidebar_current: "docs-oci-resource-media_services-stream_cdn_config"
description: |-
  Provides the Stream Cdn Config resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_stream_cdn_config
This resource provides the Stream Cdn Config resource in Oracle Cloud Infrastructure Media Services service.

Creates a new CDN Configuration.


## Example Usage

```hcl
resource "oci_media_services_stream_cdn_config" "test_stream_cdn_config" {
	#Required
	config {
		#Required
		type = var.stream_cdn_config_config_type

		#Optional
		edge_hostname = var.stream_cdn_config_config_edge_hostname
		edge_path_prefix = var.stream_cdn_config_config_edge_path_prefix
		edge_token_key = var.stream_cdn_config_config_edge_token_key
		edge_token_salt = var.stream_cdn_config_config_edge_token_salt
		is_edge_token_auth = var.stream_cdn_config_config_is_edge_token_auth
		origin_auth_secret_key_a = var.stream_cdn_config_config_origin_auth_secret_key_a
		origin_auth_secret_key_b = var.stream_cdn_config_config_origin_auth_secret_key_b
		origin_auth_secret_key_nonce_a = var.stream_cdn_config_config_origin_auth_secret_key_nonce_a
		origin_auth_secret_key_nonce_b = var.stream_cdn_config_config_origin_auth_secret_key_nonce_b
		origin_auth_sign_encryption = var.stream_cdn_config_config_origin_auth_sign_encryption
		origin_auth_sign_type = var.stream_cdn_config_config_origin_auth_sign_type
	}
	display_name = var.stream_cdn_config_display_name
	distribution_channel_id = oci_mysql_channel.test_channel.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_enabled = var.stream_cdn_config_is_enabled
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.stream_cdn_config_locks_type

		#Optional
		message = var.stream_cdn_config_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.stream_cdn_config_locks_time_created
	}
}
```

## Argument Reference

The following arguments are supported:

* `config` - (Required) (Updatable) Base fields of the StreamCdnConfig configuration object.
	* `edge_hostname` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The hostname of the CDN edge server to use when building CDN URLs.
	* `edge_path_prefix` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The path to prepend when building CDN URLs.
	* `edge_token_key` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The encryption key to use for edge token authentication.
	* `edge_token_salt` - (Applicable when type=AKAMAI_MANUAL) (Updatable) Salt to use when encrypting authentication token.
	* `is_edge_token_auth` - (Applicable when type=AKAMAI_MANUAL) (Updatable) Whether token authentication should be used at the CDN edge.
	* `origin_auth_secret_key_a` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The shared secret key A, two for errorless key rotation.
	* `origin_auth_secret_key_b` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The shared secret key B, two for errorless key rotation.
	* `origin_auth_secret_key_nonce_a` - (Applicable when type=AKAMAI_MANUAL) (Updatable) Nonce identifier for originAuthSecretKeyA (used to determine key used to sign).
	* `origin_auth_secret_key_nonce_b` - (Applicable when type=AKAMAI_MANUAL) (Updatable) Nonce identifier for originAuthSecretKeyB (used to determine key used to sign).
	* `origin_auth_sign_encryption` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The type of encryption used to compute the signature.
	* `origin_auth_sign_type` - (Applicable when type=AKAMAI_MANUAL) (Updatable) The type of data used to compute the signature.
	* `type` - (Required) (Updatable) The name of the CDN configuration type.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) CDN Config display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `distribution_channel_id` - (Required) Distribution Channel Identifier.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_enabled` - (Optional) (Updatable) Whether publishing to CDN is enabled.
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `config` - Base fields of the StreamCdnConfig configuration object.
	* `edge_hostname` - The hostname of the CDN edge server to use when building CDN URLs.
	* `edge_path_prefix` - The path to prepend when building CDN URLs.
	* `edge_token_key` - The encryption key to use for edge token authentication.
	* `edge_token_salt` - Salt to use when encrypting authentication token.
	* `is_edge_token_auth` - Whether token authentication should be used at the CDN edge.
	* `origin_auth_secret_key_a` - The shared secret key A, two for errorless key rotation.
	* `origin_auth_secret_key_b` - The shared secret key B, two for errorless key rotation.
	* `origin_auth_secret_key_nonce_a` - Nonce identifier for originAuthSecretKeyA (used to determine key used to sign).
	* `origin_auth_secret_key_nonce_b` - Nonce identifier for originAuthSecretKeyB (used to determine key used to sign).
	* `origin_auth_sign_encryption` - The type of encryption used to compute the signature.
	* `origin_auth_sign_type` - The type of data used to compute the signature.
	* `type` - The name of the CDN configuration type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The CDN Configuration identifier or display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `distribution_channel_id` - Distribution Channel Identifier.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `is_enabled` - Whether publishing to CDN is enabled.
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `state` - The current state of the CDN Configuration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the CDN Config was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the CDN Config was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Cdn Config
	* `update` - (Defaults to 20 minutes), when updating the Stream Cdn Config
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Cdn Config


## Import

StreamCdnConfigs can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_stream_cdn_config.test_stream_cdn_config "id"
```

