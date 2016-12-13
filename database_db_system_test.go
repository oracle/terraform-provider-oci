package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceDatabaseDBSystemTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DBSystem
	DeletedRes   *baremetal.DBSystem
}

func (s *ResourceDatabaseDBSystemTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	dbHomeOpts := &baremetal.DisplayNameOptions{}
	dbHomeOpts.DisplayName = "db_home_display_name"
	dbHome := baremetal.NewCreateDBHomeDetails(
		"admin_password", "db_name", "db_version", dbHomeOpts,
	)
	opts := &baremetal.LaunchDBSystemOptions{}
	opts.DisplayName = "display_name"
	opts.DatabaseEdition = baremetal.DatabaseEditionStandard
	opts.DBHome = dbHome
	opts.DiskRedundancy = baremetal.DiskRedundancyNormal
	opts.Domain = "domain.com"
	opts.Hostname = "hostname"

	s.Res = &baremetal.DBSystem{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		Shape:              "shape",
		SubnetID:           "subnet_id",
		SSHPublicKeys:      []string{"ansshkey"},
		CPUCoreCount:       2,

		DisplayName:      opts.DisplayName,
		DatabaseEdition:  opts.DatabaseEdition,
		DBHome:           opts.DBHome,
		DiskRedundancy:   opts.DiskRedundancy,
		Domain:           opts.Domain,
		Hostname:         opts.Hostname,
		ID:               "id",
		LifecycleDetails: "lifecycle_details",
		ListenerPort:     1,
		State:            baremetal.ResourceAvailable,
		TimeCreated:      s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.Client.On("CreateDBSystem",
		s.Res.AvailabilityDomain, s.Res.CompartmentID, s.Res.Shape, s.Res.SubnetID,
		s.Res.SSHPublicKeys, s.Res.CPUCoreCount, opts,
	).Return(s.Res, nil)

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceTerminated

	s.Client.On("TerminateDBSystem", "id", &baremetal.IfMatchOptions(nil)).Return(nil)

	s.ResourceName = "baremetal_database_db_system.t"
	template := `
		resource "baremetal_database_db_system" "t" {
			availability_domain = "availability_domain"
			compartment_id = "compartment_id"
			shape = "shape"
			subnet_id = "subnet_id"
			ssh_public_keys = ["ansshkey"]
			cpu_database_count = 2
			display_name = "display_name"
			database_edition = "database_edition"
      db_home {
				database {
					"admin_password" = "123456789"
					"db_name" = "db_name"
				}
				db_version = "db_version"
				display_name = "db_home_display_name"
			}
			disk_redundancy = "disk_redundancy"
			domain = "domain.com"
			hostname = "hostname"
		}
	`
	var buf bytes.Buffer
	tmpl := template.Must(template.New("config").Parse(template))
	tmpl.Execute(buf, s.Res)
	s.Config = buf.String()
	s.Config += testProviderConfig
}

func (s *ResourceDatabaseDBSystemTestSuite) TestCreateResourcedatabaseSecurityList() {
	// s.Client.On("GetDBSystem", "id").Return(s.Res, nil).Times(2)
	// s.Client.On("GetDBSystem", "id").Return(s.DeletedRes, nil)

	// resource.UnitTest(s.T(), resource.TestCase{
	// 	Providers: s.Providers,
	// 	Steps: []resource.TestStep{
	// 		resource.TestStep{
	// 			Config: s.Config,
	// 			Check: resource.ComposeTestCheckFunc(
	// 				resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
	// 				resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
	// 				resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.0.code", "1"),
	// 				resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.tcp_options.0.max", "2"),
	// 			),
	// 		},
	// 	},
	// })
}

// func (s *ResourceDatabaseDBSystemTestSuite) TestDeleteSecurityList() {
// 	s.Client.On("GetDBSystem", "id").Return(s.Res, nil).Times(2)
// 	s.Client.On("GetDBSystem", "id").Return(s.DeletedRes, nil)

// 	resource.UnitTest(s.T(), resource.TestCase{
// 		Providers: s.Providers,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: s.Config,
// 			},
// 			resource.TestStep{
// 				Config:  s.Config,
// 				Destroy: true,
// 			},
// 		},
// 	})

// 	s.Client.AssertCalled(s.T(), "DeleteDBSystem", "id", []baremetal.Options(nil))
// }

func TestResourceDatabaseDBSystemTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceDatabaseDBSystemTestSuite))
}
