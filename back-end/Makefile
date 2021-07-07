build: ailib
	@echo "Building BE"
	@go build -o be/be_server be/main/main.go

vendor: go.mod go.sum ## pull the vendor pkgs for deps
	@GO111MODULE=on go mod vendor

run: vendor ## run the server from source code
	@go run be/main/main.go

ailib: be/ai/cpplib/ailib.cpp be/ai/cpplib/ailib_impl.cpp
	@echo "Building AI lib"
	@mkdir -p ~/lib
	@clang++ -o be/ai/libailib.so be/ai/cpplib/ailib.cpp be/ai/cpplib/ailib_impl.cpp \
    -std=c++17 -O3 -Wall -Wextra -fPIC -shared
	@cp be/ai/libailib.so ~/lib/