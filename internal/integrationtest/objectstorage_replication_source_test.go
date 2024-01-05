// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ObjectStorageObjectStorageReplicationSourceDataSourceRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket_replication.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	ObjectStorageReplicationSourceResourceConfig = ObjectStorageReplicationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Required, acctest.Create, ObjectStorageReplicationPolicyRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStorageReplicationSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageReplicationSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_objectstorage_replication_sources.test_replication_sources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create replication policy first
		{
			Config: config + compartmentIdVariableStr + ObjectStorageReplicationSourceResourceConfig,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_replication_sources", "test_replication_sources", acctest.Required, acctest.Create, ObjectStorageObjectStorageReplicationSourceDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectStorageReplicationSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", replicationBucketName),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.0.policy_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.0.source_bucket_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.0.source_region_name"),
			),
		},
	})
}
