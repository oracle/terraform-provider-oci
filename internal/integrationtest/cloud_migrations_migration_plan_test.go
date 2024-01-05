// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudMigrationsAvailabilityDomain = "UsfO:US-ASHBURN-AD-1"
	CloudMigrationsVcn                = `${var.vcnId}`

	CloudMigrationsMigrationPlanRequiredOnlyResource = CloudMigrationsMigrationPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Required, acctest.Create, CloudMigrationsMigrationPlanRepresentation)

	CloudMigrationsMigrationPlanResourceConfig = CloudMigrationsMigrationPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Optional, acctest.Update, CloudMigrationsMigrationPlanRepresentation)

	CloudMigrationsCloudMigrationsMigrationPlanSingularDataSourceRepresentation = map[string]interface{}{
		"migration_plan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration_plan.test_migration_plan.id}`},
	}

	CloudMigrationsCloudMigrationsMigrationPlanDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"migration_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_migration.test_migration.id}`},
		"migration_plan_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_migration_plan.test_migration_plan.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudMigrationsMigrationPlanDataSourceFilterRepresentation}}
	CloudMigrationsMigrationPlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_migrations_migration_plan.test_migration_plan.id}`}},
	}

	CloudMigrationsMigrationPlanRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"migration_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration.test_migration.id}`},
		"strategies":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsMigrationPlanStrategiesRepresentation},
		"target_environments": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsMigrationPlanTargetEnvironmentsRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAjustmentMultiplierRepresentation},
	}
	CloudMigrationsMigrationPlanStrategiesRepresentation = map[string]interface{}{
		"resource_type": acctest.Representation{RepType: acctest.Required, Create: `CPU`, Update: `CPU`},
		"strategy_type": acctest.Representation{RepType: acctest.Required, Create: `AS_IS`, Update: `AS_IS`},
	}
	ignoreAjustmentMultiplierRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`strategies`}},
	}
	CloudMigrationsMigrationPlanTargetEnvironmentsRepresentation = map[string]interface{}{
		"subnet":                  acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsSubnetId, Update: CloudMigrationsSubnetId},
		"target_environment_type": acctest.Representation{RepType: acctest.Required, Create: `VM_TARGET_ENV`},
		"vcn":                     acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsVcn, Update: CloudMigrationsVcn},
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsAvailabilityDomain, Update: CloudMigrationsAvailabilityDomain},
		"ms_license":              acctest.Representation{RepType: acctest.Optional, Create: `msLicense`, Update: `msLicense2`},
		"preferred_shape_type":    acctest.Representation{RepType: acctest.Optional, Create: `VM`, Update: `VM_INTEL`},
		"target_compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: compartmentId},
	}

	CloudMigrationsMigrationPlanResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Required, acctest.Create, CloudMigrationsMigrationRepresentation)
)

