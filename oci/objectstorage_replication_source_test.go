// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	replicationSourceDataSourceRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket_replication.name}`},
		"namespace": Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
	}

	ReplicationSourceResourceConfig = ReplicationPolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_replication_policy", "test_replication_policy", Required, Create, replicationPolicyRepresentation)
)

func TestObjectStorageReplicationSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageReplicationSourceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_objectstorage_replication_sources.test_replication_sources"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create replication policy first
			{
				Config: config + compartmentIdVariableStr + ReplicationSourceResourceConfig,
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_replication_sources", "test_replication_sources", Required, Create, replicationSourceDataSourceRepresentation) +
					compartmentIdVariableStr + ReplicationSourceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", replicationBucketName),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
					resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.0.policy_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.0.source_bucket_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "replication_sources.0.source_region_name"),
				),
			},
		},
	})
}
