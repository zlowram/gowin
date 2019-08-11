package gowin

// https://docs.microsoft.com/en-us/windows/win32/api/winternl/ns-winternl-_peb
type Peb struct {
	Reserved1              [2]byte
	BeingDebugged          byte
	Reserved2              [1]byte
	Reserved3              [2]*uint64
	Ldr                    *PebLdrData
	ProcessParameters      *uint64 // PRTL_USER_PROCESS_PARAMETERS
	Reserved4              [3]*uint64
	AtlThunkSListPtr       *uint64
	Reserved5              *uint64
	Reserved6              uint64
	Reserved7              *uint64
	Reserved8              uint64
	AtlThunkSListPtr32     uint64
	Reserved9              [45]*uint64
	Reserved10             [96]byte
	PostProcessInitRoutine *uint64 // PPS_POST_PROCESS_INIT_ROUTINE
	Reserved11             [128]byte
	Reserved12             *uint64
	SessionId              uint64
}
