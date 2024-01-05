// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
)

func AiAnomalyDetectionProjectDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["project_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiAnomalyDetectionProjectResource(), fieldMap, readSingularAiAnomalyDetectionProject)
}

func readSingularAiAnomalyDetectionProject(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionProjectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionProjectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.GetProjectResponse
}

func (s *AiAnomalyDetectionProjectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionProjectDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetProjectRequest{}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.GetProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiAnomalyDetectionProjectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	// if s.Res.SystemTags != nil {
	// 	s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	// }

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
