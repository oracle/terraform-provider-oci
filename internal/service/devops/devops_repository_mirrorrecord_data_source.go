// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryMirrorrecordDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryMirrorrecord,
		Schema: map[string]*schema.Schema{
			"mirror_record_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"end_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enqueue_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDevopsRepositoryMirrorrecord(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryMirrorrecordDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryMirrorrecordDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetMirrorRecordResponse
}

func (s *DevopsRepositoryMirrorrecordDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryMirrorrecordDataSourceCrud) Get() error {
	request := oci_devops.GetMirrorRecordRequest{}

	if mirrorRecordType, ok := s.D.GetOkExists("mirror_record_type"); ok {
		request.MirrorRecordType = oci_devops.GetMirrorRecordMirrorRecordTypeEnum(mirrorRecordType.(string))
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetMirrorRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryMirrorrecordDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryMirrorrecordDataSource-", DevopsRepositoryMirrorrecordDataSource(), s.D))

	if s.Res.TimeEnded != nil {
		s.D.Set("end_time", s.Res.TimeEnded.String())
	}

	if s.Res.TimeEnqueued != nil {
		s.D.Set("enqueue_time", s.Res.TimeEnqueued.String())
	}

	s.D.Set("mirror_status", s.Res.MirrorStatus)

	if s.Res.TimeStarted != nil {
		s.D.Set("start_time", s.Res.TimeStarted.String())
	}

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}
