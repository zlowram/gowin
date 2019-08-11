package gowin

type UnicodeString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

type PebLdrData struct {
	Reserved1               [8]byte
	Reserved2               [3]*uint64
	InMemoryOrderModuleList *ListEntry
}

type ListEntry struct {
	Flink *ListEntry
	Blink *ListEntry
}

type LdrDataTableEntry struct {
	Reserved1          [2]*uint64
	InMemoryOrderLinks ListEntry
	Reserved2          [2]*uint64
	DllBase            *uint64
	EntryPoint         *uint64
	Reserved3          *uint64
	FullDllName        UnicodeString
	Reserved4          [8]byte
	Reserved5          [3]*uint64
	CheckSum           uint64
	Reserved6          *uint64
	TimeDateStamp      uint64
}

type ImageExportDirectory struct {
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32 // RVA
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32 // RVA
	AddressOfNames        uint32 // RVA
	AddressOfNameOrdinals uint32 // RVA
}
