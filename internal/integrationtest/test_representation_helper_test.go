// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/stretchr/testify/assert"
)

const (
	requiredCreateConfig = `{
	array_property = ["create1", "create2"]
	map_property = {
		"map_property1" = "create1"
		"map_property2" = "create2"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
	}
	string_create_only_property = "create"
	string_property = "create"
	}
`

	allCreateConfig = `{
	array_property = ["create1", "create2"]
	map_property = {
		"map_property1" = "create1"
		"map_property2" = "create2"
	}
	nested_property1 {
		array_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
		nested_nested_property {
			array_property = ["create1", "create2"]
			map_create_only_property = {
				"map_property1" = "create1"
				"map_property2" = "create2"
			}
			string_property = "create"
		}
		string_property = "create"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
		string_property = "create"
	}
	string_create_only_property = "create"
	string_property = "create"
	}
`
	requiredUpdateConfig = `{
	array_property = ["update1", "update2"]
	map_property = {
		"map_property1" = "update1"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
	}
	string_create_only_property = "create"
	string_property = "update"
	}
`
	allUpdateConfig = `{
	array_property = ["update1", "update2"]
	map_property = {
		"map_property1" = "update1"
	}
	nested_property1 {
		array_property = ["update1", "update2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		nested_nested_property {
			array_property = ["update1", "update2", "update3"]
			map_create_only_property = {
				"map_property1" = "create1"
				"map_property2" = "create2"
			}
			string_property = "update"
		}
		string_property = "update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "update"
	}
	string_create_only_property = "create"
	string_property = "update"
	}
`
	updatedGroupRequiredCreateConfig = `{
	array_property = ["create1", "create2"]
	map_property = {
		"map_property1" = "create1"
		"map_property2" = "create2"
	}
	nested_property1 {
		array_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
		nested_nested_property {
			array_property = ["create1", "create2"]
			string_property = "create"
		}
		string_property = "create"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
	}
	string_create_only_property = "create"
	string_property = "create"
	}
`
	updatedValueAllUpdateConfig = `{
	array_property = ["update1", "update2"]
	map_property = {
		"map_property1" = "update1"
	}
	nested_property1 {
		array_property = ["update1", "update2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		nested_nested_property {
			array_property = ["update1", "update2", "update3"]
			map_create_only_property = {
				"map_property1" = "create1"
				"map_property2" = "create2"
			}
			string_property = "updated_by_changes_in_the_representation"
		}
		string_property = "update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "update"
	}
	string_create_only_property = "create"
	string_property = "update"
	}
`

	updatedValueMultipleUpdateConfig = `{
	array_property = ["update1", "update2"]
	map_property = {
		"map_property1" = "update1"
	}
	nested_property1 {
		array_property = ["update1", "update2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		nested_nested_property {
			array_property = ["update1", "update2", "update3"]
			map_create_only_property = {
				"map_property1" = "create1"
				"map_property2" = "create2"
			}
			string_property = "re_updated_by_changes_in_the_representation"
		}
		string_property = "update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "update"
	}
	string_create_only_property = "create"
	string_property = "updated_update"
	}
`

	allUpdateConfigWithAdditions = `{
	added_property = "added"
	another_added_property = "addedUpdate"
	array_property = ["update1", "update2"]
	map_property = {
		"map_property1" = "update1"
	}
	nested_property1 {
		array_property = ["update1", "update2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		nested_nested_property {
			array_property = ["update1", "update2", "update3"]
			map_create_only_property = {
				"map_property1" = "create1"
				"map_property2" = "create2"
			}
			string_property = "update"
		}
		string_property = "update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "update"
	}
	string_create_only_property = "create"
	string_property = "update"
	}
`
)

