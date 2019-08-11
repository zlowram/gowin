#include "textflag.h"

TEXT ·getPebPointer(SB), NOSPLIT, $0
	MOVQ 0x30(FS), AX
	MOVQ AX, ret+0(FP)
	RET
