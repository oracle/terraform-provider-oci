// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/pkg/errors"

	"github.com/oracle/terraform-provider-oci/crud"
)

type PreauthenticatedRequestResourceCrud struct {
	crud.BaseCrud
	Id          string
	Name        string
	Namespace   string
	BucketName  string
	ObjectName  string
	AccessURI   string
	AccessType  baremetal.PARAccessType
	TimeExpires baremetal.Time
	TimeCreated baremetal.Time
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
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"access_type": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(baremetal.PARAnyObjectWrite),
					string(baremetal.PARObjectRead),
					string(baremetal.PARObjectReadWrite),
					string(baremetal.ObjectRead)}, true),
			},
			"access_uri": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"time_expires": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
		},
	}
}

func createPreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readPreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updatePreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(d, sync)
}

func deletePreauthenticatedRequest(d *schema.ResourceData, m interface{}) (e error) {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

func (s *PreauthenticatedRequestResourceCrud) ID() string {
	return s.Id
}

func (s *PreauthenticatedRequestResourceCrud) SetData() {
	s.D.Set("namespace", s.Namespace)
	s.D.Set("bucket", s.BucketName)
	s.D.Set("object", s.ObjectName)
	s.D.Set("time_expires", s.TimeExpires.Format(time.RFC3339))
	s.D.Set("access_type", s.AccessType)
	s.D.Set("access_uri", s.AccessURI)
	s.D.Set("id", s.ID)
	s.D.Set("name", s.Name)
}

func (s *PreauthenticatedRequestResourceCrud) Create() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	name := s.D.Get("name").(string)
	accessType := s.D.Get("access_type").(string)
	t, _ := time.Parse(time.RFC3339, s.D.Get("time_expires").(string))
	details := &baremetal.CreatePreauthenticatedRequestDetails{
		Name:        name,
		TimeExpires: baremetal.Time{Time: t},
		AccessType:  baremetal.PARAccessType(accessType),
	}

	object := s.D.Get("object").(string)
	if object, ok := s.D.GetOk("object"); ok {
		details.ObjectName = object.(string)
	}

	var res *baremetal.PreauthenticatedRequest
	res, e = s.Client.CreatePreauthenticatedRequest(baremetal.Namespace(namespace), bucket, details)

	if e != nil {
		return
	}

	s.AccessURI = res.AccessURI
	s.Id = res.ID
	s.TimeCreated = res.TimeCreated
	s.TimeExpires = res.TimeExpires
	s.AccessType = res.AccessType
	s.Namespace = namespace
	s.BucketName = bucket
	s.ObjectName = object
	s.Name = res.Name
	return
}

func (s *PreauthenticatedRequestResourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	parId := s.D.Get("id").(string)

	var res *baremetal.PreauthenticatedRequestSummary
	res, e = s.Client.GetPreauthenticatedRequest(baremetal.Namespace(namespace), bucket, parId,
		&baremetal.ClientRequestOptions{})

	if e != nil {
		return
	}

	s.Id = res.ID
	//Access URI is only accessible on create
	s.AccessURI = ""
	s.TimeCreated = res.TimeCreated
	s.TimeExpires = res.TimeExpires
	s.AccessType = res.AccessType
	s.Namespace = namespace
	s.BucketName = bucket
	s.ObjectName = res.ObjectName
	s.Name = res.Name

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

	return s.Client.DeletePreauthenticatedRequest(baremetal.Namespace(namespace), bucket, parId, opts)
}
