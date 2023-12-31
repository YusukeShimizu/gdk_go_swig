package greenaddress

import (
	"log"
	"testing"
	"unsafe"
)

func TestGenerateMnemonic(t *testing.T) {
	mnemonic, err := GenerateMnemonic()
	if err != nil {
		return
	}
	log.Printf("mnemnonic: %s", mnemonic)

	var rptr uintptr
	Validate_mnemonic(mnemonic, SwigcptrUint32_t((uintptr)(unsafe.Pointer(&rptr))))
	r := (int)(SwigcptrUint32_t(rptr))
	if r != GA_TRUE {
		return
	}

	mnemonic12, err := GenerateMnemonic12()
	if err != nil {
		return
	}
	log.Printf("mnemnonic12: %s", mnemonic12)

	var rptr12 uintptr
	Validate_mnemonic(mnemonic12, SwigcptrUint32_t((uintptr)(unsafe.Pointer(&rptr12))))
	r12 := (int)(SwigcptrUint32_t(rptr12))
	if r12 != GA_TRUE {
		return
	}
}
