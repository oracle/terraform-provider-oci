// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlReplicasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlReplicas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_up_to_date": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"replica_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replicas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(MysqlReplicaResource()),
			},
		},
	}
}

func readMysqlReplicas(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.ReadResource(sync)
}

type MysqlReplicasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.ReplicasClient
	Res    *oci_mysql.ListReplicasResponse
}

func (s *MysqlReplicasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlReplicasDataSourceCrud) Get() error {
	request := oci_mysql.ListReplicasRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isUpToDate, ok := s.D.GetOkExists("is_up_to_date"); ok {
		tmp := isUpToDate.(bool)
		request.IsUpToDate = &tmp
	}

	if replicaId, ok := s.D.GetOkExists("id"); ok {
		tmp := replicaId.(string)
		request.ReplicaId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.ReplicaSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListReplicas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReplicas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlReplicasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlReplicasDataSource-", MysqlReplicasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		replica := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			replica["availability_domain"] = *r.AvailabilityDomain
		}

		if r.ConfigurationId != nil {
			replica["configuration_id"] = *r.ConfigurationId
		}

		if r.DbSystemId != nil {
			replica["db_system_id"] = *r.DbSystemId
		}

		if r.DefinedTags != nil {
			replica["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			replica["description"] = *r.Description
		}

		if r.DisplayName != nil {
			replica["display_name"] = *r.DisplayName
		}

		if r.FaultDomain != nil {
			replica["fault_domain"] = *r.FaultDomain
		}

		replica["freeform_tags"] = r.FreeformTags
		replica["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			replica["id"] = *r.Id
		}

		if r.IpAddress != nil {
			replica["ip_address"] = *r.IpAddress
		}

		if r.IsDeleteProtected != nil {
			replica["is_delete_protected"] = *r.IsDeleteProtected
		}

		if r.LifecycleDetails != nil {
			replica["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MysqlVersion != nil {
			replica["mysql_version"] = *r.MysqlVersion
		}

		if r.Port != nil {
			replica["port"] = *r.Port
		}

		if r.PortX != nil {
			replica["port_x"] = *r.PortX
		}

		if r.ReplicaOverrides != nil {
			replica["replica_overrides"] = []interface{}{ReplicaOverridesToMap(r.ReplicaOverrides)}
		} else {
			replica["replica_overrides"] = nil
		}

		if r.ShapeName != nil {
			replica["shape_name"] = *r.ShapeName
		}

		replica["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			replica["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			replica["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, replica)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlReplicasDataSource().Schema["replicas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("replicas", resources); err != nil {
		return err
	}

	return nil
}
