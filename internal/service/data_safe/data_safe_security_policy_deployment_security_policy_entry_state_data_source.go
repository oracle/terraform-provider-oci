// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSecurityPolicyDeploymentSecurityPolicyEntryState,
		Schema: map[string]*schema.Schema{
			"security_policy_deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_policy_entry_state_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"security_policy_entry_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeSecurityPolicyDeploymentSecurityPolicyEntryState(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSecurityPolicyEntryStateResponse
}

func (s *DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceCrud) Get() error {
	request := oci_data_safe.GetSecurityPolicyEntryStateRequest{}

	if securityPolicyDeploymentId, ok := s.D.GetOkExists("security_policy_deployment_id"); ok {
		tmp := securityPolicyDeploymentId.(string)
		request.SecurityPolicyDeploymentId = &tmp
	}

	if securityPolicyEntryStateId, ok := s.D.GetOkExists("security_policy_entry_state_id"); ok {
		tmp := securityPolicyEntryStateId.(string)
		request.SecurityPolicyEntryStateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSecurityPolicyEntryState(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("deployment_status", s.Res.DeploymentStatus)

	if s.Res.EntryDetails != nil {
		entryDetailsArray := []interface{}{}
		if entryDetailsMap := EntryDetailsToMap(&s.Res.EntryDetails); entryDetailsMap != nil {
			entryDetailsArray = append(entryDetailsArray, entryDetailsMap)
		}
		s.D.Set("entry_details", entryDetailsArray)
	} else {
		s.D.Set("entry_details", nil)
	}

	if s.Res.SecurityPolicyEntryId != nil {
		s.D.Set("security_policy_entry_id", *s.Res.SecurityPolicyEntryId)
	}

	return nil
}
