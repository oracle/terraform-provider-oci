---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_deployment"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_deployment"
description: |-
  Provides details about a specific Security Policy Deployment in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_deployment
This data source provides details about a specific Security Policy Deployment resource in Oracle Cloud Infrastructure Data Safe service.

Gets a security policy deployment by identifier.

## Example Usage

```hcl
data "oci_data_safe_security_policy_deployment" "test_security_policy_deployment" {
	#Required
	security_policy_deployment_id = oci_data_safe_security_policy_deployment.test_security_policy_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_deployment_id` - (Required) The OCID of the security policy deployment resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the security policy deployment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security policy deployment.
* `display_name` - The display name of the security policy deployment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security policy deployment.
* `lifecycle_details` - Details about the current state of the security policy deployment in Data Safe.
* `security_policy_id` - The OCID of the security policy corresponding to the security policy deployment.
* `state` - The current state of the security policy deployment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target where the security policy is deployed.
* `time_created` - The time that the security policy deployment was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the security policy deployment was updated, in the format defined by RFC3339.

