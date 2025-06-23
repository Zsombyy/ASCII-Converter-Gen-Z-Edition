#!/bin/bash
# Windows Cross Builder with Individual Zipping

# Script made by: Cubxy Development (Zsombyy)

APP_NAME="brainrot-ascii"
BUILD_DIR="builds"
ZIP_DIR="zips"

# Coloured outputs
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Function for coloured output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_warn() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed on your device or not in PATH. Please install Go before building."
    exit 1
fi

print_status "Starting Windows cross-building for Go Application: $APP_NAME"

# Create directories if they don't exist
mkdir -p "$BUILD_DIR"
mkdir -p "$ZIP_DIR"
print_status "Created directories: $BUILD_DIR and $ZIP_DIR"

declare -a architectures=("amd64" "386" "arm64")
declare -a arch_names=("64-bit" "32-bit" "ARM64")

print_status "Building for Windows platforms"
echo

# Build counter
successful_builds=0
total_builds=${#architectures[@]}

# Build for each architecture
for i in "${!architectures[@]}"; do
    arch=${architectures[$i]}
    arch_name=${arch_names[$i]}
    output_file="$BUILD_DIR/${APP_NAME}_windows_${arch}.exe"
    
    print_status "Building for Windows $arch_name ($arch)..."
    
    # Build the executable
    env GOOS=windows GOARCH=$arch go build -o "$output_file"
    
    # Check if the build was successful
    if [ $? -eq 0 ] && [ -f "$output_file" ]; then
        file_size=$(du -h "$output_file" | cut -f1)
        print_success "Built: $output_file ($file_size)"
        
        # Create individual zip for this executable
        zip_file="$ZIP_DIR/${APP_NAME}_windows_${arch}.zip"
        if command -v zip &>/dev/null; then
            (cd "$BUILD_DIR" && zip "../$zip_file" "${APP_NAME}_windows_${arch}.exe")
            if [ $? -eq 0 ]; then
                print_success "Created zip: $zip_file"
            else
                print_error "Failed to create zip for $output_file"
            fi
        else
            print_warn "zip command not found. Skipping zip creation."
        fi
        
        ((successful_builds++))
    else
        print_error "Failed to build for Windows $arch_name ($arch)"
    fi
done

echo
print_status "Build Summary:"
echo "---------------------"
print_status "Successful builds: $successful_builds/$total_builds"

# List all built executables and zips
if [ $successful_builds -gt 0 ]; then
    echo
    print_status "Built executables:"
    ls -la "$BUILD_DIR"/*.exe 2>/dev/null | while read line; do
        echo "  $line"
    done
    
    echo
    print_status "Created zip archives:"
    ls -la "$ZIP_DIR"/*.zip 2>/dev/null | while read line; do
        echo "  $line"
    done
fi

print_status "Build completed. Executables are in '$BUILD_DIR', zips are in '$ZIP_DIR'."
