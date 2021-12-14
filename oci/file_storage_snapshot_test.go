// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v54/filestorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SnapshotRequiredOnlyResource = SnapshotResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Required, Create, snapshotRepresentation)

	SnapshotResourceConfig = SnapshotResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Optional, Update, snapshotRepresentation)

	snapshotSingularDataSourceRepresentation = map[string]interface{}{
		"snapshot_id": Representation{RepType: Required, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
	}

	snapshotDataSourceRepresentation = map[string]interface{}{
		"file_system_id": Representation{RepType: Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"id":             Representation{RepType: Optional, Create: `${oci_file_storage_snapshot.test_snapshot.id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, snapshotDataSourceFilterRepresentation}}
	snapshotDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_file_storage_snapshot.test_snapshot.id}`}},
	}

	snapshotRepresentation = map[string]interface{}{
		"file_system_id": Representation{RepType: Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"name":           Representation{RepType: Required, Create: `snapshot-1`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	SnapshotResourceDependencies = GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", Required, Create, fileSystemRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: file_storage/default
func TestFileStorageSnapshotResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageSnapshotResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_snapshot.test_snapshot"
	datasourceName := "data.oci_file_storage_snapshots.test_snapshots"
	singularDatasourceName := "data.oci_file_storage_snapshot.test_snapshot"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+SnapshotResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Optional, Create, snapshotRepresentation), "filestorage", "snapshot", t)

	ResourceTest(t, testAccCheckFileStorageSnapshotDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SnapshotResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Required, Create, snapshotRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SnapshotResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SnapshotResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Optional, Create, snapshotRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SnapshotResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Optional, Update, snapshotRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_file_storage_snapshots", "test_snapshots", Optional, Update, snapshotDataSourceRepresentation) +
				compartmentIdVariableStr + SnapshotResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Optional, Update, snapshotRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "file_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "snapshots.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.file_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "snapshots.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.is_clone_source"),
				resource.TestCheckResourceAttr(datasourceName, "snapshots.0.name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.provenance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "snapshots.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_file_storage_snapshot", "test_snapshot", Required, Create, snapshotSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SnapshotResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "snapshot_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_clone_source"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "snapshot-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provenance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SnapshotResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageSnapshotDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).fileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_snapshot" {
			noResourceFound = false
			request := oci_file_storage.GetSnapshotRequest{}

			tmp := rs.Primary.ID
			request.SnapshotId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "file_storage")

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
