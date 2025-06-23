#!/bin/bash
# Windows Cross Builder

# Script made by: Cubxy Development (Zsombyy)

APP_NAME="brainrot-ascii"
BUILD_DIR="builds"

# Coloured outputs

RED='\033[0;31m'
GREEN='\033[0;32'
BLUE='\033[0;34'
YELLOW='\033[1;33'
NC='\033[0m'

# Function for coloured output

print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SECCUESS]${NC} $1"
}

print_error() {
    echo -e "${RED} [ERROR]${NC} $1"
}

print_warn() {
    echo -e "${YELLOW} [WARNING]${NC} $1"
}
# Check if Go is installed

if ! command -v go &> /dev/null; then
print_error "Go is not installed in or device or in the path please make sure install Go before build"
exit 1
fi

print_status "Starting Windows cross-building for Go Application: $APP_NAME"

# Create build directory if it dosen't exist
if [ ! -d "$BUILD_DIR" ]; then
mkdir -p "$BUILD_DIR"
print_status "Created build directory by the scirpt: $BUILD_DIR"
fi

declare -a architectures=("amd64" "368" "arm64")
declare -a arch_names=("64-bit" "32-bit" "ARM64")

print_status "Building for Windows platfroms"
echo

# Build counter
seccessful_builds=0
total_builds=${#architectures[@]}

# Build the executeable
env GOOS=windows GOARCH=$arch go build -o "$output_file"

# Check if the build was seccuessful

if [ $? -eq 0 ] && [ -f "$output_file"]; then
file_size$=(du -h "$output_file" | cut -fi)
print_success "Built: $output_file ($file_size)"
((seccessful_builds++))
else
print_error "Failed to build for Windows $arch_name ($arhc)"
fi
done
echo
print_status "Buid Summary:"
echo "---------------------"
print_status "Seccussful builds: $seccussful_builds/$total_builds"

# Listg all built executeables
if [ $seccessful_builds -gt 0 ]; then
echo
print_status "Built executeables:"
ls -la "$BUILD_DIR"/*.exe 2>/dev/null | while read line; do
echo "  $line"
done
fi

print_status "Built completed. Executables are in the '$UILD_DIR' directory."

# Optinal Create zip archives pf all executables
read -p "Do you wanto to create a zip archive of all executables? (y/N):" -r
if [[ $REPLY =~ ^[Yy]$ ]]; then
archive_name="$APP_NAME-windows-all.zip"
if command -v zip &>/dev/null; then;
cd "$BUILD_DIR" && zip "../$archive_name" *exe && cd ..
print_success "Created archive: $archive_name"
else
print_warn "zip command not found. Skipping archive creation."
fi
fi
