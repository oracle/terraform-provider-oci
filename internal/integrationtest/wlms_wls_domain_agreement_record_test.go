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
	WlmsWlsDomainAgreementRecordDataSourceRepresentation = map[string]interface{}{
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainAgreementRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainAgreementRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	// compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_wlms_wls_domain_agreement_records.test_wls_domain_agreement_records"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_agreement_records", "test_wls_domain_agreement_records", acctest.Required, acctest.Create, WlmsWlsDomainAgreementRecordDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "agreement_record_collection.#"),
			),
		},
	})
}
