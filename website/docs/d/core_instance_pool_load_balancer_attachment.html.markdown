---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool_load_balancer_attachment"
sidebar_current: "docs-oci-datasource-core-instance_pool_load_balancer_attachment"
description: |-
  Provides details about a specific Instance Pool Load Balancer Attachment in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_pool_load_balancer_attachment
This data source provides details about a specific Instance Pool Load Balancer Attachment resource in Oracle Cloud Infrastructure Core service.

Gets information about a load balancer that is attached to the specified instance pool.


## Example Usage

```hcl
data "oci_core_instance_pool_load_balancer_attachment" "test_instance_pool_load_balancer_attachment" {
	#Required
	instance_pool_id = oci_core_instance_pool.test_instance_pool.id
	instance_pool_load_balancer_attachment_id = oci_core_instance_pool_load_balancer_attachment.test_instance_pool_load_balancer_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
* `instance_pool_load_balancer_attachment_id` - (Required) The OCID of the load balancer attachment.


## Attributes Reference

The following attributes are exported:

* `backend_set_name` - The name of the backend set on the load balancer.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attachment.
* `instance_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool of the load balancer attachment. 
* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attached to the instance pool. 
* `port` - The port value used for the backends.
* `state` - The status of the interaction between the instance pool and the load balancer.
* `vnic_selection` - Indicates which VNIC on each instance in the instance pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool. 

