// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below
	JmsFleetInstallationSiteWithAdvancedFeature = utils.GetEnvSettingWithBlankDefault("fleet_advanced_feature_ocid")

	JmsFleetInstallationSiteSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: JmsFleetInstallationSiteWithAdvancedFeature},
		"application_id":      acctest.Representation{RepType: acctest.Optional, Create: `dummy.application.id`},
		"installation_path":   acctest.Representation{RepType: acctest.Optional, Create: `installationPath`},
		"jre_distribution":    acctest.Representation{RepType: acctest.Optional, Create: `jreDistribution`},
		"jre_security_status": acctest.Representation{RepType: acctest.Optional, Create: `EARLY_ACCESS`},
		"jre_vendor":          acctest.Representation{RepType: acctest.Optional, Create: `jreVendor`},
		"jre_version":         acctest.Representation{RepType: acctest.Optional, Create: `jreVersion`},
		"os_family":           acctest.Representation{RepType: acctest.Optional, Create: []string{`LINUX`}},
		"path_contains":       acctest.Representation{RepType: acctest.Optional, Create: `installationPath`},
		"time_start":          acctest.Representation{RepType: acctest.Optional, Create: `2022-07-01T01:00:00Z`},
		"time_end":            acctest.Representation{RepType: acctest.Optional, Create: `2022-07-20T01:00:00Z`},
	}

	JmsFleetInstallationSiteDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: JmsFleetInstallationSiteWithAdvancedFeature},
		"application_id":      acctest.Representation{RepType: acctest.Optional, Create: `dummy.application.id`},
		"installation_path":   acctest.Representation{RepType: acctest.Optional, Create: `installationPath`},
		"jre_distribution":    acctest.Representation{RepType: acctest.Optional, Create: `jreDistribution`},
		"jre_security_status": acctest.Representation{RepType: acctest.Optional, Create: `UNKNOWN`},
		"jre_vendor":          acctest.Representation{RepType: acctest.Optional, Create: `jreVendor`},
		"jre_version":         acctest.Representation{RepType: acctest.Optional, Create: `jreVersion`},
		"os_family":           acctest.Representation{RepType: acctest.Optional, Create: []string{`LINUX`}},
		"path_contains":       acctest.Representation{RepType: acctest.Optional, Create: `installationPath`},
		"time_start":          acctest.Representation{RepType: acctest.Optional, Create: `2022-07-01T01:00:00Z`},
		"time_end":            acctest.Representation{RepType: acctest.Optional, Create: `2022-07-20T01:00:00Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetInstallationSiteResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetInstallationSiteResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_installation_sites.test_fleet_installation_sites"
	singularDatasourceName := "data.oci_jms_fleet_installation_site.test_fleet_installation_site"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_installation_sites",
					"test_fleet_installation_sites",
					acctest.Optional,
					acctest.Create,
					JmsFleetInstallationSiteDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "application_id", "dummy.application.id"),
				resource.TestCheckResourceAttr(datasourceName, "installation_path", "installationPath"),
				resource.TestCheckResourceAttr(datasourceName, "jre_distribution", "jreDistribution"),
				resource.TestCheckResourceAttr(datasourceName, "jre_security_status", "UNKNOWN"),
				resource.TestCheckResourceAttr(datasourceName, "jre_vendor", "jreVendor"),
				resource.TestCheckResourceAttr(datasourceName, "jre_version", "jreVersion"),
				resource.TestCheckResourceAttr(datasourceName, "os_family.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "path_contains"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_start"),

				resource.TestCheckResourceAttrSet(datasourceName, "installation_site_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "installation_site_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_installation_site",
					"test_fleet_installation_site",
					acctest.Optional,
					acctest.Create,
					JmsFleetInstallationSiteSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "application_id", "dummy.application.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "installation_path", "installationPath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_distribution", "jreDistribution"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_security_status", "EARLY_ACCESS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_vendor", "jreVendor"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_version", "jreVersion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "path_contains"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_start"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
