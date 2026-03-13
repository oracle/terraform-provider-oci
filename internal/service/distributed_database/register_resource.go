// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_distributed_database_distributed_autonomous_database", DistributedDatabaseDistributedAutonomousDatabaseResource())
	tfresource.RegisterResource("oci_distributed_database_distributed_database", DistributedDatabaseDistributedDatabaseResource())
	tfresource.RegisterResource("oci_distributed_database_distributed_database_private_endpoint", DistributedDatabaseDistributedDatabasePrivateEndpointResource())
}
