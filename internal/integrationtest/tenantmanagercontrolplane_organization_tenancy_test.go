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
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_tenantmanagercontrolplane_organization_tenancy.test_organization_tenancy.tenancy_id}`},
	}

	TenantmanagercontrolplaneOrganizationTenancyDataSourceRepresentation = map[string]interface{}{
		"organization_id": acctest.Representation{RepType: acctest.Required, Create: `${var.organization_id}`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneOrganizationTenancyDataSourceFilterRepresentation},
	}
	TenantmanagercontrolplaneOrganizationTenancyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `tenancy_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_tenantmanagercontrolplane_organization_tenancy.test_organization_tenancy.tenancy_id}`}},
	}

	TenantmanagercontrolplaneOrganizationTenancyRepresentation = map[string]interface{}{
		"admin_email":     acctest.Representation{RepType: acctest.Required, Create: `${var.organization_tenancy_admin_email}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.organization_tenancy_compartment_id}`},
		"home_region":     acctest.Representation{RepType: acctest.Required, Create: `${var.organization_tenancy_home_region}`},
		"organization_id": acctest.Representation{RepType: acctest.Required, Create: `${var.organization_id}`},
		"tenancy_name":    acctest.Representation{RepType: acctest.Required, Create: `${var.organization_tenancy_name}`},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
//
// Note: destroying an oci_tenantmanagercontrolplane_organization_tenancy only removes it from
// Terraform state; it does not terminate the child tenancy (termination must be done from within
// the child tenancy or via an Oracle Support request). The test therefore does not assert that the
// tenancy is deleted, and there is no sweeper that attempts to delete it.
func TestTenantmanagercontrolplaneOrganizationTenancyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneOrganizationTenancyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("organization_tenancy_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"organization_tenancy_compartment_id\" { default = \"%s\" }\n", compartmentId)

	organizationId := utils.GetEnvSettingWithBlankDefault("organization_id")
	organizationIdVariableStr := fmt.Sprintf("variable \"organization_id\" { default = \"%s\" }\n", organizationId)

	adminEmail := utils.GetEnvSettingWithBlankDefault("organization_tenancy_admin_email")
	adminEmailVariableStr := fmt.Sprintf("variable \"organization_tenancy_admin_email\" { default = \"%s\" }\n", adminEmail)

	homeRegion := utils.GetEnvSettingWithBlankDefault("organization_tenancy_home_region")
	homeRegionVariableStr := fmt.Sprintf("variable \"organization_tenancy_home_region\" { default = \"%s\" }\n", homeRegion)

	tenancyName := utils.GetEnvSettingWithBlankDefault("organization_tenancy_name")
	tenancyNameVariableStr := fmt.Sprintf("variable \"organization_tenancy_name\" { default = \"%s\" }\n", tenancyName)

	variableStr := compartmentIdVariableStr + organizationIdVariableStr + adminEmailVariableStr + homeRegionVariableStr + tenancyNameVariableStr

	resourceName := "oci_tenantmanagercontrolplane_organization_tenancy.test_organization_tenancy"
	datasourceName := "data.oci_tenantmanagercontrolplane_organization_tenancies.test_organization_tenancies"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_organization_tenancy.test_organization_tenancy"

	acctest.SaveConfigContent("", "", "", t)

	createConfig := config + variableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancy", "test_organization_tenancy", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancyRepresentation)
	dataSourceConfig := config + variableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancy", "test_organization_tenancy", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancyRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancies", "test_organization_tenancies", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancyDataSourceRepresentation)
	singularDataSourceConfig := config + variableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancy", "test_organization_tenancy", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancyRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_organization_tenancy", "test_organization_tenancy", acctest.Required, acctest.Create, TenantmanagercontrolplaneOrganizationTenancySingularDataSourceRepresentation)

	fmt.Printf("Create Config: %s\n", createConfig)
	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: createConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "home_region"),
				resource.TestCheckResourceAttrSet(resourceName, "organization_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "organization_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "organization_tenancy_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "organization_tenancy_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "organization_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "governance_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
