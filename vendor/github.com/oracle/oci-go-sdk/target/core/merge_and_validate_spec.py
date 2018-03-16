#!/usr/bin/env python
# This script merges together partial swagger specs. It throws an error
# if multiple specs define the same object. It also replaces $ref values
# that are not valid swagger, but which are used in the proto specs to
# avoid duplication.

import argparse
import os
import yaml

from valid_swagger import process_yaml


# Process command line arguments
def process_args():
    parser = argparse.ArgumentParser(description='Extract swagger spec.')
    parser.add_argument('--input_directory', '-i')
    parser.add_argument('--output_file', '-o', nargs='?',
                        type=argparse.FileType('w'),
                        default=None)
    parser.add_argument('--external', help='Generate External spec', action='store_true')
    return parser.parse_args()


# Only merge one level deep
def merge(combined_yaml, source_yaml):
    for key in source_yaml:
        if key not in combined_yaml:
            combined_yaml[key] = source_yaml[key]
        elif isinstance(combined_yaml[key], dict) and isinstance(source_yaml[key], dict):
            for subkey in source_yaml[key]:
                if subkey in combined_yaml[key]:
                    raise ValueError('Subkey {} is included under {} in two files'.format(subkey, key))
                else:
                    combined_yaml[key][subkey] = source_yaml[key][subkey]
        else:
            raise ValueError('Key {} in two files and is not a dictionary'.format(key))

    return combined_yaml


if (__name__ == '__main__'):
    args = process_args()
    output_file = args.output_file or file('output.yaml', 'w')
    external = args.external

    combined_yaml = {}

    for input_filename in os.listdir(args.input_directory):
        input_filename = os.path.join(args.input_directory, input_filename)
        print('Processing {}'.format(input_filename))
        if (input_filename.endswith(".yaml")):
            with open(input_filename, 'r') as input_file:
                source_yaml = yaml.load(input_file)
                combined_yaml = merge(combined_yaml, source_yaml)

    process_yaml(combined_yaml, output_file, external)
