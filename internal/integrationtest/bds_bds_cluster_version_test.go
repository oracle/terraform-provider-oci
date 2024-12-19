// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsClusterVersionDataSourceRepresentation = map[string]interface{}{}

	BdsBdsClusterVersionResourceConfig = ""
)

// issue-routing-tag: bds/default
func TestBdsBdsClusterVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsClusterVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_bds_bds_cluster_versions.test_bds_cluster_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_cluster_versions", "test_bds_cluster_versions", acctest.Required, acctest.Create, BdsBdsClusterVersionDataSourceRepresentation) +
				compartmentIdVariableStr + BdsBdsClusterVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "bds_cluster_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_cluster_versions.0.bds_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_cluster_versions.0.odh_version"),
			),
		},
	})
}
