// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"log"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/pkg/errors"
	"time"
)

type PreauthenticatedRequestResourceCrud struct {
	crud.BaseCrud
	Namespace  string
	BucketName string
	Res        *baremetal.PreauthenticatedRequest
	Summary    *baremetal.PreauthenticatedRequestSummary
}

func PreauthenticatedRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPreauthenticatedRequest,
		Read:     readPreauthenticatedRequest,
		Update:   updatePreauthenticatedRequest,
		Delete:   deletePreauthenticatedRequest,
		Schema:   preauthenticatedRequestSchema,
	}
}

func createPreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readPreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updatePreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deletePreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}

func (s *PreauthenticatedRequestResourceCrud) ID() string {
	return s.Res.ID
}

func (s *PreauthenticatedRequestResourceCrud) SetData() {
	log.Printf("=======================\n%v\n===================", s.Res)
	s.D.Set("namespace", s.Namespace)
	s.D.Set("bucket", s.BucketName)
	s.D.Set("object", s.Res.ObjectName)
	s.D.Set("time_expires", s.Res.TimeExpires.Format(time.RFC3339))
	s.D.Set("access_type", s.Res.AccessType)
}

func (s *PreauthenticatedRequestResourceCrud) Create() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	accessType := s.D.Get("access_type").(string)
	t, _ := time.Parse(time.RFC3339, s.D.Get("time_expires").(string))
	details := &baremetal.CreatePreauthenticatedRequestDetails{
		ObjectName:  object,
		TimeExpires: baremetal.Time{Time: t},
		AccessType:  baremetal.PARAccessType(accessType),
	}

	s.Res, e = s.Client.CreatePreauthenticatedRequest(baremetal.Namespace(namespace), bucket, details)
	return
}

func (s *PreauthenticatedRequestResourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	parId := s.D.Get("id").(string)
	s.Summary, e = s.Client.GetPreauthenticatedRequest(baremetal.Namespace(namespace), bucket, parId,
		&baremetal.ClientRequestOptions{})

	return
}

func (s *PreauthenticatedRequestResourceCrud) Update() (e error) {
	e = errors.New("Update operation is not supported for PreauthenticatedRequest. Create a new " +
		"PreauthenticatedRequest if you need to make modifications ")
	return
}

func (s *PreauthenticatedRequestResourceCrud) Delete() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	parId := s.D.Get("id").(string)
	opts := &baremetal.ClientRequestOptions{}

	e = s.Client.DeletePreauthenticatedRequest(baremetal.Namespace(namespace), bucket, parId, opts)
	return
}
