// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	FileStorageSnapshotRequiredOnlyResource = FileStorageSnapshotResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Required, acctest.Create, FileStorageSnapshotRepresentation)

	FileStorageSnapshotResourceConfig = FileStorageSnapshotResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Optional, acctest.Update, FileStorageSnapshotRepresentation)

	FileStorageFileStorageSnapshotSingularDataSourceRepresentation = map[string]interface{}{
		"snapshot_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
	}

	FileStorageFileStorageSnapshotDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"file_system_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"filesystem_snapshot_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id}`},
		"id":                            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
		"state":                         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageSnapshotDataSourceFilterRepresentation}}
	FileStorageSnapshotDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_snapshot.test_snapshot.id}`}},
	}

	FileStorageSnapshotRepresentation = map[string]interface{}{
		"file_system_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `snapshot-1`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"expiration_time": acctest.Representation{RepType: acctest.Optional, Create: ExpirationTimeCreate, Update: ExpirationTimeUpdate},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	ExpirationTimeCreate = time.Now().Add(10 * time.Hour).Format(TimeFormat)
	ExpirationTimeUpdate = time.Now().Add(11 * time.Hour).Format(TimeFormat)

	FileStorageSnapshotResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_filesystem_snapshot_policy", "test_filesystem_snapshot_policy", acctest.Required, acctest.Create, FileStorageFilesystemSnapshotPolicyRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: file_storage/default
func TestFileStorageSnapshotResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageSnapshotResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_snapshot.test_snapshot"
	datasourceName := "data.oci_file_storage_snapshots.test_snapshots"
	singularDatasourceName := "data.oci_file_storage_snapshot.test_snapshot"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageSnapshotResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Optional, acctest.Create, FileStorageSnapshotRepresentation), "filestorage", "snapshot", t)

	acctest.ResourceTest(t, testAccCheckFileStorageSnapshotDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Required, acctest.Create, FileStorageSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageSnapshotResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileStorageSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Optional, acctest.Create, FileStorageSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "expiration_time", ExpirationTimeCreate),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + FileStorageSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Optional, acctest.Update, FileStorageSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "expiration_time", ExpirationTimeUpdate),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_snapshots", "test_snapshots", acctest.Optional, acctest.Update, FileStorageFileStorageSnapshotDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Optional, acctest.Update, FileStorageSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "filesystem_snapshot_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "snapshots.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "snapshots.0.expiration_time", ExpirationTimeUpdate),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.file_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "snapshots.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.is_clone_source"),
				resource.TestCheckResourceAttr(datasourceName, "snapshots.0.name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.provenance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.snapshot_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.snapshot_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", acctest.Required, acctest.Create, FileStorageFileStorageSnapshotSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "snapshot_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "expiration_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_clone_source"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provenance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "snapshot_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "snapshot_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + FileStorageSnapshotRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageSnapshotDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_snapshot" {
			noResourceFound = false
			request := oci_file_storage.GetSnapshotRequest{}

			tmp := rs.Primary.ID
			request.SnapshotId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetSnapshot(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.SnapshotLifecycleStateDeleted): true,
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
