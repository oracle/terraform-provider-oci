// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsGroup,
		Read:     readIdentityDomainsGroup,
		Update:   updateIdentityDomainsGroup,
		Delete:   deleteIdentityDomainsGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_delete": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      membersHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"date_added": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"membership_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"non_unique_display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextension_oci_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"defined_tags": {
							Type:             schema.TypeList,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"freeform_tags": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
						"tag_slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensiondynamic_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"membership_rule": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"membership_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensiongroup_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"creation_mechanism": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"owners": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// Computed
						"app_roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"admin_role": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"app_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"app_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"legacy_group_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"grants": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"app_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"grant_mechanism": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"password_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"synced_from_app": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionposix_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"gid_number": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionrequestable_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"requestable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"urnietfparamsscimschemasoracleidcsextensiondbcs_group": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain_level_schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain_level_schema_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"domain_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"instance_level_schema": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_level_schema_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"db_instance_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
		},
	}
}

func createIdentityDomainsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGroupResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGroupResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "groups")
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

func updateIdentityDomainsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGroupResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsGroupResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.Group
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsGroupResourceCrud) Create() error {
	request := oci_identity_domains.CreateGroupRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if members, ok := s.D.GetOkExists("members"); ok {
		set := members.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.GroupMembers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := membersHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToGroupMembers(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	if nonUniqueDisplayName, ok := s.D.GetOkExists("non_unique_display_name"); ok {
		tmp := nonUniqueDisplayName.(string)
		request.NonUniqueDisplayName = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionOCITags, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextension_oci_tags"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionOCITags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextension_oci_tags", 0)
			tmp, err := s.mapToExtensionOCITags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiondynamicGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiondynamic_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiondynamicGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiondynamic_group", 0)
			tmp, err := s.mapToExtensionDynamicGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiongroupGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiongroup_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiongroupGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiongroup_group", 0)
			tmp, err := s.mapToExtensionGroupGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionposixGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionposix_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionposixGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionposix_group", 0)
			tmp, err := s.mapToExtensionPosixGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionrequestableGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionrequestable_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionrequestableGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionrequestable_group", 0)
			tmp, err := s.mapToExtensionRequestableGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Group
	return nil
}

func (s *IdentityDomainsGroupResourceCrud) Get() error {
	request := oci_identity_domains.GetGroupRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	tmp := s.D.Id()
	request.GroupId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	groupId, err := parseGroupCompositeId(s.D.Id())
	if err == nil {
		request.GroupId = &groupId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Group
	return nil
}

func (s *IdentityDomainsGroupResourceCrud) Update() error {
	request := oci_identity_domains.PutGroupRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	tmp := s.D.Id()
	request.GroupId = &tmp

	tmp = s.D.Id()
	request.Id = &tmp

	if members, ok := s.D.GetOkExists("members"); ok {
		set := members.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.GroupMembers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := membersHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToGroupMembers(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	if nonUniqueDisplayName, ok := s.D.GetOkExists("non_unique_display_name"); ok {
		tmp := nonUniqueDisplayName.(string)
		request.NonUniqueDisplayName = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionOCITags, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextension_oci_tags"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionOCITags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextension_oci_tags", 0)
			tmp, err := s.mapToExtensionOCITags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiondynamicGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiondynamic_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiondynamicGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiondynamic_group", 0)
			tmp, err := s.mapToExtensionDynamicGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiongroupGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiongroup_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiongroupGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiongroup_group", 0)
			tmp, err := s.mapToExtensionGroupGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionposixGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionposix_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionposixGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionposix_group", 0)
			tmp, err := s.mapToExtensionPosixGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionrequestableGroup, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionrequestable_group"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionrequestableGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionrequestable_group", 0)
			tmp, err := s.mapToExtensionRequestableGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Group
	return nil
}

func (s *IdentityDomainsGroupResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteGroupRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.GroupId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteGroup(context.Background(), request)
	return err
}

func (s *IdentityDomainsGroupResourceCrud) SetData() error {

	groupId, err := parseGroupCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(groupId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	members := []interface{}{}
	for _, item := range s.Res.Members {
		members = append(members, GroupMembersToMap(item))
	}
	s.D.Set("members", schema.NewSet(membersHashCodeForSets, members))

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.NonUniqueDisplayName != nil {
		s.D.Set("non_unique_display_name", *s.Res.NonUniqueDisplayName)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", []interface{}{ExtensionOCITagsToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_group", []interface{}{ExtensionDbcsGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondynamic_group", []interface{}{ExtensionDynamicGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondynamic_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiongroup_group", []interface{}{ExtensionGroupGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiongroup_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_group", []interface{}{ExtensionPosixGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionposix_group", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_group", []interface{}{ExtensionRequestableGroupToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_group", nil)
	}

	return nil
}

//func GetGroupCompositeId(groupId string) string {
//	groupId = url.PathEscape(groupId)
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/groups/" + groupId
//	return compositeId
//}

func parseGroupCompositeId(compositeId string) (groupId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/groups/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	groupId, _ = url.PathUnescape(parts[3])

	return
}

func ExtensionDbcsGroupToMap(obj *oci_identity_domains.ExtensionDbcsGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainLevelSchema != nil {
		result["domain_level_schema"] = string(*obj.DomainLevelSchema)
	}

	domainLevelSchemaNames := []interface{}{}
	for _, item := range obj.DomainLevelSchemaNames {
		domainLevelSchemaNames = append(domainLevelSchemaNames, GroupExtDomainLevelSchemaNamesToMap(item))
	}
	result["domain_level_schema_names"] = domainLevelSchemaNames

	if obj.InstanceLevelSchema != nil {
		result["instance_level_schema"] = string(*obj.InstanceLevelSchema)
	}

	instanceLevelSchemaNames := []interface{}{}
	for _, item := range obj.InstanceLevelSchemaNames {
		instanceLevelSchemaNames = append(instanceLevelSchemaNames, GroupExtInstanceLevelSchemaNamesToMap(item))
	}
	result["instance_level_schema_names"] = instanceLevelSchemaNames

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToExtensionDynamicGroup(fieldKeyFormat string) (oci_identity_domains.ExtensionDynamicGroup, error) {
	result := oci_identity_domains.ExtensionDynamicGroup{}

	if membershipRule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "membership_rule")); ok {
		tmp := membershipRule.(string)
		if tmp != "" {
			result.MembershipRule = &tmp
		}
	}

	if membershipType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "membership_type")); ok {
		result.MembershipType = oci_identity_domains.ExtensionDynamicGroupMembershipTypeEnum(membershipType.(string))
	}

	return result, nil
}

func ExtensionDynamicGroupToMap(obj *oci_identity_domains.ExtensionDynamicGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MembershipRule != nil {
		result["membership_rule"] = string(*obj.MembershipRule)
	}

	result["membership_type"] = string(obj.MembershipType)

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToExtensionGroupGroup(fieldKeyFormat string) (oci_identity_domains.ExtensionGroupGroup, error) {
	result := oci_identity_domains.ExtensionGroupGroup{}

	if appRoles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "app_roles")); ok {
		interfaces := appRoles.([]interface{})
		tmp := make([]oci_identity_domains.GroupExtAppRoles, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "app_roles"), stateDataIndex)
			converted, err := s.mapToGroupExtAppRoles(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "app_roles")) {
			result.AppRoles = tmp
		}
	}

	if creationMechanism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "creation_mechanism")); ok {
		result.CreationMechanism = oci_identity_domains.ExtensionGroupGroupCreationMechanismEnum(creationMechanism.(string))
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if grants, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grants")); ok {
		interfaces := grants.([]interface{})
		tmp := make([]oci_identity_domains.GroupExtGrants, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "grants"), stateDataIndex)
			converted, err := s.mapToGroupExtGrants(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "grants")) {
			result.Grants = tmp
		}
	}

	if owners, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "owners")); ok {
		interfaces := owners.([]interface{})
		tmp := make([]oci_identity_domains.GroupExtOwners, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "owners"), stateDataIndex)
			converted, err := s.mapToGroupExtOwners(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "owners")) {
			result.Owners = tmp
		}
	}

	if passwordPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_policy")); ok {
		if tmpList := passwordPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "password_policy"), 0)
			tmp, err := s.mapToGroupExtPasswordPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert password_policy, encountered error: %v", err)
			}
			result.PasswordPolicy = &tmp
		}
	}

	if syncedFromApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "synced_from_app")); ok {
		if tmpList := syncedFromApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "synced_from_app"), 0)
			tmp, err := s.mapToGroupExtSyncedFromApp(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert synced_from_app, encountered error: %v", err)
			}
			result.SyncedFromApp = &tmp
		}
	}

	return result, nil
}

