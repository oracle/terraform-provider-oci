package acctest

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	"github.com/terraform-providers/terraform-provider-oci/internal/provider"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/zclconf/go-cty/cty"
)

func newTerraformStateWithValue(name, key, value string) *terraform.State {
	instanceState := terraform.NewInstanceStateShimmedFromValue(cty.Value{}, 0)
	instanceState.Attributes = make(map[string]string)
	instanceState.Attributes[key] = value
	state := terraform.NewState()
	state.RootModule().Resources = make(map[string]*terraform.ResourceState)
	state.RootModule().Resources[name] = &terraform.ResourceState{
		Primary: instanceState,
	}
	return state
}

func getJsonMapString(key, value string) string {
	result := map[string]string{key: value}
	resultB, _ := json.Marshal(result)
	return string(resultB)
}
func TestUnitFromInstanceState(t *testing.T) {
	stateWithKey := terraform.NewState()
	stateWithKey.RootModule().Resources = make(map[string]*terraform.ResourceState)
	stateWithKey.RootModule().Resources["name"] = &terraform.ResourceState{}

	stateWithPrimary := terraform.NewState()
	stateWithPrimary.RootModule().Resources["name"] = &terraform.ResourceState{
		Primary: terraform.NewInstanceStateShimmedFromValue(cty.Value{}, 0),
	}

	stateWithAttribute := newTerraformStateWithValue("name", "key", "value")
	type args struct {
		s    *terraform.State
		name string
		key  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test empty state",
			args: args{
				s:    terraform.NewState(),
				name: "name",
				key:  "key",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test state with key, No primary instance",
			args: args{
				s:    stateWithKey,
				name: "name",
				key:  "key",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test state with no attribute,Attribute '%s' not found",
			args: args{
				s:    stateWithPrimary,
				name: "name",
				key:  "key",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test state with attribute,no error",
			args: args{
				s:    stateWithAttribute,
				name: "name",
				key:  "key",
			},
			want:    "value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromInstanceState(tt.args.s, tt.args.name, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInstanceState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromInstanceState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitCheckJsonStringsEqual(t *testing.T) {
	type args struct {
		expectedJsonString string
		actualJsonString   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test equal string",
			args: args{
				expectedJsonString: "name",
				actualJsonString:   "name",
			},
			wantErr: false,
		},
		{
			name: "Test invalid Json from expected",
			args: args{
				expectedJsonString: "invalid json",
				actualJsonString:   getJsonMapString("key", "value"),
			},
			wantErr: true,
		},
		{
			name: "Test invalid Json from actual",
			args: args{
				expectedJsonString: getJsonMapString("key", "value"),
				actualJsonString:   "invalid json",
			},
			wantErr: true,
		},
		{
			name: "Test not equal Json string",
			args: args{
				expectedJsonString: getJsonMapString("key", "value1"),
				actualJsonString:   getJsonMapString("key", "value"),
			},
			wantErr: true,
		},
		{
			name: "Test equal Json string",
			args: args{
				expectedJsonString: getJsonMapString("key", "value"),
				actualJsonString:   getJsonMapString("key", "value"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckJsonStringsEqual(tt.args.expectedJsonString, tt.args.actualJsonString); (err != nil) != tt.wantErr {
				t.Errorf("CheckJsonStringsEqual() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnitCheckJsonResourceAttr(t *testing.T) {
	type args struct {
		state        *terraform.State
		name         string
		key          string
		expectedJson string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with empty state",
			args: args{
				state:        terraform.NewState(),
				name:         "name",
				key:          "key",
				expectedJson: "error",
			},
			wantErr: true,
		},
		{
			name: "Test with not equal json",
			args: args{
				state:        newTerraformStateWithValue("name", "key", "value1"),
				name:         "name",
				key:          "key",
				expectedJson: "key:value}",
			},
			wantErr: true,
		},
		{
			name: "Test with equal json",
			args: args{
				state:        newTerraformStateWithValue("name", "key", "value"),
				name:         "name",
				key:          "key",
				expectedJson: getJsonMapString("key", "value"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TestCheckJsonResourceAttr(tt.args.name, tt.args.key, tt.args.expectedJson)
			if err := got(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("TestCheckJsonResourceAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnitCheckResourceAttributesEqual(t *testing.T) {
	stateWithTwoKey := newTerraformStateWithValue("name", "key", "value")
	instanceState := stateWithTwoKey.RootModule().Resources["name"]
	instanceState.Primary.Attributes["key2"] = "value2"
	type args struct {
		name1 string
		key1  string
		name2 string
		key2  string
		state *terraform.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test empty state",
			args: args{
				state: terraform.NewState(),
			},
			wantErr: true,
		},
		{
			name: "Test empty val1",
			args: args{
				name1: "invalid",
				key1:  "invalid",
				state: newTerraformStateWithValue("name", "key", "value"),
			},
			wantErr: true,
		},
		{
			name: "Test empty val2",
			args: args{
				name1: "name",
				key1:  "key",
				name2: "invalid",
				key2:  "invalid",
				state: newTerraformStateWithValue("name", "key", "value"),
			},
			wantErr: true,
		},
		{
			name: "Test not equal value",
			args: args{
				name1: "name",
				key1:  "key",
				name2: "name",
				key2:  "key2",
				state: stateWithTwoKey,
			},
			wantErr: true,
		},
		{
			name: "Test equal value",
			args: args{
				name1: "name",
				key1:  "key",
				name2: "name",
				key2:  "key",
				state: newTerraformStateWithValue("name", "key", "value"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TestCheckResourceAttributesEqual(tt.args.name1, tt.args.key1, tt.args.name2, tt.args.key2)
			if err := got(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("TestCheckResourceAttributesEqual() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnitCheckAttributeBase64Encoded(t *testing.T) {
	type args struct {
		name                string
		key                 string
		state               *terraform.State
		expectBase64Encoded bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with empty state",
			args: args{
				state: terraform.NewState(),
			},
			wantErr: true,
		},
		{
			name: "Test with content not base64, expected base64",
			args: args{
				state:               newTerraformStateWithValue("name", "key", "value"),
				name:                "name",
				key:                 "key",
				expectBase64Encoded: true,
			},
			wantErr: true,
		},
		{
			name: "Test with content base64, expected not base64",
			args: args{
				state:               newTerraformStateWithValue("name", "key", base64.StdEncoding.EncodeToString([]byte("value"))),
				name:                "name",
				key:                 "key",
				expectBase64Encoded: false,
			},
			wantErr: true,
		},
		{
			name: "Test with content base64, expected base64",
			args: args{
				state:               newTerraformStateWithValue("name", "key", base64.StdEncoding.EncodeToString([]byte("value"))),
				name:                "name",
				key:                 "key",
				expectBase64Encoded: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TestCheckAttributeBase64Encoded(tt.args.name, tt.args.key, tt.args.expectBase64Encoded)
			if err := got(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("TestCheckAttributeBase64Encoded() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnitCheckResourceSetContainsElementWithProperties(t *testing.T) {
	stateNotPrimary := terraform.NewState()
	stateNotPrimary.RootModule().Resources = make(map[string]*terraform.ResourceState)
	stateNotPrimary.RootModule().Resources["name"] = &terraform.ResourceState{}
	type args struct {
		name              string
		setKey            string
		properties        map[string]string
		presentProperties []string
		state             *terraform.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with empty state",
			args: args{
				name:  "name",
				state: terraform.NewState(),
			},
			wantErr: true,
		},
		{
			name: "Test state with key, No primary instance",
			args: args{
				name:  "name",
				state: stateNotPrimary,
			},
			wantErr: true,
		},
		{
			name: "Test state with key and property",
			args: args{
				name:       "name",
				state:      newTerraformStateWithValue("name", "key", "value"),
				setKey:     "key",
				properties: map[string]string{"key1": "value1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckResourceSetContainsElementWithProperties(tt.args.name, tt.args.setKey, tt.args.properties, tt.args.presentProperties)
			if err := got(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("CheckResourceSetContainsElementWithProperties() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestUnitCheckResourceSetContainsElementWithPropertiesContainingNestedSets(t *testing.T) {
	stateNotPrimary := terraform.NewState()
	stateNotPrimary.RootModule().Resources = make(map[string]*terraform.ResourceState)
	stateNotPrimary.RootModule().Resources["name"] = &terraform.ResourceState{}
	type args struct {
		name              string
		setKey            string
		properties        map[string]interface{}
		presentProperties []string
		state             *terraform.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with empty state",
			args: args{
				name:  "name",
				state: terraform.NewState(),
			},
			wantErr: true,
		},
		{
			name: "Test state with key, No primary instance",
			args: args{
				name:  "name",
				state: stateNotPrimary,
			},
			wantErr: true,
		},
		{
			name: "Test state with key and property",
			args: args{
				name:       "name",
				state:      newTerraformStateWithValue("name", "key", "value"),
				setKey:     "key",
				properties: map[string]interface{}{"key1": "value1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckResourceSetContainsElementWithPropertiesContainingNestedSets(tt.args.name, tt.args.setKey, tt.args.properties, tt.args.presentProperties)
			if err := got(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("CheckResourceSetContainsElementWithPropertiesContainingNestedSets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnitCommonTestVariables(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test content return as is",
			want: `
	variable "tenancy_ocid" {
		default = "` + utils.GetEnvSettingWithBlankDefault("tenancy_ocid") + `"
	}

	variable "ssh_public_key" {
		default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
	}

	variable "region" {
		default = "` + utils.GetEnvSettingWithBlankDefault("region") + `"
	}

	`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonTestVariables(); got != tt.want {
				t.Errorf("CommonTestVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitComposeAggregateTestCheckFuncWrapper(t *testing.T) {
	type args struct {
		fs []resource.TestCheckFunc
	}
	tests := []struct {
		name string
		args args
		want resource.TestCheckFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComposeAggregateTestCheckFuncWrapper(tt.args.fs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComposeAggregateTestCheckFuncWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestUnitConditionShouldRetry(t *testing.T) {
//	type args struct {
//		timeout                time.Duration
//		shouldWait             ShouldWaitFunc
//		service                string
//		disableNotFoundRetries bool
//		optionals              []interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want func(response oci_common.OCIOperationResponse) bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ConditionShouldRetry(tt.args.timeout, tt.args.shouldWait, tt.args.service, tt.args.disableNotFoundRetries, tt.args.optionals...); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ConditionShouldRetry() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestUnitGenerateDataSourceFromRepresentationMap(t *testing.T) {
	auditEventDataSourceRepresentation := map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"end_time":       Representation{RepType: Required, Create: `${timestamp()}`},
		"start_time":     Representation{RepType: Required, Create: `${timeadd(timestamp(), "-1m")}`},
	}

	type args struct {
		resourceType       string
		resourceName       string
		representationType RepresentationType
		representationMode RepresentationMode
		representations    map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive test",
			args: args{
				resourceType:       "oci_audit_events",
				resourceName:       "test_audit_events",
				representationType: Required,
				representationMode: Create,
				representations:    auditEventDataSourceRepresentation,
			},
			want: `
data "oci_audit_events" "test_audit_events" {
compartment_id = "${var.compartment_id}"
end_time = "${timestamp()}"
start_time = "${timeadd(timestamp(), "-1m")}"
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDataSourceFromRepresentationMap(tt.args.resourceType, tt.args.resourceName, tt.args.representationType, tt.args.representationMode, tt.args.representations); got != tt.want {
				t.Errorf("GenerateDataSourceFromRepresentationMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitGenerateResourceFromMap(t *testing.T) {
	auditEventDataSourceRepresentation := map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"end_time":       Representation{RepType: Required, Create: `${timestamp()}`, Update: `update`},
		"start_time":     Representation{RepType: Required, Create: `${timeadd(timestamp(), "-1m")}`},
	}
	type args struct {
		representationType RepresentationType
		representationMode RepresentationMode
		representations    map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive test with Create",
			args: args{
				representationType: Required,
				representationMode: Create,
				representations:    auditEventDataSourceRepresentation,
			},
			want: `{
compartment_id = "${var.compartment_id}"
end_time = "${timestamp()}"
start_time = "${timeadd(timestamp(), "-1m")}"
}
`,
		},
		{
			name: "Positive test with Update",
			args: args{
				representationType: Required,
				representationMode: Update,
				representations:    auditEventDataSourceRepresentation,
			},
			want: `{
compartment_id = "${var.compartment_id}"
end_time = "update"
start_time = "${timeadd(timestamp(), "-1m")}"
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateResourceFromMap(tt.args.representationType, tt.args.representationMode, tt.args.representations); got != tt.want {
				t.Errorf("GenerateResourceFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitGenerateResourceFromRepresentationMap(t *testing.T) {
	auditEventResourceRepresentation := map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"end_time":       Representation{RepType: Required, Create: `${timestamp()}`},
		"start_time":     Representation{RepType: Required, Create: `${timeadd(timestamp(), "-1m")}`},
	}
	type args struct {
		resourceType       string
		resourceName       string
		representationType RepresentationType
		representationMode RepresentationMode
		representations    map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive test",
			args: args{
				resourceType:       "oci_audit_events",
				resourceName:       "test_audit_events",
				representationType: Required,
				representationMode: Create,
				representations:    auditEventResourceRepresentation,
			},
			want: `
resource "oci_audit_events" "test_audit_events" {
compartment_id = "${var.compartment_id}"
end_time = "${timestamp()}"
start_time = "${timeadd(timestamp(), "-1m")}"
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateResourceFromRepresentationMap(tt.args.resourceType, tt.args.resourceName, tt.args.representationType, tt.args.representationMode, tt.args.representations); got != tt.want {
				t.Errorf("GenerateResourceFromRepresentationMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestUnitGenericTestStepPreConfiguration(t *testing.T) {
//	type args struct {
//		stepNumber int
//	}
//	tests := []struct {
//		name string
//		args args
//		want func()
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := GenericTestStepPreConfiguration(tt.args.stepNumber); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GenericTestStepPreConfiguration() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestUnitGetMultipleUpdatedRepresenationCopy(t *testing.T) {
	auditEventResourceRepresentation := map[string]interface{}{
		"end_time": Representation{RepType: Required, Create: `create`},
	}
	type args struct {
		propertyNames   []string
		newValues       []interface{}
		representations map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Test positive case",
			args: args{
				propertyNames:   []string{"end_time"},
				newValues:       []interface{}{"new"},
				representations: auditEventResourceRepresentation,
			},
			want: map[string]interface{}{"end_time": "new"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMultipleUpdatedRepresenationCopy(tt.args.propertyNames, tt.args.newValues, tt.args.representations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMultipleUpdatedRepresenationCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitGetRepresentationCopyWithMultipleRemovedProperties(t *testing.T) {
	auditEventResourceRepresentation := map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"end_time":       Representation{RepType: Required, Create: `create`},
	}
	type args struct {
		propertyNames  []string
		representation map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Test with possible case",
			args: args{
				propertyNames:  []string{"compartment_id"},
				representation: auditEventResourceRepresentation,
			},
			want: map[string]interface{}{"end_time": Representation{RepType: Required, Create: `create`}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRepresentationCopyWithMultipleRemovedProperties(tt.args.propertyNames, tt.args.representation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepresentationCopyWithMultipleRemovedProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitGetTestClients(t *testing.T) {
	r := &schema.Resource{
		Schema: provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.Set("auth", globalvar.AuthAPIKeySetting)
	type args struct {
		data *schema.ResourceData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test positive case",
			args: args{
				data: d,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := GetTestClients(tt.args.data)
			assert.NotEmpty(t, client)
		})
	}
}

func TestUnitGetUpdatedRepresentationCopy(t *testing.T) {
	auditEventResourceRepresentation := map[string]interface{}{
		"end_time": Representation{RepType: Required, Create: `create`},
	}
	type args struct {
		propertyNameStr string
		newValue        interface{}
		representations map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Test positive case",
			args: args{
				propertyNameStr: "end_time",
				newValue:        "new",
				representations: auditEventResourceRepresentation,
			},
			want: map[string]interface{}{"end_time": "new"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUpdatedRepresentationCopy(tt.args.propertyNameStr, tt.args.newValue, tt.args.representations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpdatedRepresentationCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitPreCheck(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test positive case",
			args: args{
				t: t,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreCheck(t)
		})
	}
}

func TestUnitProviderConfigTest(t *testing.T) {
	ProviderConfigTest(t, true, true, globalvar.AuthAPIKeySetting, "", nil)
}

func TestUnitProviderTestConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test positive case",
			want: `
	# Need to have this block even though it's empty; for import testing
	provider "oci" {
	}
	` + CommonTestVariables(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProviderTestConfig(); got != tt.want {
				t.Errorf("ProviderTestConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitProviderTestCopy(t *testing.T) {
	type ConfigFunc func(d *schema.ResourceData) (interface{}, error)
	type args struct {
		configfn schema.ConfigureFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test positive case",
			args: args{
				configfn: func(d *schema.ResourceData) (interface{}, error) {
					return GetTestClients(d), nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := ProviderTestCopy(tt.args.configfn)
			assert.NotEmpty(t, client)
		})
	}
}

func TestUnitRepresentationCopyWithNewProperties(t *testing.T) {
	type args struct {
		representations map[string]interface{}
		newProperties   map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepresentationCopyWithNewProperties(tt.args.representations, tt.args.newProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepresentationCopyWithNewProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitRepresentationCopyWithRemovedNestedProperties(t *testing.T) {
	type args struct {
		propertyNameStr string
		representation  map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepresentationCopyWithRemovedNestedProperties(tt.args.propertyNameStr, tt.args.representation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepresentationCopyWithRemovedNestedProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitRepresentationCopyWithRemovedProperties(t *testing.T) {
	type args struct {
		representations   map[string]interface{}
		removedProperties []string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepresentationCopyWithRemovedProperties(tt.args.representations, tt.args.removedProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepresentationCopyWithRemovedProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitResourceTest(t *testing.T) {
	type args struct {
		t                *testing.T
		checkDestroyFunc resource.TestCheckFunc
		steps            []resource.TestStep
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUnitSaveConfigContent(t *testing.T) {
	type args struct {
		content  string
		service  string
		resource string
		t        *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUnitTokenizeWithHttpReplay(t *testing.T) {
	type args struct {
		defaultString string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 TokenFn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := TokenizeWithHttpReplay(tt.args.defaultString)
			if got != tt.want {
				t.Errorf("TokenizeWithHttpReplay() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TokenizeWithHttpReplay() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

//func TestUnitWaitTillCondition(t *testing.T) {
//	type args struct {
//		testAccProvider        *schema.Provider
//		resourceId             *string
//		shouldWait             ShouldWaitFunc
//		timeout                time.Duration
//		fetchOperationFunc     FetchOperationFunc
//		service                string
//		disableNotFoundRetries bool
//	}
//	tests := []struct {
//		name string
//		args args
//		want func()
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := WaitTillCondition(tt.args.testAccProvider, tt.args.resourceId, tt.args.shouldWait, tt.args.timeout, tt.args.fetchOperationFunc, tt.args.service, tt.args.disableNotFoundRetries); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("WaitTillCondition() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestUnitWriteToFile(t *testing.T) {
	type args struct {
		content  string
		service  string
		resource string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteToFile(tt.args.content, tt.args.service, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnit_apply(t *testing.T) {
	type args struct {
		template string
		values   map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apply(tt.args.template, tt.args.values); got != tt.want {
				t.Errorf("apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_cloneRepresentation(t *testing.T) {
	type args struct {
		representations map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneRepresentation(tt.args.representations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_setEnvSetting(t *testing.T) {
	type args struct {
		s string
		v string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setEnvSetting(tt.args.s, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("setEnvSetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnit_updateNestedRepresentation(t *testing.T) {
	type args struct {
		currIndex       int
		propertyNames   []string
		newValue        interface{}
		representations map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateNestedRepresentation(tt.args.currIndex, tt.args.propertyNames, tt.args.newValue, tt.args.representations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateNestedRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_updateNestedRepresentationRemoveProperty(t *testing.T) {
	type args struct {
		currIndex      int
		propertyNames  []string
		representation map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateNestedRepresentationRemoveProperty(tt.args.currIndex, tt.args.propertyNames, tt.args.representation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateNestedRepresentationRemoveProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
