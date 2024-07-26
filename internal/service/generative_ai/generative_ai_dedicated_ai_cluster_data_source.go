// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiDedicatedAiClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dedicated_ai_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiDedicatedAiClusterResource(), fieldMap, readSingularGenerativeAiDedicatedAiCluster)
}

func readSingularGenerativeAiDedicatedAiCluster(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiDedicatedAiClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiDedicatedAiClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetDedicatedAiClusterResponse
}

func (s *GenerativeAiDedicatedAiClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiDedicatedAiClusterDataSourceCrud) Get() error {
	request := oci_generative_ai.GetDedicatedAiClusterRequest{}

	if dedicatedAiClusterId, ok := s.D.GetOkExists("dedicated_ai_cluster_id"); ok {
		tmp := dedicatedAiClusterId.(string)
		request.DedicatedAiClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetDedicatedAiCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiDedicatedAiClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Capacity != nil {
		capacityArray := []interface{}{}
		if capacityMap := DedicatedAiClusterCapacityToMap(&s.Res.Capacity); capacityMap != nil {
			capacityArray = append(capacityArray, capacityMap)
		}
		s.D.Set("capacity", capacityArray)
	} else {
		s.D.Set("capacity", nil)
	}

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UnitCount != nil {
		s.D.Set("unit_count", *s.Res.UnitCount)
	}

	s.D.Set("unit_shape", s.Res.UnitShape)

	return nil
}
