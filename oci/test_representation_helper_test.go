// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"strings"
	"testing"

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
	string_create_only_property = "Create"
	string_property = "Create"
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
			string_property = "Create"
		}
		string_property = "Create"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
		string_property = "Create"
	}
	string_create_only_property = "Create"
	string_property = "Create"
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
	string_create_only_property = "Create"
	string_property = "Update"
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
			string_property = "Update"
		}
		string_property = "Update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "Update"
	}
	string_create_only_property = "Create"
	string_property = "Update"
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
			string_property = "Create"
		}
		string_property = "Create"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "create1"
			"map_property2" = "create2"
		}
	}
	string_create_only_property = "Create"
	string_property = "Create"
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
		string_property = "Update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "Update"
	}
	string_create_only_property = "Create"
	string_property = "Update"
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
		string_property = "Update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "Update"
	}
	string_create_only_property = "Create"
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
			string_property = "Update"
		}
		string_property = "Update"
	}
	nested_property2 {
		array_create_only_property = ["create1", "create2"]
		map_property = {
			"map_property1" = "update1"
			"map_property2" = "update2"
		}
		string_property = "Update"
	}
	string_create_only_property = "Create"
	string_property = "Update"
	}
`
)

// issue-routing-tag: terraform/default
func TestUnitGenerateResourceRepresentationFromMap(t *testing.T) {
	assert := assert.New(t)

	nested2Map := map[string]interface{}{
		"string_property":          Representation{RepType: Required, Create: "Create", Update: "Update"},
		"array_property":           Representation{RepType: Required, Create: []string{"create1", "create2"}, Update: []string{"update1", "update2", "update3"}},
		"map_create_only_property": Representation{RepType: Optional, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}},
	}

	nestedMap1 := map[string]interface{}{
		"string_property":        Representation{RepType: Required, Create: "Create", Update: "Update"},
		"array_property":         Representation{RepType: Required, Create: []string{"create1", "create2"}, Update: []string{"update1", "update2"}},
		"map_property":           Representation{RepType: Required, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, Update: map[string]string{"map_property1": "update1", "map_property2": "update2"}},
		"nested_nested_property": RepresentationGroup{Required, nested2Map},
	}

	nestedMap2 := map[string]interface{}{
		"string_property":            Representation{RepType: Optional, Create: "Create", Update: "Update"},
		"array_create_only_property": Representation{RepType: Required, Create: []string{"create1", "create2"}},
		"map_property":               Representation{RepType: Required, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, Update: map[string]string{"map_property1": "update1", "map_property2": "update2"}},
	}

	testMap := map[string]interface{}{
		"string_property":             Representation{RepType: Required, Create: "Create", Update: "Update"},
		"string_create_only_property": Representation{RepType: Required, Create: "Create"},
		"array_property":              Representation{RepType: Required, Create: []string{"create1", "create2"}, Update: []string{"update1", "update2"}},
		"map_property":                Representation{RepType: Required, Create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, Update: map[string]string{"map_property1": "update1"}},
		"nested_property1":            RepresentationGroup{Optional, nestedMap1},
		"nested_property2":            RepresentationGroup{Required, nestedMap2},
	}

	assert.Equal(strings.Replace(requiredCreateConfig, "\t", "", -1),
		generateResourceFromMap(Required, Create, testMap), `"Required properties with Create values" Representation is wrong`)
	assert.Equal(strings.Replace(allCreateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Create, testMap), `"All properties with Create values" Representation is wrong`)
	assert.Equal(strings.Replace(requiredUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Required, Update, testMap), `"Required properties with Update values" Representation is wrong`)
	assert.Equal(strings.Replace(allUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update, testMap), `"All properties with Update values" Representation is wrong`)
	//make nested_property1 Required, will add nested_nested_property to Required Representation
	assert.Equal(strings.Replace(updatedGroupRequiredCreateConfig, "\t", "", -1),
		generateResourceFromMap(Required, Create, GetUpdatedRepresentationCopy("nested_property1", RepresentationGroup{Required, nestedMap1}, testMap)),
		`"Updated Required properties with Create values" Representation is wrong`)
	//change the value for the nested_nested_property in the representation
	assert.Equal(strings.Replace(updatedValueAllUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update, GetUpdatedRepresentationCopy("nested_property1.nested_nested_property.string_property", Representation{RepType: Required, Create: "updated_by_changes_in_the_representation"}, testMap)),
		`"Updated All properties with Update values" Representation is wrong`)
	//Update multiple values in the representation
	assert.Equal(strings.Replace(updatedValueMultipleUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update,
			GetMultipleUpdatedRepresenationCopy(
				[]string{"string_property", "nested_property1.nested_nested_property.string_property"},
				[]interface{}{Representation{RepType: Required, Create: "updated_create", Update: "updated_update"}, Representation{RepType: Required, Create: "re_updated_by_changes_in_the_representation"}},
				testMap)),
		`"Updated Multiple properties with Update values" Representation is wrong`)
	//add new properties to the representation
	assert.Equal(strings.Replace(allUpdateConfigWithAdditions, "\t", "", -1),
		generateResourceFromMap(Optional, Update, RepresentationCopyWithNewProperties(testMap, map[string]interface{}{
			"added_property":         Representation{RepType: Required, Create: "added"},
			"another_added_property": Representation{RepType: Optional, Create: "added", Update: "addedUpdate"},
		})),
		"Adding new properties to representation is wrong")
	//verify that the representation is not changed after the updates
	assert.Equal(strings.Replace(allUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update, testMap), `"All properties with Update values" Representation is wrong after the updates for the map`)

}
