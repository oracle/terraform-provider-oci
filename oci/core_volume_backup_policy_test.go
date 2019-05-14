// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	volumeBackupPolicyDataSourceRepresentation = map[string]interface{}{}

	VolumeBackupPolicyResourceConfig = ""
)

func TestCoreVolumeBackupPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupPolicyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_volume_backup_policies.test_volume_backup_policies"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_volume_backup_policies", "test_volume_backup_policies", Required, Create, volumeBackupPolicyDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeBackupPolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.schedules.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.time_created"),
				),
			},
		},
	})
}
