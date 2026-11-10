// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/oracle/oci-go-sdk/v65/common"
	sdk "github.com/oracle/oci-go-sdk/v65/distributeddatabase"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func sdkRetryToken() string { return common.RetryToken() }

const (
	pendingWorkRequestIDField        = "pending_work_request_id"
	pendingWorkRequestOperationField = "pending_work_request_operation"
	pendingWorkRequestActionField    = "pending_work_request_action"
	pendingOperationalStateField     = "pending_operational_state"
)

func configureMainDatabaseWorkRequestRecoverySchema(resource *schema.Resource) {
	resource.Schema[pendingWorkRequestIDField] = &schema.Schema{Type: schema.TypeString, Computed: true}
	resource.Schema[pendingWorkRequestOperationField] = &schema.Schema{Type: schema.TypeString, Computed: true}
	resource.Schema[pendingWorkRequestActionField] = &schema.Schema{Type: schema.TypeString, Computed: true}
	resource.Schema[pendingOperationalStateField] = &schema.Schema{Type: schema.TypeString, Computed: true}
}

type pendingMainDatabaseWorkRequest struct {
	ID               string
	Operation        string
	Action           string
	OperationalState string
}

func recordPendingMainDatabaseWorkRequest(d *schema.ResourceData, workRequestID, operation, action, operationalState string) error {
	if err := d.Set(pendingWorkRequestIDField, workRequestID); err != nil {
		return err
	}
	if err := d.Set(pendingWorkRequestOperationField, operation); err != nil {
		return err
	}
	if err := d.Set(pendingWorkRequestActionField, action); err != nil {
		return err
	}
	return d.Set(pendingOperationalStateField, operationalState)
}

func pendingMainDatabaseWorkRequestFromState(d *schema.ResourceData) (pendingMainDatabaseWorkRequest, bool) {
	id := strings.TrimSpace(actionString(d.Get(pendingWorkRequestIDField)))
	if id == "" {
		return pendingMainDatabaseWorkRequest{}, false
	}
	return pendingMainDatabaseWorkRequest{
		ID:               id,
		Operation:        actionString(d.Get(pendingWorkRequestOperationField)),
		Action:           actionString(d.Get(pendingWorkRequestActionField)),
		OperationalState: actionString(d.Get(pendingOperationalStateField)),
	}, true
}

func clearPendingMainDatabaseWorkRequest(d *schema.ResourceData) error {
	for _, field := range []string{pendingWorkRequestIDField, pendingWorkRequestOperationField, pendingWorkRequestActionField, pendingOperationalStateField} {
		if err := d.Set(field, nil); err != nil {
			return err
		}
	}
	return nil
}

func recordDistributedDatabaseActionSubmission(d *schema.ResourceData, kind databaseActionKind, autonomous bool, workID string) error {
	d.SetId(distributedDatabaseActionID("DistributedDatabase"+actionOperationName(kind)+"Action-", d, actionTargetField(autonomous), string(kind)))
	if err := d.Set("work_request_id", workID); err != nil {
		return err
	}
	return d.Set("action_pending", true)
}

