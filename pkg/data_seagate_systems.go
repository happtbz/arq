// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapseagate_systems = map[string]*DeviceData{
    "Exos CORVAULT 4U106": {
        Manufacturer: "Seagate Systems",
        Model: "Exos CORVAULT 4U106",
        Slug: "seagate-systems-exos-corvault-4u106",
        UHeight: 4,
        PartNumber: "R4106I2500T002",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 131.5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "PSU1", Label: "", Position: "1" },
            { Name: "PSU2", Label: "", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
            { Name: "A", Label: "" },
            { Name: "B", Label: "" },
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "Exos CORVAULT SAS Controller": {
        Manufacturer: "Seagate Systems",
        Model: "Exos CORVAULT SAS Controller",
        Slug: "seagate-systems-exos-corvault-sas-controller",
        UHeight: 0,
        PartNumber: "R4106I2500T002-controller",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "manufacturer-do-not-use", Type: "usb-micro-a", Label: "", Poe: false },
            { Name: "serial", Type: "usb-micro-a", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "SAS 0", Label: "", Type: "other", MgmtOnly: false },
            { Name: "SAS 1", Label: "", Type: "other", MgmtOnly: false },
            { Name: "SAS 2", Label: "", Type: "other", MgmtOnly: false },
            { Name: "SAS 3", Label: "", Type: "other", MgmtOnly: false },
            { Name: "management", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
