---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_object_lifecycle_policy"
sidebar_current: "docs-oci-resource-objectstorage-object_lifecycle_policy"
description: |-
  Provides the Object Lifecycle Policy resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_object_lifecycle_policy
This resource provides the Object Lifecycle Policy resource in Oracle Cloud Infrastructure Object Storage service.

Creates or replaces the object lifecycle policy for the bucket.


## Example Usage

```hcl
resource "oci_objectstorage_object_lifecycle_policy" "test_object_lifecycle_policy" {
	#Required
	bucket = var.object_lifecycle_policy_bucket
	namespace = var.object_lifecycle_policy_namespace

	#Optional
	rules {
		#Required
		action = var.object_lifecycle_policy_rules_action
		is_enabled = var.object_lifecycle_policy_rules_is_enabled
		name = var.object_lifecycle_policy_rules_name
		time_amount = var.object_lifecycle_policy_rules_time_amount
		time_unit = var.object_lifecycle_policy_rules_time_unit

		#Optional
		object_name_filter {

			#Optional
			exclusion_patterns = var.object_lifecycle_policy_rules_object_name_filter_exclusion_patterns
			inclusion_patterns = var.object_lifecycle_policy_rules_object_name_filter_inclusion_patterns
			inclusion_prefixes = var.object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes
		}
		target = var.object_lifecycle_policy_rules_target
	}
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `rules` - (Optional) (Updatable) The bucket's set of lifecycle policy rules.
	* `action` - (Required) (Updatable) The action of the object lifecycle policy rule. Rules using the action 'ARCHIVE' move objects from Standard and InfrequentAccess storage tiers into the [Archive storage tier](https://docs.cloud.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm). Rules using the action 'INFREQUENT_ACCESS' move objects from Standard storage tier into the Infrequent Access Storage tier. Objects that are already in InfrequentAccess tier or in Archive tier are left untouched. Rules using the action 'DELETE' permanently delete objects from buckets. Rules using 'ABORT' abort the uncommitted multipart-uploads and permanently delete their parts from buckets. 
	* `is_enabled` - (Required) (Updatable) A Boolean that determines whether this rule is currently enabled.
	* `name` - (Required) (Updatable) The name of the lifecycle rule to be applied.
	* `object_name_filter` - (Optional) (Updatable) A filter that compares object names to a set of prefixes or patterns to determine if a rule applies to a given object. The filter can contain include glob patterns, exclude glob patterns and inclusion prefixes. The inclusion prefixes property is kept for backward compatibility. It is recommended to use inclusion patterns instead of prefixes. Exclusions take precedence over inclusions. 
		* `exclusion_patterns` - (Optional) (Updatable) An array of glob patterns to match the object names to exclude. An empty array is ignored. Exclusion patterns take precedence over inclusion patterns. A Glob pattern is a sequence of characters to match text. Any character that appears in the pattern, other than the special pattern characters described below, matches itself. Glob patterns must be between 1 and 1024 characters.

			The special pattern characters have the following meanings:

			\           Escapes the following character
			*           Matches any string of characters. ?           Matches any single character . [...]       Matches a group of characters. A group of characters can be: A set of characters, for example: [Zafg9@]. This matches any character in the brackets. A range of characters, for example: [a-z]. This matches any character in the range. [a-f] is equivalent to [abcdef]. For character ranges only the CHARACTER-CHARACTER pattern is supported. [ab-yz] is not valid [a-mn-z] is not valid Character ranges can not start with ^ or : To include a '-' in the range, make it the first or last character. 
		* `inclusion_patterns` - (Optional) (Updatable) An array of glob patterns to match the object names to include. An empty array includes all objects in the bucket. Exclusion patterns take precedence over inclusion patterns. A Glob pattern is a sequence of characters to match text. Any character that appears in the pattern, other than the special pattern characters described below, matches itself. Glob patterns must be between 1 and 1024 characters.

			The special pattern characters have the following meanings:

			\           Escapes the following character
			*           Matches any string of characters. ?           Matches any single character . [...]       Matches a group of characters. A group of characters can be: A set of characters, for example: [Zafg9@]. This matches any character in the brackets. A range of characters, for example: [a-z]. This matches any character in the range. [a-f] is equivalent to [abcdef]. For character ranges only the CHARACTER-CHARACTER pattern is supported. [ab-yz] is not valid [a-mn-z] is not valid Character ranges can not start with ^ or : To include a '-' in the range, make it the first or last character. 
		* `inclusion_prefixes` - (Optional) (Updatable) An array of object name prefixes that the rule will apply to. An empty array means to include all objects. 
	* `target` - (Optional) (Updatable) The target of the object lifecycle policy rule. The values of target can be either "objects", "multipart-uploads" or "previous-object-versions". This field when declared as "objects" is used to specify ARCHIVE, INFREQUENT_ACCESS or DELETE rule for objects. This field when declared as "previous-object-versions" is used to specify ARCHIVE, INFREQUENT_ACCESS or DELETE rule for previous versions of existing objects. This field when declared as "multipart-uploads" is used to specify the ABORT (only) rule for uncommitted multipart-uploads. 
	* `time_amount` - (Required) (Updatable) Specifies the age of objects to apply the rule to. The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified time. 
	* `time_unit` - (Required) (Updatable) The unit that should be used to interpret timeAmount.  Days are defined as starting and ending at midnight UTC. Years are defined as 365.2425 days long and likewise round up to the next midnight UTC. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `rules` - The live lifecycle policy on the bucket.

	For an example of this value, see the [PutObjectLifecyclePolicy API documentation](https://docs.cloud.oracle.com/iaas/api/#/en/objectstorage/20160918/ObjectLifecyclePolicy/PutObjectLifecyclePolicy). 
	* `action` - The action of the object lifecycle policy rule. Rules using the action 'ARCHIVE' move objects from Standard and InfrequentAccess storage tiers into the [Archive storage tier](https://docs.cloud.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm). Rules using the action 'INFREQUENT_ACCESS' move objects from Standard storage tier into the Infrequent Access Storage tier. Objects that are already in InfrequentAccess tier or in Archive tier are left untouched. Rules using the action 'DELETE' permanently delete objects from buckets. Rules using 'ABORT' abort the uncommitted multipart-uploads and permanently delete their parts from buckets. 
	* `is_enabled` - A Boolean that determines whether this rule is currently enabled.
	* `name` - The name of the lifecycle rule to be applied.
	* `object_name_filter` - A filter that compares object names to a set of prefixes or patterns to determine if a rule applies to a given object. The filter can contain include glob patterns, exclude glob patterns and inclusion prefixes. The inclusion prefixes property is kept for backward compatibility. It is recommended to use inclusion patterns instead of prefixes. Exclusions take precedence over inclusions. 
		* `exclusion_patterns` - An array of glob patterns to match the object names to exclude. An empty array is ignored. Exclusion patterns take precedence over inclusion patterns. A Glob pattern is a sequence of characters to match text. Any character that appears in the pattern, other than the special pattern characters described below, matches itself. Glob patterns must be between 1 and 1024 characters.

			The special pattern characters have the following meanings:

			\           Escapes the following character
			*           Matches any string of characters. ?           Matches any single character . [...]       Matches a group of characters. A group of characters can be: A set of characters, for example: [Zafg9@]. This matches any character in the brackets. A range of characters, for example: [a-z]. This matches any character in the range. [a-f] is equivalent to [abcdef]. For character ranges only the CHARACTER-CHARACTER pattern is supported. [ab-yz] is not valid [a-mn-z] is not valid Character ranges can not start with ^ or : To include a '-' in the range, make it the first or last character. 
		* `inclusion_patterns` - An array of glob patterns to match the object names to include. An empty array includes all objects in the bucket. Exclusion patterns take precedence over inclusion patterns. A Glob pattern is a sequence of characters to match text. Any character that appears in the pattern, other than the special pattern characters described below, matches itself. Glob patterns must be between 1 and 1024 characters.

			The special pattern characters have the following meanings:

			\           Escapes the following character
			*           Matches any string of characters. ?           Matches any single character . [...]       Matches a group of characters. A group of characters can be: A set of characters, for example: [Zafg9@]. This matches any character in the brackets. A range of characters, for example: [a-z]. This matches any character in the range. [a-f] is equivalent to [abcdef]. For character ranges only the CHARACTER-CHARACTER pattern is supported. [ab-yz] is not valid [a-mn-z] is not valid Character ranges can not start with ^ or : To include a '-' in the range, make it the first or last character. 
	    * `inclusion_prefixes` - An array of object name prefixes that the rule will apply to. An empty array means to include all objects. 
	* `target` - The target of the object lifecycle policy rule. The values of target can be either "objects", "multipart-uploads" or "previous-object-versions". This field when declared as "objects" is used to specify ARCHIVE, INFREQUENT_ACCESS or DELETE rule for objects. This field when declared as "previous-object-versions" is used to specify ARCHIVE, INFREQUENT_ACCESS or DELETE rule for previous versions of existing objects. This field when declared as "multipart-uploads" is used to specify the ABORT (only) rule for uncommitted multipart-uploads. 
	* `time_amount` - Specifies the age of objects to apply the rule to. The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified time. 
	* `time_unit` - The unit that should be used to interpret timeAmount.  Days are defined as starting and ending at midnight UTC. Years are defined as 365.2425 days long and likewise round up to the next midnight UTC. 
* `time_created` - The date and time the object lifecycle policy was created, as described in [RFC 3339](https://tools.ietf.org/html/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Object Lifecycle Policy
	* `update` - (Defaults to 20 minutes), when updating the Object Lifecycle Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Object Lifecycle Policy


## Import

ObjectLifecyclePolicies can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy "n/{namespaceName}/b/{bucketName}/l" 
```

