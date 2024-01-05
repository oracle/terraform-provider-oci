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
	awrHubSourceAwrhubsourcesManagementRepresentation = map[string]interface{}{
		"awr_hub_source_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.awr_hub_source_id}`},
		"enable_awrhubsource": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	AwrHubSourceAwrhubsourcesManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubSourceAwrhubsourcesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubSourceAwrhubsourcesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// The AWR Hub Source should have minimum 1 object uploaded, so that the state will be ACCEPTING. The enable/disable
	// operation is permitted only if there are uploads. There are no terraform provider for uploading objects.
	// SO going with a input for the source id which has objects already uploaded.
	awrHubSourceId := utils.GetEnvSettingWithBlankDefault("awr_hub_source_ocid")
	awrHubSourceIdVariableStr := fmt.Sprintf("variable \"awr_hub_source_id\" { default = \"%s\" }\n", awrHubSourceId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opsi_awr_hub_source_awrhubsources_management.test_awr_hub_source_awrhubsources_management"
	parentResourceName := "oci_opsi_awr_hub_source_awrhubsources_management.test_awr_hub_source_awrhubsources_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+awrHubSourceIdVariableStr+AwrHubSourceAwrhubsourcesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source_awrhubsources_management", "test_awr_hub_source_awrhubsources_management", acctest.Required, acctest.Create, awrHubSourceAwrhubsourcesManagementRepresentation), "operationsinsights", "awrHubSourceAwrhubsourcesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + awrHubSourceIdVariableStr + AwrHubSourceAwrhubsourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source_awrhubsources_management", "test_awr_hub_source_awrhubsources_management", acctest.Required, acctest.Create, awrHubSourceAwrhubsourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_source_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + awrHubSourceIdVariableStr + AwrHubSourceAwrhubsourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source_awrhubsources_management", "test_awr_hub_source_awrhubsources_management", acctest.Required, acctest.Create, awrHubSourceAwrhubsourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_awrhubsource", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + awrHubSourceIdVariableStr + AwrHubSourceAwrhubsourcesManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + awrHubSourceIdVariableStr + AwrHubSourceAwrhubsourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source_awrhubsources_management", "test_awr_hub_source_awrhubsources_management", acctest.Optional, acctest.Create, awrHubSourceAwrhubsourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_source_id"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + awrHubSourceIdVariableStr + AwrHubSourceAwrhubsourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source_awrhubsources_management", "test_awr_hub_source_awrhubsources_management", acctest.Optional, acctest.Update, awrHubSourceAwrhubsourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_source_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + awrHubSourceIdVariableStr + AwrHubSourceAwrhubsourcesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source_awrhubsources_management", "test_awr_hub_source_awrhubsources_management", acctest.Optional, acctest.Update, awrHubSourceAwrhubsourcesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_awrhubsource", "false"),
			),
		},
	})
}
