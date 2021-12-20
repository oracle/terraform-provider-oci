// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	replicationSourceDataSourceRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket_replication.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	ReplicationSourceResourceConfig = ReplicationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", acctest.Required, acctest.Create, replicationPolicyRepresentation)
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
			Config: config + compartmentIdVariableStr + ReplicationSourceResourceConfig,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_replication_sources", "test_replication_sources", acctest.Required, acctest.Create, replicationSourceDataSourceRepresentation) +
				compartmentIdVariableStr + ReplicationSourceResourceConfig,
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
