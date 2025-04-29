---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_software_update_action"
sidebar_current: "docs-oci-resource-bds-bds_instance_software_update_action"
description: |-
    Install the specified software update to this Big Data cluster
---

# oci_bds_bds_instance_software_update
This resource installs the specified software update to the nodes of the Oracle Cloud Infrastructure Big Data Service cluster.

Install Software Update of the specified SoftwareUpdateId to this BDS cluster's nodes.


## Example Usage
```hcl
resource "oci_bds_bds_instance_software_update_action" "test_bds_instance_software_update_action" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	software_update_key = var.bds_instance_software_update_action_software_update_key
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `software_update_key` - (Required) The unique identifier of the software update, software_update_key can be retrieved from the output of get and list datasources

