---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_peers"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_peers"
description: |-
  Provides the list of Deployment Peers in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_peers
This data source provides the list of Deployment Peers in Oracle Cloud Infrastructure Golden Gate service.

Lists the local and remote peers in a deployment.


## Example Usage

```hcl
data "oci_golden_gate_deployment_peers" "test_deployment_peers" {
	#Required
	deployment_id = oci_golden_gate_deployment.test_deployment.id

	#Optional
	display_name = var.deployment_peer_display_name
	state = var.deployment_peer_state
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) A unique Deployment identifier. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `state` - (Optional) A filter to return only the resources that match the 'lifecycleState' given. 


## Attributes Reference

The following attributes are exported:

* `deployment_peer_collection` - The list of deployment_peer_collection.

### DeploymentPeer Reference

The following attributes are exported:

* `items` - An array of DeploymentPeers. 
	* `availability_domain` - The availability domain of a placement.
	* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
	* `display_name` - An object's Display Name. 
	* `fault_domain` - The fault domain of a placement.
	* `peer_role` - The type of the deployment role. 
	* `peer_type` - The type of the deployment peer. 
	* `region` - The name of the region. e.g.: us-ashburn-1 If the region is not provided, backend will default to the default region. 
	* `state` - Possible lifecycle states for deployment peer.
	* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
	* `time_role_changed` - The time of the last role change. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
	* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

