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
	TenantmanagercontrolplaneDomainSingularDataSourceRepresentation = map[string]interface{}{
		"domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.domain_id}`},
	}

	TenantmanagercontrolplaneDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"domain_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.domain_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.domain_name}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `${var.domain_state}`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `${var.domain_status}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneDomainDataSourceFilterRepresentation}}
	TenantmanagercontrolplaneDomainDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.domain_id}`}},
	}

	TenantmanagercontrolplaneDomainResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	domainId := utils.GetEnvSettingWithBlankDefault("domain_id")
	domainIdVariableStr := fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainId)

	domainName := utils.GetEnvSettingWithBlankDefault("domain_name")
	domainNameVariableStr := fmt.Sprintf("variable \"domain_name\" { default = \"%s\" }\n", domainName)

	domainState := utils.GetEnvSettingWithBlankDefault("domain_state")
	domainStateVariableStr := fmt.Sprintf("variable \"domain_state\" { default = \"%s\" }\n", domainState)

	domainStatus := utils.GetEnvSettingWithBlankDefault("domain_status")
	domainStatusVariableStr := fmt.Sprintf("variable \"domain_status\" { default = \"%s\" }\n", domainStatus)

	datasourceName := "data.oci_tenantmanagercontrolplane_domains.test_domains"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_domain.test_domain"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + compartmentIdVariableStr + domainIdVariableStr + domainNameVariableStr + domainStateVariableStr + domainStatusVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_domains", "test_domains", acctest.Optional, acctest.Create, TenantmanagercontrolplaneDomainDataSourceRepresentation)
	singularDataSourceConfig := config + domainIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_domain", "test_domain", acctest.Required, acctest.Create, TenantmanagercontrolplaneDomainSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "domain_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),
				resource.TestCheckResourceAttrSet(datasourceName, "status"),

				resource.TestCheckResourceAttrSet(datasourceName, "domain_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "domain_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "txt_record"),
			),
		},
	})
}
