// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
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
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_volume_backups" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
      volume_id = "volume_id"
    }
  `
	s.Config += testProviderConfig
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
	opts := &baremetal.ListBackupsOptions{}
	opts.VolumeID = "volume_id"
	opts.Limit = 1
	opts.Page = "page"

	s.Client.On("ListVolumeBackups", "compartment_id", opts).Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "volume_id", "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumeBackups", "compartment_id", opts)
}

func (s *ResourceCoreVolumeBackupsTestSuite) TestReadVolumeBackupsWithPagination() {
	opts := &baremetal.ListBackupsOptions{}
	opts.Limit = 1
	opts.Page = "page"
	opts.VolumeID = "volume_id"

	listVal := *s.List
	list := &listVal
	list.NextPage = "nextpage"
	s.Client.On("ListVolumeBackups", "compartment_id", opts).Return(list, nil)

	opts2 := &baremetal.ListBackupsOptions{}
	opts2.VolumeID = "volume_id"
	opts2.Limit = 1
	opts2.Page = "nextpage"

	list2Val := *s.List
	list2 := &list2Val
	b3 := s.List.VolumeBackups[0]
	b3.ID = "id3"
	b4 := s.List.VolumeBackups[1]
	b4.ID = "id4"
	list2.VolumeBackups = []baremetal.VolumeBackup{b3, b4}
	s.Client.On("ListVolumeBackups", "compartment_id", opts2).Return(list2, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumeBackups", "compartment_id", opts2)
}

func TestResourceCoreVolumeBackupsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeBackupsTestSuite))
}
