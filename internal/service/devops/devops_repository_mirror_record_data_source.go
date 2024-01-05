// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryMirrorRecordDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryMirrorRecord,
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
			"mirror_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_enqueued": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
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

func readSingularDevopsRepositoryMirrorRecord(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryMirrorRecordDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryMirrorRecordDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetMirrorRecordResponse
}

func (s *DevopsRepositoryMirrorRecordDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryMirrorRecordDataSourceCrud) Get() error {
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

func (s *DevopsRepositoryMirrorRecordDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryMirrorRecordDataSource-", DevopsRepositoryMirrorRecordDataSource(), s.D))

	s.D.Set("mirror_status", s.Res.MirrorStatus)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeEnqueued != nil {
		s.D.Set("time_enqueued", s.Res.TimeEnqueued.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}

func RepositoryMirrorRecordSummaryToMap(obj oci_devops.RepositoryMirrorRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["mirror_status"] = string(obj.MirrorStatus)

	if obj.TimeCompleted != nil {
		result["time_completed"] = obj.TimeCompleted.String()
	}

	if obj.TimeEnqueued != nil {
		result["time_enqueued"] = obj.TimeEnqueued.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}
