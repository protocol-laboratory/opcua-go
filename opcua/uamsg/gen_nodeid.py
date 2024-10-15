import csv
import os


def camel_to_snake(name):
    import re
    s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
    return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()


def generate_nodeid_code_by_type(file_name):
    go_code_by_type = {}

    with open(file_name, newline='') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            param1, param2, param3 = row[0], int(row[1]), row[2]

            if param2 < 256:
                param_type = "TwoByte"
            else:
                param_type = "FourByte"

            file_type = camel_to_snake(param3)

            if file_type not in go_code_by_type:
                go_code_by_type[file_type] = f'''package uamsg

// https://github.com/OPCFoundation/UA-Nodeset/blob/latest/Schema/NodeIds.csv
var (
'''

            go_code_by_type[file_type] += f"    {param3}{param1} NodeId = NodeId{{{param_type}, 0, uint16({param2})}}\n"

    for file_type in go_code_by_type:
        go_code_by_type[file_type] += ")\n"

    return go_code_by_type


def save_go_code_by_type(go_code_by_type):
    if not os.path.exists("generated_nodeids"):
        os.makedirs("generated_nodeids")

    for file_type, go_code in go_code_by_type.items():
        output_file = f'std_nodeids_{file_type}.go'
        with open(output_file, 'w') as file:
            file.write(go_code)
        print(f"Go code has been saved to {output_file}")


csv_file = 'NodeIds.csv'
go_code_by_type = generate_nodeid_code_by_type(csv_file)

save_go_code_by_type(go_code_by_type)
