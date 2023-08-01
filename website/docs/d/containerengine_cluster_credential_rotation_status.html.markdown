---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_credential_rotation_status"
sidebar_current: "docs-oci-datasource-containerengine-cluster_credential_rotation_status"
description: |-
  Provides details about a specific Cluster Credential Rotation Status in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_credential_rotation_status
This data source provides details about a specific Cluster Credential Rotation Status resource in Oracle Cloud Infrastructure Container Engine service.

Get cluster credential rotation status.

## Example Usage

```hcl
data "oci_containerengine_cluster_credential_rotation_status" "test_cluster_credential_rotation_status" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `status` - Credential rotation status of a kubernetes cluster IN_PROGRESS: Issuing new credentials to kubernetes cluster control plane and worker nodes or retiring old credentials from kubernetes cluster control plane and worker nodes. WAITING: Waiting for customer to invoke the complete rotation action or the automcatic complete rotation action. COMPLETED: New credentials are functional on kuberentes cluster. 
* `status_details` - Details of a kuberenetes cluster credential rotation status: ISSUING_NEW_CREDENTIALS: Credential rotation is in progress. Starting to issue new credentials to kubernetes cluster control plane and worker nodes. NEW_CREDENTIALS_ISSUED: New credentials are added. At this stage cluster has both old and new credentials and is awaiting old credentials retirement. RETIRING_OLD_CREDENTIALS: Retirement of old credentials is in progress. Starting to remove old credentials from kubernetes cluster control plane and worker nodes. COMPLETED: Credential rotation is complete. Old credentials are retired. 
* `time_auto_completion_scheduled` - The time by which retirement of old credentials should start.

