// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	WlmsWlsDomainApplicablePatchDataSourceRepresentation = map[string]interface{}{
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainApplicablePatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainApplicablePatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_wls_domain_applicable_patches.test_wls_domain_applicable_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_applicable_patches", "test_wls_domain_applicable_patches", acctest.Required, acctest.Create, WlmsWlsDomainApplicablePatchDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "applicable_patch_collection.#"),
			),
		},
	})
}
