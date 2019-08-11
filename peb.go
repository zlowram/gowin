package gowin

import (
	"path/filepath"
	"strings"
	"unsafe"
)

func GetPeb() *Peb {
	return (*Peb)(unsafe.Pointer(uintptr(getPebPointer())))
}

func (p *Peb) GetLoadedModules() map[string]*Module {
	modules := make(map[string]*Module)

	peb := GetPeb()
	moduleList := peb.Ldr.InMemoryOrderModuleList
	firstModule := moduleList
	for moduleList.Flink != firstModule.Blink {
		moduleList = moduleList.Flink
		ldrDataTableEntry := getLdrDataTableEntry(moduleList)
		fullDllName := ldrDataTableEntry.FullDllName.String()
		dllName := strings.ToLower(filepath.Base(fullDllName))
		dllBase := ldrDataTableEntry.DllBase
		modules[dllName] = NewModule(
			dllName,
			uint64(uintptr(unsafe.Pointer(dllBase))),
		)
	}

	return modules
}

func getLdrDataTableEntry(moduleListEntry *ListEntry) *LdrDataTableEntry {
	moduleListEntryPointer := uintptr(unsafe.Pointer(moduleListEntry))
	return (*LdrDataTableEntry)(unsafe.Pointer(moduleListEntryPointer - uintptr(0x10)))
}
