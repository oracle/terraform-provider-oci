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
	IotIotDomainGroupConfigureDataAccessRepresentation = map[string]interface{}{
		"db_allow_listed_vcn_ids": acctest.Representation{RepType: acctest.Required, Create: []string{}},
		"iot_domain_group_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_group_id}`},
	}
)

// issue-routing-tag: iot/default
func TestIotIotDomainGroupConfigureDataAccessResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotDomainGroupConfigureDataAccessResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotDomainGroupId := utils.GetEnvSettingWithBlankDefault("iot_domain_group_ocid")
	iotDomainGroupIdVariableStr := fmt.Sprintf("variable \"iot_domain_group_id\" { default = \"%s\" }\n", iotDomainGroupId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_ocid")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	resourceName := "oci_iot_iot_domain_group_configure_data_access.test_iot_domain_group_configure_data_access"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+iotDomainGroupIdVariableStr+vcnIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group_configure_data_access", "test_iot_domain_group_configure_data_access", acctest.Required, acctest.Create, IotIotDomainGroupConfigureDataAccessRepresentation), "iot", "iotDomainGroupConfigureDataAccess", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + iotDomainGroupIdVariableStr + vcnIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group_configure_data_access", "test_iot_domain_group_configure_data_access", acctest.Required, acctest.Create, IotIotDomainGroupConfigureDataAccessRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "db_allow_listed_vcn_ids.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_group_id"),
			),
		},
	})
}
