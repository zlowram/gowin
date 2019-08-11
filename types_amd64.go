package gowin

// https://docs.microsoft.com/en-us/windows/win32/api/winternl/ns-winternl-_peb#remarks
type Peb struct {
	Reserved1              [2]byte
	BeingDebugged          byte
	Reserved2              [21]byte
	Ldr                    *PebLdrData
	ProcessParameters      uint64
	Reserved3              [520]byte
	PostProcessInitRoutine uint64
	Reserved4              [136]byte
	SessionId              uint64
}