func ExtensionGroupGroupToMap(obj *oci_identity_domains.ExtensionGroupGroup) map[string]interface{} {
	result := map[string]interface{}{}

	appRoles := []interface{}{}
	for _, item := range obj.AppRoles {
		appRoles = append(appRoles, GroupExtAppRolesToMap(item))
	}
	result["app_roles"] = appRoles

	result["creation_mechanism"] = string(obj.CreationMechanism)

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	grants := []interface{}{}
	for _, item := range obj.Grants {
		grants = append(grants, GroupExtGrantsToMap(item))
	}
	result["grants"] = grants

	owners := []interface{}{}
	for _, item := range obj.Owners {
		owners = append(owners, GroupExtOwnersToMap(item))
	}
	result["owners"] = owners

	if obj.PasswordPolicy != nil {
		result["password_policy"] = []interface{}{GroupExtPasswordPolicyToMap(obj.PasswordPolicy)}
	}

	if obj.SyncedFromApp != nil {
		result["synced_from_app"] = []interface{}{GroupExtSyncedFromAppToMap(obj.SyncedFromApp)}
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToExtensionOCITags(fieldKeyFormat string) (oci_identity_domains.ExtensionOciTags, error) {
	result := oci_identity_domains.ExtensionOciTags{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		interfaces := definedTags.([]interface{})
		tmp := make([]oci_identity_domains.DefinedTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "defined_tags"), stateDataIndex)
			converted, err := s.mapTodefinedTags(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "defined_tags")) {
			result.DefinedTags = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		interfaces := freeformTags.([]interface{})
		tmp := make([]oci_identity_domains.FreeformTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "freeform_tags"), stateDataIndex)
			converted, err := s.mapTofreeformTags(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "freeform_tags")) {
			result.FreeformTags = tmp
		}
	}

	if tagSlug, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_slug")); ok {
		result.TagSlug = &tagSlug
	}

	return result, nil
}

