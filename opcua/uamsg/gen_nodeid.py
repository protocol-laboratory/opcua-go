import csv


def generate_nodeid_code(file_name):
    go_code = '''package uamsg

// https://github.com/OPCFoundation/UA-Nodeset/blob/latest/Schema/NodeIds.csv

var (
'''

    with open(file_name, newline='') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            param1, param2, param3 = row[0], int(row[1]), row[2]
            if param2 < 256:
                param_type = "TwoByte"
            else:
                param_type = "FourByte"

            go_code += f"    {param3}{param1} NodeId = NodeId{{{param_type}, 0, uint16({param2})}}\n"

    go_code += ")\n"
    return go_code


def save_go_code(go_code, output_file):
    with open(output_file, 'w') as file:
        file.write(go_code)
    print(f"Go code has been saved to {output_file}")


# https://github.com/OPCFoundation/UA-Nodeset/blob/latest/Schema/NodeIds.csv
csv_file = 'NodeIds.csv'
go_code = generate_nodeid_code(csv_file)

output_file = 'std_nodeids.go'
save_go_code(go_code, output_file)
