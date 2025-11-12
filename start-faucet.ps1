# Lightning Testnet Faucet Startup Script
# This script helps you start the Lightning Faucet easily

Write-Host "===========================================================" -ForegroundColor Cyan
Write-Host "  Lightning Network Testnet Faucet - Quick Start" -ForegroundColor Cyan
Write-Host "===========================================================" -ForegroundColor Cyan
Write-Host ""

# Check if LND is running
Write-Host "[1/3] Checking if LND is running..." -ForegroundColor Yellow
$lndRunning = Get-Process -Name "lnd" -ErrorAction SilentlyContinue
if ($null -eq $lndRunning) {
    Write-Host "❌ LND is not running!" -ForegroundColor Red
    Write-Host ""
    Write-Host "Please start LND first in a separate terminal:" -ForegroundColor Yellow
    Write-Host "  .\lnd.exe" -ForegroundColor White
    Write-Host ""
    Write-Host "Then unlock your wallet:" -ForegroundColor Yellow
    Write-Host "  .\lncli.exe --network=testnet unlock" -ForegroundColor White
    Write-Host ""
    exit 1
} else {
    Write-Host "✅ LND is running (PID: $($lndRunning.Id))" -ForegroundColor Green
}

# Check if wallet is unlocked by testing connection
Write-Host "[2/3] Checking if LND wallet is unlocked..." -ForegroundColor Yellow
try {
    $info = & .\lncli.exe --network=testnet getinfo 2>&1
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ LND wallet is unlocked" -ForegroundColor Green
        
        # Parse and display node info
        $infoJson = $info | ConvertFrom-Json
        Write-Host ""
        Write-Host "Node Information:" -ForegroundColor Cyan
        Write-Host "  Public Key: $($infoJson.identity_pubkey)" -ForegroundColor White
        Write-Host "  Alias: $($infoJson.alias)" -ForegroundColor White
        Write-Host "  Synced: $($infoJson.synced_to_chain)" -ForegroundColor White
        Write-Host "  Block Height: $($infoJson.block_height)" -ForegroundColor White
        Write-Host "  Active Channels: $($infoJson.num_active_channels)" -ForegroundColor White
        Write-Host ""
        
        if (-not $infoJson.synced_to_chain) {
            Write-Host "⚠️  WARNING: Node is not fully synced yet!" -ForegroundColor Yellow
            Write-Host "   The faucet may not work properly until sync completes." -ForegroundColor Yellow
            Write-Host ""
        }
    } else {
        throw "LND connection failed"
    }
} catch {
    Write-Host "❌ Cannot connect to LND or wallet is locked" -ForegroundColor Red
    Write-Host ""
    Write-Host "Please unlock your wallet:" -ForegroundColor Yellow
    Write-Host "  .\lncli.exe --network=testnet unlock" -ForegroundColor White
    Write-Host ""
    exit 1
}

# Start the faucet
Write-Host "[3/3] Starting Lightning Faucet..." -ForegroundColor Yellow
Write-Host ""
Write-Host "Faucet Configuration:" -ForegroundColor Cyan
Write-Host "  - Network: Bitcoin Testnet" -ForegroundColor White
Write-Host "  - Port: 8080" -ForegroundColor White
Write-Host "  - URL: http://localhost:8080" -ForegroundColor White
Write-Host "  - Min Channel Size: 50,000 sats" -ForegroundColor White
Write-Host "  - Max Channel Size: 16,777,216 sats" -ForegroundColor White
Write-Host ""
Write-Host "===========================================================" -ForegroundColor Cyan
Write-Host "  Press Ctrl+C to stop the faucet" -ForegroundColor Yellow
Write-Host "===========================================================" -ForegroundColor Cyan
Write-Host ""

# Start the faucet
& .\lightning-faucet.exe --port=8080 --lnd_ip=127.0.0.1 --network=bitcoin --net=testnet --nodes=localhost:10009

