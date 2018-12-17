---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_attachments"
sidebar_current: "docs-oci-datasource-core-drg_attachments"
description: |-
  Provides the list of Drg Attachments in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_attachments
This data source provides the list of Drg Attachments in Oracle Cloud Infrastructure Core service.

Lists the `DrgAttachment` objects for the specified compartment. You can filter the
results by VCN or DRG.


## Example Usage

```hcl
data "oci_core_drg_attachments" "test_drg_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	drg_id = "${oci_core_drg.test_drg.id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `drg_id` - (Optional) The OCID of the DRG.
* `vcn_id` - (Optional) The OCID of the VCN.


## Attributes Reference

The following attributes are exported:

* `drg_attachments` - The list of drg_attachments.

### DrgAttachment Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DRG attachment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG.
* `id` - The DRG attachment's Oracle ID (OCID).
* `route_table_id` - The OCID of the route table the DRG attachment is using. For information about why you would associate a route table with a DRG attachment, see [Advanced Scenario: Transit Routing](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm). 
* `state` - The DRG attachment's current state.
* `time_created` - The date and time the DRG attachment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN.

