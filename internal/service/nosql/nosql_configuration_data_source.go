// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package nosql

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NosqlConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NosqlConfigurationResource(), fieldMap, readSingularNosqlConfiguration)
}

func readSingularNosqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NosqlClient()

	return tfresource.ReadResource(sync)
}

type NosqlConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_nosql.NosqlClient
	Res    *oci_nosql.GetConfigurationResponse
}

func (s *NosqlConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NosqlConfigurationDataSourceCrud) Get() error {
	request := oci_nosql.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "nosql")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NosqlConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	var resourceId string
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		resourceId = GetConfigurationCompositeId(tmp)
	} else {
		resourceId = tfresource.GenerateDataSourceHashID("NosqlConfigurationDataSource-", NosqlConfigurationDataSource(), s.D)
	}

	s.D.SetId(resourceId)
	switch v := (s.Res.Configuration).(type) {
	case oci_nosql.HostedConfiguration:
		s.D.Set("environment", "HOSTED")

		if v.KmsKey != nil {
			s.D.Set("kms_key", []interface{}{KmsKeyToMap(v.KmsKey)})
		} else {
			s.D.Set("kms_key", nil)
		}
	case oci_nosql.MultiTenancyConfiguration:
		s.D.Set("environment", "MULTI_TENANCY")
	default:
		log.Printf("[WARN] Received 'environment' of unknown type %v", s.Res.Configuration)
		return nil
	}

	return nil
}
