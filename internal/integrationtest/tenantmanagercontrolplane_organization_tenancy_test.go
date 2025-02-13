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
	TenantmanagercontrolplaneOrganizationTenancySingularDataSourceRepresentation = map[string]interface{}{
		"organization_id": acctest.Representation{RepType: acctest.Required, Create: `${var.organization_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
	}

	TenantmanagercontrolplaneOrganizationTenancyDataSourceRepresentation = map[string]interface{}{
		"organization_id": acctest.Representation{RepType: acctest.Required, Create: `${var.organization_id}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneOrganizationTenancyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneOrganizationTenancyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	organizationId := utils.GetEnvSettingWithBlankDefault("organization_id")
	organizationIdVariableStr := fmt.Sprintf("variable \"organization_id\" { default = \"%s\" }\n", organizationId)

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_id")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	datasourceName := "data.oci_tenantmanagercontrolplane_organization_tenancies.test_organization_tenancies"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_organization_tenancy.test_organization_tenancy"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + organizationIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancies", "test_organization_tenancies", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancyDataSourceRepresentation)
	singularDataSourceConfig := config + organizationIdVariableStr + tenancyIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancy", "test_organization_tenancy", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancySingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "organization_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "organization_tenancy_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "organization_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_approved_for_transfer"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_joined"),
			),
		},
	})
}
