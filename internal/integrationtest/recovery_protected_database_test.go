// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RecoveryProtectedDatabaseRequiredOnlyResource = RecoveryProtectedDatabaseResourceStaticDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Required, acctest.Create, RecoveryProtectedDatabaseRepresentation)

	RecoveryProtectedDatabaseResourceConfig = RecoveryProtectedDatabaseResourceStaticDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Optional, acctest.Update, RecoveryProtectedDatabaseRepresentation)

	RecoveryProtectedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"protected_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_recovery_protected_database.test_protected_database.id}`},
	}

	RecoveryProtectedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                         acctest.Representation{RepType: acctest.Optional, Create: `${oci_recovery_protected_database.test_protected_database.id}`},
		"protection_policy_id":       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_recovery_protection_policy.test_protection_policy.id}`},
		"recovery_service_subnet_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id}`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: RecoveryProtectedDatabaseDataSourceFilterRepresentation}}
	RecoveryProtectedDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_recovery_protected_database.test_protected_database.id}`}},
	}

	RecoveryProtectedDatabaseRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_unique_name":           acctest.Representation{RepType: acctest.Required, Create: `dbUniqueName`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_secret#11`, Update: `BEstrO0ng_secret#12`},
		"protection_policy_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_recovery_protection_policy.test_protection_policy.id}`},
		"recovery_service_subnets": acctest.RepresentationGroup{RepType: acctest.Required, Group: RecoveryProtectedDatabaseRecoveryServiceSubnetsRepresentation},
		"database_id":              acctest.Representation{RepType: acctest.Optional, Create: `DummyDatabaseID_` + utils.RandomString(10, utils.Charset)},
		"database_size":            acctest.Representation{RepType: acctest.Optional, Create: `XS`, Update: `S`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deletion_schedule":        acctest.Representation{RepType: acctest.Optional, Create: `DELETE_AFTER_72_HOURS`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_redo_logs_shipped":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"subscription_id":          acctest.Representation{RepType: acctest.Optional, Create: `ocid1.organizationssubscription.oc1..amaaaaaa6jqx4paaa2rxk42owtrtvwkhauvoqb2equbymlvdrlv5tclvvvta`, Update: `ocid1.organizationssubscription.oc1..amaaaaaa6jqx4paaa2rxk42owtrtvwkhauvoqb2equbymlvdrlv5tclvvita`},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: recoveryIgnoreDefinedTagsRepresentation},
	}
	RecoveryProtectedDatabaseRecoveryServiceSubnetsRepresentation = map[string]interface{}{
		"recovery_service_subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id}`},
	}
	RecoveryProtectedDatabaseResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Required, acctest.Create, RecoveryProtectionPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_recovery_service_subnet", "test_recovery_service_subnet", acctest.Required, acctest.Create, RecoveryRecoveryServiceSubnetRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	recoveryServiceSubnetId            = utils.GetEnvSettingWithBlankDefault("recovery_service_subnet_id")
	protectionPolicyId                 = utils.GetEnvSettingWithBlankDefault("protection_policy_id")
	recoveryServiceSubnetIdVariableStr = fmt.Sprintf("variable \"recovery_service_subnet_id\" { default = \"%s\" }\n", recoveryServiceSubnetId)
	protectionPolicyIdVariableStr      = fmt.Sprintf("variable \"protection_policy_id\" { default = \"%s\" }\n", protectionPolicyId)

	recoveryServiceSubnetDependency = recoveryServiceSubnetIdVariableStr + `
	data "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet" {
		recovery_service_subnet_id = "${var.recovery_service_subnet_id}"
	}
	`
	recoveryProtectionPolicyDependency = protectionPolicyIdVariableStr + `
	data "oci_recovery_protection_policy" "test_protection_policy" {
		protection_policy_id = "${var.protection_policy_id}"
	}
	`
	RecoveryProtectedDatabaseResourceStaticDependencies = recoveryServiceSubnetDependency +
		recoveryProtectionPolicyDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	recoveryIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: recovery/default
func TestRecoveryProtectedDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRecoveryProtectedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_recovery_protected_database.test_protected_database"
	datasourceName := "data.oci_recovery_protected_databases.test_protected_databases"
	singularDatasourceName := "data.oci_recovery_protected_database.test_protected_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RecoveryProtectedDatabaseResourceStaticDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Optional, acctest.Create, RecoveryProtectedDatabaseRepresentation), "recovery", "protectedDatabase", t)

	acctest.ResourceTest(t, testAccCheckRecoveryProtectedDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RecoveryProtectedDatabaseResourceStaticDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Required, acctest.Create, RecoveryProtectedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", "dbUniqueName"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_secret#11"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "recovery_service_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "recovery_service_subnets.0.recovery_service_subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RecoveryProtectedDatabaseResourceStaticDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RecoveryProtectedDatabaseResourceStaticDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Optional, acctest.Create, RecoveryProtectedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_size", "XS"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", "dbUniqueName"),
				resource.TestCheckResourceAttr(resourceName, "deletion_schedule", "DELETE_AFTER_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logs_shipped", "false"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_secret#11"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "recovery_service_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "recovery_service_subnets.0.recovery_service_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vpc_user_name"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RecoveryProtectedDatabaseResourceStaticDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RecoveryProtectedDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_size", "XS"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", "dbUniqueName"),
				resource.TestCheckResourceAttr(resourceName, "deletion_schedule", "DELETE_AFTER_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logs_shipped", "false"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_secret#11"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "recovery_service_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "recovery_service_subnets.0.recovery_service_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vpc_user_name"),

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
			Config: config + compartmentIdVariableStr + RecoveryProtectedDatabaseResourceStaticDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Optional, acctest.Update, RecoveryProtectedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_size", "S"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", "dbUniqueName"),
				resource.TestCheckResourceAttr(resourceName, "deletion_schedule", "DELETE_AFTER_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_redo_logs_shipped", "true"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_secret#12"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "recovery_service_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "recovery_service_subnets.0.recovery_service_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vpc_user_name"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protected_databases", "test_protected_databases", acctest.Optional, acctest.Update, RecoveryProtectedDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryProtectedDatabaseResourceStaticDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Optional, acctest.Update, RecoveryProtectedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "protection_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "recovery_service_subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "protected_database_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "protected_database_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Required, acctest.Create, RecoveryProtectedDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryProtectedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protected_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_size", "S"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_unique_name", "dbUniqueName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_read_only_resource"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_redo_logs_shipped", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metrics.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_locked_date_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recovery_service_subnets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_service_subnets.0.state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vpc_user_name"),
			),
		},
		// verify resource import
		{
			Config:            config + RecoveryProtectedDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"deletion_schedule",
				"password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckRecoveryProtectedDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_recovery_protected_database" {
			noResourceFound = false
			request := oci_recovery.GetProtectedDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ProtectedDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")

			response, err := client.GetProtectedDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_recovery.LifecycleStateDeleted):         true,
					string(oci_recovery.LifecycleStateDeleteScheduled): true,
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
	if !acctest.InSweeperExcludeList("RecoveryProtectedDatabase") {
		resource.AddTestSweepers("RecoveryProtectedDatabase", &resource.Sweeper{
			Name:         "RecoveryProtectedDatabase",
			Dependencies: acctest.DependencyGraph["protectedDatabase"],
			F:            sweepRecoveryProtectedDatabaseResource,
		})
	}
}

func sweepRecoveryProtectedDatabaseResource(compartment string) error {
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()
	protectedDatabaseIds, err := getRecoveryProtectedDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, protectedDatabaseId := range protectedDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[protectedDatabaseId]; !ok {
			deleteProtectedDatabaseRequest := oci_recovery.DeleteProtectedDatabaseRequest{}

			deleteProtectedDatabaseRequest.ProtectedDatabaseId = &protectedDatabaseId

			deleteProtectedDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")
			_, error := databaseRecoveryClient.DeleteProtectedDatabase(context.Background(), deleteProtectedDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ProtectedDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", protectedDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &protectedDatabaseId, RecoveryProtectedDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				RecoveryProtectedDatabaseSweepResponseFetchOperation, "recovery", true)
		}
	}
	return nil
}

func getRecoveryProtectedDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProtectedDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()

	listProtectedDatabasesRequest := oci_recovery.ListProtectedDatabasesRequest{}
	listProtectedDatabasesRequest.CompartmentId = &compartmentId
	listProtectedDatabasesRequest.LifecycleState = oci_recovery.ListProtectedDatabasesLifecycleStateActive
	listProtectedDatabasesResponse, err := databaseRecoveryClient.ListProtectedDatabases(context.Background(), listProtectedDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ProtectedDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, protectedDatabase := range listProtectedDatabasesResponse.Items {
		id := *protectedDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProtectedDatabaseId", id)
	}
	return resourceIds, nil
}

func RecoveryProtectedDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if protectedDatabaseResponse, ok := response.Response.(oci_recovery.GetProtectedDatabaseResponse); ok {
		return protectedDatabaseResponse.LifecycleState != oci_recovery.LifecycleStateDeleted
	}
	return false
}

func RecoveryProtectedDatabaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseRecoveryClient().GetProtectedDatabase(context.Background(), oci_recovery.GetProtectedDatabaseRequest{
		ProtectedDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
