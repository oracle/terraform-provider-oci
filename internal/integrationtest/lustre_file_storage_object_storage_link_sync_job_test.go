// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LustreFileStorageObjectStorageLinkSyncJobSingularDataSourceRepresentation = map[string]interface{}{
		"object_storage_link_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_object_storage_link.test_object_storage_link.id}`},
		"sync_job_id":            acctest.Representation{RepType: acctest.Required, Create: ``},
	}

	LustreFileStorageObjectStorageLinkSyncJobDataSourceRepresentation = map[string]interface{}{
		"object_storage_link_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_object_storage_link.test_object_storage_link.id}`},
	}

	LustreFileStorageObjectStorageLinkSyncJobResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Required, acctest.Create, LustreFileStorageLustreFileSystemRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link", "test_object_storage_link", acctest.Required, acctest.Create, LustreFileStorageObjectStorageLinkRepresentation)
)

// issue-routing-tag: lustre_file_storage/default
func TestLustreFileStorageObjectStorageLinkSyncJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLustreFileStorageObjectStorageLinkSyncJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_lustre_file_storage_object_storage_link_sync_jobs.test_object_storage_link_sync_jobs"
	singularDatasourceName := "data.oci_lustre_file_storage_object_storage_link_sync_job.test_object_storage_link_sync_job"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link_sync_jobs", "test_object_storage_link_sync_jobs", acctest.Required, acctest.Create, LustreFileStorageObjectStorageLinkSyncJobDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageObjectStorageLinkSyncJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "sync_job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sync_job_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_object_storage_link_sync_job", "test_object_storage_link_sync_job", acctest.Required, acctest.Create, LustreFileStorageObjectStorageLinkSyncJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageObjectStorageLinkSyncJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_link_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sync_job_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "bytes_transferred"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_overwrite"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lustre_file_system_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "objects_transferred"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skipped_error_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_objects_scanned"),
			),
		},
	})
}
