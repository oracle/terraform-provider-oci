// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	//"context"
	//"encoding/json"
	//"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowJobFactDataSource() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This data source has been deprecated and is no longer supported.",
		Read:               readSingularMediaServicesMediaWorkflowJobFact,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"media_workflow_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"detail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularMediaServicesMediaWorkflowJobFact(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobFactDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaWorkflowJobFactDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	//Res    *oci_media_services.GetMediaWorkflowJobFactResponse
}

func (s *MediaServicesMediaWorkflowJobFactDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaWorkflowJobFactDataSourceCrud) Get() error {
	//request := oci_media_services.GetMediaWorkflowJobFactRequest{}
	//
	//if key, ok := s.D.GetOkExists("key"); ok {
	//	intValue := 0
	//	_, _ = fmt.Sscan(key.(string), &intValue)
	//	request.Key = &intValue
	//}
	//
	//if mediaWorkflowJobId, ok := s.D.GetOkExists("media_workflow_job_id"); ok {
	//	tmp := mediaWorkflowJobId.(string)
	//	request.MediaWorkflowJobId = &tmp
	//}
	//
	//request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")
	//request.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 1
	//
	//response, err := s.Client.GetMediaWorkflowJobFact(context.Background(), request)
	//if err != nil && response.RawResponse.StatusCode != 404 {
	//	return err
	//}
	//
	//s.Res = &response
	return nil
}

func (s *MediaServicesMediaWorkflowJobFactDataSourceCrud) SetData() error {
	//if s.Res == nil {
	//	return nil
	//}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesMediaWorkflowJobFactDataSource-", MediaServicesMediaWorkflowJobFactDataSource(), s.D))

	//if s.Res.Detail != nil {
	//	jsonStr, err := json.Marshal(s.Res.Detail)
	//	if err == nil {
	//		s.D.Set("detail", string(jsonStr))
	//	}
	//}
	//
	//if s.Res.Name != nil {
	//	s.D.Set("name", *s.Res.Name)
	//}
	//
	//if s.Res.Type != nil {
	//	s.D.Set("type", *s.Res.Type)
	//}

	return nil
}
