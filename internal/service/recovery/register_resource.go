// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_recovery_protected_database", RecoveryProtectedDatabaseResource())
	tfresource.RegisterResource("oci_recovery_protection_policy", RecoveryProtectionPolicyResource())
	tfresource.RegisterResource("oci_recovery_recovery_service_subnet", RecoveryRecoveryServiceSubnetResource())
}
