---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_versions"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_versions"
description: |-
  Provides the list of Deployment Versions in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_versions
This data source provides the list of Deployment Versions in Oracle Cloud Infrastructure Golden Gate service.

Returns the list of available deployment versions.


## Example Usage

```hcl
data "oci_golden_gate_deployment_versions" "test_deployment_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	deployment_type = var.deployment_version_deployment_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `deployment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources. 
* `deployment_type` - (Optional) The type of deployment, the value determines the exact 'type' of the service executed in the deployment. Default value is DATABASE_ORACLE. 


## Attributes Reference

The following attributes are exported:

* `deployment_version_collection` - The list of deployment_version_collection.

### DeploymentVersion Reference

The following attributes are exported:

* `items` - Array of DeploymentVersionSummary. 
	* `deployment_type` - The type of deployment, the value determines the exact 'type' of service executed in the Deployment. NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.  Its use is discouraged in favor of the equivalent 'DATABASE_ORACLE' value. 
	* `is_security_fix` - Indicates if OGG release contains security fix. 
	* `ogg_version` - Version of OGG 
	* `release_type` - The type of release. 
	* `time_released` - The time the resource was released. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
	* `time_supported_until` - The time until OGG version is supported. After this date has passed OGG version will not be available anymore. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

