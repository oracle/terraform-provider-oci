// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"time"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PreauthenticatedRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPreauthenticatedRequest,
		Read:     readPreauthenticatedRequest,
		Delete:   deletePreauthenticatedRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"access_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeObjectread),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeObjectwrite),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeObjectreadwrite),
					string(oci_object_storage.PreauthenticatedRequestSummaryAccessTypeAnyobjectwrite),
				}, true),
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_expires": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createPreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.CreateResource(d, sync)
}

func readPreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
}

func deletePreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &PreauthenticatedRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type PreauthenticatedRequestResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.PreauthenticatedRequest
	DisableNotFoundRetries bool
}

func (s *PreauthenticatedRequestResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PreauthenticatedRequestResourceCrud) Create() error {
	request := oci_object_storage.CreatePreauthenticatedRequestRequest{}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		request.AccessType = oci_object_storage.CreatePreauthenticatedRequestDetailsAccessTypeEnum(accessType.(string))
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if timeExpires, ok := s.D.GetOkExists("time_expires"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpires.(string))
		if err != nil {
			return err
		}

		request.TimeExpires = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.CreatePreauthenticatedRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PreauthenticatedRequest
	return nil
}

func (s *PreauthenticatedRequestResourceCrud) Get() error {
	request := oci_object_storage.GetPreauthenticatedRequestRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	tmp := s.D.Id()
	request.ParId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetPreauthenticatedRequest(context.Background(), request)
	if err != nil {
		return err
	}

	// Some contortions follow since GETs actually return a PreauthenticatedRequestSummary, but s.Res is typed as a
	// PreauthenticatedRequest

	s.Res = &oci_object_storage.PreauthenticatedRequest{
		Id:          response.PreauthenticatedRequestSummary.Id,
		AccessType:  oci_object_storage.PreauthenticatedRequestAccessTypeEnum(string(response.PreauthenticatedRequestSummary.AccessType)),
		Name:        response.PreauthenticatedRequestSummary.Name,
		ObjectName:  response.PreauthenticatedRequestSummary.ObjectName,
		TimeCreated: response.PreauthenticatedRequestSummary.TimeCreated,
		TimeExpires: response.PreauthenticatedRequestSummary.TimeExpires,
	}

	return nil
}

func (s *PreauthenticatedRequestResourceCrud) Delete() error {
	request := oci_object_storage.DeletePreauthenticatedRequestRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	tmp := s.D.Id()
	request.ParId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeletePreauthenticatedRequest(context.Background(), request)
	return err
}

func (s *PreauthenticatedRequestResourceCrud) SetData() {
	s.D.Set("access_type", s.Res.AccessType)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires)
	}

	if s.Res.AccessUri != nil {
		s.D.Set("access_uri", *s.Res.AccessUri)
	}
}
