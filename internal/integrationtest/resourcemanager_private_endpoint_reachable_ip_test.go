// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	privateEndpointReachableIpSingularDataSourceRepresentation = map[string]interface{}{
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resourcemanager_private_endpoint.test_private_endpoint.id}`},
		"private_ip":          acctest.Representation{RepType: acctest.Required, Create: `privateIp`},
	}

	PrivateEndpointReachableIpResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, resourceManagerprivateEndpointRepresentation)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_private_endpoint_reachable_ip", "test_private_endpoint_reachable_ip", acctest.Required, acctest.Create, privateEndpointReachableIpSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PrivateEndpointReachableIpResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_ip", "privateIp"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "reachable_ip"),
			),
		},
	})
}
