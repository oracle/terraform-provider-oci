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
	TenantmanagercontrolplaneDomainGovernanceSingularDataSourceRepresentation = map[string]interface{}{
		"domain_governance_id": acctest.Representation{RepType: acctest.Required, Create: `${var.domain_governance_id}`},
	}

	TenantmanagercontrolplaneDomainGovernanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"domain_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.domain_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `${var.domain_governance_state}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneDomainGovernanceDataSourceFilterRepresentation}}
	TenantmanagercontrolplaneDomainGovernanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.domain_governance_id}`}},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneDomainGovernanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneDomainGovernanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	domainId := utils.GetEnvSettingWithBlankDefault("domain_id")
	domainIdVariableStr := fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainId)

	domainGovernanceId := utils.GetEnvSettingWithBlankDefault("domain_governance_id")
	domainGovernanceIdVariableStr := fmt.Sprintf("variable \"domain_governance_id\" { default = \"%s\" }\n", domainGovernanceId)

	state := utils.GetEnvSettingWithBlankDefault("domain_governance_state")
	stateVariableStr := fmt.Sprintf("variable \"domain_governance_state\" { default = \"%s\" }\n", state)

	datasourceName := "data.oci_tenantmanagercontrolplane_domain_governances.test_domain_governances"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_domain_governance.test_domain_governance"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + compartmentIdVariableStr + domainIdVariableStr + stateVariableStr + domainGovernanceIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_domain_governances", "test_domain_governances", acctest.Optional, acctest.Create, TenantmanagercontrolplaneDomainGovernanceDataSourceRepresentation)
	singularDataSourceConfig := config + domainGovernanceIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_domain_governance", "test_domain_governance", acctest.Required, acctest.Create, TenantmanagercontrolplaneDomainGovernanceSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "domain_governance_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_governance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_governance_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_email"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
