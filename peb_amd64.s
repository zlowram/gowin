#include "textflag.h"

TEXT ·getPebPointer(SB), NOSPLIT, $0
	MOVQ 0x60(GS), AX
	MOVQ AX, ret+0(FP)
	RET
