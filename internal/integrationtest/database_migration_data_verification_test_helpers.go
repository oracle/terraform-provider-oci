// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License 2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const databaseMigrationDataVerificationMigrationIDEnv = "data_verification_migration_id"

type databaseMigrationDataVerificationTestInputs struct {
	variableConfig string
	migrationID    string
	compartmentID  string
}

// databaseMigrationDataVerificationTestInputsFor returns the Terraform variables
// for BDV acceptance tests and verifies that the supplied migration is ready for
// data verification. BDV tests intentionally consume an existing migration and
// never create or mutate one.
func databaseMigrationDataVerificationTestInputsFor(t *testing.T) databaseMigrationDataVerificationTestInputs {
	t.Helper()

	migrationID := utils.GetEnvSettingWithBlankDefault(databaseMigrationDataVerificationMigrationIDEnv)
	if migrationID == "" {
		t.Fatalf("%s must be set to an existing migration OCID whose Data Verification state is SUCCEEDED", databaseMigrationDataVerificationMigrationIDEnv)
	}

	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	if compartmentID == "" {
		t.Fatalf("compartment_ocid must be set for Database Migration Data Verification acceptance tests")
	}

	request := oci_database_migration.GetDataVerificationDetailRequest{MigrationId: &migrationID}
	response, err := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient().GetDataVerificationDetail(context.Background(), request)
	if err != nil {
		t.Fatalf("migration %q is not ready for Data Verification acceptance tests: GetDataVerificationDetail failed: %v; use a migration OCID after the migration has reached a state where Data Verification is available", migrationID, err)
	}

	if response.LifecycleState != oci_database_migration.DataVerificationDetailLifecycleStateSucceeded {
		t.Fatalf("migration %q is not ready for Data Verification acceptance tests: Data Verification state is %q, want %q", migrationID, response.LifecycleState, oci_database_migration.DataVerificationDetailLifecycleStateSucceeded)
	}

	return databaseMigrationDataVerificationTestInputs{
		variableConfig: fmt.Sprintf(
			"variable \"compartment_id\" { default = %q }\nvariable \"data_verification_migration_id\" { default = %q }\n",
			compartmentID,
			migrationID,
		),
		migrationID:   migrationID,
		compartmentID: compartmentID,
	}
}
