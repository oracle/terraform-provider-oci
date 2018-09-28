// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	volumeBackupPolicyDataSourceRepresentation = map[string]interface{}{
		"filter": RepresentationGroup{Required, volumeBackupPolicyDataSourceFilterRepresentation},
	}
	volumeBackupPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `display_name`},
		"values": Representation{repType: Required, create: []string{`silver`}},
	}

	VolumeBackupPolicyResourceConfig = ""
)

func TestCoreVolumeBackupPolicyResource_basic(t *testing.T) {
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

					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.display_name", "silver"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backup_policies.0.schedules.#", "3"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backup_policies.0.time_created"),
				),
			},
		},
	})
}
