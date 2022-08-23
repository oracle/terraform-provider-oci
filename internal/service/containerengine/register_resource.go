// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_containerengine_cluster", ContainerengineClusterResource())
	tfresource.RegisterResource("oci_containerengine_node_pool", ContainerengineNodePoolResource())
}