// issue-routing-tag: cloud_migrations/default
func TestCloudMigrationsMigrationPlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudMigrationsMigrationPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcnId")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcnId\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnetId")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnetId\" { default = \"%s\" }\n", subnetId)

	variableStr := compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_migrations_migration_plan.test_migration_plan"
	datasourceName := "data.oci_cloud_migrations_migration_plans.test_migration_plans"
	singularDatasourceName := "data.oci_cloud_migrations_migration_plan.test_migration_plan"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudMigrationsMigrationPlanResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Optional, acctest.Create, CloudMigrationsMigrationPlanRepresentation), "cloudmigrations", "migrationPlan", t)

	acctest.ResourceTest(t, testAccCheckCloudMigrationsMigrationPlanDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + CloudMigrationsMigrationPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Required, acctest.Create, CloudMigrationsMigrationPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variableStr + CloudMigrationsMigrationPlanResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variableStr + CloudMigrationsMigrationPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Optional, acctest.Create, CloudMigrationsMigrationPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "strategies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "strategies.0.resource_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "strategies.0.strategy_type", "AS_IS"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_environments.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.ms_license", "msLicense"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.preferred_shape_type", "VM"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.subnet", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "target_environments.0.target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.target_environment_type", "VM_TARGET_ENV"),
				resource.TestCheckResourceAttrSet(resourceName, "target_environments.0.vcn"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + variableStr + compartmentIdUVariableStr + CloudMigrationsMigrationPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudMigrationsMigrationPlanRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "strategies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "strategies.0.resource_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "strategies.0.strategy_type", "AS_IS"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_environments.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.ms_license", "msLicense"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.preferred_shape_type", "VM"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.subnet", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "target_environments.0.target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.target_environment_type", "VM_TARGET_ENV"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.vcn", vcnId),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + variableStr + CloudMigrationsMigrationPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Optional, acctest.Update, CloudMigrationsMigrationPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "strategies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "strategies.0.resource_type", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "strategies.0.strategy_type", "AS_IS"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.availability_domain", CloudMigrationsAvailabilityDomain),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.ms_license", "msLicense2"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.preferred_shape_type", "VM_INTEL"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.subnet", subnetId),
				resource.TestCheckResourceAttrSet(resourceName, "target_environments.0.target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.target_environment_type", "VM_TARGET_ENV"),
				resource.TestCheckResourceAttr(resourceName, "target_environments.0.vcn", vcnId),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration_plans", "test_migration_plans", acctest.Optional, acctest.Update, CloudMigrationsCloudMigrationsMigrationPlanDataSourceRepresentation) +
				variableStr + CloudMigrationsMigrationPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Optional, acctest.Update, CloudMigrationsMigrationPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "migration_plan_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "migration_plan_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "migration_plan_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsMigrationPlanSingularDataSourceRepresentation) +
				variableStr + CloudMigrationsMigrationPlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_plan_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "migration_plan_stats.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "strategies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "strategies.0.strategy_type", "AS_IS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.0.availability_domain", CloudMigrationsAvailabilityDomain),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.0.ms_license", "msLicense2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.0.preferred_shape_type", "VM_INTEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.0.subnet", subnetId),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.0.target_environment_type", "VM_TARGET_ENV"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_environments.0.vcn", vcnId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudMigrationsMigrationPlanRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudMigrationsMigrationPlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_migrations_migration_plan" {
			noResourceFound = false
			request := oci_cloud_migrations.GetMigrationPlanRequest{}

			tmp := rs.Primary.ID
			request.MigrationPlanId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")

			response, err := client.GetMigrationPlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_migrations.MigrationPlanLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudMigrationsMigrationPlan") {
		resource.AddTestSweepers("CloudMigrationsMigrationPlan", &resource.Sweeper{
			Name:         "CloudMigrationsMigrationPlan",
			Dependencies: acctest.DependencyGraph["migrationPlan"],
			F:            sweepCloudMigrationsMigrationPlanResource,
		})
	}
}

func sweepCloudMigrationsMigrationPlanResource(compartment string) error {
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()
	migrationPlanIds, err := getCloudMigrationsMigrationPlanIds(compartment)
	if err != nil {
		return err
	}
	for _, migrationPlanId := range migrationPlanIds {
		if ok := acctest.SweeperDefaultResourceId[migrationPlanId]; !ok {
			deleteMigrationPlanRequest := oci_cloud_migrations.DeleteMigrationPlanRequest{}

			deleteMigrationPlanRequest.MigrationPlanId = &migrationPlanId

			deleteMigrationPlanRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")
			_, error := migrationClient.DeleteMigrationPlan(context.Background(), deleteMigrationPlanRequest)
			if error != nil {
				fmt.Printf("Error deleting MigrationPlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", migrationPlanId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &migrationPlanId, CloudMigrationsMigrationPlanSweepWaitCondition, time.Duration(3*time.Minute),
				CloudMigrationsMigrationPlanSweepResponseFetchOperation, "cloud_migrations", true)
		}
	}
	return nil
}

func getCloudMigrationsMigrationPlanIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MigrationPlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()

	listMigrationPlansRequest := oci_cloud_migrations.ListMigrationPlansRequest{}
	listMigrationPlansRequest.CompartmentId = &compartmentId
	listMigrationPlansRequest.LifecycleState = oci_cloud_migrations.MigrationPlanLifecycleStateActive
	listMigrationPlansResponse, err := migrationClient.ListMigrationPlans(context.Background(), listMigrationPlansRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MigrationPlan list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, migrationPlan := range listMigrationPlansResponse.Items {
		id := *migrationPlan.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MigrationPlanId", id)
	}
	return resourceIds, nil
}

func CloudMigrationsMigrationPlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if migrationPlanResponse, ok := response.Response.(oci_cloud_migrations.GetMigrationPlanResponse); ok {
		return migrationPlanResponse.LifecycleState != oci_cloud_migrations.MigrationPlanLifecycleStateDeleted
	}
	return false
}

func CloudMigrationsMigrationPlanSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MigrationClient().GetMigrationPlan(context.Background(), oci_cloud_migrations.GetMigrationPlanRequest{
		MigrationPlanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
