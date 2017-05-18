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
)

type ResourceObjectstorageObjectTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Object
}

func (s *ResourceObjectstorageObjectTestSuite) SetupTest() {
	s.Client = GetTestProvider()

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
		resource "baremetal_objectstorage_object" "t" {
			namespace = "namespaceID"
			bucket = "bucketID"
			object = "objectID"
			content = "bodyContent"
			metadata = {
				"foo" = "bar"
			}
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_objectstorage_object.t"
	metadata := map[string]string{
		"foo": "bar",
	}
	s.Res = &baremetal.Object{
		Body: []byte("bodyContent"),
	}
	s.Res.Namespace = baremetal.Namespace("namespaceID")
	s.Res.Metadata = metadata
	s.Res.ID = "objectID"
	s.Res.Bucket = "bucketID"
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	opts := &baremetal.PutObjectOptions{}
	opts.Metadata = metadata
	s.Client.On("PutObject", s.Res.Namespace, s.Res.Bucket,
		s.Res.ID, s.Res.Body, opts).Return(s.Res, nil).Once()
	s.Client.On("DeleteObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.DeleteObjectOptions{}).Return(&baremetal.DeleteObject{}, nil)
}

func (s *ResourceObjectstorageObjectTestSuite) TestCreateResourceObjectstorageObject() {
	s.Client.On("GetObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.GetObjectOptions{}).Return(s.Res, nil).Times(2)
	s.Client.On("GetObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.GetObjectOptions{}).Return(nil, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Res.Bucket),
					resource.TestCheckResourceAttr(s.ResourceName, "object", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", string(s.Res.Namespace)),
					resource.TestCheckResourceAttr(s.ResourceName, "content", string(s.Res.Body)),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageObjectTestSuite) TestUpdateResourceObjectstorageObject() {
	s.Client.On("GetObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.GetObjectOptions{}).Return(s.Res, nil).Times(2)

	config := `
		resource "baremetal_objectstorage_object" "t" {
			object = "objectID"
			bucket = "bucketID"
			namespace = "namespaceID"
			content = "bodyContent2"
			metadata = {
				"foo" = "bar"
			}
		}
	`
	config += testProviderConfig()
	metadata := map[string]string{
		"foo": "bar",
	}

	res := &baremetal.Object{
		Body: []byte("bodyContent2"),
	}
	res.Namespace = baremetal.Namespace("namespaceID")
	res.Metadata = metadata
	res.ID = "objectID"
	res.Bucket = "bucketID"
	res.ETag = "etag2"
	res.RequestID = "opcrequestid2"

	opts := &baremetal.PutObjectOptions{}
	opts.Metadata = metadata

	s.Client.On("PutObject", res.Namespace, res.Bucket,
		res.ID, res.Body, opts).Return(res, nil).Once()

	s.Client.On("GetObject", res.Namespace, res.Bucket, res.ID, &baremetal.GetObjectOptions{}).Return(res, nil)
	s.Client.On("DeleteObject", res.Namespace, res.Bucket, res.ID, &baremetal.DeleteObjectOptions{}).Return(&baremetal.DeleteObject{}, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "content", string(res.Body)),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageObjectTestSuite) TestDeleteResourceObjectstorageObject() {
	s.Client.On("GetObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.GetObjectOptions{}).Return(s.Res, nil).Times(2)
	s.Client.On("GetObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.GetObjectOptions{}).Return(nil, nil)
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})
	s.Client.AssertCalled(s.T(), "DeleteObject", s.Res.Namespace, s.Res.Bucket, s.Res.ID, &baremetal.DeleteObjectOptions{})
}

func TestResourceobjectstorageObjectTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageObjectTestSuite))
}
