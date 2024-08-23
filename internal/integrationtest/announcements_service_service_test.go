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
	AnnouncementsServiceServiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"comms_manager_name": acctest.Representation{RepType: acctest.Optional, Create: `CN`},
		"platform_type":      acctest.Representation{RepType: acctest.Optional, Create: `IAAS`},
	}

	AnnouncementsServiceServiceResourceConfig = ""
)

// issue-routing-tag: announcements_service/default
func TestAnnouncementsServiceServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnnouncementsServiceServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_announcements_service_services.test_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_announcements_service_services", "test_services", acctest.Required, acctest.Create, AnnouncementsServiceServiceDataSourceRepresentation) +
				compartmentIdVariableStr + AnnouncementsServiceServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "services_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "services_collection.0.items.#"),
			),
		},
	})
}

func TestAnnouncementsServiceServiceResource_WithOptionalParams(t *testing.T) {
	httpreplay.SetScenario("TestAnnouncementsServiceServiceResource_WithOptionalParams")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_announcements_service_services.test_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_announcements_service_services", "test_services", acctest.Optional, acctest.Create, AnnouncementsServiceServiceDataSourceRepresentation) +
				compartmentIdVariableStr + AnnouncementsServiceServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "comms_manager_name", "CN"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "platform_type", "IAAS"),

				resource.TestCheckResourceAttrSet(datasourceName, "services_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "services_collection.0.items.#"),
			),
		},
	})
}
