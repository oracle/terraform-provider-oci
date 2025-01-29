---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_recipient_invitation"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-recipient_invitation"
description: |-
  Provides details about a specific Recipient Invitation in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_recipient_invitation
This data source provides details about a specific Recipient Invitation resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the recipient invitation.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_recipient_invitation" "test_recipient_invitation" {
	#Required
	recipient_invitation_id = oci_tenantmanagercontrolplane_recipient_invitation.test_recipient_invitation.id
}
```

## Argument Reference

The following arguments are supported:

* `recipient_invitation_id` - (Required) OCID of the recipient invitation to retrieve.


## Attributes Reference

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

