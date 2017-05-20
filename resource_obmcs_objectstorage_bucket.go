// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"time"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func BucketResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createBucket,
		Read:     readBucket,
		Update:   updateBucket,
		Delete:   deleteBucket,
		Schema:   bucketSchema,
	}
}

func createBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deleteBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}

type BucketResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Bucket
}

func (s *BucketResourceCrud) ID() string {
	return string(s.Res.Namespace) + "/" + s.Res.Name
}

func (s *BucketResourceCrud) ExtraWaitPostCreateDelete() time.Duration {
	return time.Duration(10 * time.Second)
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
		metadata := resourceObjectStorageMapToMetadata(rawMetadata.(map[string]interface{}))
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
		metadata := resourceObjectStorageMapToMetadata(rawMetadata.(map[string]interface{}))
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
