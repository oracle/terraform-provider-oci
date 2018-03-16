#!/usr/bin/env python
# This script takes the Core Services API Spec and removes the refs that aren't
# valid swagger.
#
# IMPORTANT NOTE: This script will not do the right thing if there are
# extra blank characters at the ends of lines, or if unicode characters
# (like "smart quotes") are embedded in the descriptions.

import argparse
from collections import OrderedDict
import copy
import yaml

INTERNAL_ONLY = "x-obmcs-internal-only"

# For a prettier spec, we want to dump strings differently depending
# on whether they contain newline characters. This makes a new type to
# wrap strings we want to print as literals, i.e., with the leading |
# symbol
class literal_str(str):
    pass


def represent_literal_str(dumper, data):
    scalar = yaml.representer.SafeRepresenter.represent_str(dumper, data)
    scalar.style = '|'
    return scalar


yaml.add_representer(literal_str, represent_literal_str)


# For human readability we want to output the yaml file in a fixed
# order.
def represent_ordereddict(dumper, data):
    value = []

    for item_key, item_value in data.items():
        node_key = dumper.represent_data(item_key)
        node_value = dumper.represent_data(item_value)

        value.append((node_key, node_value))

    return yaml.nodes.MappingNode(u'tag:yaml.org,2002:map', value)


yaml.add_representer(OrderedDict, represent_ordereddict)


# Any string that has newlines in it needs to output as a literal for
# HTML codegen to work. Theoretically this should be more complicated
# to deal with lists also, but our specs don't have lists that have
# strings we need to convert.
def convert_to_literal_strs(o):
    if isinstance(o, dict):
        for key in o:
            if isinstance(o[key], str) and '\n' in o[key]:
                o[key] = literal_str(o[key])
            elif isinstance(o[key], dict):
                convert_to_literal_strs(o[key])


# Process command line arguments
def process_args():
    parser = argparse.ArgumentParser(description='Extract swagger spec.')
    parser.add_argument('input_file', type=argparse.FileType('r'))
    parser.add_argument('--output_file', nargs='?',
                        type=argparse.FileType('w'),
                        default=None)
    parser.add_argument('--external', help='Generate External spec', action='store_true')
    return parser.parse_args()


def remove_x_descriptions(source_yaml, d):
    if isinstance(d, dict):
        if '$ref' in d and d['$ref'][:17] == '#/x-descriptions/':
            ref = d['$ref'][17:]
            for k in list(source_yaml['x-descriptions'][ref]):
                d[k] = source_yaml['x-descriptions'][ref][k]
            del d['$ref']
        for k in list(d):
            remove_x_descriptions(source_yaml, d[k])
    elif isinstance(d, list):
        for item in d:
            remove_x_descriptions(source_yaml, item)


def remove_nested_error_response_refs(source_yaml):
    for k in list(source_yaml['responses']):
        if '$ref' in source_yaml['responses'][k]:
            if source_yaml['responses'][k]['$ref'] == '#/responses/Error':
                for error_key in source_yaml['responses']['Error']:
                    source_yaml['responses'][k][error_key] = copy.deepcopy(source_yaml['responses']['Error'][error_key])
                del source_yaml['responses'][k]['$ref']
    del source_yaml['responses']['Error']


def remove_internal_params_from_paths(source_yaml):
    for k in list(source_yaml['paths']):
        for method in source_yaml['paths'][k]:
            if 'parameters' in source_yaml['paths'][k][method]:
                source_yaml['paths'][k][method]['parameters'] = [x for x in source_yaml['paths'][k][method]['parameters'] if
                                                                 not ('$ref' in x and x['$ref'][:13] == '#/parameters/' and
                                                                      INTERNAL_ONLY in source_yaml['parameters'][x['$ref'][13:]] and
                                                                      source_yaml['parameters'][x['$ref'][13:]][INTERNAL_ONLY])]


def remove_internal_params(source_yaml):
    source_yaml['parameters'] = {x: source_yaml['parameters'][x] for x in source_yaml['parameters'] if
                                 not (INTERNAL_ONLY in source_yaml['parameters'][x] and
                                      source_yaml['parameters'][x][INTERNAL_ONLY])}


def remove_internal_fields(source_yaml):
    for k in list(source_yaml['definitions']):
        if 'properties' not in source_yaml['definitions'][k]:
            continue
        source_yaml['definitions'][k]['properties'] = {
            x: source_yaml['definitions'][k]['properties'][x] for x in source_yaml['definitions'][k]['properties'] if
            not (INTERNAL_ONLY in source_yaml['definitions'][k]['properties'][x]
                 and source_yaml['definitions'][k]['properties'][x][INTERNAL_ONLY])}


def remove_internal_definitions(source_yaml):
    for defn in list(source_yaml['definitions']):
        if INTERNAL_ONLY in source_yaml['definitions'][defn]:
            del source_yaml['definitions'][defn]


def remove_internal_apis(source_yaml):
    for path in list(source_yaml['paths']):
        path_details = source_yaml['paths'][path]

        if INTERNAL_ONLY in path_details and path_details[INTERNAL_ONLY]:
            del source_yaml['paths'][path]
            continue

        for method in list(path_details):
            method_details = source_yaml['paths'][path][method]

            if INTERNAL_ONLY in method_details and method_details[INTERNAL_ONLY]:
                del source_yaml['paths'][path][method]
                continue


def process_yaml(source_yaml, output_file, external):
    # Build the output yaml in order
    ordered_output = OrderedDict()
    ordered_output['swagger'] = source_yaml['swagger']
    ordered_output['info'] = source_yaml['info']
    for field in ['schemes', 'consumes', 'produces', 'basePath']:
        if field in source_yaml:
            ordered_output[field] = source_yaml[field]

    remove_x_descriptions(source_yaml, source_yaml)
    remove_nested_error_response_refs(source_yaml)
    if (external):
        remove_internal_apis(source_yaml)
        remove_internal_definitions(source_yaml)
        remove_internal_params_from_paths(source_yaml)
        remove_internal_params(source_yaml)
        remove_internal_fields(source_yaml)

    ordered_output['paths'] = source_yaml['paths']
    ordered_output['definitions'] = source_yaml['definitions']
    ordered_output['parameters'] = source_yaml['parameters']
    ordered_output['responses'] = source_yaml['responses']

    convert_to_literal_strs(ordered_output)

    yaml.dump(ordered_output, output_file, default_flow_style=False)


if (__name__ == '__main__'):
    args = process_args()
    input_file = args.input_file
    output_file = args.output_file or file('output.yaml', 'w')
    external = args.external

    # Read yaml file
    try:
        source_yaml = yaml.load(input_file)
    except:
        print("ERROR: Failed to load/process {}".format(input_file))
        raise

    process_yaml(source_yaml, output_file, external)
