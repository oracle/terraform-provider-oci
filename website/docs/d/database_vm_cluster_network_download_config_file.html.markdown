---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_network_download_config_file"
sidebar_current: "docs-oci-datasource-database-vm_cluster_network_download_config_file"
description: |-
  Provides details about a specific Vm Cluster Network Download Config File in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_network_download_config_file
This data source provides details about a specific Vm Cluster Network Download Config File resource in Oracle Cloud Infrastructure Database service.

Downloads the configuration file for the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_vm_cluster_network_download_config_file" "test_vm_cluster_network_download_config_file" {
	#Required
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	vm_cluster_network_id = oci_database_vm_cluster_network.test_vm_cluster_network.id

	#Optional
	base64_encode_content = "false"
}
```

## Argument Reference

The following arguments are supported:

* `base64_encode_content` - (Optional) Encodes the downloaded txt config in base64. It is recommended to set this to `true` to avoid corrupting the file in Terraform state. The default value is `false`.
* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `vm_cluster_network_id` - (Required) The VM cluster network [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded config file for exadata infrastructure. If `base64_encode_content` is set to `true`, then this content will be base64 encoded.

