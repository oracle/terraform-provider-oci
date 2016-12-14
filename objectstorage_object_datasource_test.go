package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.ListObjects
}

func (s *DatasourceObjectstorageObjectTestSuite) SetupTest() {
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

	s.Config = `
		data "baremetal_objectstorage_objectlist" "t" {
			namespace = "namespaceID"
			bucket = "bucketID"
			prefix = "testprefix"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_objectstorage_objectlist.t"
	s.Res = &baremetal.ListObjects{

		Objects: []baremetal.ObjectSummary{
			{
				Name: "testprefix-1",
			},
			{
				Name: "testprefix-2",
			},
		},
	}
}

func (s *DatasourceObjectstorageObjectTestSuite) TestObjectstorageListObjects() {
	opts := &baremetal.ListObjectsOptions{Prefix: "testprefix"}
	s.Client.On("ListObjects", baremetal.Namespace("namespaceID"), "bucketID", opts).Return(s.Res, nil).Once()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "prefix", "testprefix"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", "bucketID"),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", "namespaceID"),
				),
			},
		},
	})
	s.Client.AssertCalled(s.T(), "ListObjects", baremetal.Namespace("namespaceID"), "bucketID", opts)
}


func TestDatasourceobjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
