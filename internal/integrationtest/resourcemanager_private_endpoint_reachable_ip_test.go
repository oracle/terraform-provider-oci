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
	ResourcemanagerResourcemanagerPrivateEndpointReachableIpSingularDataSourceRepresentation = map[string]interface{}{
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resourcemanager_private_endpoint.test_private_endpoint.id}`},
		"private_ip":          acctest.Representation{RepType: acctest.Required, Create: `privateIp`},
	}

	ResourcemanagerPrivateEndpointReachableIpResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, ResourceManagerprivateEndpointRepresentation)
)

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerPrivateEndpointReachableIpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourcemanagerPrivateEndpointReachableIpResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_resourcemanager_private_endpoint_reachable_ip.test_private_endpoint_reachable_ip"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_private_endpoint_reachable_ip", "test_private_endpoint_reachable_ip", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerPrivateEndpointReachableIpSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourcemanagerPrivateEndpointReachableIpResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_ip", "privateIp"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
			),
		},
	})
}
