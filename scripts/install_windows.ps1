# PowerShell installation script for Gitzy
# Safely installs gitzy.exe to user's bin directory and updates PATH

$ErrorActionPreference = "Stop"

# Installation directory
$TargetDir = "$env:USERPROFILE\bin"

# Create directory if it doesn't exist
if (!(Test-Path $TargetDir)) {
    New-Item -ItemType Directory -Path $TargetDir -Force | Out-Null
}

# Copy executable
$Source = "gitzy.exe"
$Destination = "$TargetDir\gitzy.exe"

if (!(Test-Path $Source)) {
    Write-Error "gitzy.exe not found in current directory"
    exit 1
}

Copy-Item $Source $Destination -Force
Write-Host "Installed gitzy.exe to $TargetDir" -ForegroundColor Green

# Update user PATH safely
$UserPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
if ($UserPath -notlike "*$TargetDir*") {
    $NewPath = if ($UserPath) { "$UserPath;$TargetDir" } else { $TargetDir }
    [Environment]::SetEnvironmentVariable('PATH', $NewPath, 'User')
    Write-Host "Added $TargetDir to user PATH" -ForegroundColor Green
} else {
    Write-Host "$TargetDir already in PATH" -ForegroundColor Yellow
}

Write-Host "`nInstallation complete!" -ForegroundColor Green
Write-Host "Please restart your terminal to use 'gitzy' command." -ForegroundColor Cyan