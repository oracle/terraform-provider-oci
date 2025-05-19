// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryProtectedDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRecoveryProtectedDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protection_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recovery_service_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protected_database_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(RecoveryProtectedDatabaseResource()),
						},
					},
				},
			},
		},
	}
}

func readRecoveryProtectedDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

type RecoveryProtectedDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_recovery.DatabaseRecoveryClient
	Res    *oci_recovery.ListProtectedDatabasesResponse
}

func (s *RecoveryProtectedDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RecoveryProtectedDatabasesDataSourceCrud) Get() error {
	request := oci_recovery.ListProtectedDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if protectionPolicyId, ok := s.D.GetOkExists("protection_policy_id"); ok {
		tmp := protectionPolicyId.(string)
		request.ProtectionPolicyId = &tmp
	}

	if recoveryServiceSubnetId, ok := s.D.GetOkExists("recovery_service_subnet_id"); ok {
		tmp := recoveryServiceSubnetId.(string)
		request.RecoveryServiceSubnetId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_recovery.ListProtectedDatabasesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "recovery")

	response, err := s.Client.ListProtectedDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProtectedDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RecoveryProtectedDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RecoveryProtectedDatabasesDataSource-", RecoveryProtectedDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	protectedDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProtectedDatabaseSummaryToMap(item))
	}
	protectedDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RecoveryProtectedDatabasesDataSource().Schema["protected_database_collection"].Elem.(*schema.Resource).Schema)
		protectedDatabase["items"] = items
	}

	resources = append(resources, protectedDatabase)
	if err := s.D.Set("protected_database_collection", resources); err != nil {
		return err
	}

	return nil
}
