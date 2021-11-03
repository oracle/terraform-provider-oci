// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ContainerConfigurationResourceConfig = ContainerConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", Optional, Update, containerConfigurationRepresentation)

	containerConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	containerConfigurationRepresentation = map[string]interface{}{
		"compartment_id":                      Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"is_repository_created_on_first_push": Representation{RepType: Required, Create: `false`, Update: `true`},
	}

	ContainerConfigurationResourceDependencies = ""
)

// issue-routing-tag: artifacts/default
func TestArtifactsContainerConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsContainerConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_artifacts_container_configuration.test_container_configuration"

	singularDatasourceName := "data.oci_artifacts_container_configuration.test_container_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ContainerConfigurationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", Required, Create, containerConfigurationRepresentation), "artifacts", "containerConfiguration", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", Required, Create, containerConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_repository_created_on_first_push", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ContainerConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", Optional, Update, containerConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "is_repository_created_on_first_push", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_artifacts_container_configuration", "test_container_configuration", Required, Create, containerConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerConfigurationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(singularDatasourceName, "is_repository_created_on_first_push", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
			),
		},
	})
}
