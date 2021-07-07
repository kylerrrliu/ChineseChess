#include <iostream>

#include "ailib.h"
#include "ailib_impl.hpp"

void* LIB_NewCNChessAI(int smartness) {
    auto lib = new CNChessAI(smartness);
    return lib;
}

CNChessAI* AsCNChessAI(void* cnChessAI) {
    return reinterpret_cast<CNChessAI*>(cnChessAI);
}

int LIB_Plus(void* cnChessAI, int value) {
    return value + AsCNChessAI(cnChessAI)->_smartness;
}