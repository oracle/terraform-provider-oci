// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	wlsDomainDisplayName                          = utils.GetEnvSettingWithBlankDefault("wls_domain_display_name")
	wlsDomainMWType                               = utils.GetEnvSettingWithBlankDefault("wls_domain_mw_type")
	wlsVersion                                    = utils.GetEnvSettingWithBlankDefault("wls_mw_version")
	WlmsWlsDomainSingularDataSourceRepresentation = map[string]interface{}{
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}

	WlmsWlsDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: wlmsCompartmentOcid},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: wlsDomainDisplayName},
		"id":                     acctest.Representation{RepType: acctest.Optional, Create: wlsDomainOcid},
		"middleware_type":        acctest.Representation{RepType: acctest.Optional, Create: wlsDomainMWType},
		"patch_readiness_status": acctest.Representation{RepType: acctest.Optional, Create: `OK`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"weblogic_version":       acctest.Representation{RepType: acctest.Optional, Create: wlsVersion},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	// compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	datasourceName := "data.oci_wlms_wls_domains.test_wls_domains"
	singularDatasourceName := "data.oci_wlms_wls_domain.test_wls_domain"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domains", "test_wls_domains", acctest.Optional, acctest.Create, WlmsWlsDomainDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "wls_domain_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain", "test_wls_domain", acctest.Required, acctest.Create, WlmsWlsDomainSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_accepted_terms_and_conditions"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "middleware_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_readiness_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "weblogic_version"),
			),
		},
	})
}
