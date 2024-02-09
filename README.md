# bethesda-modutils

Tools for working with Bethesda's ESM, ESP, and ESL file formats.

These file formats are used to store data, configuration values, and scripts for
Bethesda games, as well as mods for those games. These formats are based on the
NetImmerse NIF format, which has been in use since 2002. As a result, these tools
can be used with files for The Elder Scrolls, Fallout, and Starfield.

## To Do

- Add parsers for unknown Field types
- Add `espdiff` utility

## Library Usage

Every ESM, ESP, and ESL file used in Bethesda games has three key data structures:

- Groups, which are collections of Records
- Records, which are ccollections of Fields
- Fields, which contain arbitrary data based on the Field type

These types are modeled in `group.go`, `record.go`, and `field.go`, respectively.

At the top level, each mod file contains a metadata Record, and a collection of Groups.
This structure is modeled in `mod_file.go`.
See [Reference Documentation](#reference-documentation) below for more detail.

To parse and load a mod file, use `modutils.LoadModFile()`:
```go
mod, err := modutils.LoadModFile(pathToModFile)
if err != nil {
	// Handle error.
}
r := mod.Metadata // Access the metadata record.
g :- mod.Groups // Access the top-level groups.
```

You will then have access to all Groups, Records, and Fields within the mod.

## Command-Line Tools

There are also two command-line tools in this repo:

- `espcat`, for printing a mod's metadata and field values
- `espdiff`, for comparing two mod files (TBD)

## Reference Documentation

- [UESP: Mod File Format](https://en.uesp.net/wiki/Skyrim_Mod:Mod_File_Format)
- [UESP: Mod File Format/TES4](https://en.uesp.net/wiki/Skyrim_Mod:Mod_File_Format/TES4)
