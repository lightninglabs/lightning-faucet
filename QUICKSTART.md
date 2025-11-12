# Quick Start Guide - Lightning Testnet Faucet

This is a simplified guide to get your Lightning Testnet Faucet running quickly.

## Prerequisites

✅ All required software has been installed:
- Go 1.25.3
- LND v0.19.3-beta  
- Lightning Faucet (built successfully)

## Quick Start (3 Steps)

### Terminal 1: Start LND

```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet
.\lnd.exe
```

**First time only:** You'll be prompted to create a wallet. Follow the on-screen instructions.

### Terminal 2: Unlock Wallet

```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet
.\lncli.exe --network=testnet unlock
```

Enter your wallet password when prompted.

### Terminal 3: Start the Faucet

**Option A - Use the startup script:**
```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet
.\start-faucet.ps1
```

**Option B - Manual start:**
```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet
.\lightning-faucet.exe --port=8080 --lnd_ip=127.0.0.1 --network=bitcoin --net=testnet
```

### Access the Faucet

Open your browser: **http://localhost:8080**

## First Time Setup Checklist

If this is your first time running the faucet:

1. ✅ Start LND
2. ✅ Create a new wallet (save your seed phrase!)
3. ⏳ Wait for blockchain sync (this takes time - check with `.\lncli.exe --network=testnet getinfo`)
4. 💰 Fund your wallet with testnet Bitcoin:
   - Get your address: `.\lncli.exe --network=testnet newaddress p2wkh`
   - Visit: https://testnet-faucet.mempool.co/
   - Send at least 0.01 tBTC to your address
5. ✅ Start the faucet

## Common Issues

**"Cannot connect to LND"**
- Make sure LND is running in Terminal 1
- Make sure you unlocked your wallet in Terminal 2

**"Wallet is locked"**
- Run: `.\lncli.exe --network=testnet unlock`

**"Not synced yet"**
- Wait for LND to sync with testnet (check: `.\lncli.exe --network=testnet getinfo`)
- Look for `"synced_to_chain": true`

## Need More Help?

See the complete [SETUP_GUIDE.md](SETUP_GUIDE.md) for detailed instructions.

## Stopping Everything

1. Press `Ctrl+C` in the faucet terminal (Terminal 3)
2. Run: `.\lncli.exe --network=testnet stop` (Terminal 2)
3. Press `Ctrl+C` in the LND terminal if needed (Terminal 1)

