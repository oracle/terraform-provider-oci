// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPolicy,
		Read:     readPolicy,
		Update:   updatePolicy,
		Delete:   deletePolicy,
		Schema: map[string]*schema.Schema{
			// Required
			// The legacy provider required this and the API requires. Do not make it optional or swap tenancy OCID in behind the scenes
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"statements": {
				Type:             schema.TypeList,
				Required:         true,
				MinItems:         1,
				DiffSuppressFunc: ignorePolicyFormatDiff,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"version_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
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
			"ETag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policyHash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastUpdateETag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// @Deprecated: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Deprecated: crud.FieldDeprecated("time_modified"),
				Computed:   true,
			},
		},
	}
}

func createPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updatePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deletePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type PolicyResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Policy
	DisableNotFoundRetries bool
}

func (s *PolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.PolicyLifecycleStateCreating),
	}
}

func (s *PolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.PolicyLifecycleStateActive),
	}
}

func (s *PolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.PolicyLifecycleStateDeleting),
	}
}

func (s *PolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.PolicyLifecycleStateDeleted),
	}
}

func (s *PolicyResourceCrud) Create() error {
	request := oci_identity.CreatePolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.Statements = []string{}
	if statements, ok := s.D.GetOkExists("statements"); ok {
		interfaces := statements.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.Statements = tmp
	}

	if versionDate, ok := s.D.GetOkExists("version_date"); ok {
		// one off change to go sdk policy.go:57 create_policy_details.go:36 and update_policy_details.go:29 to treat this as a string
		tmp := versionDate.(string)
		request.VersionDate = &tmp
	}
	// TODO: pending spec/sdk versionDate solution to support basic `date` types, some form of this code may be required here and in `Update()` below
	//if versionDate, ok := s.D.GetOkExists("version_date"); ok {
	//	const scheme = `2006-01-02`
	//	tmp, err := time.Parse(scheme, versionDate.(string))
	//	if err != nil {
	//		return err
	//	}
	//	request.VersionDate = &oci_common.SDKTime{Time: tmp}
	//}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreatePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Policy

	// if the response was successful, store off policy hash and etag
	statements := toStringArray(s.D.Get("statements").([]interface{}))
	s.D.Set("policyHash", getMD5Hash(statements))
	s.D.Set("ETag", response.Etag)
	s.D.Set("lastUpdateETag", response.Etag)

	return nil
}

func (s *PolicyResourceCrud) Get() error {
	request := oci_identity.GetPolicyRequest{}

	tmp := s.D.Id()
	request.PolicyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Policy

	// update etag on a successful get
	s.D.Set("ETag", response.Etag)

	return nil
}

func (s *PolicyResourceCrud) Update() error {
	request := oci_identity.UpdatePolicyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.PolicyId = &tmp

	request.Statements = []string{}
	if statements, ok := s.D.GetOkExists("statements"); ok {
		interfaces := statements.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.Statements = tmp
	}

	// TODO: see comment "pending spec/sdk versionDate solution" above
	if versionDate, ok := s.D.GetOkExists("version_date"); ok {
		tmp := versionDate.(string)
		request.VersionDate = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdatePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Policy

	// if the response was successful, store off policy hash and etag
	statements := toStringArray(s.D.Get("statements").([]interface{}))
	s.D.Set("policyHash", getMD5Hash(statements))
	s.D.Set("ETag", response.Etag)
	s.D.Set("lastUpdateETag", response.Etag)

	return nil
}

func (s *PolicyResourceCrud) Delete() error {
	request := oci_identity.DeletePolicyRequest{}

	tmp := s.D.Id()
	request.PolicyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeletePolicy(context.Background(), request)
	return err
}

func (s *PolicyResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("statements", s.Res.Statements)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VersionDate != nil {
		s.D.Set("version_date", *s.Res.VersionDate)
	}

}

func ignorePolicyFormatDiff(k string, old string, new string, d *schema.ResourceData) bool {
	oldHash := getOrDefault(d, "policyHash", "")
	newHash := getMD5Hash(toStringArray(d.Get("statements")))
	oldETag := getOrDefault(d, "lastUpdateETag", "")
	currentETag := getOrDefault(d, "ETag", "")
	suppressDiff := strings.EqualFold(oldHash, newHash) && strings.EqualFold(oldETag, currentETag)
	return suppressDiff
}

func getOrDefault(d *schema.ResourceData, key string, defaultValue string) string {
	valueString := defaultValue
	if value, ok := d.GetOkExists(key); ok {
		valueString = value.(string)
	}
	return valueString
}

func getMD5Hash(values []string) string {
	hash := md5.Sum([]byte(strings.Join(values, "#")))
	return hex.EncodeToString(hash[:])
}

func toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}
