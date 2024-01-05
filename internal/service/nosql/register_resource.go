// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_nosql_index", NosqlIndexResource())
	tfresource.RegisterResource("oci_nosql_table", NosqlTableResource())
	tfresource.RegisterResource("oci_nosql_table_replica", NosqlTableReplicaResource())
}
