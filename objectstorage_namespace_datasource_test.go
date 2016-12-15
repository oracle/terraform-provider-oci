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

type DatasourceObjectstorageNamespaceTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Namespace
}

func (s *DatasourceObjectstorageNamespaceTestSuite) SetupTest() {
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
		data "baremetal_objectstorage_namespace" "t" {}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_objectstorage_namespace.t"
	namespace := baremetal.Namespace("namespaceID")
	s.Res = &namespace
}

func (s *DatasourceObjectstorageNamespaceTestSuite) TestObjectstorageNamespace() {
	s.Client.On("GetNamespace").Return(s.Res, nil).Once()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", "namespaceID"),
				),
			},
		},
	})
	s.Client.AssertCalled(s.T(), "GetNamespace")
}

func TestDatasourceobjectstorageNamespaceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
