#!/usr/bin/env python3
"""
Universal Dependencies Installer
Installs Rust and dependencies (tar, flate2) with distro family support
Supports: Arch, openSUSE, Debian, Fedora families
"""

import os
import sys
import subprocess
import tempfile
import shutil
from pathlib import Path
from typing import List, Optional, Dict, Tuple

class Colors:
    RED = '\033[0;31m'
    GREEN = '\033[0;32m'
    YELLOW = '\033[1;33m'
    BLUE = '\033[0;34m'
    CYAN = '\033[0;36m'
    BOLD = '\033[1m'
    NC = '\033[0m'  # No Color

class Logger:
    @staticmethod
    def info(msg: str):
        print(f"{Colors.BLUE}[INFO]{Colors.NC} {msg}")
    
    @staticmethod
    def success(msg: str):
        print(f"{Colors.GREEN}[SUCCESS]{Colors.NC} {msg}")
    
    @staticmethod
    def warning(msg: str):
        print(f"{Colors.YELLOW}[WARNING]{Colors.NC} {msg}")
    
    @staticmethod
    def error(msg: str):
        print(f"{Colors.RED}[ERROR]{Colors.NC} {msg}")
    
    @staticmethod
    def menu(msg: str):
        print(f"{Colors.CYAN}[MENU]{Colors.NC} {msg}")

class DistroDetector:
    DISTRO_FAMILIES = {
        'arch': {
            'distros': ['arch', 'manjaro', 'endeavouros', 'artix', 'garuda'],
            'package_manager': 'pacman',
            'dev_tools': ['base-devel'],
            'install_cmd': 'sudo pacman -Sy --noconfirm'
        },
        'opensuse': {
            'distros': ['opensuse', 'opensuse-leap', 'opensuse-tumbleweed', 'sled', 'sles'],
            'package_manager': 'zypper',
            'dev_tools': [],  # Special handling for patterns
            'install_cmd': 'sudo zypper install -y'
        },
        'debian': {
            'distros': ['debian', 'ubuntu', 'mint', 'elementary', 'kali', 'pop', 'zorin', 'deepin'],
            'package_manager': 'apt-get',
            'dev_tools': ['build-essential'],
            'install_cmd': 'sudo apt-get update && sudo apt-get install -y'
        },
        'fedora': {
            'distros': ['fedora', 'centos', 'rhel', 'rocky', 'almalinux', 'oracle'],
            'package_manager': 'dnf',
            'dev_tools': [],  # Special handling for groups
            'install_cmd': 'sudo dnf install -y'
        }
    }

    @staticmethod
    def command_exists(command: str) -> bool:
        """Check if a command exists in PATH"""
        return shutil.which(command) is not None

    @staticmethod
    def read_os_release() -> Dict[str, str]:
        """Read /etc/os-release file"""
        os_release = {}
        try:
            with open('/etc/os-release', 'r') as f:
                for line in f:
                    if '=' in line:
                        key, value = line.strip().split('=', 1)
                        os_release[key] = value.strip('"')
        except FileNotFoundError:
            pass
        return os_release

    @classmethod
    def detect_distro_family(cls) -> Optional[str]:
        """Auto-detect the distribution family"""
        os_release = cls.read_os_release()
        
        if not os_release:
            return None
        
        distro_id = os_release.get('ID', '').lower()
        id_like = os_release.get('ID_LIKE', '').lower()
        
        # Check direct matches first
        for family, info in cls.DISTRO_FAMILIES.items():
            if distro_id in info['distros']:
                return family
        
        # Check ID_LIKE for derivatives
        for family, info in cls.DISTRO_FAMILIES.items():
            if any(distro in id_like for distro in info['distros']):
                return family
        
        # Special cases for ID_LIKE
        if 'arch' in id_like:
            return 'arch'
        elif 'suse' in id_like:
            return 'opensuse'
        elif 'debian' in id_like:
            return 'debian'
        elif any(x in id_like for x in ['fedora', 'rhel']):
            return 'fedora'
        
        return None

    @classmethod
    def choose_distro_family(cls) -> str:
        """Interactive distro family selection"""
        detected = cls.detect_distro_family()
        
        if detected:
            Logger.info(f"Auto-detected distro family: {detected}")
            choice = input("Use auto-detected family? [Y/n]: ").strip().lower()
            if choice not in ['n', 'no']:
                return detected
        
        Logger.menu("Please choose your Linux distribution family:")
        Logger.menu("1) Arch Linux family (Arch, Manjaro, EndeavourOS, etc.)")
        Logger.menu("2) openSUSE family (openSUSE Leap, Tumbleweed, SLED, etc.)")
        Logger.menu("3) Debian family (Debian, Ubuntu, Mint, Pop!_OS, etc.)")
        Logger.menu("4) Fedora family (Fedora, CentOS, RHEL, Rocky Linux, etc.)")
        Logger.menu("5) Other/Manual (will attempt generic installation)")
        
        while True:
            try:
                choice = input("Enter your choice (1-5): ").strip()
                if choice == '1':
                    return 'arch'
                elif choice == '2':
                    return 'opensuse'
                elif choice == '3':
                    return 'debian'
                elif choice == '4':
                    return 'fedora'
                elif choice == '5':
                    return 'other'
                else:
                    Logger.warning("Invalid choice. Please enter 1-5.")
            except KeyboardInterrupt:
                Logger.error("\nInstallation cancelled by user.")
                sys.exit(1)

