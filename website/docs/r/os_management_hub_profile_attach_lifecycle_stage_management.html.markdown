---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile_attach_lifecycle_stage_management"
sidebar_current: "docs-oci-resource-os_management_hub-profile_attach_lifecycle_stage_management"
description: |-
  Provides the Profile Attach Lifecycle Stage Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_profile_attach_lifecycle_stage_management
This resource provides the Profile Attach Lifecycle Stage Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Attaches the specified lifecycle stage to a profile.


## Example Usage

```hcl
resource "oci_os_management_hub_profile_attach_lifecycle_stage_management" "test_profile_attach_lifecycle_stage_management" {
	#Required
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id
	profile_id = oci_os_management_hub_profile.test_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_stage_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage that the instance will be associated with.
* `profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Profile Attach Lifecycle Stage Management
	* `update` - (Defaults to 20 minutes), when updating the Profile Attach Lifecycle Stage Management
	* `delete` - (Defaults to 20 minutes), when destroying the Profile Attach Lifecycle Stage Management


## Import

ProfileAttachLifecycleStageManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_profile_attach_lifecycle_stage_management.test_profile_attach_lifecycle_stage_management "id"
```

