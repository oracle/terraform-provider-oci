// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

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

type ObjectstorageBucketSummaryTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	TimeCreated  time.Time
}

func (s *ObjectstorageBucketSummaryTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_objectstorage_bucket_summaries" "t" {
      compartment_id = "compartmentid"
      namespace = "namespace"
      limit = 2
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_objectstorage_bucket_summaries.t"
	s.TimeCreated = time.Now()
}

func (s *ObjectstorageBucketSummaryTestSuite) TestReadBucketSummaries() {
	namespace := baremetal.Namespace("namespace")

	opts := &baremetal.ListBucketsOptions{}
	opts.Page = "page"
	opts.Limit = 2

	res := &baremetal.ListBuckets{}
	res.NextPage = "nextpage"
	res.BucketSummaries = []baremetal.BucketSummary{
		{
			Namespace:     "namespace",
			Name:          "name0",
			CompartmentID: "compartmentID",
			CreatedBy:     "created_by",
			TimeCreated:   s.TimeCreated,
			ETag:          "etag",
		},
		{
			Namespace:     "namespace",
			Name:          "name1",
			CompartmentID: "compartmentID",
			CreatedBy:     "created_by",
			TimeCreated:   s.TimeCreated,
			ETag:          "etag",
		},
	}

	s.Client.On(
		"ListBuckets", "compartmentid", namespace, opts,
	).Return(res, nil)

	opts2 := &baremetal.ListBucketsOptions{}
	opts2.Page = "nextpage"
	opts2.Limit = 2
	s.Client.On(
		"ListBuckets", "compartmentid", namespace, opts2,
	).Return(
		&baremetal.ListBuckets{
			BucketSummaries: []baremetal.BucketSummary{
				{
					Namespace:     "namespace",
					Name:          "name2",
					CompartmentID: "compartmentID",
					CreatedBy:     "created_by",
					TimeCreated:   s.TimeCreated,
					ETag:          "etag",
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.0.name", "name0"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.2.name", "name2"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket_summaries.#", "3"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(
		s.T(), "ListBuckets", "compartmentid", namespace, opts2,
	)

}

func TestObjectstorageBucketSummaryTestSuite(t *testing.T) {
	suite.Run(t, new(ObjectstorageBucketSummaryTestSuite))
}
