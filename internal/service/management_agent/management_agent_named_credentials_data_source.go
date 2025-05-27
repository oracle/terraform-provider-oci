// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentNamedCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentNamedCredentials,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"management_agent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"named_credential_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ManagementAgentNamedCredentialResource()),
						},
					},
				},
			},
		},
	}
}

func readManagementAgentNamedCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentNamedCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentNamedCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListNamedCredentialsResponse
}

func (s *ManagementAgentNamedCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentNamedCredentialsDataSourceCrud) Get() error {
	request := oci_management_agent.ListNamedCredentialsRequest{}

	if id, ok := s.D.GetOkExists("id"); ok {
		interfaces, ok := id.([]interface{})
		if !ok {
			return errors.New("invalid id type; expected array of strings")
		}
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				if str, ok := interfaces[i].(string); ok {
					tmp[i] = str
				} else {
					return errors.New("invalid id type; expected string")
				}
			}
		}
		if len(tmp) != 0 || s.D.HasChange("id") {
			request.Id = tmp
		}
	}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces, ok := name.([]interface{})
		if !ok {
			return errors.New("invalid name type; expected array of strings")
		}
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				if str, ok := interfaces[i].(string); ok {
					tmp[i] = str
				} else {
					return errors.New("invalid name type; expected string")
				}
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Name = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		stateArr, ok := state.([]interface{})
		if !ok {
			return errors.New("invalid state type; expected array of strings")
		}
		lifecycleStates := make([]oci_management_agent.NamedCredentialLifecycleStateEnum, len(stateArr))
		for i, s := range stateArr {
			stateStr, ok := s.(string)
			if !ok {
				return errors.New("invalid state type; expected string")
			}
			lifecycleStates[i] = oci_management_agent.NamedCredentialLifecycleStateEnum(stateStr)
		}
		request.LifecycleState = lifecycleStates
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		typeArr, ok := type_.([]interface{})
		if !ok {
			return errors.New("invalid Type type; expected array of strings")
		}
		tmp := make([]string, len(typeArr))
		for i := range typeArr {
			if typeArr[i] != nil {
				if str, ok := typeArr[i].(string); ok {
					tmp[i] = str
				} else {
					return errors.New("invalid Type type; expected string")
				}
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.ListNamedCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNamedCredentials(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentNamedCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentNamedCredentialsDataSource-", ManagementAgentNamedCredentialsDataSource(), s.D))
	resources := []map[string]interface{}{}
	namedCredential := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NamedCredentialSummaryToMap(item))
	}
	namedCredential["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ManagementAgentNamedCredentialsDataSource().Schema["named_credential_collection"].Elem.(*schema.Resource).Schema)
		namedCredential["items"] = items
	}

	resources = append(resources, namedCredential)
	if err := s.D.Set("named_credential_collection", resources); err != nil {
		return err
	}

	return nil
}