class PackageInstaller:
    def __init__(self, distro_family: str):
        self.distro_family = distro_family
        self.detector = DistroDetector()

    def run_command(self, command: str, shell: bool = True) -> Tuple[bool, str]:
        """Run a shell command and return success status and output"""
        try:
            result = subprocess.run(
                command, 
                shell=shell, 
                capture_output=True, 
                text=True, 
                check=True
            )
            return True, result.stdout
        except subprocess.CalledProcessError as e:
            return False, e.stderr

    def install_packages(self, packages: List[str]) -> bool:
        """Install packages using the appropriate package manager"""
        if not packages:
            return True
        
        if self.distro_family == 'other':
            if self.detector.command_exists('brew'):
                cmd = f"brew install {' '.join(packages)}"
            else:
                Logger.error("Cannot install packages automatically for this system.")
                return False
        else:
            family_info = DistroDetector.DISTRO_FAMILIES[self.distro_family]
            cmd = f"{family_info['install_cmd']} {' '.join(packages)}"
        
        Logger.info(f"Installing packages: {', '.join(packages)}")
        success, output = self.run_command(cmd)
        
        if success:
            Logger.success(f"Successfully installed: {', '.join(packages)}")
        else:
            Logger.error(f"Failed to install packages: {output}")
        
        return success

    def install_dev_tools(self) -> bool:
        """Install development tools for the distro family"""
        Logger.info("Installing development tools...")
        
        if self.distro_family == 'arch':
            return self.install_packages(['base-devel'])
        
        elif self.distro_family == 'opensuse':
            cmd = "sudo zypper install -y -t pattern devel_basis"
            success, output = self.run_command(cmd)
            if success:
                Logger.success("Development tools installed")
            else:
                Logger.error(f"Failed to install development tools: {output}")
            return success
        
        elif self.distro_family == 'debian':
            return self.install_packages(['build-essential'])
        
        elif self.distro_family == 'fedora':
            # Try dnf first, fallback to yum
            if self.detector.command_exists('dnf'):
                cmd = 'sudo dnf groupinstall -y "Development Tools"'
            else:
                cmd = 'sudo yum groupinstall -y "Development Tools"'
            
            success, output = self.run_command(cmd)
            if success:
                Logger.success("Development tools installed")
            else:
                Logger.error(f"Failed to install development tools: {output}")
            return success
        
        elif self.distro_family == 'other':
            Logger.warning("Skipping development tools installation for unknown system")
            return True
        
        return False

    def install_system_dependencies(self) -> bool:
        """Install system dependencies (tar, curl/wget, dev tools)"""
        Logger.info("Installing system dependencies...")
        
        packages_needed = []
        
        # Check for tar
        if not self.detector.command_exists('tar'):
            packages_needed.append('tar')
        else:
            Logger.success("tar is already installed")
        
        # Check for curl or wget
        if not self.detector.command_exists('curl') and not self.detector.command_exists('wget'):
            packages_needed.append('curl')
        else:
            Logger.success("curl or wget is already available")
        
        # Install missing packages
        if packages_needed:
            if not self.install_packages(packages_needed):
                return False
        
        # Install development tools
        return self.install_dev_tools()

