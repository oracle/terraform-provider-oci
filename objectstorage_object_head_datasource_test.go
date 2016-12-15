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

type DatasourceObjectstorageObjectHeadTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.HeadObject
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) SetupTest() {
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
		data "baremetal_objectstorage_objecthead" "t" {
			namespace = "namespaceID"
			bucket = "bucketID"
			object = "object"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "data.baremetal_objectstorage_objecthead.t"
	s.Res = &baremetal.HeadObject{
		Namespace: baremetal.Namespace("namespaceID"),
		Bucket: "bucketID",
		ID: "object",
	}
	metadata := map[string]string{"foo": "bar"}
	s.Res.Metadata = metadata
	s.Res.ContentLength = 123
	s.Res.ContentType = "type"
}

func (s *DatasourceObjectstorageObjectHeadTestSuite) TestObjectstorageHeadObject() {
	opts := &baremetal.HeadObjectOptions{}
	s.Client.On("HeadObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, opts).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "object", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Res.Bucket),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", string(s.Res.Namespace)),
					resource.TestCheckResourceAttr(s.ResourceName, "metadata.foo", "bar"),
					//resource.TestCheckResourceAttr(s.ResourceName, "content-length", s.Res.ContentLength),
					resource.TestCheckResourceAttr(s.ResourceName, "content-type", s.Res.ContentType),
				),
			},
		},
	})
	s.Client.AssertCalled(s.T(), "HeadObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, opts)
}

func TestDatasourceobjectstorageObjectHeadTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceObjectstorageObjectHeadTestSuite))
}
