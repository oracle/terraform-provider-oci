// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v56/apmsynthetics"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsScriptDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["script_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmSyntheticsScriptResource(), fieldMap, readSingularApmSyntheticsScript)
}

func readSingularApmSyntheticsScript(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsScriptDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsScriptDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.GetScriptResponse
}

func (s *ApmSyntheticsScriptDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsScriptDataSourceCrud) Get() error {
	request := oci_apm_synthetics.GetScriptRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if scriptId, ok := s.D.GetOkExists("script_id"); ok {
		tmp := scriptId.(string)

		scriptId, apmDomainId, err := parseScriptCompositeId(tmp)
		if err == nil {
			request.ScriptId = &scriptId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.GetScript(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmSyntheticsScriptDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Content != nil {
		s.D.Set("content", *s.Res.Content)
	}

	if s.Res.ContentFileName != nil {
		s.D.Set("content_file_name", *s.Res.ContentFileName)
	}

	if s.Res.ContentSizeInBytes != nil {
		s.D.Set("content_size_in_bytes", *s.Res.ContentSizeInBytes)
	}

	s.D.Set("content_type", s.Res.ContentType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MonitorStatusCountMap != nil {
		s.D.Set("monitor_status_count_map", []interface{}{MonitorStatusCountMapToMap(s.Res.MonitorStatusCountMap)})
	} else {
		s.D.Set("monitor_status_count_map", nil)
	}

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, ScriptParameterInfoToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeUploaded != nil {
		s.D.Set("time_uploaded", s.Res.TimeUploaded.String())
	}

	return nil
}
