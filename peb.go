package gowin

import (
	"errors"
	"path/filepath"
	"strings"
	"unsafe"
)

func PebAddress() *Peb {
	return (*Peb)(unsafe.Pointer(uintptr(pebPointer())))
}

func (p *Peb) Module(name string) (module *Module, err error) {
	module, ok := p.Modules()[name]
	if !ok {
		return nil, errors.New("Gowin.Peb.Module: module is not loaded")
	}
	return module, nil
}

func (p *Peb) Modules() map[string]*Module {
	modules := make(map[string]*Module)

	moduleList := p.Ldr.InMemoryOrderModuleList
	firstModule := moduleList
	for moduleList.Flink != firstModule.Blink {
		moduleList = moduleList.Flink
		entry := ldrDataTableEntry(moduleList)
		fullDllName := entry.FullDllName.String()
		dllName := strings.ToLower(filepath.Base(fullDllName))
		dllBase := entry.DllBase
		modules[dllName] = NewModule(
			dllName,
			uint64(uintptr(unsafe.Pointer(dllBase))),
		)
	}

	return modules
}

func ldrDataTableEntry(moduleListEntry *ListEntry) *LdrDataTableEntry {
	moduleListEntryPointer := uintptr(unsafe.Pointer(moduleListEntry))
	return (*LdrDataTableEntry)(unsafe.Pointer(moduleListEntryPointer - uintptr(0x10)))
}