func (s *IdentityDomainsGroupResourceCrud) mapToExtensionPosixGroup(fieldKeyFormat string) (oci_identity_domains.ExtensionPosixGroup, error) {
	result := oci_identity_domains.ExtensionPosixGroup{}

	if gidNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gid_number")); ok {
		tmp := gidNumber.(int)
		result.GidNumber = &tmp
	}

	return result, nil
}

func ExtensionPosixGroupToMap(obj *oci_identity_domains.ExtensionPosixGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GidNumber != nil {
		result["gid_number"] = int(*obj.GidNumber)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToExtensionRequestableGroup(fieldKeyFormat string) (oci_identity_domains.ExtensionRequestableGroup, error) {
	result := oci_identity_domains.ExtensionRequestableGroup{}

	if requestable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "requestable")); ok {
		tmp := requestable.(bool)
		result.Requestable = &tmp
	}

	return result, nil
}

func ExtensionRequestableGroupToMap(obj *oci_identity_domains.ExtensionRequestableGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Requestable != nil {
		result["requestable"] = bool(*obj.Requestable)
	}

	return result
}

func GroupToMap(obj oci_identity_domains.Group) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	members := []interface{}{}
	for _, item := range obj.Members {
		members = append(members, GroupMembersToMap(item))
	}
	result["members"] = members

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.NonUniqueDisplayName != nil {
		result["non_unique_display_name"] = string(*obj.NonUniqueDisplayName)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["schemas"] = obj.Schemas

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		result["urnietfparamsscimschemasoracleidcsextension_oci_tags"] = []interface{}{ExtensionOCITagsToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsGroup != nil {
		result["urnietfparamsscimschemasoracleidcsextensiondbcs_group"] = []interface{}{ExtensionDbcsGroupToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsGroup)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup != nil {
		result["urnietfparamsscimschemasoracleidcsextensiondynamic_group"] = []interface{}{ExtensionDynamicGroupToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDynamicGroup)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup != nil {
		result["urnietfparamsscimschemasoracleidcsextensiongroup_group"] = []interface{}{ExtensionGroupGroupToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionGroupGroup)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup != nil {
		result["urnietfparamsscimschemasoracleidcsextensionposix_group"] = []interface{}{ExtensionPosixGroupToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionPosixGroup)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup != nil {
		result["urnietfparamsscimschemasoracleidcsextensionrequestable_group"] = []interface{}{ExtensionRequestableGroupToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableGroup)}
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtAppRoles(fieldKeyFormat string) (oci_identity_domains.GroupExtAppRoles, error) {
	result := oci_identity_domains.GroupExtAppRoles{}

	if adminRole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_role")); ok {
		tmp := adminRole.(bool)
		result.AdminRole = &tmp
	}

	if appId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "app_id")); ok {
		tmp := appId.(string)
		result.AppId = &tmp
	}

	if appName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "app_name")); ok {
		tmp := appName.(string)
		result.AppName = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if legacyGroupName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "legacy_group_name")); ok {
		tmp := legacyGroupName.(string)
		result.LegacyGroupName = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.GroupExtAppRolesTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GroupExtAppRolesToMap(obj oci_identity_domains.GroupExtAppRoles) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminRole != nil {
		result["admin_role"] = bool(*obj.AdminRole)
	}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	if obj.AppName != nil {
		result["app_name"] = string(*obj.AppName)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.LegacyGroupName != nil {
		result["legacy_group_name"] = string(*obj.LegacyGroupName)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtDomainLevelSchemaNames(fieldKeyFormat string) (oci_identity_domains.GroupExtDomainLevelSchemaNames, error) {
	result := oci_identity_domains.GroupExtDomainLevelSchemaNames{}

	if domainName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_name")); ok {
		tmp := domainName.(string)
		result.DomainName = &tmp
	}

	if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
		tmp := schemaName.(string)
		result.SchemaName = &tmp
	}

	return result, nil
}

func GroupExtDomainLevelSchemaNamesToMap(obj oci_identity_domains.GroupExtDomainLevelSchemaNames) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtGrants(fieldKeyFormat string) (oci_identity_domains.GroupExtGrants, error) {
	result := oci_identity_domains.GroupExtGrants{}

	if appId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "app_id")); ok {
		tmp := appId.(string)
		result.AppId = &tmp
	}

	if grantMechanism, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grant_mechanism")); ok {
		result.GrantMechanism = oci_identity_domains.GroupExtGrantsGrantMechanismEnum(grantMechanism.(string))
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GroupExtGrantsToMap(obj oci_identity_domains.GroupExtGrants) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	result["grant_mechanism"] = string(obj.GrantMechanism)

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtInstanceLevelSchemaNames(fieldKeyFormat string) (oci_identity_domains.GroupExtInstanceLevelSchemaNames, error) {
	result := oci_identity_domains.GroupExtInstanceLevelSchemaNames{}

	if dbInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_instance_id")); ok {
		tmp := dbInstanceId.(string)
		result.DbInstanceId = &tmp
	}

	if schemaName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_name")); ok {
		tmp := schemaName.(string)
		result.SchemaName = &tmp
	}

	return result, nil
}

