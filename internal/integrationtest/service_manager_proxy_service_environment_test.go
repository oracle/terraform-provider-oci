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
	ServiceManagerProxyServiceManagerProxyServiceEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.service_environment_id}`},
	}

	ServiceManagerProxyServiceEnvironmentResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_service_manager_proxy_service_environment", "test_service_environment", acctest.Required, acctest.Create, ServiceManagerProxyServiceManagerProxyServiceEnvironmentSingularDataSourceRepresentation)

	ServiceManagerProxyServiceManagerProxyServiceEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ServiceEnvironmentsResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_service_manager_proxy_service_environments", "test_service_environments", acctest.Required, acctest.Create, ServiceManagerProxyServiceManagerProxyServiceEnvironmentDataSourceRepresentation)
)

// issue-routing-tag: service_manager_proxy/default
func TestServiceManagerProxyServiceEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceManagerProxyServiceEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	serviceEnvironmentId := utils.GetEnvSettingWithBlankDefault("service_environment_id")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	serviceEnvironmentIdVariableStr := fmt.Sprintf("variable \"service_environment_id\" { default = \"%s\" }\n", serviceEnvironmentId)

	datasourceName := "data.oci_service_manager_proxy_service_environments.test_service_environments"
	singularDatasourceName := "data.oci_service_manager_proxy_service_environment.test_service_environment"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_manager_proxy_service_environments", "test_service_environments", acctest.Required, acctest.Create, ServiceManagerProxyServiceManagerProxyServiceEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_environment_collection.0.items.0.service_definition.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_environment_collection.0.items.0.service_definition.0.type"),
				resource.TestCheckResourceAttr(datasourceName, "service_environment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_manager_proxy_service_environment", "test_service_environment", acctest.Required, acctest.Create, ServiceManagerProxyServiceManagerProxyServiceEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + serviceEnvironmentIdVariableStr,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "console_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_definition.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
			),
		},
	})
}
