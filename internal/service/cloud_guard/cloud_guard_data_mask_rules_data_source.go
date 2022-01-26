// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardDataMaskRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardDataMaskRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_mask_rule_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iam_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_mask_rule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudGuardDataMaskRuleResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardDataMaskRules(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataMaskRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardDataMaskRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListDataMaskRulesResponse
}

func (s *CloudGuardDataMaskRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardDataMaskRulesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListDataMaskRulesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_cloud_guard.ListDataMaskRulesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataMaskRuleStatus, ok := s.D.GetOkExists("data_mask_rule_status"); ok {
		request.DataMaskRuleStatus = oci_cloud_guard.ListDataMaskRulesDataMaskRuleStatusEnum(dataMaskRuleStatus.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if iamGroupId, ok := s.D.GetOkExists("iam_group_id"); ok {
		tmp := iamGroupId.(string)
		request.IamGroupId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.ListDataMaskRulesLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		tmp := targetType.(string)
		request.TargetType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListDataMaskRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataMaskRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardDataMaskRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardDataMaskRulesDataSource-", CloudGuardDataMaskRulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	dataMaskRule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataMaskRuleSummaryToMap(item))
	}
	dataMaskRule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardDataMaskRulesDataSource().Schema["data_mask_rule_collection"].Elem.(*schema.Resource).Schema)
		dataMaskRule["items"] = items
	}

	resources = append(resources, dataMaskRule)
	if err := s.D.Set("data_mask_rule_collection", resources); err != nil {
		return err
	}

	return nil
}
