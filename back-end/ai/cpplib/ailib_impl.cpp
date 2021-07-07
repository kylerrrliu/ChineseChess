#include "ailib_impl.hpp"
#include <iostream>

CNChessAI::CNChessAI(int smartness) : _smartness(smartness) {
    std::cout<<"Setting smartness as "<< smartness <<std::endl;
}

int CNChessAI::Plus(int value) {
    return value + _smartness;
}