func databaseActionError(kind databaseActionKind, id, work string, err error) error {
	prefix := fmt.Sprintf("%s action for database %q", actionOperationName(kind), id)
	if work != "" {
		prefix += fmt.Sprintf(", work request %q", work)
	}
	return fmt.Errorf("%s: %w", prefix, err)
}
func createDistributedDatabaseOperationWithContext(ctx context.Context, d *schema.ResourceData, meta interface{}, kind databaseActionKind, autonomous bool) diag.Diagnostics {
	if err := runDistributedDatabaseOperation(ctx, d, meta, kind, autonomous, true, false); err != nil {
		return tfresource.HandleDiagError(meta, err)
	}
	return nil
}
func updateDistributedDatabaseOperationWithContext(ctx context.Context, d *schema.ResourceData, meta interface{}, kind databaseActionKind, autonomous bool) diag.Diagnostics {
	if err := runDistributedDatabaseOperation(ctx, d, meta, kind, autonomous, false, false); err != nil {
		return tfresource.HandleDiagError(meta, err)
	}
	return nil
}
func readDistributedDatabaseOperationWithContext(ctx context.Context, d *schema.ResourceData, meta interface{}, kind databaseActionKind, autonomous bool) diag.Diagnostics {
	if err := runDistributedDatabaseOperation(ctx, d, meta, kind, autonomous, false, true); err != nil {
		return tfresource.HandleDiagError(meta, err)
	}
	return nil
}
func runDistributedDatabaseOperation(ctx context.Context, d *schema.ResourceData, meta interface{}, kind databaseActionKind, autonomous, creating, read bool) error {
	id := actionString(d.Get(actionTargetField(autonomous)))
	if strings.TrimSpace(id) == "" {
		return databaseActionError(kind, id, "", fmt.Errorf("%s is required", actionTargetField(autonomous)))
	}
	clients := meta.(*client.OracleClients)
	var err error
	if autonomous {
		sync := &DistributedDatabaseDistributedAutonomousDatabasesPatchResourceCrud{Client: clients.DistributedAutonomousDbServiceClient(), WorkRequestClient: clients.DistributedDatabaseDistributedDbWorkRequestServiceClient(), Kind: kind}
		sync.D = d
		if read {
			if d.Get("action_pending").(bool) {
				err = sync.resumeAction(ctx, schema.TimeoutRead)
			} else {
				err = sync.GetWithContext(ctx)
			}
		} else if creating {
			err = sync.CreateWithContext(ctx)
		} else {
			err = sync.UpdateWithContext(ctx)
		}
	} else {
		sync := &DistributedDatabaseDistributedDatabasesPatchResourceCrud{Client: clients.DistributedDbServiceClient(), WorkRequestClient: clients.DistributedDatabaseDistributedDbWorkRequestServiceClient(), Kind: kind}
		sync.D = d
		if read {
			if d.Get("action_pending").(bool) {
				err = sync.resumeAction(ctx, schema.TimeoutRead)
			} else {
				err = sync.GetWithContext(ctx)
			}
		} else if creating {
			err = sync.CreateWithContext(ctx)
		} else {
			err = sync.UpdateWithContext(ctx)
		}
	}
	if err != nil {
		// Only a direct parent GET during refresh may clear stale action state.
		// Pending work-request and create/update errors must remain recoverable.
		if read && !d.Get("action_pending").(bool) {
			if serviceError, ok := common.IsServiceError(err); ok && serviceError.GetHTTPStatusCode() == 404 {
				d.SetId("")
				return nil
			}
		}
		return databaseActionError(kind, id, actionString(d.Get("work_request_id")), err)
	}
	return nil
}

// Each public resource submits operations of one type. Mutable operation lists use
// Update; the target database is immutable for the lifetime of the action.
func distributedDatabaseOperationResource(kind databaseActionKind, autonomous bool) *schema.Resource {
	idField := "distributed_database_id"
	if autonomous {
		idField = "distributed_autonomous_database_id"
	}
	block := map[string]*schema.Schema{
		"selection": {Type: schema.TypeString, Required: true, ValidateFunc: validation.StringIsNotWhiteSpace},
	}
	switch kind {
	case databaseActionInsert:
		block["value"] = &schema.Schema{Type: schema.TypeString, Required: true, Sensitive: true, ValidateFunc: validateActionJSON, DiffSuppressFunc: operationJSONEquivalent, Description: "Sensitive JSON document matching the selected topology's create payload."}
	case databaseActionInsertDg:
		block["replica_name"] = &schema.Schema{Type: schema.TypeString, Required: true, ValidateFunc: validation.StringMatch(databaseActionNamePattern, "must be a valid replica name")}
		block["database"] = operationDataGuardDatabaseSchema(autonomous)
	case databaseActionRemoveDg:
		block["replica_name"] = &schema.Schema{Type: schema.TypeString, Required: true, ValidateFunc: validation.StringMatch(databaseActionNamePattern, "must be a valid replica name")}
	}
	resource := &schema.Resource{
		Description: "Executes the configured database operations. Use additional resource instances for additional operations. Destroy only removes Terraform action state.",
		Timeouts:    &schema.ResourceTimeout{Create: tfresource.GetTimeoutDuration("12h"), Read: tfresource.GetTimeoutDuration("12h"), Update: tfresource.GetTimeoutDuration("12h"), Delete: tfresource.GetTimeoutDuration("12h")},
		Schema: map[string]*schema.Schema{
			idField:             {Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validation.StringIsNotWhiteSpace},
			string(kind):        {Type: schema.TypeList, Required: true, MinItems: 1, MaxItems: 3, Elem: &schema.Resource{Schema: block}, Description: "One to three operations submitted together in a single request. Changing this list executes all configured operations."},
			"work_request_id":   {Type: schema.TypeString, Computed: true},
			"action_pending":    {Type: schema.TypeBool, Computed: true, Description: "An accepted action still requires completion polling. Retries resume its saved work request."},
			"state":             {Type: schema.TypeString, Computed: true},
			"lifecycle_details": {Type: schema.TypeString, Computed: true},
			"retry_token":       {Type: schema.TypeString, Computed: true},
		},
	}
	resource.CustomizeDiff = func(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
		if !d.NewValueKnown(string(kind)) {
			return nil
		}
		_, err := operationActionInstructions(d, kind, autonomous)
		return err
	}
	resource.CreateContext = func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		return createDistributedDatabaseOperationWithContext(ctx, d, meta, kind, autonomous)
	}
	resource.UpdateContext = func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		if !d.HasChange(string(kind)) {
			return nil
		}
		return updateDistributedDatabaseOperationWithContext(ctx, d, meta, kind, autonomous)
	}
	resource.ReadContext = func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		return readDistributedDatabaseOperationWithContext(ctx, d, meta, kind, autonomous)
	}
	resource.DeleteContext = deleteDistributedDatabaseDistributedDatabasesConfigureShardingWithContext
	return resource
}