class RustInstaller:
    def __init__(self):
        self.detector = DistroDetector()
        self.cargo_env = Path.home() / '.cargo' / 'env'

    def is_rust_installed(self) -> bool:
        """Check if Rust is already installed"""
        return (self.detector.command_exists('rustc') and 
                self.detector.command_exists('cargo'))

    def install_rust(self) -> bool:
        """Install Rust using rustup"""
        Logger.info("Checking Rust installation...")
        
        if self.is_rust_installed():
            try:
                result = subprocess.run(['rustc', '--version'], 
                                      capture_output=True, text=True)
                Logger.success(f"Rust is already installed: {result.stdout.strip()}")
                return True
            except:
                pass
        
        Logger.info("Installing Rust...")
        
        # Download and run rustup installer
        if self.detector.command_exists('curl'):
            cmd = 'curl --proto "=https" --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y'
        elif self.detector.command_exists('wget'):
            cmd = 'wget -qO- https://sh.rustup.rs | sh -s -- -y'
        else:
            Logger.error("Neither curl nor wget found after installation attempt.")
            return False
        
        try:
            subprocess.run(cmd, shell=True, check=True)
            Logger.success("Rust installed successfully")
            
            # Source cargo environment
            if self.cargo_env.exists():
                cargo_bin = Path.home() / '.cargo' / 'bin'
                os.environ['PATH'] = f"{cargo_bin}:{os.environ.get('PATH', '')}"
            
            return True
        except subprocess.CalledProcessError:
            Logger.error("Failed to install Rust")
            return False

    def install_rust_dependencies(self) -> bool:
        """Install Rust crate dependencies (tar, flate2)"""
        Logger.info("Installing Rust dependencies...")
        
        # Ensure cargo is available
        if not self.detector.command_exists('cargo'):
            if self.cargo_env.exists():
                # Try to source cargo env
                cargo_bin = Path.home() / '.cargo' / 'bin'
                os.environ['PATH'] = f"{cargo_bin}:{os.environ.get('PATH', '')}"
                
                if not self.detector.command_exists('cargo'):
                    Logger.error("Cargo not found. Please restart your shell or source ~/.cargo/env")
                    return False
            else:
                Logger.error("Cargo not found and ~/.cargo/env doesn't exist")
                return False
        
        # Create temporary project
        with tempfile.TemporaryDirectory() as temp_dir:
            temp_path = Path(temp_dir)
            
            Logger.info("Creating temporary Cargo project...")
            
            try:
                # Initialize cargo project
                subprocess.run(['cargo', 'init', '--name', 'temp_deps_installer', '--bin'], 
                             cwd=temp_path, check=True, capture_output=True)
                
                # Add dependencies to Cargo.toml
                cargo_toml = temp_path / 'Cargo.toml'
                with open(cargo_toml, 'a') as f:
                    f.write('\n[dependencies]\n')
                    f.write('tar = "*"\n')
                    f.write('flate2 = "*"\n')
                
                Logger.info("Installing tar and flate2 crates...")
                subprocess.run(['cargo', 'build'], 
                             cwd=temp_path, check=True, capture_output=True)
                
                Logger.success("Rust dependencies installed successfully")
                return True
                
            except subprocess.CalledProcessError as e:
                Logger.error(f"Failed to install Rust dependencies: {e}")
                return False

