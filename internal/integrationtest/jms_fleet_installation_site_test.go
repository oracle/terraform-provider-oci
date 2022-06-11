// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"

	"terraform-provider-oci/internal/utils"
)

var (
	JmsJmsFleetInstallationSiteSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
		"application_id":      acctest.Representation{RepType: acctest.Optional, Create: `dummy.application.id`},
		"installation_path":   acctest.Representation{RepType: acctest.Optional, Create: `installationPath`},
		"jre_distribution":    acctest.Representation{RepType: acctest.Optional, Create: `jreDistribution`},
		"jre_security_status": acctest.Representation{RepType: acctest.Optional, Create: `UNKNOWN`},
		"jre_vendor":          acctest.Representation{RepType: acctest.Optional, Create: `jreVendor`},
		"jre_version":         acctest.Representation{RepType: acctest.Optional, Create: `jreVersion`},
		"os_family":           acctest.Representation{RepType: acctest.Optional, Create: []string{`LINUX`}},
	}

	JmsJmsFleetInstallationSiteDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
		"application_id":      acctest.Representation{RepType: acctest.Optional, Create: `dummy.application.id`},
		"installation_path":   acctest.Representation{RepType: acctest.Optional, Create: `installationPath`},
		"jre_distribution":    acctest.Representation{RepType: acctest.Optional, Create: `jreDistribution`},
		"jre_security_status": acctest.Representation{RepType: acctest.Optional, Create: `UNKNOWN`},
		"jre_vendor":          acctest.Representation{RepType: acctest.Optional, Create: `jreVendor`},
		"jre_version":         acctest.Representation{RepType: acctest.Optional, Create: `jreVersion`},
		"os_family":           acctest.Representation{RepType: acctest.Optional, Create: []string{`LINUX`}},
	}

	fleetForInstallationSiteRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Created Fleet for Installation Site`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet for Installation Site`},
		"inventory_log":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetInventoryLogRepresentation},
		"operation_log":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetInventoryLogRepresentation},
	}

	FleetInstallationSiteResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, fleetForInstallationSiteRepresentation)
)

// issue-routing-tag: jms/default
func TestJmsFleetInstallationSiteResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetInstallationSiteResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_jms_fleet_installation_sites.test_fleet_installation_sites"
	singularDatasourceName := "data.oci_jms_fleet_installation_site.test_fleet_installation_site"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_fleet_installation_sites", "test_fleet_installation_sites", acctest.Optional, acctest.Create, JmsJmsFleetInstallationSiteDataSourceRepresentation) +
				compartmentIdVariableStr + FleetInstallationSiteResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "application_id", "dummy.application.id"),
				resource.TestCheckResourceAttr(datasourceName, "installation_path", "installationPath"),
				resource.TestCheckResourceAttr(datasourceName, "jre_distribution", "jreDistribution"),
				resource.TestCheckResourceAttr(datasourceName, "jre_security_status", "UNKNOWN"),
				resource.TestCheckResourceAttr(datasourceName, "jre_vendor", "jreVendor"),
				resource.TestCheckResourceAttr(datasourceName, "jre_version", "jreVersion"),
				resource.TestCheckResourceAttr(datasourceName, "os_family.#", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "installation_site_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "installation_site_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_fleet_installation_site", "test_fleet_installation_site", acctest.Optional, acctest.Create, JmsJmsFleetInstallationSiteSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetInstallationSiteResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "application_id", "dummy.application.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "installation_path", "installationPath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_distribution", "jreDistribution"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_security_status", "UNKNOWN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_vendor", "jreVendor"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_version", "jreVersion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family.#", "1"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetInstallationSite") {
		resource.AddTestSweepers("JmsFleetInstallationSite", &resource.Sweeper{
			Name:         "JmsFleetInstallationSite",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
