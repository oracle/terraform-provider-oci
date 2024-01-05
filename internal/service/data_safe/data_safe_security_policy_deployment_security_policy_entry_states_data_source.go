// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityPolicyDeploymentSecurityPolicyEntryStates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"deployment_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_policy_deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_policy_entry_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_policy_entry_state_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"deployment_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entry_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"entry_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_generated": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_status_updated": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_policy_deployment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_policy_entry_id": {
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

func readDataSafeSecurityPolicyDeploymentSecurityPolicyEntryStates(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSecurityPolicyEntryStatesResponse
}

func (s *DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSourceCrud) Get() error {
	request := oci_data_safe.ListSecurityPolicyEntryStatesRequest{}

	if deploymentStatus, ok := s.D.GetOkExists("deployment_status"); ok {
		request.DeploymentStatus = oci_data_safe.ListSecurityPolicyEntryStatesDeploymentStatusEnum(deploymentStatus.(string))
	}

	if securityPolicyDeploymentId, ok := s.D.GetOkExists("security_policy_deployment_id"); ok {
		tmp := securityPolicyDeploymentId.(string)
		request.SecurityPolicyDeploymentId = &tmp
	}

	if securityPolicyEntryId, ok := s.D.GetOkExists("security_policy_entry_id"); ok {
		tmp := securityPolicyEntryId.(string)
		request.SecurityPolicyEntryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSecurityPolicyEntryStates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityPolicyEntryStates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSource-", DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityPolicyDeploymentSecurityPolicyEntryState := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SecurityPolicyEntryStateSummaryToMap(item))
	}
	securityPolicyDeploymentSecurityPolicyEntryState["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSource().Schema["security_policy_entry_state_collection"].Elem.(*schema.Resource).Schema)
		securityPolicyDeploymentSecurityPolicyEntryState["items"] = items
	}

	resources = append(resources, securityPolicyDeploymentSecurityPolicyEntryState)
	if err := s.D.Set("security_policy_entry_state_collection", resources); err != nil {
		return err
	}

	return nil
}

func EntryDetailsToMap(obj *oci_data_safe.EntryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.FirewallPolicyEntryDetails:
		result["entry_type"] = "FIREWALL_POLICY"

		if v.TimeGenerated != nil {
			result["time_generated"] = v.TimeGenerated.Format(time.RFC3339Nano)
		}

		if v.TimeStatusUpdated != nil {
			result["time_status_updated"] = v.TimeStatusUpdated.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'entry_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SecurityPolicyEntryStateSummaryToMap(obj oci_data_safe.SecurityPolicyEntryStateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["deployment_status"] = string(obj.DeploymentStatus)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SecurityPolicyDeploymentId != nil {
		result["security_policy_deployment_id"] = string(*obj.SecurityPolicyDeploymentId)
	}

	if obj.SecurityPolicyEntryId != nil {
		result["security_policy_entry_id"] = string(*obj.SecurityPolicyEntryId)
	}

	return result
}
