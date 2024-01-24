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
	IdentityDomainsIdentityDomainsAccountMgmtInfoSingularDataSourceRepresentation = map[string]interface{}{
		"account_mgmt_info_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_account_mgmt_infos.test_account_mgmt_infos.account_mgmt_infos.0.id}`},
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":       acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsAccountMgmtInfoDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"account_mgmt_info_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"account_mgmt_info_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	IdentityDomainsAccountMgmtInfoResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsAccountMgmtInfoResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsAccountMgmtInfoResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_account_mgmt_infos.test_account_mgmt_infos"
	singularDatasourceName := "data.oci_identity_domains_account_mgmt_info.test_account_mgmt_info"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_account_mgmt_infos", "test_account_mgmt_infos", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAccountMgmtInfoDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAccountMgmtInfoResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "account_mgmt_infos.#"),
				resource.TestCheckResourceAttr(datasourceName, "account_mgmt_infos.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_account_mgmt_infos", "test_account_mgmt_infos", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsAccountMgmtInfoDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_account_mgmt_info", "test_account_mgmt_info", acctest.Optional, acctest.Create, IdentityDomainsIdentityDomainsAccountMgmtInfoSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsAccountMgmtInfoResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "account_mgmt_info_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "account_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "active"),
				resource.TestCheckResourceAttr(singularDatasourceName, "app.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_key"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "do_not_back_fill_grants"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "do_not_perform_action_on_target"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "favorite"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_account"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "last_accessed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_owners.#", "0"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_class.#", "0"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_context"),
				resource.TestCheckResourceAttr(singularDatasourceName, "owner.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "preview_only"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "sync_response"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "sync_situation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sync_timestamp"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tags.#", "0"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "uid"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_wallet_artifact.#", "0"),
			),
		},
	})
}
