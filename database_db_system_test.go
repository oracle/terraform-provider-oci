package main

import (
	"bytes"
	"testing"
	"time"

	"text/template"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatabaseDBSystemTestSuite struct {
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

func (s *DatabaseDBSystemTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) { return s.Client, nil },
	)
	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}

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

		DisplayName:     opts.DisplayName,
		DatabaseEdition: opts.DatabaseEdition,
		DBHome:          opts.DBHome,
		DiskRedundancy:  opts.DiskRedundancy,
		Domain:          opts.Domain,
		Hostname:        opts.Hostname,

		ID:               "id",
		LifecycleDetails: "lifecycle_details",
		ListenerPort:     1,
		State:            baremetal.ResourceAvailable,
		TimeCreated:      s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"
	s.Client.On("LaunchDBSystem",
		s.Res.AvailabilityDomain, s.Res.CompartmentID, s.Res.Shape, s.Res.SubnetID,
		s.Res.SSHPublicKeys, s.Res.CPUCoreCount, opts,
	).Return(s.Res, nil)

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceTerminated
	s.Client.On("TerminateDBSystem", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	tmpl := `
		resource "baremetal_database_db_system" "t" {
			availability_domain = "{{.AvailabilityDomain}}"
			compartment_id = "{{.CompartmentID}}"
			shape = "{{.Shape}}"
			subnet_id = "{{.SubnetID}}"
			ssh_public_keys = ["{{index .SSHPublicKeys 0}}"]
			cpu_core_count = {{.CPUCoreCount}}
			display_name = "{{.DisplayName}}"
			database_edition = "{{.DatabaseEdition}}"
			db_home {
				database {
					"admin_password" = "{{.DBHome.Database.AdminPassword}}"
					"db_name" = "{{.DBHome.Database.DBName}}"
				}
				db_version = "{{.DBHome.DBVersion}}"
				display_name = "{{.DBHome.DisplayName}}"
			}
			disk_redundancy = "{{.DiskRedundancy}}"
			domain = "{{.Domain}}"
			hostname = "{{.Hostname}}"
		}
	`
	var buf bytes.Buffer
	parsed := template.Must(template.New("config").Parse(tmpl))
	parsed.Execute(&buf, s.Res)
	s.Config = buf.String()
	s.Config += testProviderConfig

	s.ResourceName = "baremetal_database_db_system.t"
}

func (s *DatabaseDBSystemTestSuite) TestCreateDBSystem() {
	s.Client.On("GetDBSystem", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDBSystem", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "db_home.0.db_version", s.Res.DBHome.DBVersion),
					resource.TestCheckResourceAttr(s.ResourceName, "db_home.0.database.0.db_name", s.Res.DBHome.Database.DBName),
				),
			},
		},
	})
}

func (s *DatabaseDBSystemTestSuite) TestTerminateDBSystem() {
	s.Client.On("GetDBSystem", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDBSystem", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "TerminateDBSystem", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestDatabaseDBSystemTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBSystemTestSuite))
}