func operationDataGuardDatabaseSchema(autonomous bool) *schema.Schema {
	resource := DistributedDatabaseDistributedDatabaseResource()
	if autonomous {
		resource = DistributedDatabaseDistributedAutonomousDatabaseResource()
	}
	catalog := resource.Schema["catalog_details"].Elem.(*schema.Resource)
	replica := catalog.Schema["data_guard_replicas"].Elem.(*schema.Resource)
	database := cloneActionSchema(replica.Schema["database"])
	database.Required, database.Optional, database.Computed = true, false, false
	database.MinItems, database.MaxItems = 1, 1
	fields := database.Elem.(*schema.Resource).Schema
	var requestOnly func(*schema.Schema)
	requestOnly = func(field *schema.Schema) {
		field.ForceNew = false
		field.DiffSuppressFunc = nil
		field.ConflictsWith = nil
		field.RequiredWith = nil
		field.ExactlyOneOf = nil
		field.AtLeastOneOf = nil
		if field.Optional {
			field.Computed = false
		}
		if nested, ok := field.Elem.(*schema.Resource); ok {
			for name, child := range nested.Schema {
				if child.Computed && !child.Optional {
					delete(nested.Schema, name)
				} else {
					requestOnly(child)
				}
			}
		}
	}
	requestOnly(database)
	// Reuse placement schemas, but restrict the action to documented DG fields.
	allowed := map[string]bool{"source": true, "vm_cluster_id": true, "db_home_id": true, "availability_domain": true, "exascale_db_storage_vault_id": true, "vm_cluster_details": true, "db_storage_vault_details": true, "admin_password": true, "admin_password_secret_id": true, "admin_password_secret_version_number": true, "protection_mode": true, "transport_type": true}
	if autonomous {
		allowed = map[string]bool{"source": true, "cloud_autonomous_vm_cluster_id": true, "protection_mode": true, "fast_start_fail_over_lag_limit_in_seconds": true, "standby_maintenance_buffer_in_days": true}
	}
	for name, field := range fields {
		if !allowed[name] {
			delete(fields, name)
			continue
		}
		// Placement requirements depend on source. The request validator checks
		// the selected SDK variant rather than requiring every variant's inputs.
		field.Required, field.Optional, field.Computed = name == "source", name != "source", false
	}
	for _, name := range []string{"admin_password", "admin_password_secret_id"} {
		if field := fields[name]; field != nil {
			field.Required = false
			field.Optional = true
			field.Computed = false
			field.Sensitive = name == "admin_password"
		}
	}
	if field := fields["protection_mode"]; field != nil {
		field.Default = "MAXIMUM_PERFORMANCE"
	}
	if field := fields["transport_type"]; field != nil {
		field.Default = "ASYNC"
	}

	return database
}

func actionTargetField(autonomous bool) string {
	if autonomous {
		return "distributed_autonomous_database_id"
	}
	return "distributed_database_id"
}

var databaseTopologyRoots = []string{"catalog_details", "user_shard_spaces", "system_raft_clusters", "composite_raft_shard_spaces", "system_data_guard_databases", "composite_data_guard_shard_spaces"}

type databaseTopologyValues map[string][]interface{}

func initialDatabaseTopology(d *schema.ResourceData) databaseTopologyValues {
	result := databaseTopologyValues{}
	for _, root := range databaseTopologyRoots {
		result[root] = actionList(d.Get(root))
	}
	return result
}

