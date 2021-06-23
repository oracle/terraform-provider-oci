---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_bucket"
sidebar_current: "docs-oci-resource-objectstorage-bucket"
description: |-
  Provides the Bucket resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_bucket
This resource provides the Bucket resource in Oracle Cloud Infrastructure Object Storage service.

Creates a bucket in the given namespace with a bucket name and optional user-defined metadata. Avoid entering
confidential information in bucket names.


## Example Usage

```hcl
resource "oci_objectstorage_bucket" "test_bucket" {
	#Required
	compartment_id = var.compartment_id
	name = var.bucket_name
	namespace = var.bucket_namespace

	#Optional
	access_type = var.bucket_access_type
	auto_tiering = var.bucket_auto_tiering
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	kms_key_id = oci_kms_key.test_key.id
	metadata = var.bucket_metadata
	object_events_enabled = var.bucket_object_events_enabled
	storage_tier = var.bucket_storage_tier
	retention_rules {
        display_name = var.retention_rule_display_name
        duration {
            #Required
            time_amount = var.retention_rule_duration_time_amount
            time_unit = var.retention_rule_duration_time_unit
        }
        time_rule_locked = var.retention_rule_time_rule_locked
    }
	versioning = var.bucket_versioning
}
```

## Argument Reference

The following arguments are supported:

* `access_type` - (Optional) (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations. 
* `auto_tiering` - (Optional) (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering `Disabled`. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to `InfrequentAccess` are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects. 
* `compartment_id` - (Required) (Updatable) The ID of the compartment in which to create the bucket.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `kms_key_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key. 
* `metadata` - (Optional) (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
* `name` - (Required) The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `object_events_enabled` - (Optional) (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
* `retention_rules` - (Optional) (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
    * `display_name` - (Required) A user-specified name for the retention rule. Names can be helpful in identifying retention rules. The name should be unique. This attribute is a forcenew attribute  
    * `duration` - (Optional) (Updatable) 
        * `time_amount` - (Required) (Updatable) The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp. 
        * `time_unit` - (Required) (Updatable) The unit that should be used to interpret timeAmount.
    * `time_rule_locked` - (Optional) (Updatable) The date and time as per [RFC 3339](https://tools.ietf.org/html/rfc3339) after which this rule is locked and can only be deleted by deleting the bucket. Once a rule is locked, only increases in the duration are allowed and no other properties can be changed. This property cannot be updated for rules that are in a locked state. Specifying it when a duration is not specified is considered an error. 
* `storage_tier` - (Optional) The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created. 
* `versioning` - (Optional) (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning `Disabled`. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_type` - The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations. 
* `approximate_count` - The approximate number of objects in the bucket. Count statistics are reported periodically. You will see a lag between what is displayed and the actual object count. 
* `approximate_size` - The approximate total size in bytes of all objects in the bucket. Size statistics are reported periodically. You will see a lag between what is displayed and the actual size of the bucket. 
* `auto_tiering` - The auto tiering status on the bucket. A bucket is created with auto tiering `Disabled` by default. For auto tiering `InfrequentAccess`, objects are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects. 
* `bucket_id` - The OCID of the bucket which is a Oracle assigned unique identifier for this resource type (bucket). `bucket_id` cannot be used for bucket lookup.
* `compartment_id` - The compartment ID in which the bucket is authorized.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the bucket.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `etag` - The entity tag (ETag) for the bucket.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_read_only` - Whether or not this bucket is read only. By default, `isReadOnly` is set to `false`. This will be set to 'true' when this bucket is configured as a destination in a replication policy. 
* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key. 
* `metadata` - Arbitrary string keys and values for user-defined metadata.
* `name` - The name of the bucket. Avoid entering confidential information. Example: my-new-bucket1 
* `namespace` - The Object Storage namespace in which the bucket resides.
* `object_events_enabled` - Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm). 
* `object_lifecycle_policy_etag` - The entity tag (ETag) for the live object lifecycle policy on the bucket.
* `retention_rules` - User specified list of retention rules for the bucket. 
    * `display_name` - User specified name for the retention rule.
    * `duration` - 
        * `time_amount` - The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp. 
        * `time_unit` - The unit that should be used to interpret timeAmount.
    * `retention_rule_id` - Unique identifier for the retention rule.
    * `time_created` - The date and time that the retention rule was created as per [RFC3339](https://tools.ietf.org/html/rfc3339).
    * `time_modified` - The date and time that the retention rule was modified as per [RFC3339](https://tools.ietf.org/html/rfc3339).
    * `time_rule_locked` - The date and time as per [RFC 3339](https://tools.ietf.org/html/rfc3339) after which this rule becomes locked. and can only be deleted by deleting the bucket. 
* `replication_enabled` - Whether or not this bucket is a replication source. By default, `replicationEnabled` is set to `false`. This will be set to 'true' when you create a replication policy for the bucket. 
* `storage_tier` - The storage tier type assigned to the bucket. A bucket is set to `Standard` tier by default, which means objects uploaded or copied to the bucket will be in the standard storage tier. When the `Archive` tier type is set explicitly for a bucket, objects uploaded or copied to the bucket will be stored in archive storage. The `storageTier` property is immutable after bucket is created. 
* `time_created` - The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
* `versioning` - The versioning status on the bucket. A bucket is created with versioning `Disabled` by default. For versioning `Enabled`, objects are protected from overwrites and deletes, by maintaining their version history. When versioning is `Suspended`, the previous versions will still remain but new versions will no longer be created when overwitten or deleted. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bucket
	* `update` - (Defaults to 20 minutes), when updating the Bucket
	* `delete` - (Defaults to 20 minutes), when destroying the Bucket


## Import

Buckets can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_bucket.test_bucket "n/{namespaceName}/b/{bucketName}" 
```

