// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
)

func CertificatesManagementAssociationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCertificatesManagementAssociation,
		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"associated_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"association_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificates_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCertificatesManagementAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.GetAssociationResponse
}

func (s *CertificatesManagementAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementAssociationDataSourceCrud) Get() error {
	request := oci_certificates_management.GetAssociationRequest{}

	if associationId, ok := s.D.GetOkExists("association_id"); ok {
		tmp := associationId.(string)
		request.AssociationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.GetAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesManagementAssociationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssociatedResourceId != nil {
		s.D.Set("associated_resource_id", *s.Res.AssociatedResourceId)
	}

	s.D.Set("association_type", s.Res.AssociationType)

	if s.Res.CertificatesResourceId != nil {
		s.D.Set("certificates_resource_id", *s.Res.CertificatesResourceId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func AssociationSummaryToMap(obj oci_certificates_management.AssociationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssociatedResourceId != nil {
		result["associated_resource_id"] = string(*obj.AssociatedResourceId)
	}

	result["association_type"] = string(obj.AssociationType)

	if obj.CertificatesResourceId != nil {
		result["certificates_resource_id"] = string(*obj.CertificatesResourceId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
