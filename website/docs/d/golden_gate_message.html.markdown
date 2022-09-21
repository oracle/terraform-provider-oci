---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_message"
sidebar_current: "docs-oci-datasource-golden_gate-message"
description: |-
  Provides details about a specific Message in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_message
This data source provides details about a specific Message resource in Oracle Cloud Infrastructure Golden Gate service.

Lists the DeploymentMessages for a deployment. The sorting order is not important. By default first will be Upgrade message, next Exception message and then Storage Utilization message.


## Example Usage

```hcl
data "oci_golden_gate_message" "test_message" {
	#Required
	deployment_id = oci_golden_gate_deployment.test_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) A unique Deployment identifier. 


## Attributes Reference

The following attributes are exported:

* `items` - An array of DeploymentMessages. 
	* `deployment_message` - The deployment Message in plain text with optional HTML anchor tags. 
	* `deployment_message_status` - The deployment Message Status. 
	* `id` - The deployment Message Id. 

