// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardDataMaskRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_mask_rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardDataMaskRuleResource(), fieldMap, readSingularCloudGuardDataMaskRule)
}

func readSingularCloudGuardDataMaskRule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataMaskRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardDataMaskRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetDataMaskRuleResponse
}

func (s *CloudGuardDataMaskRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardDataMaskRuleDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetDataMaskRuleRequest{}

	if dataMaskRuleId, ok := s.D.GetOkExists("data_mask_rule_id"); ok {
		tmp := dataMaskRuleId.(string)
		request.DataMaskRuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetDataMaskRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardDataMaskRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("data_mask_categories", s.Res.DataMaskCategories)

	s.D.Set("data_mask_rule_status", s.Res.DataMaskRuleStatus)

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

	if s.Res.IamGroupId != nil {
		s.D.Set("iam_group_id", *s.Res.IamGroupId)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetSelected != nil {
		targetSelectedArray := []interface{}{}
		if targetSelectedMap := TargetSelectedToMap(&s.Res.TargetSelected); targetSelectedMap != nil {
			targetSelectedArray = append(targetSelectedArray, targetSelectedMap)
		}
		s.D.Set("target_selected", targetSelectedArray)
	} else {
		s.D.Set("target_selected", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
