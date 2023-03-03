// Package ns implements NoxScript 4 used in OpenNox maps.
//
// NS4 is based on NoxScript 3 but provides a more Go-friendly object-oriented API.
//
// These changes include:
//
// - Most functions became methods on the corresponding object types.
// - Functions have a more generic signature, using interfaces where appropriate.
// - Similar methods were renamed, e.g. ObjectOn and WaypointOn is now called Enable for both types.
// - Type safety was improved. Package defines more names types for enums.
// - All enum constants were moved to their own sub-packages.
//
// Package is compatible with NS3: methods like AsObj convert NS3 objects into NS4 objects.
package ns
