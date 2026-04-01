# TACO Installer for Windows
# Usage: irm https://raw.githubusercontent.com/maddenmanel/taco/main/install.ps1 | iex
#
# This script downloads the TACO binary, unblocks it (removes the
# "downloaded from internet" mark), and adds it to your PATH.
# No SmartScreen warning. No admin rights required.

param(
    [string]$Version = ""
)

$ErrorActionPreference = "Stop"
$Repo = "maddenmanel/taco"
$BinName = "taco.exe"
$InstallDir = "$env:USERPROFILE\.taco\bin"

function Write-Step($msg) { Write-Host "  $msg" -ForegroundColor Cyan }
function Write-Ok($msg)   { Write-Host "  ✓ $msg" -ForegroundColor Green }
function Write-Fail($msg) { Write-Host "  ✗ $msg" -ForegroundColor Red; exit 1 }

Write-Host ""
Write-Host "  🌮 TACO Installer" -ForegroundColor White
Write-Host ""

# Detect architecture
$Arch = if ([System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture -eq "Arm64") { "arm64" } else { "amd64" }
Write-Ok "Platform: windows/$Arch"

# Get latest version if not specified
if (-not $Version) {
    Write-Step "Fetching latest version..."
    try {
        $Release = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest"
        $Version = $Release.tag_name
    } catch {
        Write-Fail "Could not fetch release info. Check https://github.com/$Repo/releases"
    }
}

Write-Ok "Version: $Version"

# Build download URL
$FileName = "taco-windows-$Arch.exe"
$Url = "https://github.com/$Repo/releases/download/$Version/$FileName"

# Create install dir
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

# Download
$Dest = Join-Path $InstallDir $BinName
Write-Step "Downloading $FileName..."
Invoke-WebRequest -Uri $Url -OutFile $Dest -UseBasicParsing

# KEY STEP: Remove the "downloaded from internet" flag
# This prevents Windows SmartScreen from blocking the binary.
Unblock-File -Path $Dest
Write-Ok "Unblocked (SmartScreen bypass)"

# Add to user PATH if not already present
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    $env:Path += ";$InstallDir"
    Write-Ok "Added $InstallDir to PATH"
} else {
    Write-Ok "PATH already configured"
}

Write-Host ""
Write-Host "  🌮 TACO installed successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "  Get started (open a new terminal):" -ForegroundColor White
Write-Host "    taco add deepseek --key=`"sk-your-key`""
Write-Host "    taco use deepseek"
Write-Host "    taco --help"
Write-Host ""
Write-Host "  To uninstall: taco uninstall" -ForegroundColor Gray
Write-Host ""
