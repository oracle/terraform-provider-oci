// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below
	JmsFleetDrsFileWithAdvancedFeature = utils.GetEnvSettingWithBlankDefault("fleet_advanced_feature_ocid")

	JmsFleetDrsFileDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetDrsFileWithAdvancedFeature},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetDrsFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetDrsFileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_drs_files.test_fleet_drs_files"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_drs_files",
					"test_fleet_drs_files",
					acctest.Required,
					acctest.Create,
					JmsFleetDrsFileDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "drs_file_collection.#"),
			),
		},
	})
}
