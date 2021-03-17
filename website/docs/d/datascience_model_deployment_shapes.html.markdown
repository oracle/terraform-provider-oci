---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_deployment_shapes"
sidebar_current: "docs-oci-datasource-datascience-model_deployment_shapes"
description: |-
  Provides the list of Model Deployment Shapes in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_deployment_shapes
This data source provides the list of Model Deployment Shapes in Oracle Cloud Infrastructure Datascience service.

Lists the valid model deployment shapes.

## Example Usage

```hcl
data "oci_datascience_model_deployment_shapes" "test_model_deployment_shapes" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `model_deployment_shapes` - The list of model_deployment_shapes.

### ModelDeploymentShape Reference

The following attributes are exported:

* `core_count` - The number of cores associated with this model deployment shape. 
* `memory_in_gbs` - The amount of memory in GBs associated with this model deployment shape. 
* `name` - The name of the model deployment shape. 

