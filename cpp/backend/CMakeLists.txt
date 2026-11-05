cmake_minimum_required(VERSION 3.15)
project(ServerCpp)

set(CMAKE_CXX_STANDARD 17)
set(SOURCE_FILES src/main.cpp)
set(SOURCE_FILES src/server/Server.cpp)

add_executable(
    ${PROJECT_NAME}
    src/main.cpp
    src/server/src/Server.cpp
)