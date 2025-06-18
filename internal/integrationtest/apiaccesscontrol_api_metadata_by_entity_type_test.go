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
	ApiaccesscontrolApiMetadataByEntityTypeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `CLOUDEXADATAINFRASTRUCTURE`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	ApiaccesscontrolApiMetadataByEntityTypeResourceConfig = ""
)

// issue-routing-tag: apiaccesscontrol/default
func TestApiaccesscontrolApiMetadataByEntityTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApiaccesscontrolApiMetadataByEntityTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_apiaccesscontrol_api_metadata_by_entity_types.test_api_metadata_by_entity_types"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_api_metadata_by_entity_types", "test_api_metadata_by_entity_types", acctest.Optional, acctest.Create, ApiaccesscontrolApiMetadataByEntityTypeDataSourceRepresentation) +
				compartmentIdVariableStr + ApiaccesscontrolApiMetadataByEntityTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "CLOUDEXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "api_metadata_by_entity_type_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "api_metadata_by_entity_type_collection.0.items.#"),
			),
		},
	})
}
