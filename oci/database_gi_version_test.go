// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	giVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"shape":          Representation{RepType: Required, Create: `ExadataCC.Quarter3.100`},
	}

	GiVersionResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseGiVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseGiVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_gi_versions.test_gi_versions"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_gi_versions", "test_gi_versions", Required, Create, giVersionDataSourceRepresentation) +
				compartmentIdVariableStr + GiVersionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "gi_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "gi_versions.0.version"),
			),
		},
	})
}
