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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FileStorageReplicationRequiredOnlyResource = FileStorageReplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Required, acctest.Create, FileStorageReplicationRepresentation)

	FileStorageReplicationSingularDataSourceRepresentation = map[string]interface{}{
		"replication_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_replication.test_replication.id}`},
	}

	FileStorageReplicationDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `replication-policy-1`, Update: `displayName2`},
		"file_system_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system_source.id}`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_replication.test_replication.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageReplicationDataSourceFilterRepresentation}}
	FileStorageReplicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_replication.test_replication.id}`}},
	}

	FileStorageReplicationRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system_source.id}`},
		"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system_target.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `replication-policy-1`, Update: `displayName2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"replication_interval": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `16`},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageReplicationRepresentationWithFullLock = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system_source.id}`},
		"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system_target.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `replication-policy-1`, Update: `displayName2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageReplicationFullLocksRepresentation},
		"is_lock_override":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"replication_interval": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `16`},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageReplicationRepresentationWithDeleteLock = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system_source.id}`},
		"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system_target.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `replication-policy-1`, Update: `displayName2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageReplicationDeleteLocksRepresentation},
		"is_lock_override":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"replication_interval": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `16`},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	FileStorageReplicationFullLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	FileStorageReplicationDeleteLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `DELETE`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	FileStorageReplicationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system_source", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system_target", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	FileStorageReplicationResourceConfig = FileStorageReplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Update, FileStorageReplicationRepresentation)
)

// issue-routing-tag: file_storage/default
func TestFileStorageReplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageReplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_replication.test_replication"
	datasourceName := "data.oci_file_storage_replications.test_replications"
	singularDatasourceName := "data.oci_file_storage_replication.test_replication"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageReplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Create, FileStorageReplicationRepresentation), "filestorage", "replication", t)

	acctest.ResourceTest(t, testAccCheckFileStorageReplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Required, acctest.Create, FileStorageReplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageReplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileStorageReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Create, FileStorageReplicationRepresentationWithDeleteLock),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "replication-policy-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(resourceName, "replication_interval", "15"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileStorageReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FileStorageReplicationRepresentationWithDeleteLock, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "replication-policy-1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(resourceName, "replication_interval", "15"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
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
			Config: config + compartmentIdVariableStr + FileStorageReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Update, FileStorageReplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(resourceName, "replication_interval", "16"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_replications", "test_replications", acctest.Optional, acctest.Update, FileStorageReplicationDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Update, FileStorageReplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "replications.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "replications.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "replications.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.locks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(datasourceName, "replications.0.locks.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.locks.0.type", "DELETE"),
				resource.TestCheckResourceAttrSet(datasourceName, "replications.0.recovery_point_time"),
				resource.TestCheckResourceAttr(datasourceName, "replications.0.replication_interval", "16"),
				resource.TestCheckResourceAttrSet(datasourceName, "replications.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "replications.0.time_created"),
			),
		},
		// verify singular datasource
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, replicationCheckWaitCondition, time.Duration(10*time.Minute),
				FileStorageReplicationSweepResponseFetchOperation, "file_storage", true),
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Required, acctest.Create, FileStorageReplicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageReplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delta_progress"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delta_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_snapshot_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_point_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_interval", "16"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_target_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageReplicationResourceDependencies,
		},
		// verify Create with optionals and FULL lock
		{
			Config: config + compartmentIdVariableStr + FileStorageReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_replication", "test_replication", acctest.Optional, acctest.Create, FileStorageReplicationRepresentationWithFullLock),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),

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
		// verify resource import
		{
			Config:                  config + FileStorageReplicationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"is_lock_override"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageReplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_replication" {
			noResourceFound = false
			request := oci_file_storage.GetReplicationRequest{}

			tmp := rs.Primary.ID
			request.ReplicationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetReplication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.ReplicationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FileStorageReplication") {
		resource.AddTestSweepers("FileStorageReplication", &resource.Sweeper{
			Name:         "FileStorageReplication",
			Dependencies: acctest.DependencyGraph["replication"],
			F:            sweepFileStorageReplicationResource,
		})
	}
}

func sweepFileStorageReplicationResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	replicationIds, err := getFileStorageReplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, replicationId := range replicationIds {
		if ok := acctest.SweeperDefaultResourceId[replicationId]; !ok {
			deleteReplicationRequest := oci_file_storage.DeleteReplicationRequest{}

			deleteReplicationRequest.ReplicationId = &replicationId

			deleteReplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteReplication(context.Background(), deleteReplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting Replication %s %s, It is possible that the resource is already deleted. Please verify manually \n", replicationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &replicationId, FileStorageReplicationSweepWaitCondition, time.Duration(3*time.Minute),
				FileStorageReplicationSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileStorageReplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ReplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listReplicationsRequest := oci_file_storage.ListReplicationsRequest{}
	listReplicationsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for Replication list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listReplicationsRequest.AvailabilityDomain = &availabilityDomainName

		listReplicationsRequest.LifecycleState = oci_file_storage.ListReplicationsLifecycleStateActive
		listReplicationsResponse, err := fileStorageClient.ListReplications(context.Background(), listReplicationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Replication list for compartment id testing : %s , %s \n", compartmentId, err)
		}
		for _, replication := range listReplicationsResponse.Items {
			id := *replication.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ReplicationId", id)
		}

	}
	return resourceIds, nil
}

func FileStorageReplicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if replicationResponse, ok := response.Response.(oci_file_storage.GetReplicationResponse); ok {
		return replicationResponse.LifecycleState != oci_file_storage.ReplicationLifecycleStateDeleted
	}
	return false
}

func replicationCheckWaitCondition(response common.OCIOperationResponse) bool {
	if replicationCheckResponse, ok := response.Response.(oci_file_storage.GetReplicationResponse); ok {
		return replicationCheckResponse.DeltaStatus != oci_file_storage.ReplicationDeltaStatusApplying
	}
	return false
}

func FileStorageReplicationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetReplication(context.Background(), oci_file_storage.GetReplicationRequest{
		ReplicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
