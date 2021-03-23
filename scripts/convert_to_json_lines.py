#!/usr/bin/env python
import sys
try:
    import ujson as json
except ImportError:
    import json


SEP = '\x1b'


def convert_data_file(src_file, dst_file, columns):
    with open(src_file, 'r') as src_fp, open(dst_file, 'w') as dst_fp:
        for line in src_fp:
            fields = dict(zip(columns, line.strip().split(SEP)))
            dst_fp.write(json.dumps(fields))
            dst_fp.write('\n')


if __name__ == '__main__':
    args = sys.argv[1:]
    if len(args) != 3:
        sys.stderr.write('''Usage:
python convert_to_json_lines.py <source data file> <dest data file> <field1,field2,field3>
''')
        sys.exit(1)
    src_file, dst_file, columns = args
    convert_data_file(src_file, dst_file, columns.split(','))
