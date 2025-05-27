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
	ManagementAgentManagementAgentNamedCredentialsMetadataSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"management_agent_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.managed_agent_id}`},
	}

	ManagementAgentManagementAgentNamedCredentialsMetadataResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, CloudBridgeAgentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Required, acctest.Create, ManagementAgentManagementAgentRepresentation)
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentNamedCredentialsMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentNamedCredentialsMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//  1. List all agents in compartment (env:TF_VAR_compartment_ocid) with status=ACTIVE and displayName=terraformTest
	managementAgentIds, err := getManagementAgentIds(compartmentId)
	if err != nil {
		t.Errorf("Failed to get agents in compartment %s", err)
	}
	if len(managementAgentIds) == 0 {
		t.Errorf("Failed to find any active agents in compartment %s", compartmentId)
	}
	managementAgentId := managementAgentIds[0]
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_named_credentials_metadata.test_management_agent_named_credentials_metadata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_named_credentials_metadata", "test_management_agent_named_credentials_metadata", acctest.Required, acctest.Create, ManagementAgentManagementAgentNamedCredentialsMetadataSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.minimum_agent_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.properties.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.properties.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.properties.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.properties.0.value_category.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata.0.properties.0.is_required"),
			),
		},
	})
}
