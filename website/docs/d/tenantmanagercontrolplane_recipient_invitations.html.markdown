---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_recipient_invitations"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-recipient_invitations"
description: |-
  Provides the list of Recipient Invitations in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_recipient_invitations
This data source provides the list of Recipient Invitations in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Return a (paginated) list of recipient invitations.


## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_recipient_invitations" "test_recipient_invitations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	sender_tenancy_id = oci_identity_tenancy.test_tenancy.id
	state = var.recipient_invitation_state
	status = var.recipient_invitation_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `sender_tenancy_id` - (Optional) The tenancy that sent the invitation.
* `state` - (Optional) The lifecycle state of the resource.
* `status` - (Optional) The status of the recipient invitation.


## Attributes Reference

The following attributes are exported:

* `recipient_invitation_collection` - The list of recipient_invitation_collection.

### RecipientInvitation Reference

The following attributes are exported:

* `compartment_id` - OCID of the recipient tenancy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-created name to describe the invitation. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID of the recipient invitation.
* `recipient_email_address` - Email address of the recipient.
* `sender_invitation_id` - OCID of the corresponding sender invitation.
* `sender_tenancy_id` - OCID of the sender tenancy.
* `state` - Lifecycle state of the recipient invitation.
* `status` - Status of the recipient invitation.
* `subjects` - The list of subjects the invitation contains.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date and time when the recipient invitation was created.
* `time_updated` - Date and time when the recipient invitation was last updated.