// issue-routing-tag: terraform/default
func TestUnitGenerateResourceRepresentationFromMap(t *testing.T) {
	assert := assert.New(t)

	nested2Map := map[string]interface{}{
		"string_property":          acctest.Representation{RepType: acctest.Required, Create: "create", Update: "update"},
		"array_property":           acctest.Representation{RepType: acctest.Required, Create: []string{"create1", "create2"}, Update: []string{"update1", "update2", "update3"}},
		"map_create_only_property": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}},
	}

	nestedMap1 := map[string]interface{}{
		"string_property":        acctest.Representation{RepType: acctest.Required, Create: "create", Update: "update"},
		"array_property":         acctest.Representation{RepType: acctest.Required, Create: []string{"create1", "create2"}, Update: []string{"update1", "update2"}},
		"map_property":           acctest.Representation{RepType: acctest.Required, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, Update: map[string]string{"map_property1": "update1", "map_property2": "update2"}},
		"nested_nested_property": acctest.RepresentationGroup{RepType: acctest.Required, Group: nested2Map},
	}

	nestedMap2 := map[string]interface{}{
		"string_property":            acctest.Representation{RepType: acctest.Optional, Create: "create", Update: "update"},
		"array_create_only_property": acctest.Representation{RepType: acctest.Required, Create: []string{"create1", "create2"}},
		"map_property":               acctest.Representation{RepType: acctest.Required, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, Update: map[string]string{"map_property1": "update1", "map_property2": "update2"}},
	}

	testMap := map[string]interface{}{
		"string_property":             acctest.Representation{RepType: acctest.Required, Create: "create", Update: "update"},
		"string_create_only_property": acctest.Representation{RepType: acctest.Required, Create: "create"},
		"array_property":              acctest.Representation{RepType: acctest.Required, Create: []string{"create1", "create2"}, Update: []string{"update1", "update2"}},
		"map_property":                acctest.Representation{RepType: acctest.Required, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, Update: map[string]string{"map_property1": "update1"}},
		"nested_property1":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: nestedMap1},
		"nested_property2":            acctest.RepresentationGroup{RepType: acctest.Required, Group: nestedMap2},
	}

	assert.Equal(strings.Replace(requiredCreateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Required, acctest.Create, testMap), `"Required properties with Create values" Representation is wrong`)
	assert.Equal(strings.Replace(allCreateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Optional, acctest.Create, testMap), `"All properties with Create values" Representation is wrong`)
	assert.Equal(strings.Replace(requiredUpdateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Required, acctest.Update, testMap), `"Required properties with Update values" Representation is wrong`)
	assert.Equal(strings.Replace(allUpdateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Optional, acctest.Update, testMap), `"All properties with Update values" Representation is wrong`)
	//make nested_property1 Required, will add nested_nested_property to Required Representation
	assert.Equal(strings.Replace(updatedGroupRequiredCreateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("nested_property1", acctest.RepresentationGroup{RepType: acctest.Required, Group: nestedMap1}, testMap)),
		`"Updated Required properties with Create values" Representation is wrong`)
	//change the value for the nested_nested_property in the representation
	assert.Equal(strings.Replace(updatedValueAllUpdateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy("nested_property1.nested_nested_property.string_property", acctest.Representation{RepType: acctest.Required, Create: "updated_by_changes_in_the_representation"}, testMap)),
		`"Updated All properties with Update values" Representation is wrong`)
	//Update multiple values in the representation
	assert.Equal(strings.Replace(updatedValueMultipleUpdateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Optional, acctest.Update,
			acctest.GetMultipleUpdatedRepresenationCopy(
				[]string{"string_property", "nested_property1.nested_nested_property.string_property"},
				[]interface{}{acctest.Representation{RepType: acctest.Required, Create: "updated_create", Update: "updated_update"}, acctest.Representation{RepType: acctest.Required, Create: "re_updated_by_changes_in_the_representation"}},
				testMap)),
		`"Updated Multiple properties with Update values" Representation is wrong`)
	//add new properties to the representation
	assert.Equal(strings.Replace(allUpdateConfigWithAdditions, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(testMap, map[string]interface{}{
			"added_property":         acctest.Representation{RepType: acctest.Required, Create: "added"},
			"another_added_property": acctest.Representation{RepType: acctest.Optional, Create: "added", Update: "addedUpdate"},
		})),
		"Adding new properties to representation is wrong")
	//verify that the representation is not changed after the updates
	assert.Equal(strings.Replace(allUpdateConfig, "\t", "", -1),
		acctest.GenerateResourceFromMap(acctest.Optional, acctest.Update, testMap), `"All properties with Update values" Representation is wrong after the updates for the map`)

}
