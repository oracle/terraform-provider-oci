---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_sender_invitation"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-sender_invitation"
description: |-
  Provides details about a specific Sender Invitation in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_sender_invitation
This data source provides details about a specific Sender Invitation resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Gets information about the sender invitation.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_sender_invitation" "test_sender_invitation" {
	#Required
	sender_invitation_id = oci_tenantmanagercontrolplane_sender_invitation.test_sender_invitation.id
}
```

## Argument Reference

The following arguments are supported:

* `sender_invitation_id` - (Required) OCID of the sender invitation to retrieve.


## Attributes Reference

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