class ShellConfigurer:
    def __init__(self):
        self.home = Path.home()
        self.cargo_bin = self.home / '.cargo' / 'bin'

    def get_shell_config_files(self) -> List[Path]:
        """Get list of shell configuration files to update"""
        configs = []
        
        # Bash
        bashrc = self.home / '.bashrc'
        bash_profile = self.home / '.bash_profile'
        if bashrc.exists():
            configs.append(bashrc)
        elif bash_profile.exists():
            configs.append(bash_profile)
        
        # Zsh
        zshrc = self.home / '.zshrc'
        if zshrc.exists():
            configs.append(zshrc)
        
        # Fish
        fish_config = self.home / '.config' / 'fish' / 'config.fish'
        if fish_config.exists():
            configs.append(fish_config)
        
        return configs

    def update_shell_configs(self):
        """Update shell configuration files to include cargo in PATH"""
        configs = self.get_shell_config_files()
        
        if not configs:
            Logger.warning("No shell configuration files found to update")
            return
        
        for config_file in configs:
            try:
                # Check if cargo path already exists
                with open(config_file, 'r') as f:
                    content = f.read()
                
                if '.cargo/bin' in content:
                    Logger.success(f"Cargo path already configured in {config_file}")
                    continue
                
                # Add cargo to PATH based on shell type
                if config_file.name == 'config.fish':
                    path_line = 'set -gx PATH $HOME/.cargo/bin $PATH\n'
                else:
                    path_line = 'export PATH="$HOME/.cargo/bin:$PATH"\n'
                
                with open(config_file, 'a') as f:
                    f.write(f'\n# Added by Rust installer\n{path_line}')
                
                Logger.success(f"Added cargo to PATH in {config_file}")
                
            except Exception as e:
                Logger.warning(f"Failed to update {config_file}: {e}")

def main():
    """Main installation function"""
    try:
        Logger.info("Starting Universal Dependencies Installer")
        Logger.info("Target dependencies: Rust, tar, flate2")
        Logger.info("Supported distro families: Arch, openSUSE, Debian, Fedora")
        
        # Check if running as root
        if os.geteuid() == 0:
            Logger.warning("Running as root. This may cause issues with Rust installation.")
            response = input("Continue anyway? [y/N]: ").strip().lower()
            if response not in ['y', 'yes']:
                Logger.info("Installation cancelled.")
                return
        
        # Choose distro family
        distro_family = DistroDetector.choose_distro_family()
        Logger.info(f"Selected distro family: {distro_family}")
        
        # Install system dependencies
        package_installer = PackageInstaller(distro_family)
        if not package_installer.install_system_dependencies():
            Logger.error("Failed to install system dependencies")
            sys.exit(1)
        
        # Install Rust
        rust_installer = RustInstaller()
        if not rust_installer.install_rust():
            Logger.error("Failed to install Rust")
            sys.exit(1)
        
        # Install Rust dependencies
        if not rust_installer.install_rust_dependencies():
            Logger.error("Failed to install Rust dependencies")
            sys.exit(1)
        
        # Update shell configurations
        shell_configurer = ShellConfigurer()
        shell_configurer.update_shell_configs()
        
        # Final success message
        Logger.success("All dependencies installed successfully!")
        Logger.info("You can now use Rust with tar and flate2 crates in your projects")
        
        # Show versions if available
        detector = DistroDetector()
        if detector.command_exists('rustc'):
            Logger.info("Installed versions:")
            try:
                rustc_version = subprocess.run(['rustc', '--version'], 
                                             capture_output=True, text=True, check=True)
                cargo_version = subprocess.run(['cargo', '--version'], 
                                             capture_output=True, text=True, check=True)
                print(f"  {rustc_version.stdout.strip()}")
                print(f"  {cargo_version.stdout.strip()}")
            except:
                pass
        
        Logger.warning("Please restart your shell or source your shell's config file to use cargo")
        
    except KeyboardInterrupt:
        Logger.error("\nInstallation cancelled by user.")
        sys.exit(1)
    except Exception as e:
        Logger.error(f"Unexpected error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
