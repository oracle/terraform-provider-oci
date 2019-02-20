// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

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

func TestGenerateResourceRepresentationFromMap(t *testing.T) {
	assert := assert.New(t)

	nested2Map := map[string]interface{}{
		"string_property":          Representation{repType: Required, create: "create", update: "update"},
		"array_property":           Representation{repType: Required, create: []string{"create1", "create2"}, update: []string{"update1", "update2", "update3"}},
		"map_create_only_property": Representation{repType: Optional, create: map[string]string{"map_property1": "create1", "map_property2": "create2"}},
	}

	nestedMap1 := map[string]interface{}{
		"string_property":        Representation{repType: Required, create: "create", update: "update"},
		"array_property":         Representation{repType: Required, create: []string{"create1", "create2"}, update: []string{"update1", "update2"}},
		"map_property":           Representation{repType: Required, create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, update: map[string]string{"map_property1": "update1", "map_property2": "update2"}},
		"nested_nested_property": RepresentationGroup{Required, nested2Map},
	}

	nestedMap2 := map[string]interface{}{
		"string_property":            Representation{repType: Optional, create: "create", update: "update"},
		"array_create_only_property": Representation{repType: Required, create: []string{"create1", "create2"}},
		"map_property":               Representation{repType: Required, create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, update: map[string]string{"map_property1": "update1", "map_property2": "update2"}},
	}

	testMap := map[string]interface{}{
		"string_property":             Representation{repType: Required, create: "create", update: "update"},
		"string_create_only_property": Representation{repType: Required, create: "create"},
		"array_property":              Representation{repType: Required, create: []string{"create1", "create2"}, update: []string{"update1", "update2"}},
		"map_property":                Representation{repType: Required, create: map[string]string{"map_property1": "create1", "map_property2": "create2"}, update: map[string]string{"map_property1": "update1"}},
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
		generateResourceFromMap(Required, Create, getUpdatedRepresentationCopy("nested_property1", RepresentationGroup{Required, nestedMap1}, testMap)),
		`"Updated Required properties with Create values" Representation is wrong`)
	//change the value for the nested_nested_property in the representation
	assert.Equal(strings.Replace(updatedValueAllUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update, getUpdatedRepresentationCopy("nested_property1.nested_nested_property.string_property", Representation{repType: Required, create: "updated_by_changes_in_the_representation"}, testMap)),
		`"Updated All properties with Update values" Representation is wrong`)
	//update multiple values in the representation
	assert.Equal(strings.Replace(updatedValueMultipleUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update,
			getMultipleUpdatedRepresenationCopy(
				[]string{"string_property", "nested_property1.nested_nested_property.string_property"},
				[]interface{}{Representation{repType: Required, create: "updated_create", update: "updated_update"}, Representation{repType: Required, create: "re_updated_by_changes_in_the_representation"}},
				testMap)),
		`"Updated Multiple properties with Update values" Representation is wrong`)
	//add new properties to the representation
	assert.Equal(strings.Replace(allUpdateConfigWithAdditions, "\t", "", -1),
		generateResourceFromMap(Optional, Update, representationCopyWithNewProperties(testMap, map[string]interface{}{
			"added_property":         Representation{repType: Required, create: "added"},
			"another_added_property": Representation{repType: Optional, create: "added", update: "addedUpdate"},
		})),
		"Adding new properties to representation is wrong")
	//verify that the representation is not changed after the updates
	assert.Equal(strings.Replace(allUpdateConfig, "\t", "", -1),
		generateResourceFromMap(Optional, Update, testMap), `"All properties with Update values" Representation is wrong after the updates for the map`)

}
