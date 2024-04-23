// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cluster_placement_groups

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_cluster_placement_groups_cluster_placement_group", ClusterPlacementGroupsClusterPlacementGroupResource())
}
