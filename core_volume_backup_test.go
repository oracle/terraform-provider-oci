package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VolumeBackup
	DeletedRes   *baremetal.VolumeBackup
	Opts         []baremetal.Options
}

func (s *ResourceCoreVolumeBackupTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)
	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}

	s.ResourceName = "baremetal_core_volume_backup.t"
	s.Config = `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "volume_id"
			display_name = "display_name"
		}
	`
	s.Config += testProviderConfig

	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.Res = &baremetal.VolumeBackup{
		CompartmentID:       "compartment_id",
		DisplayName:         "display_name",
		ID:                  "id",
		State:               baremetal.ResourceAvailable,
		SizeInMBs:           1,
		TimeCreated:         s.TimeCreated,
		TimeRequestReceived: s.TimeCreated,
		UniqueSizeInMBs:     1,
		VolumeID:            "volume_id",
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := baremetal.Options{DisplayName: "display_name"}
	s.Opts = []baremetal.Options{opts}
	s.Client.On("CreateVolumeBackup", "volume_id", s.Opts).Return(s.Res, nil)
	s.Client.On("DeleteVolumeBackup", "id", []baremetal.Options(nil)).Return(nil)
}

func (s *ResourceCoreVolumeBackupTestSuite) TestCreateVolumeBackup() {
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", fmt.Sprintf("%v", s.Res.SizeInMBs)),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "time_request_received", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "unique_size_in_mbs", fmt.Sprintf("%v", s.Res.UniqueSizeInMBs)),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_id", s.Res.VolumeID),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeBackupTestSuite) TestCreateVolumeBackupWithoutDisplayName() {
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.DeletedRes, nil)

	s.Config = `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "volume_id"
		}
	`
	s.Config += testProviderConfig

	opts := baremetal.Options{}
	s.Client.On("CreateVolumeBackup", "volume_id", []baremetal.Options{opts}).
		Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreVolumeBackupTestSuite) TestUpdateVolumeBackupDisplayName() {
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "volume_id"
			display_name = "new_display_name"
		}
	`
	config += testProviderConfig

	resVal := *s.Res
	res := &resVal
	res.DisplayName = "new_display_name"

	deletedResVal := *res
	deletedRes := &deletedResVal
	deletedRes.State = baremetal.ResourceTerminated

	opts := baremetal.Options{DisplayName: "new_display_name"}
	s.Client.On("UpdateVolumeBackup", "id", []baremetal.Options{opts}).Return(res, nil)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(deletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreVolumeBackupTestSuite) TestUpdateVolumeIDForcesNewVolumeBackup() {
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(3)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.DeletedRes, nil).Once()

	config := `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "new_volume_id"
		}
  `
	config += testProviderConfig

	resVal := *s.Res
	res := &resVal
	res.VolumeID = "new_volume_id"

	deletedResVal := *res
	deletedRes := &deletedResVal
	deletedRes.State = baremetal.ResourceTerminated

	opts := baremetal.Options{}
	s.Client.On("CreateVolumeBackup", res.VolumeID, []baremetal.Options{opts}).Return(res, nil)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(deletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "volume_id", res.VolumeID),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeBackupTestSuite) TestDeleteVolumeBackup() {
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id", []baremetal.Options(nil)).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteVolumeBackup", "id", []baremetal.Options(nil))
}

func TestResourceCoreVolumeBackupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeBackupTestSuite))
}