func GroupExtInstanceLevelSchemaNamesToMap(obj oci_identity_domains.GroupExtInstanceLevelSchemaNames) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbInstanceId != nil {
		result["db_instance_id"] = string(*obj.DbInstanceId)
	}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtOwners(fieldKeyFormat string) (oci_identity_domains.GroupExtOwners, error) {
	result := oci_identity_domains.GroupExtOwners{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.GroupExtOwnersTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GroupExtOwnersToMap(obj oci_identity_domains.GroupExtOwners) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtPasswordPolicy(fieldKeyFormat string) (oci_identity_domains.GroupExtPasswordPolicy, error) {
	result := oci_identity_domains.GroupExtPasswordPolicy{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if priority, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "priority")); ok {
		tmp := priority.(int)
		result.Priority = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GroupExtPasswordPolicyToMap(obj *oci_identity_domains.GroupExtPasswordPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupExtSyncedFromApp(fieldKeyFormat string) (oci_identity_domains.GroupExtSyncedFromApp, error) {
	result := oci_identity_domains.GroupExtSyncedFromApp{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.GroupExtSyncedFromAppTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GroupExtSyncedFromAppToMap(obj *oci_identity_domains.GroupExtSyncedFromApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapToGroupMembers(fieldKeyFormat string) (oci_identity_domains.GroupMembers, error) {
	result := oci_identity_domains.GroupMembers{}

	if dateAdded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_added")); ok {
		tmp := dateAdded.(string)
		result.DateAdded = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if membershipOcid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "membership_ocid")); ok {
		tmp := membershipOcid.(string)
		result.MembershipOcid = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.GroupMembersTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func GroupMembersToMap(obj oci_identity_domains.GroupMembers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DateAdded != nil {
		result["date_added"] = string(*obj.DateAdded)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.MembershipOcid != nil {
		result["membership_ocid"] = string(*obj.MembershipOcid)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsGroupResourceCrud) mapTofreeformTags(fieldKeyFormat string) (oci_identity_domains.FreeformTags, error) {
	result := oci_identity_domains.FreeformTags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *IdentityDomainsGroupResourceCrud) mapTodefinedTags(fieldKeyFormat string) (oci_identity_domains.DefinedTags, error) {
	result := oci_identity_domains.DefinedTags{}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *IdentityDomainsGroupResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func membersHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if ocid, ok := m["ocid"]; ok && ocid != "" {
		buf.WriteString(fmt.Sprintf("%v-", ocid))
	}
	if type_, ok := m["type"]; ok && type_ != "" {
		buf.WriteString(fmt.Sprintf("%v-", type_))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return utils.GetStringHashcode(buf.String())
}
