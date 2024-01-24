// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ArtifactsContainerConfigurationRequiredOnlyResource = ArtifactsContainerConfigurationResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", acctest.Required, acctest.Create, ArtifactscontainerConfigurationRepresentation)

	ArtifactsContainerConfigurationResourceConfig = ArtifactsContainerConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", acctest.Optional, acctest.Update, ArtifactscontainerConfigurationRepresentation)

	ArtifactsArtifactscontainerConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	ArtifactscontainerConfigurationRepresentation = map[string]interface{}{
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_repository_created_on_first_push": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}

	ArtifactsContainerConfigurationResourceDependencies = ""
)

// issue-routing-tag: artifacts/default
func TestArtifactsContainerConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsContainerConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_artifacts_container_configuration.test_container_configuration"

	singularDatasourceName := "data.oci_artifacts_container_configuration.test_container_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ArtifactsContainerConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", acctest.Required, acctest.Create, ArtifactscontainerConfigurationRepresentation), "artifacts", "containerConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ArtifactsContainerConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", acctest.Required, acctest.Create, ArtifactscontainerConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_repository_created_on_first_push", "false"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ArtifactsContainerConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", acctest.Optional, acctest.Update, ArtifactscontainerConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_repository_created_on_first_push", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", acctest.Required, acctest.Create, ArtifactsArtifactscontainerConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ArtifactsContainerConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(singularDatasourceName, "is_repository_created_on_first_push", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
			),
		},
		// verify resource import
		{
			Config:                  config + ArtifactsContainerConfigurationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
