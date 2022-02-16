// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database_management "github.com/oracle/oci-go-sdk/v58/databasemanagement"
)

func DatabaseManagementManagedDatabasesResetDatabaseParameterResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementManagedDatabasesResetDatabaseParameter,
		Read:     readDatabaseManagementManagedDatabasesResetDatabaseParameter,
		Delete:   deleteDatabaseManagementManagedDatabasesResetDatabaseParameter,
		Schema: map[string]*schema.Schema{
			// Required
			"credentials": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDatabaseManagementManagedDatabasesResetDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasesResetDatabaseParameterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementManagedDatabasesResetDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseManagementManagedDatabasesResetDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementManagedDatabasesResetDatabaseParameterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Req                    *oci_database_management.ResetDatabaseParametersRequest
	Res                    *oci_database_management.UpdateDatabaseParametersResult
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementManagedDatabasesResetDatabaseParameterResourceCrud) ID() string {
	return *s.Req.ManagedDatabaseId
}

func (s *DatabaseManagementManagedDatabasesResetDatabaseParameterResourceCrud) Create() error {
	request := oci_database_management.ResetDatabaseParametersRequest{}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", 0)
			tmp, err := s.mapToDatabaseCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Credentials = &tmp
		}
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_database_management.ParameterScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.ResetDatabaseParameters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UpdateDatabaseParametersResult
	s.Req = &request
	return nil
}

func (s *DatabaseManagementManagedDatabasesResetDatabaseParameterResourceCrud) SetData() error {
	s.D.Set("status", s.Res.Status)

	return nil
}

func (s *DatabaseManagementManagedDatabasesResetDatabaseParameterResourceCrud) mapToDatabaseCredentials(fieldKeyFormat string) (oci_database_management.DatabaseCredentials, error) {
	result := oci_database_management.DatabaseCredentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
		result.Role = oci_database_management.DatabaseCredentialsRoleEnum(role.(string))
	}

	if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
		tmp := secretId.(string)
		result.SecretId = &tmp
	}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	return result, nil
}

func DatabaseCredentialsToMap(obj *oci_database_management.DatabaseCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	result["role"] = string(obj.Role)

	if obj.SecretId != nil {
		result["secret_id"] = string(*obj.SecretId)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}
