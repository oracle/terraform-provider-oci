// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type BucketResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Bucket
}

func (s *BucketResourceCrud) ID() string {
	return string(s.Res.Namespace) + "/" + s.Res.Name
}

func (s *BucketResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("name", s.Res.Name)
	s.D.Set("namespace", s.Res.Namespace)
	s.D.Set("metadata", s.Res.Metadata)
	s.D.Set("created_by", s.Res.CreatedBy)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *BucketResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	opts := &baremetal.CreateBucketOptions{}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}
	s.Res, e = s.Client.CreateBucket(compartmentID, name, baremetal.Namespace(namespace), opts)
	return
}

func (s *BucketResourceCrud) Get() (e error) {
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	s.Res, e = s.Client.GetBucket(name, baremetal.Namespace(namespace))
	return
}

func (s *BucketResourceCrud) Update() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	opts := &baremetal.UpdateBucketOptions{}
	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}

	s.Res, e = s.Client.UpdateBucket(compartmentID, name, baremetal.Namespace(namespace), opts)
	return
}

func (s *BucketResourceCrud) Delete() (e error) {
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	return s.Client.DeleteBucket(name, baremetal.Namespace(namespace), nil)
}
