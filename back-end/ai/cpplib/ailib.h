#pragma once
#ifdef __cplusplus
extern "C" {
#endif

void* LIB_NewCNChessAI(int smartness);
int LIB_Plus(void* cnChessAI, int value);

#ifdef __cplusplus
}  // extern "C"
#endif