// Main-resource creation inputs remain unchanged when separate action resources
// add members. Full remote membership remains available through observed outputs.
func configureInitialDatabaseTopology(resource *schema.Resource) {
	for _, root := range databaseTopologyRoots {
		configured := resource.Schema[root]
		if configured == nil {
			continue
		}
		observed := cloneActionSchema(configured)
		makeActionSchemaComputed(observed)
		resource.Schema["observed_"+root] = observed
	}
}
func makeActionSchemaComputed(field *schema.Schema) {
	field.Required, field.Optional, field.Computed, field.ForceNew = false, false, true, false
	field.MinItems, field.MaxItems = 0, 0
	field.Default = nil
	field.ValidateFunc = nil
	field.DiffSuppressFunc = nil
	field.ConflictsWith = nil
	field.RequiredWith = nil
	field.ExactlyOneOf = nil
	field.AtLeastOneOf = nil
	if nested, ok := field.Elem.(*schema.Resource); ok {
		for _, child := range nested.Schema {
			makeActionSchemaComputed(child)
		}
	}
}
func preserveInitialDatabaseTopology(d *schema.ResourceData, initial databaseTopologyValues, autonomous bool) error {
	imported := true
	for _, root := range databaseTopologyRoots {
		if len(initial[root]) > 0 {
			imported = false
			break
		}
	}
	for _, root := range databaseTopologyRoots {
		var observed interface{} = d.Get(root)
		if autonomous && root != "catalog_details" {
			observed = d.Get("observed_" + root)
		}
		if err := d.Set("observed_"+root, observed); err != nil {
			return err
		}
		if imported {
			if err := d.Set(root, observed); err != nil {
				return err
			}
		} else {
			if err := d.Set(root, initial[root]); err != nil {
				return err
			}
		}
	}
	return nil
}

func cloneActionSchema(field *schema.Schema) *schema.Schema {
	copy := *field
	if nested, ok := field.Elem.(*schema.Resource); ok {
		element := *nested
		element.Schema = map[string]*schema.Schema{}
		for name, child := range nested.Schema {
			element.Schema[name] = cloneActionSchema(child)
		}
		copy.Elem = &element
	}
	return &copy
}
func actionList(value interface{}) []interface{} { result, _ := value.([]interface{}); return result }
func actionMap(value interface{}) map[string]interface{} {
	result, _ := value.(map[string]interface{})
	return result
}
func actionString(value interface{}) string { result, _ := value.(string); return result }
func actionFirst(value interface{}) map[string]interface{} {
	items := actionList(value)
	if len(items) == 0 {
		return nil
	}
	return actionMap(items[0])
}

type databaseActionKind string

const (
	databaseActionInsert   databaseActionKind = "insert"
	databaseActionRemove   databaseActionKind = "remove"
	databaseActionInsertDg databaseActionKind = "insert_dg"
	databaseActionRemoveDg databaseActionKind = "remove_dg"
)

