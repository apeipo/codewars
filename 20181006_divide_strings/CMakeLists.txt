cmake_minimum_required(VERSION 3.10)
project(ctest)

set(CMAKE_CXX_FLAGS " -std=c++11 -O3")

add_executable(divide_string main.cpp divide_strings.h)