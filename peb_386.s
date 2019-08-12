#include "textflag.h"

TEXT ·pebPointer(SB), NOSPLIT, $0
	MOVQ 0x30(FS), AX
	MOVQ AX, ret+0(FP)
	RET
