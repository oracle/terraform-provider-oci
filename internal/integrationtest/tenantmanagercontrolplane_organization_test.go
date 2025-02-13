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
	TenantmanagercontrolplaneOrganizationSingularDataSourceRepresentation = map[string]interface{}{
		"organization_id": acctest.Representation{RepType: acctest.Required, Create: `${var.organization_id}`},
	}

	TenantmanagercontrolplaneOrganizationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneOrganizationDataSourceFilterRepresentation}}

	TenantmanagercontrolplaneOrganizationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.organization_id}`}},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneOrganizationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneOrganizationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	organizationId := utils.GetEnvSettingWithBlankDefault("organization_id")
	organizationIdVariableStr := fmt.Sprintf("variable \"organization_id\" { default = \"%s\" }\n", organizationId)

	datasourceName := "data.oci_tenantmanagercontrolplane_organizations.test_organizations"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_organization.test_organization"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + compartmentIdVariableStr + organizationIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_organizations", "test_organizations", acctest.Optional, acctest.Update, TenantmanagercontrolplaneOrganizationDataSourceRepresentation)
	singularDataSourceConfig := config + organizationIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization", "test_organization", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "organization_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "organization_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
