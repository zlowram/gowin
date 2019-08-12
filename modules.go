package gowin

import (
	"debug/pe"
	"errors"
	"unsafe"

	"github.com/zlowram/memread"
)

type Module struct {
	Name string
	Addr uint64
}

type Export struct {
	Ordinal int
	Addr    uint64
}

func NewModule(name string, addr uint64) *Module {
	return &Module{name, addr}
}

func (m *Module) Export(name string) (export *Export, err error) {
	exports, err := m.Exports()
	if err != nil {
		return nil, err
	}
	export, ok := exports[name]
	if !ok {
		err = errors.New("Gowin.Module.Export: export not found")
	}
	return export, err
}

func (m *Module) Exports() (exports map[string]*Export, err error) {
	memoryReader := memread.NewReader(m.Addr)
	peFile, err := pe.NewFile(memoryReader)
	if err != nil {
		return nil, err
	}

	imageDirectoryRva := peFile.OptionalHeader.(*pe.OptionalHeader64).DataDirectory[pe.IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress
	imageExportDirectory := (*ImageExportDirectory)(unsafe.Pointer(uintptr(uint64(imageDirectoryRva) + m.Addr)))
	addressOfFunctions := (*uint32)(unsafe.Pointer(uintptr(uint64(imageExportDirectory.AddressOfFunctions) + m.Addr)))
	addressOfNames := (*uint32)(unsafe.Pointer(uintptr(uint64(imageExportDirectory.AddressOfNames) + m.Addr)))
	numberOfNames := int(imageExportDirectory.NumberOfNames)
	addressOfNameOrdinals := (*uint32)(unsafe.Pointer(uintptr(uint64(imageExportDirectory.AddressOfNameOrdinals) + m.Addr)))

	exports = make(map[string]*Export, numberOfNames)
	for i := 0; i < numberOfNames; i++ {
		nameStringRva := *(*uint32)(unsafe.Pointer((uintptr(unsafe.Pointer(addressOfNames)) + uintptr(i*4))))
		nameStringAddr := m.Addr + uint64(nameStringRva)
		reader := memread.NewReader(nameStringAddr)
		name := NewCString(reader)
		ordinal := *(*uint16)(unsafe.Pointer((uintptr(unsafe.Pointer(addressOfNameOrdinals)) + uintptr(i*2))))
		functionRva := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(addressOfFunctions)) + uintptr(ordinal*uint16(4))))
		functionAddr := m.Addr + uint64(functionRva)
		exports[name] = &Export{
			Ordinal: i + int(imageExportDirectory.Base),
			Addr:    functionAddr,
		}
	}

	return exports, nil
}
