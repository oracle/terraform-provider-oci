---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_sender_invitations"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-sender_invitations"
description: |-
  Provides the list of Sender Invitations in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_sender_invitations
This data source provides the list of Sender Invitations in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Return a (paginated) list of sender invitations.


## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_sender_invitations" "test_sender_invitations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.sender_invitation_display_name
	recipient_tenancy_id = oci_identity_tenancy.test_tenancy.id
	state = var.sender_invitation_state
	status = var.sender_invitation_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `recipient_tenancy_id` - (Optional) The tenancy that the invitation is addressed to.
* `state` - (Optional) The lifecycle state of the resource.
* `status` - (Optional) The status of the sender invitation.


## Attributes Reference

The following attributes are exported:

* `sender_invitation_collection` - The list of sender_invitation_collection.

### SenderInvitation Reference

The following attributes are exported:

* `compartment_id` - OCID of the sender tenancy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-created name to describe the invitation. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID of the sender invitation.
* `recipient_email_address` - Email address of the recipient.
* `recipient_invitation_id` - OCID of the corresponding recipient invitation.
* `recipient_tenancy_id` - OCID of the recipient tenancy.
* `state` - Lifecycle state of the sender invitation.
* `status` - Status of the sender invitation.
* `subjects` - The list of subjects the invitation contains.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date and time when the sender invitation was created.
* `time_updated` - Date and time when the sender invitation was last updated.

