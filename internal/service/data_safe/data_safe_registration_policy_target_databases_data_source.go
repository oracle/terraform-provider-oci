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

func DataSafeRegistrationPolicyTargetDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeRegistrationPolicyTargetDatabasesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"membership_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registration_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registration_policy_target_database_summary_collection": {
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
									"discovered_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"discovered_resource_type": {
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

func readDataSafeRegistrationPolicyTargetDatabasesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeRegistrationPolicyTargetDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeRegistrationPolicyTargetDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListRegistrationPolicyTargetDatabasesResponse
}

func (s *DataSafeRegistrationPolicyTargetDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeRegistrationPolicyTargetDatabasesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListRegistrationPolicyTargetDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if membershipStatus, ok := s.D.GetOkExists("membership_status"); ok {
		request.MembershipStatus = oci_data_safe.ListRegistrationPolicyTargetDatabasesMembershipStatusEnum(membershipStatus.(string))
	}

	if registrationPolicyId, ok := s.D.GetOkExists("registration_policy_id"); ok {
		tmp := registrationPolicyId.(string)
		request.RegistrationPolicyId = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListRegistrationPolicyTargetDatabases(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRegistrationPolicyTargetDatabases(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeRegistrationPolicyTargetDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeRegistrationPolicyTargetDatabasesDataSource-", DataSafeRegistrationPolicyTargetDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	registrationPolicyTargetDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RegistrationPolicyTargetDatabaseSummaryToMap(item))
	}
	registrationPolicyTargetDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeRegistrationPolicyTargetDatabasesDataSource().Schema["registration_policy_target_database_summary_collection"].Elem.(*schema.Resource).Schema)
		registrationPolicyTargetDatabase["items"] = items
	}

	resources = append(resources, registrationPolicyTargetDatabase)
	if err := s.D.Set("registration_policy_target_database_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func RegistrationPolicyTargetDatabaseSummaryToMap(obj oci_data_safe.RegistrationPolicyTargetDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DiscoveredResourceId != nil {
		result["discovered_resource_id"] = string(*obj.DiscoveredResourceId)
	}

	result["discovered_resource_type"] = string(obj.DiscoveredResourceType)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDatabaseId != nil {
		result["target_database_id"] = string(*obj.TargetDatabaseId)
	}

	return result
}
