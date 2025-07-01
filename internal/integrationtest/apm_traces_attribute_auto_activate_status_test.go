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
	ApmTracesAttributeAutoActivateStatusSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"data_key_type": acctest.Representation{RepType: acctest.Required, Create: `PRIVATE_DATA_KEY`},
	}

	//ApmTracesAttributeAutoActivateStatusResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_traces/default
func TestApmTracesAttributeAutoActivateStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesAttributeAutoActivateStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	//This is a manual test. It requires apm_domain_id and trace_key as environment variables.
	apmDomainId := utils.GetEnvSettingWithBlankDefault("apm_domain_id")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	if apmDomainId == "" {
		t.Skip("Set apm_domain_id and data_key_type to run this test")
	}

	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)

	singularDatasourceName := "data.oci_apm_traces_attribute_auto_activate_status.test_attribute_auto_activate_status"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_attribute_auto_activate_status", "test_attribute_auto_activate_status", acctest.Required, acctest.Create, ApmTracesAttributeAutoActivateStatusSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_key_type", "PRIVATE_DATA_KEY"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
