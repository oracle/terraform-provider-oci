// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client"
)

type ResourceCoreVolumeBackupsTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListVolumeBackups
}

func (s *ResourceCoreVolumeBackupsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
data "baremetal_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}
resource "baremetal_core_volume" "t" {
	availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
	size_in_mbs = 262144
}
resource "baremetal_core_volume_backup" "t" {
	volume_id = "${baremetal_core_volume.t.id}"
	display_name = "display_name"
}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_volume_backups.t"

	b1 := baremetal.VolumeBackup{
		CompartmentID:       "compartment_id",
		DisplayName:         "display_name",
		ID:                  "id1",
		State:               baremetal.ResourceAvailable,
		SizeInMBs:           1,
		TimeCreated:         baremetal.Time{Time: time.Now()},
		TimeRequestReceived: baremetal.Time{Time: time.Now()},
		UniqueSizeInMBs:     1,
		VolumeID:            "volume_id",
	}

	b2 := b1
	b2.ID = "id2"

	s.List = &baremetal.ListVolumeBackups{
		VolumeBackups: []baremetal.VolumeBackup{b1, b2},
	}
}

func (s *ResourceCoreVolumeBackupsTestSuite) TestReadVolumeBackups() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "baremetal_core_volume_backups" "t" {
					compartment_id = "${var.compartment_id}"
					limit = 1
					volume_id = "${baremetal_core_volume.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "1"),
				),
			},
		},
	},
	)
}

func TestResourceCoreVolumeBackupsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeBackupsTestSuite))
}
