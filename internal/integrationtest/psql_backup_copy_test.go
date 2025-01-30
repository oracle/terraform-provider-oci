// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsqlBackupCopyRequiredOnlyResource = PsqlBackupResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Required, acctest.Create, PsqlBackupCopyRepresentation)

	PsqlBackupCopyRepresentation = map[string]interface{}{
		"source_backup_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlBackupCopySourceDetailsRepresentation},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"retention_period":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignorePsqlBackupSystemTagsChangesRepresentation},
	}

	PsqlBackupCopySourceDetailsRepresentation = map[string]interface{}{}

	ignorePsqlBackupSystemTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`}},
	}

	PsqlBackupCopyResourceConfig = PsqlBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Optional, acctest.Update, PsqlBackupCopyRepresentation)

	PsqlBackupCopySingularDataSourceRepresentation = map[string]interface{}{
		"backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_psql_backup.test_psql_backup_copy.id}`},
	}
)

// issue-routing-tag: psql/default
func TestPsqlBackupCopyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlBackupCopyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceBackupId := utils.GetEnvSettingWithBlankDefault("source_backup_id")
	sourceBackupIdVariableStr := fmt.Sprintf("variable \"source_backup_id\" { default = \"%s\" }\n", sourceBackupId)

	sourceBackupRegion := utils.GetEnvSettingWithBlankDefault("source_region")
	sourceBackupRegionVariableStr := fmt.Sprintf("variable \"source_region\" { default = \"%s\" }\n", sourceBackupRegion)

	//compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	//compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_psql_backup.test_psql_backup_copy"
	//singularDatasourceName := "data.oci_psql_backup.test_psql_backup_copy"

	PsqlBackupCopySourceDetailsRepresentation = map[string]interface{}{
		"source_region":    acctest.Representation{RepType: acctest.Required, Create: sourceBackupRegion},
		"source_backup_id": acctest.Representation{RepType: acctest.Required, Create: sourceBackupId},
	}

	PsqlBackupCopyRepresentation = acctest.GetUpdatedRepresentationCopy("source_backup_details", acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlBackupCopySourceDetailsRepresentation}, PsqlBackupCopyRepresentation)

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Required, acctest.Create, PsqlBackupCopyRepresentation), "psql", "backup", t)

	acctest.ResourceTest(t, testAccCheckPsqlBackupCopyDestroy, []resource.TestStep{
		// verify Create

		{
			Config: config + compartmentIdVariableStr + sourceBackupIdVariableStr + sourceBackupRegionVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Required, acctest.Create, PsqlBackupCopyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.0.source_backup_id", sourceBackupId),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.0.source_region", sourceBackupRegion),
				resource.TestCheckResourceAttrSet(resourceName, "retention_period"),
				resource.TestCheckResourceAttrSet(resourceName, "source_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + sourceBackupIdVariableStr + sourceBackupRegionVariableStr,
		},

		{
			Config: config + compartmentIdVariableStr + sourceBackupIdVariableStr + sourceBackupRegionVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Optional, acctest.Create, PsqlBackupCopyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.0.source_backup_id", sourceBackupId),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.0.source_region", sourceBackupRegion),
				resource.TestCheckResourceAttr(resourceName, "retention_period", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		{
			Config: config + compartmentIdVariableStr + sourceBackupIdVariableStr + sourceBackupRegionVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Optional, acctest.Update, PsqlBackupCopyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.0.source_backup_id", sourceBackupId),
				resource.TestCheckResourceAttr(resourceName, "source_backup_details.0.source_region", sourceBackupRegion),
				resource.TestCheckResourceAttr(resourceName, "retention_period", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "source_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		/*
			// verify singular datasource
			{
				Config: config + compartmentIdVariableStr + sourceBackupIdVariableStr + sourceBackupRegionVariableStr +
					acctest.GenerateDataSourceFromRepresentationMap("oci_psql_backup", "test_psql_backup_copy", acctest.Optional, acctest.Create, PsqlBackupCopySingularDataSourceRepresentation) +
					PsqlBackupCopyResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_size"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_backup_details.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_system_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "retention_period", "11"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "source_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created_precise"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},

		*/

		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + sourceBackupIdVariableStr + sourceBackupRegionVariableStr + PsqlBackupCopyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"source_backup_details",
				"db_system_details",
				"retention_period",
				"last_accepted_request_token",
				"last_completed_request_token",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckPsqlBackupCopyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PostgresqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_psql_backup" {
			noResourceFound = false
			request := oci_psql.GetBackupRequest{}

			tmp := rs.Primary.ID
			request.BackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")

			response, err := client.GetBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_psql.BackupLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("PsqlBackupCopy") {
		resource.AddTestSweepers("PsqlBackupCopy", &resource.Sweeper{
			Name:         "PsqlBackupCopy",
			Dependencies: acctest.DependencyGraph["backup"],
			F:            sweepPsqlBackupResource,
		})
	}
}
