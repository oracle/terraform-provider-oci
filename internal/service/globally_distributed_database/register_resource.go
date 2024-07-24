// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globally_distributed_database

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_globally_distributed_database_private_endpoint", GloballyDistributedDatabasePrivateEndpointResource())
	tfresource.RegisterResource("oci_globally_distributed_database_sharded_database", GloballyDistributedDatabaseShardedDatabaseResource())
}
