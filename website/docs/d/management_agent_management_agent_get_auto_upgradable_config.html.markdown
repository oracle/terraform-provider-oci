---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_get_auto_upgradable_config"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_get_auto_upgradable_config"
description: |-
  Provides details about a specific Management Agent Get Auto Upgradable Config in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_get_auto_upgradable_config
This data source provides details about a specific Management Agent Get Auto Upgradable Config resource in Oracle Cloud Infrastructure Management Agent service.

Get the AutoUpgradable configuration for all agents in a tenancy.
The supplied compartmentId must be a tenancy root.


## Example Usage

```hcl
data "oci_management_agent_management_agent_get_auto_upgradable_config" "test_management_agent_get_auto_upgradable_config" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.


## Attributes Reference

The following attributes are exported:

* `is_agent_auto_upgradable` - true if the agents can be upgraded automatically; false if they must be upgraded manually.

