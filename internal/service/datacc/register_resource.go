// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_datacc_infrastructure", DataccInfrastructureResource())
	tfresource.RegisterResource("oci_datacc_vm_cluster_network", DataccVmClusterNetworkResource())
	tfresource.RegisterResource("oci_datacc_vm_instance", DataccVmInstanceResource())
}
