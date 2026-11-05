cmake_minimum_required(VERSION 3.15)
project(AllegroTest VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

add_executable(AllegroTest src/main.cpp)

find_package(allegro 5 REQUIRED)
find_package(allegro_font 5 REQUIRED)
find_package(allegro_image 5 REQUIRED)
find_package(allegro_primitives 5 REQUIRED)
find_package(allegro_audio 5 REQUIRED)
find_package(allegro_acodec 5 REQUIRED)

target_link_libraries(${PROJECT_NAME} PRIVATE
	allegro::allegro
	allegro::allegro_font
	allegro::allegro_image
	allegro::allegro_primitives
	allegro::allegro_audio
	allegro::allegro_acodec
)

target_include_directories(AllegroTest PRIVATE
	"${CMAKE_CURRENT_SOURCE_DIR}/src"
)
