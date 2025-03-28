// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreChangesZprConfigurationRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags", "freeform_tags"}},
	}

	ZprConfigurationRequiredOnlyResource = ZprConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_zpr_configuration", "test_configuration", acctest.Required, acctest.Create, ZprConfigurationRepresentation)

	ZprConfigurationResourceConfig = ZprConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_zpr_configuration", "test_configuration", acctest.Optional, acctest.Update, ZprConfigurationRepresentation)

	ZprConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	ZprConfigurationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Required, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Required, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesZprConfigurationRepresentation},
	}

	ZprConfigurationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: zpr/default
func TestZprConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestZprConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_zpr_configuration.test_configuration"
	singularDatasourceName := "data.oci_zpr_configuration.test_configuration"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ZprConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_zpr_configuration", "test_configuration", acctest.Optional, acctest.Create, ZprConfigurationRepresentation), "zpr", "configuration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Configuration can only be created once
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ZprConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_zpr_configuration", "test_configuration", acctest.Required, acctest.Create, ZprConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "zpr_status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId, err := acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_zpr_configuration", "test_configuration", acctest.Required, acctest.Create, ZprConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ZprConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				// TODO: Service bug - Tags are not returned in GET call - DATASEC-3045
				//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "zpr_status", "ENABLED"),
			),
		},

		// verify resource import
		{
			Config:                  config + ZprConfigurationRequiredOnlyResource,
			ImportState:             true,
			ImportStateIdFunc:       getZprConfigurationConfigurationId(resourceName),
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags", "freeform_tags"},
			ResourceName:            resourceName,
		},

		// delete
		{
			Config: config + compartmentIdVariableStr + ZprConfigurationResourceDependencies,
		},
	})
}

func getZprConfigurationConfigurationId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		importStateID := fmt.Sprintf("%s/%s", rs.Primary.Attributes["compartment_id"], rs.Primary.Attributes["id"])
		return importStateID, nil
	}
}
