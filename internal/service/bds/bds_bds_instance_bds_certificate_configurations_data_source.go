// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceBdsCertificateConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceBdsCertificateConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bds_certificate_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceBdsCertificateConfigurationResource()),
			},
		},
	}
}

func readBdsBdsInstanceBdsCertificateConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceBdsCertificateConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceBdsCertificateConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsCertificateConfigurationsResponse
}

func (s *BdsBdsInstanceBdsCertificateConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceBdsCertificateConfigurationsDataSourceCrud) Get() error {
	request := oci_bds.ListBdsCertificateConfigurationsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.BdsCertificateConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsCertificateConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsCertificateConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceBdsCertificateConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceBdsCertificateConfigurationsDataSource-", BdsBdsInstanceBdsCertificateConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceBdsCertificateConfiguration := map[string]interface{}{
			"bds_instance_id": *r.BdsInstanceId,
		}

		if r.CertificateAuthorityId != nil {
			bdsInstanceBdsCertificateConfiguration["certificate_authority_id"] = *r.CertificateAuthorityId
		}

		if r.CompartmentId != nil {
			bdsInstanceBdsCertificateConfiguration["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			bdsInstanceBdsCertificateConfiguration["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			bdsInstanceBdsCertificateConfiguration["display_name"] = *r.DisplayName
		}

		bdsInstanceBdsCertificateConfiguration["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			bdsInstanceBdsCertificateConfiguration["id"] = *r.Id
		}

		if r.IsDefaultConfiguration != nil {
			bdsInstanceBdsCertificateConfiguration["is_default_configuration"] = *r.IsDefaultConfiguration
		}

		bdsInstanceBdsCertificateConfiguration["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceBdsCertificateConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastRefreshedOrGenerated != nil {
			bdsInstanceBdsCertificateConfiguration["time_last_refreshed_or_generated"] = r.TimeLastRefreshedOrGenerated.String()
		}

		if r.TimeUpdated != nil {
			bdsInstanceBdsCertificateConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		bdsInstanceBdsCertificateConfiguration["type"] = r.Type

		resources = append(resources, bdsInstanceBdsCertificateConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceBdsCertificateConfigurationsDataSource().Schema["bds_certificate_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bds_certificate_configurations", resources); err != nil {
		return err
	}

	return nil
}
