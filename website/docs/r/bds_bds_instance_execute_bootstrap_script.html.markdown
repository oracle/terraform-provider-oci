---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_execute_bootstrap_script_action"
sidebar_current: "docs-oci-resource-bds_bds_instance_execute_bootstrap_script_action"
description: |-
  Enables execution of a Bootstrap Action on an existing BDS instance in the Oracle Cloud Infrastructure Big Data Service
---

# oci_bds_bds_instance_execute_bootstrap_script_action

Executes a bootstrap script on an existing Oracle Cloud Infrastructure (OCI) Big Data Service (BDS) instance. This resource triggers the OCI API to run an arbitrary bootstrap script, allowing custom setup or initialization logic to be applied to a specified BDS cluster. This is a one-time action resource; it does not represent a long-lived resource in the BDS control plane.
Api doc link for the resource: https://docs.oracle.com/en-us/iaas/api/#/en/bigdata/20190531/BdsInstance/ExecuteBootstrapScript

**Example Usage**

```hcl
resource "oci_bds_bds_instance_execute_bootstrap_script_action" "example" {
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.cluster_admin_password

  # Optional: Specify a bootstrap script to run
  bootstrap_script_url   = var.bootstrap_script_url
}
```

## Argument Reference

The following arguments are supported:

- `bds_instance_id` (**Required**, string)  
  The OCID of the target BDS instance on which to execute the bootstrap script.

- `cluster_admin_password` (Optional, string, sensitive)  
  The cluster admin password for the BDS instance. This field is marked sensitive and will not be displayed in logs or plan output.

- `secret_id` (Optional, string, sensitive)  
  The secret id for the BDS instance. This field is marked sensitive and will not be displayed in logs or plan output.

- `bootstrap_script_url` (Optional, string)  
  The Object Storage URL of a script to execute on the BDS cluster. If not specified, no custom script will be run.

## Attributes Reference

This resource exports the following attribute:

- `id` – The identifier representing the completed execution of the bootstrap script action.

## Import

This resource supports importing by ID, but since it represents a completed one-time action (not a persistent resource), import is generally not applicable beyond tracking execution history.