var databaseActionNamePattern = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{0,39}$`)

func actionOperationName(kind databaseActionKind) string {
	return strings.ToUpper(strings.ReplaceAll(string(kind), "_", ""))
}

func DistributedDatabaseDistributedDatabasesInsertResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionInsert, false)
}

func DistributedDatabaseDistributedDatabasesRemoveResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionRemove, false)
}

func DistributedDatabaseDistributedDatabasesInsertDgResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionInsertDg, false)
}

func DistributedDatabaseDistributedDatabasesRemoveDgResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionRemoveDg, false)
}

func DistributedDatabaseDistributedAutonomousDatabasesInsertResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionInsert, true)
}

func DistributedDatabaseDistributedAutonomousDatabasesRemoveResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionRemove, true)
}

func DistributedDatabaseDistributedAutonomousDatabasesInsertDgResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionInsertDg, true)
}

func DistributedDatabaseDistributedAutonomousDatabasesRemoveDgResource() *schema.Resource {
	return distributedDatabaseOperationResource(databaseActionRemoveDg, true)
}

type operationActionData interface{ Get(string) interface{} }

func validateActionJSON(value interface{}, key string) ([]string, []error) {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(value.(string)), &obj); err != nil || obj == nil || len(obj) == 0 {
		return nil, []error{fmt.Errorf("%s must be a nonempty valid JSON object", key)}
	}
	return nil, nil
}
func operationJSONEquivalent(_ string, old, new string, _ *schema.ResourceData) bool {
	var a, b interface{}
	da, db := json.NewDecoder(strings.NewReader(old)), json.NewDecoder(strings.NewReader(new))
	da.UseNumber()
	db.UseNumber()
	if da.Decode(&a) != nil || db.Decode(&b) != nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}

func operationActionInstructions(d operationActionData, kind databaseActionKind, autonomous bool) ([]sdk.PatchInstruction, error) {
	blocks := actionList(d.Get(string(kind)))
	if len(blocks) < 1 || len(blocks) > 3 {
		return nil, fmt.Errorf("%s requires one to three operation blocks per resource instance; got %d", actionOperationName(kind), len(blocks))
	}
	instructions := make([]sdk.PatchInstruction, 0, len(blocks))
	for index, raw := range blocks {
		block := actionMap(raw)
		selection := actionString(block["selection"])
		path := fmt.Sprintf("%s.%d", kind, index)
		if strings.TrimSpace(selection) == "" {
			return nil, fmt.Errorf("%s.selection is required", path)
		}
		switch kind {
		case databaseActionRemove:
			instructions = append(instructions, sdk.PatchRemoveInstruction{Selection: &selection})
		case databaseActionInsert:
			text := actionString(block["value"])
			if err := validateInsertActionValue(selection, text, autonomous); err != nil {
				return nil, fmt.Errorf("%s.value: %w", path, err)
			}
			var compact bytes.Buffer
			if err := json.Compact(&compact, []byte(text)); err != nil {
				return nil, fmt.Errorf("%s.value must contain valid JSON", path)
			}
			var value interface{} = compact.String()
			instructions = append(instructions, sdk.PatchInsertInstruction{Selection: &selection, Value: &value})
		case databaseActionInsertDg, databaseActionRemoveDg:
			name := actionString(block["replica_name"])
			if !databaseActionNamePattern.MatchString(name) {
				return nil, fmt.Errorf("%s.replica_name is required and must be a valid replica name", path)
			}
			payload := map[string]interface{}{"subOperation": actionOperationName(kind)}
			if kind == databaseActionRemoveDg {
				payload["dataGuardReplica"] = map[string]interface{}{"name": name}
			} else {
				field := operationDataGuardDatabaseSchema(autonomous)
				values := actionList(block["database"])
				if len(values) != 1 {
					return nil, fmt.Errorf("%s.database requires exactly one block", path)
				}
				database := operationTFObject(actionMap(values[0]), field.Elem.(*schema.Resource).Schema, fmt.Sprintf("%s.%d.database.0", kind, index), func(path string) bool { return actionInputPresent(d, path) })
				if err := validateActionDataGuard(database, autonomous); err != nil {
					return nil, fmt.Errorf("%s.database: %w", path, err)
				}
				payload["dataGuardReplicas"] = []interface{}{map[string]interface{}{"name": name, "database": database}}
			}
			encoded, err := json.Marshal(payload)
			if err != nil {
				return nil, fmt.Errorf("%s: cannot encode Data Guard payload", path)
			}
			var value interface{} = string(encoded)
			instructions = append(instructions, sdk.PatchMergeInstruction{Selection: &selection, Value: &value})
		default:
			return nil, fmt.Errorf("unsupported action kind")
		}
	}
	return instructions, nil
}

// Convert only allow-listed request fields. Schema describes which nested blocks
// are singular objects and which fields remain arrays; no response metadata leaks.
func operationTFObject(values map[string]interface{}, fields map[string]*schema.Schema, path string, present func(string) bool) map[string]interface{} {
	result := map[string]interface{}{}
	for name, field := range fields {
		value, ok := values[name]
		if !ok || value == nil {
			continue
		}
		if !field.Required && field.Default == nil && !present(path+"."+name) {
			continue
		}
		if nested, ok := field.Elem.(*schema.Resource); ok {
			items := actionList(value)
			if len(items) == 0 {
				continue
			}
			converted := []interface{}{}
			for index, raw := range items {
				converted = append(converted, operationTFObject(actionMap(raw), nested.Schema, fmt.Sprintf("%s.%s.%d", path, name, index), present))
			}
			if field.MaxItems == 1 {
				value = converted[0]
			} else {
				value = converted
			}
		} else if set, ok := value.(*schema.Set); ok {
			value = set.List()
		}
		pieces := strings.Split(name, "_")
		api := pieces[0]
		for _, piece := range pieces[1:] {
			if piece != "" {
				api += strings.ToUpper(piece[:1]) + piece[1:]
			}
		}
		// Existing API spelling uses a capital O inside FailOver.
		switch name {
		case "fast_start_fail_over_lag_limit_in_seconds":
			api = "fastStartFailOverLagLimitInSeconds"
		case "enabled_ecpu_count":
			api = "enabledECpuCount"
		case "total_ecpu_count":
			api = "totalECpuCount"
		case "autoscale_limit_in_gbs":
			api = "autoscaleLimitInGBs"
		}
		result[api] = value
	}
	return result
}

func validateActionDataGuard(database map[string]interface{}, autonomous bool) error {
	source := actionString(database["source"])
	if autonomous {
		if source != "ADBD_EXISTING_CLUSTER" {
			return fmt.Errorf("source must be ADBD_EXISTING_CLUSTER")
		}
		if strings.TrimSpace(actionString(database["cloudAutonomousVmClusterId"])) == "" {
			return fmt.Errorf("cloud_autonomous_vm_cluster_id is required")
		}
		body, _ := json.Marshal(database)
		var model sdk.CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails
		if json.Unmarshal(body, &model) != nil {
			return fmt.Errorf("invalid Data Guard database fields")
		}
		if _, err := model.ValidateEnumValue(); err != nil {
			return fmt.Errorf("invalid protection_mode")
		}
		return nil
	}
	password := strings.TrimSpace(actionString(database["adminPassword"]))
	secret := strings.TrimSpace(actionString(database["adminPasswordSecretId"]))
	if (password == "") == (secret == "") {
		return fmt.Errorf("supply exactly one of admin_password or admin_password_secret_id")
	}
	if version, ok := database["adminPasswordSecretVersionNumber"]; ok {
		number, ok := version.(int)
		if !ok || number < 1 || number > 60 || secret == "" {
			return fmt.Errorf("admin_password_secret_version_number requires a secret and must be 1–60")
		}
	}
	var model interface{}
	switch source {
	case "XS_EXISTING_CLUSTER":
		model = &sdk.CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails{}
	case "XS_NEW_CLUSTER":
		model = &sdk.CreateDistributedDatabaseCatalogWithExadbXsNewClusterDetails{}
	case "XS_NEW_VAULT_AND_CLUSTER":
		model = &sdk.CreateDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterDetails{}
	case "XD_EXISTING_CLUSTER":
		model = &sdk.CreateDistributedDatabaseCatalogDatabaseWithExadbXdDetails{}
	case "EXISTING_DB_HOME":
		model = &sdk.CreateDistributedDatabaseCatalogDatabaseWithDbHomeDetails{}
	default:
		return fmt.Errorf("source is unsupported for this database service")
	}
	raw, _ := json.Marshal(database)
	if json.Unmarshal(raw, model) != nil {
		return fmt.Errorf("invalid database field types")
	}
	return validateActionSDKModel(reflect.ValueOf(model), "database", false)
}

func validateInsertActionValue(selection, text string, autonomous bool) error {
	var value map[string]interface{}
	if json.Unmarshal([]byte(text), &value) != nil || len(value) == 0 {
		return fmt.Errorf("must be a nonempty valid JSON object")
	}
	// INSERT value follows the create shape for the selected topology. A single
	// instruction can contain multiple databases; the three-item limit counts instructions.
	var model interface{}
	root := strings.Split(selection, ".")[0]
	root = strings.Split(root, "[")[0]
	if autonomous {
		switch root {
		case "catalogDetails":
			model = &sdk.CreateAutonomousCatalogDetails{}
		case "userShardSpaces":
			model = &sdk.CreateAutonomousUserShardSpaceDetails{}
		case "systemRaftClusters":
			model = &sdk.CreateAutonomousSystemRaftClusterDetails{}
		case "compositeRaftShardSpaces":
			model = &sdk.CreateAutonomousCompositeRaftShardSpaceDetails{}
		case "systemDataGuardDatabases":
			model = &sdk.CreateAutonomousSystemDataGuardDatabaseDetails{}
		case "compositeDataGuardShardSpaces":
			model = &sdk.CreateAutonomousCompositeDataGuardShardSpaceDetails{}
		}
	} else {
		switch root {
		case "catalogDetails":
			model = &sdk.CreateDistributedDatabaseCatalogDetails{}
		case "userShardSpaces":
			model = &sdk.CreateUserShardSpaceDetails{}
		case "systemRaftClusters":
			model = &sdk.CreateSystemRaftClusterDetails{}
		case "compositeRaftShardSpaces":
			model = &sdk.CreateCompositeRaftShardSpaceDetails{}
		case "systemDataGuardDatabases":
			model = &sdk.CreateSystemDataGuardDatabaseDetails{}
		case "compositeDataGuardShardSpaces":
			model = &sdk.CreateCompositeDataGuardShardSpaceDetails{}
		}
	}
	if model == nil {
		return fmt.Errorf("selection must target a supported database topology")
	}
	if err := json.Unmarshal([]byte(text), model); err != nil {
		return fmt.Errorf("payload fields do not match the selected topology")
	}
	if err := validateActionSDKModel(reflect.ValueOf(model), "value", true); err != nil {
		return err
	}
	return validateActionJSONSources(value, autonomous)
}

func validateActionJSONSources(value interface{}, autonomous bool) error {
	switch v := value.(type) {
	case map[string]interface{}:
		if raw, ok := v["source"]; ok {
			source := actionString(raw)
			valid := source == "ADBD_EXISTING_CLUSTER"
			if !autonomous {
				valid = source == "XS_EXISTING_CLUSTER" || source == "XS_NEW_CLUSTER" || source == "XS_NEW_VAULT_AND_CLUSTER" || source == "XD_EXISTING_CLUSTER" || source == "EXISTING_DB_HOME"
			}
			if !valid {
				return fmt.Errorf("unsupported database source")
			}
		}
		for _, child := range v {
			if err := validateActionJSONSources(child, autonomous); err != nil {
				return err
			}
		}
	case []interface{}:
		for _, child := range v {
			if err := validateActionJSONSources(child, autonomous); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateActionSDKModel(value reflect.Value, path string, insert bool) error {
	for value.IsValid() && (value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface) {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}
	if !value.IsValid() {
		return nil
	}
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			if err := validateActionSDKModel(value.Index(i), fmt.Sprintf("%s.%d", path, i), insert); err != nil {
				return err
			}
		}
		return nil
	}
	if value.Kind() != reflect.Struct {
		return nil
	}
	if validator, ok := value.Interface().(interface{ ValidateEnumValue() (bool, error) }); ok {
		if _, err := validator.ValidateEnumValue(); err != nil {
			return fmt.Errorf("%s contains an invalid enum value", path)
		}
	}
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		child := value.Field(i)
		if field.PkgPath != "" {
			continue
		}
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name == "" || name == "-" {
			continue
		}
		conditionalRaftField := strings.Contains(typ.Name(), "RaftCluster") && (name == "replicationFactor" || name == "witnessCount" || name == "ruMode")
		createOnlyCompositeRaftField := insert && (typ.Name() == "CreateAutonomousCompositeRaftShardSpaceDetails" || typ.Name() == "CreateCompositeRaftShardSpaceDetails") && (name == "chunkCount" || name == "replicationUnitCount")
		if field.Tag.Get("mandatory") == "true" && !conditionalRaftField && !createOnlyCompositeRaftField {
			missing := child.IsZero()
			if child.Kind() == reflect.Slice {
				missing = child.Len() == 0
			}
			if child.Kind() == reflect.Ptr && !child.IsNil() && child.Elem().Kind() == reflect.String {
				missing = strings.TrimSpace(child.Elem().String()) == ""
			}
			if missing {
				return fmt.Errorf("%s.%s is required", path, name)
			}
		}
		if err := validateActionSDKModel(child, path+"."+name, insert); err != nil {
			return err
		}
	}
	return nil
}

// Raw configuration distinguishes explicit zero values from omitted fields.
// GetOkExists is the fallback for legacy SDK data built without RawConfig.
func actionInputPresent(d operationActionData, path string) bool {
	if reader, ok := d.(interface{ GetRawConfig() cty.Value }); ok {
		value := reader.GetRawConfig()
		if !value.IsNull() && value.IsKnown() {
			for _, part := range strings.Split(path, ".") {
				if !value.IsKnown() || value.IsNull() {
					return false
				}
				if value.Type().IsObjectType() {
					if !value.Type().HasAttribute(part) {
						return false
					}
					value = value.GetAttr(part)
				} else {
					index, err := strconv.Atoi(part)
					if err != nil || !value.CanIterateElements() || index >= value.LengthInt() {
						return false
					}
					value = value.Index(cty.NumberIntVal(int64(index)))
				}
			}
			return !value.IsNull()
		}
	}
	if reader, ok := d.(interface {
		GetOkExists(string) (interface{}, bool)
	}); ok {
		value, present := reader.GetOkExists(path)
		// Legacy diffs can expose zero-valued attributes from prior state even
		// when absent from config. Modern protocol callers use RawConfig above.
		if prior, ok := d.(interface {
			Id() string
			HasChange(string) bool
		}); ok && prior.Id() != "" && !prior.HasChange(path) && value != nil && reflect.ValueOf(value).IsZero() {
			return false
		}
		return present
	}
	return false
}

// RAFT sizing is mandatory only when inserting a new cluster. The specification
// explicitly allows smaller name+databases payloads for an existing cluster.
func validateInsertActionTargets(d *schema.ResourceData, live interface{}) error {
	encoded, err := json.Marshal(live)
	if err != nil {
		return fmt.Errorf("cannot inspect the current database topology")
	}
	var database map[string]interface{}
	if json.Unmarshal(encoded, &database) != nil {
		return fmt.Errorf("cannot inspect current database topology")
	}
	find := func(items interface{}, name string) map[string]interface{} {
		list, _ := items.([]interface{})
		for _, item := range list {
			entry := actionMap(item)
			if actionString(entry["name"]) == name {
				return entry
			}
		}
		return nil
	}
	requireNew := func(cluster map[string]interface{}, existing map[string]interface{}) error {
		if existing != nil {
			return nil
		}
		for _, field := range []string{"replicationFactor", "witnessCount", "ruMode"} {
			value, ok := cluster[field]
			if !ok || value == nil || value == "" {
				return fmt.Errorf("%s is required when INSERT creates a new RAFT cluster", field)
			}
		}
		return nil
	}
	for _, raw := range actionList(d.Get("insert")) {
		block := actionMap(raw)
		var value map[string]interface{}
		if json.Unmarshal([]byte(actionString(block["value"])), &value) != nil {
			return fmt.Errorf("invalid INSERT JSON")
		}
		selection := actionString(block["selection"])
		if selection == "systemRaftClusters" {
			if err := requireNew(value, find(database["systemRaftClusters"], actionString(value["name"]))); err != nil {
				return err
			}
		} else if selection == "compositeRaftShardSpaces" {
			space := find(database["compositeRaftShardSpaces"], actionString(value["name"]))
			for _, rawCluster := range actionList(value["raftClusters"]) {
				cluster := actionMap(rawCluster)
				if err := requireNew(cluster, find(space["raftClusters"], actionString(cluster["name"]))); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Database home placement is required only for the EXISTING_DB_HOME source.
func configureDatabaseHomeSchema(resource *schema.Resource) {
	var visit func(map[string]*schema.Schema)
	visit = func(fields map[string]*schema.Schema) {
		if field := fields["db_home_id"]; field != nil && fields["source"] != nil {
			field.Required, field.Optional, field.Computed = false, true, true
		}
		for _, field := range fields {
			if nested, ok := field.Elem.(*schema.Resource); ok {
				visit(nested.Schema)
			}
		}
	}
	visit(resource.Schema)
	resource.CustomizeDiff = customdiff.Sequence(validateDatabaseHomeDiff, resource.CustomizeDiff)
}

func validateDatabaseHomeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	var visit func(cty.Value, string) error
	visit = func(value cty.Value, path string) error {
		if value.IsNull() || !value.IsKnown() {
			return nil
		}
		kind := value.Type()
		if kind.IsObjectType() {
			if kind.HasAttribute("source") {
				source := value.GetAttr("source")
				if source.IsKnown() && !source.IsNull() && source.AsString() == "EXISTING_DB_HOME" {
					if !kind.HasAttribute("db_home_id") {
						return fmt.Errorf("%s.db_home_id is required when source is EXISTING_DB_HOME", path)
					}
					home := value.GetAttr("db_home_id")
					if home.IsKnown() && (home.IsNull() || strings.TrimSpace(home.AsString()) == "") {
						return fmt.Errorf("%s.db_home_id is required when source is EXISTING_DB_HOME", path)
					}
				}
			}
			for name, child := range value.AsValueMap() {
				if err := visit(child, path+"."+name); err != nil {
					return err
				}
			}
		} else if kind.IsListType() || kind.IsSetType() || kind.IsTupleType() {
			iterator := value.ElementIterator()
			index := 0
			for iterator.Next() {
				_, child := iterator.Element()
				if err := visit(child, fmt.Sprintf("%s.%d", path, index)); err != nil {
					return err
				}
				index++
			}
		}
		return nil
	}
	return visit(d.GetRawConfig(), "configuration")
}
