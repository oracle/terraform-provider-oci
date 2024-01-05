// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
)

func ApmConfigConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmConfigConfigResource(), fieldMap, readSingularApmConfigConfig)
}

func readSingularApmConfigConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ApmConfigConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.ReadResource(sync)
}

type ApmConfigConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_config.ConfigClient
	Res    *oci_apm_config.GetConfigResponse
}

func (s *ApmConfigConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmConfigConfigDataSourceCrud) Get() error {
	request := oci_apm_config.GetConfigRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if configId, ok := s.D.GetOkExists("config_id"); ok {
		tmp := configId.(string)

		configId, apmDomainId, err := parseConfigCompositeId(tmp)
		if err == nil {
			request.ConfigId = &configId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", tmp)
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_config")

	response, err := s.Client.GetConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmConfigConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())

	switch v := (s.Res.Config).(type) {
	case oci_apm_config.ApdexRules:
		s.D.Set("config_type", "APDEX")

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		rules := []interface{}{}
		for _, item := range v.Rules {
			rules = append(rules, ApdexToMap(item))
		}
		s.D.Set("rules", rules)

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.MetricGroup:
		s.D.Set("config_type", "METRIC_GROUP")

		dimensions := []interface{}{}
		for _, item := range v.Dimensions {
			dimensions = append(dimensions, DimensionToMap(item))
		}
		s.D.Set("dimensions", dimensions)

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.FilterId != nil {
			s.D.Set("filter_id", *v.FilterId)
		}

		metrics := []interface{}{}
		for _, item := range v.Metrics {
			metrics = append(metrics, MetricToMap(item))
		}
		s.D.Set("metrics", metrics)

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.Options:
		s.D.Set("config_type", "OPTIONS")

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Options != nil {
			s.D.Set("options", optionsToMap(v.Options))
		} else {
			s.D.Set("options", nil)
		}

		if v.Group != nil {
			s.D.Set("group", *v.Group)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	case oci_apm_config.SpanFilter:
		s.D.Set("config_type", "SPAN_FILTER")

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.FilterText != nil {
			s.D.Set("filter_text", *v.FilterText)
		}

		if v.CreatedBy != nil {
			s.D.Set("created_by", *v.CreatedBy)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Etag != nil {
			s.D.Set("etag", *v.Etag)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		inUseBy := []interface{}{}
		for _, item := range v.InUseBy {
			inUseBy = append(inUseBy, SpanFilterReferenceToMap(item))
		}
		s.D.Set("in_use_by", inUseBy)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.UpdatedBy != nil {
			s.D.Set("updated_by", *v.UpdatedBy)
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", s.Res.Config)
		return nil
	}

	return nil
}
