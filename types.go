package gowin

// PROCESSINFOCLASS Enumeration
// Used in NTQueryInformationProcess
// Add docs.
const (
	ProcessBasicInformation     = 0
	ProcessDebugPort            = 7
	ProcessWow64Information     = 26
	ProcessImageFileName        = 27
	ProcessBreakOnTermination   = 29
	ProcessSubsystemInformation = 75
)

// Add docs.
type SystemInfo struct {
	ProcessorArchitecture     uint16
	Reserved                  uint16
	PageSize                  uint32
	MinimumApplicationAddress uintptr
	MaximumApplicationAddress uintptr
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

// Add docs.
type MemoryStatusEx struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

// Add docs.
type UnicodeString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

// Add docs.
type PebLdrData struct {
	Reserved1               [8]byte
	Reserved2               [3]*uint64
	InMemoryOrderModuleList *ListEntry
}

// Add docs.
type ListEntry struct {
	Flink *ListEntry
	Blink *ListEntry
}

// Add docs.
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

// Add docs.
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
