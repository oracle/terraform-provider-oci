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
	IotIotDomainConfigureDirectDataAccessRepresentation = map[string]interface{}{
		"iot_domain_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `DIRECT`},
		"db_allow_listed_identity_group_names": acctest.Representation{RepType: acctest.Required, Create: []string{}},
	}

	IotIotDomainConfigureOrdsDataAccessRepresentation = map[string]interface{}{
		"iot_domain_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"type":                            acctest.Representation{RepType: acctest.Required, Create: `ORDS`},
		"db_allowed_identity_domain_host": acctest.Representation{RepType: acctest.Required, Create: `${var.identity_host}`},
	}
)

// issue-routing-tag: iot/default
func TestIotIotDomainConfigureDataAccessResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotDomainConfigureDataAccessResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	identityHost := utils.GetEnvSettingWithBlankDefault("identity_host")
	identityHostVariableStr := fmt.Sprintf("variable \"identity_host\" { default = \"%s\" }\n", identityHost)

	resourceName := "oci_iot_iot_domain_configure_data_access.test_iot_domain_configure_data_access"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+iotDomainIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_configure_data_access", "test_iot_domain_configure_data_access", acctest.Required, acctest.Create, IotIotDomainConfigureDirectDataAccessRepresentation), "iot", "iotDomainConfigureDataAccess", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + iotDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_configure_data_access", "test_iot_domain_configure_data_access", acctest.Required, acctest.Create, IotIotDomainConfigureDirectDataAccessRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "db_allow_listed_identity_group_names.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "DIRECT"),
			),
		},
		{
			Config: config + iotDomainIdVariableStr + identityHostVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_configure_data_access", "test_iot_domain_configure_data_access", acctest.Required, acctest.Create, IotIotDomainConfigureOrdsDataAccessRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_allowed_identity_domain_host"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORDS"),
			),
		},
	})
}
