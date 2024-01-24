// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func KmsReplicationStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularKmsReplicationStatus,
		Schema: map[string]*schema.Schema{
			"replication_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"replica_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularKmsReplicationStatus(d *schema.ResourceData, m interface{}) error {
	sync := &KmsReplicationStatusDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}

	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type KmsReplicationStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.GetReplicationStatusResponse
}

func (s *KmsReplicationStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsReplicationStatusDataSourceCrud) Get() error {
	request := oci_kms.GetReplicationStatusRequest{}

	if replicationId, ok := s.D.GetOkExists("replication_id"); ok {
		tmp := replicationId.(string)
		request.ReplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.GetReplicationStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsReplicationStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsReplicationStatusDataSource-", KmsReplicationStatusDataSource(), s.D))

	replicaDetails := []interface{}{}
	for _, item := range s.Res.ReplicaDetails {
		replicaDetails = append(replicaDetails, KmsReplicaDetailsToMap(item))
	}
	s.D.Set("replica_details", replicaDetails)

	return nil
}

func KmsReplicaDetailsToMap(obj oci_kms.ReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["status"] = string(obj.Status)

	return result
}
