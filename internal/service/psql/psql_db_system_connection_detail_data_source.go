// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlDbSystemConnectionDetailDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularPsqlDbSystemConnectionDetail,
		Schema: map[string]*schema.Schema{
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"ca_certificate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoint": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"fqdn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"primary_db_endpoint": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"fqdn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularPsqlDbSystemConnectionDetail(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemConnectionDetailDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlDbSystemConnectionDetailDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.GetConnectionDetailsResponse
}

func (s *PsqlDbSystemConnectionDetailDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlDbSystemConnectionDetailDataSourceCrud) Get() error {
	request := oci_psql.GetConnectionDetailsRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.GetConnectionDetails(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsqlDbSystemConnectionDetailDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlDbSystemConnectionDetailDataSource-", PsqlDbSystemConnectionDetailDataSource(), s.D))

	if s.Res.CaCertificate != nil {
		s.D.Set("ca_certificate", *s.Res.CaCertificate)
	}

	instanceEndpoints := []interface{}{}
	for _, item := range s.Res.InstanceEndpoints {
		instanceEndpoints = append(instanceEndpoints, DbInstanceEndpointToMap(item))
	}
	s.D.Set("instance_endpoints", instanceEndpoints)

	if s.Res.PrimaryDbEndpoint != nil {
		s.D.Set("primary_db_endpoint", []interface{}{EndpointToMap(s.Res.PrimaryDbEndpoint)})
	} else {
		s.D.Set("primary_db_endpoint", nil)
	}

	return nil
}

func DbInstanceEndpointToMap(obj oci_psql.DbInstanceEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbInstanceId != nil {
		result["db_instance_id"] = string(*obj.DbInstanceId)
	}

	if obj.Endpoint != nil {
		result["endpoint"] = []interface{}{EndpointToMap(obj.Endpoint)}
	}

	return result
}

func EndpointToMap(obj *oci_psql.Endpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Fqdn != nil {
		result["fqdn"] = string(*obj.Fqdn)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}
