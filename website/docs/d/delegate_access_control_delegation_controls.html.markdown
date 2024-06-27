---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegation_controls"
sidebar_current: "docs-oci-datasource-delegate_access_control-delegation_controls"
description: |-
  Provides the list of Delegation Controls in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_delegation_controls
This data source provides the list of Delegation Controls in Oracle Cloud Infrastructure Delegate Access Control service.

Lists the Delegation Controls in the compartment.


## Example Usage

```hcl
data "oci_delegate_access_control_delegation_controls" "test_delegation_controls" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.delegation_control_display_name
	resource_id = oci_cloud_guard_resource.test_resource.id
	resource_type = var.delegation_control_resource_type
	state = var.delegation_control_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Delegation Control resources that match the given display name.
* `resource_id` - (Optional) A filter to return Delegation Control resources that match the given resource ID.
* `resource_type` - (Optional) A filter to return only resources that match the given resource type.
* `state` - (Optional) A filter to return only Delegation Control resources whose lifecycleState matches the given Delegation Control lifecycle state.


## Attributes Reference

The following attributes are exported:

* `delegation_control_summary_collection` - The list of delegation_control_summary_collection.

### DelegationControl Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the Delegation Control.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `delegation_subscription_ids` - List of Delegation Subscription OCID that are allowed for this Delegation Control. The allowed subscriptions will determine the available Service Provider Actions. Only support operators for the allowed subscriptions are allowed to create Delegated Resource Access Request.
* `description` - Description of the Delegation Control.
* `display_name` - Name of the Delegation Control. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Delegation Control.
* `is_auto_approve_during_maintenance` - Set to true to allow all Delegated Resource Access Request to be approved automatically during maintenance.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `notification_message_format` - The format of the Oracle Cloud Infrastructure Notification messages for this Delegation Control.
* `notification_topic_id` - The OCID of the Oracle Cloud Infrastructure Notification topic to publish messages related to this Delegation Control.
* `num_approvals_required` - number of approvals required.
* `pre_approved_service_provider_action_names` - List of pre-approved Service Provider Action names. The list of pre-defined Service Provider Actions can be obtained from the ListServiceProviderActions API. Delegated Resource Access Requests associated with a resource governed by this Delegation Control will be automatically approved if the Delegated Resource Access Request only contain Service Provider Actions in the pre-approved list. 
* `resource_ids` - The OCID of the selected resources that this Delegation Control is applicable to.
* `resource_type` - Resource type for which the Delegation Control is applicable to.
* `state` - The current lifecycle state of the Delegation Control.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Delegation Control was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_deleted` - Time when the Delegation Control was deleted expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format, e.g. '2020-05-22T21:10:29.600Z'. Note a deleted Delegation Control still stays in the system, so that you can still audit Service Provider Actions associated with Delegated Resource Access Requests raised on target resources governed by the deleted Delegation Control. 
* `time_updated` - Time when the Delegation Control was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `vault_id` - The OCID of the Oracle Cloud Infrastructure Vault that will store the secrets containing the SSH keys to access the resource governed by this Delegation Control by Delegate Access Control Service. This property is required when resourceType is CLOUDVMCLUSTER. Delegate Access Control Service will generate the SSH keys and store them as secrets in the Oracle Cloud Infrastructure Vault.
* `vault_key_id` - The OCID of the Master Encryption Key in the Oracle Cloud Infrastructure Vault specified by vaultId. This key will be used to encrypt the SSH keys to access the resource governed by this Delegation Control by Delegate Access Control Service. This property is required when resourceType is CLOUDVMCLUSTER.

