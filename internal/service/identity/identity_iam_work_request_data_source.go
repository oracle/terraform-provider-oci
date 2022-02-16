// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityIamWorkRequestDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityIamWorkRequest,
		Schema: map[string]*schema.Schema{
			"iam_work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"percent_complete": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularIdentityIamWorkRequest(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIamWorkRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityIamWorkRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetIamWorkRequestResponse
}

func (s *IdentityIamWorkRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityIamWorkRequestDataSourceCrud) Get() error {
	request := oci_identity.GetIamWorkRequestRequest{}

	if iamWorkRequestId, ok := s.D.GetOkExists("iam_work_request_id"); ok {
		tmp := iamWorkRequestId.(string)
		request.IamWorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.GetIamWorkRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityIamWorkRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("operation_type", s.Res.OperationType)

	if s.Res.PercentComplete != nil {
		s.D.Set("percent_complete", *s.Res.PercentComplete)
	}

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, IamWorkRequestResourceToMap(item))
	}
	s.D.Set("resources", resources)

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}

func IamWorkRequestResourceToMap(obj oci_identity.IamWorkRequestResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_type"] = string(obj.ActionType)

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.EntityUri != nil {
		result["entity_uri"] = string(*obj.EntityUri)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}
