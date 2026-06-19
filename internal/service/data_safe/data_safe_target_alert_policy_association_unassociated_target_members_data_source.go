// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"target_alert_policy_association_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_alert_policy_unassociated_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"not_applied_reason": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"target_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListTargetAlertPolicyUnassociatedMembersResponse
}

func (s *DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListTargetAlertPolicyUnassociatedMembersRequest{}

	if targetAlertPolicyAssociationId, ok := s.D.GetOkExists("target_alert_policy_association_id"); ok {
		tmp := targetAlertPolicyAssociationId.(string)
		request.TargetAlertPolicyAssociationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListTargetAlertPolicyUnassociatedMembers(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTargetAlertPolicyUnassociatedMembers(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSource-", DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSource(), s.D))
	resources := []map[string]interface{}{}
	targetAlertPolicyAssociationUnassociatedTargetMember := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TargetAlertPolicyUnassociatedSummaryToMap(item))
	}
	targetAlertPolicyAssociationUnassociatedTargetMember["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSource().Schema["target_alert_policy_unassociated_collection"].Elem.(*schema.Resource).Schema)
		targetAlertPolicyAssociationUnassociatedTargetMember["items"] = items
	}

	resources = append(resources, targetAlertPolicyAssociationUnassociatedTargetMember)
	if err := s.D.Set("target_alert_policy_unassociated_collection", resources); err != nil {
		return err
	}

	return nil
}

func TargetAlertPolicyUnassociatedSummaryToMap(obj oci_data_safe.TargetAlertPolicyUnassociatedSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.NotAppliedReason != nil {
		result["not_applied_reason"] = string(*obj.NotAppliedReason)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDatabaseId != nil {
		result["target_database_id"] = string(*obj.TargetDatabaseId)
	}

	return result
}
