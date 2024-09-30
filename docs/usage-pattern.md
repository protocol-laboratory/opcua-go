## Two Usage Methods for the `opcua-go` Library

The `opcua-go` library offers flexibility in how you manage OPC UA variables and address spaces, allowing integration in various system architectures. This article outlines two main usage methods for the library, providing clarity on when and how each approach can be applied.

### Method 1: `opcua-go` Maintains Variables and Address Space

Considering you have no existing system managing variables and data, or you want to build a standalone OPC UA server that handles all data and variables internally, this method is ideal.

In this approach, the `opcua-go` library itself is responsible for managing all OPC UA variables and the address space. This means that the variables, data, and node structure exist solely within the library. Any client or server communicating with this OPC UA instance will reference the variables and address space maintained by the library.

### Method 2: Variables and Address Space Managed by an External System

Considering you already have an existing system that manages variables and data, or you have an IoT platform, SCADA system, or other software that handles the data, you want to use this library to expose the data via OPC UA protocol.

In this scenario, the `opcua-go` library acts as a bridge between your existing system and the OPC UA clients.

This library may maintain variables and address space for the OPC UA server, but the data is synchronized with the external system(by notification or polling). The external system is the source of truth for the data, and the library ensures that the OPC UA server reflects the latest values and structure.
