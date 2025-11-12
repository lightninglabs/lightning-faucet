# Lightning Testnet Faucet Setup Guide

This guide will help you set up and run the Lightning Network Testnet Faucet on Windows.

## Prerequisites

✅ **Installed Components:**
- Go 1.25.3 (installed via winget)
- Lightning Network Daemon (lnd) v0.19.3-beta
- Lightning Faucet (built from source)

## Directory Structure

```
C:\Users\jafil\Documents\GitHub\lightning-faucet\
├── lightning-faucet.exe     # Main faucet application
├── lnd.exe                   # Lightning Network Daemon
├── lncli.exe                 # LND Command Line Interface
├── static\                   # Web frontend files
└── ...

C:\Users\jafil\AppData\Local\Lnd\
└── lnd.conf                  # LND configuration file
```

## Step-by-Step Setup

### Step 1: Start LND on Testnet

Open a PowerShell terminal and navigate to the faucet directory:

```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet
```

Start the Lightning Network Daemon:

```powershell
.\lnd.exe
```

**First Time Setup:** 
- You'll be prompted to create a wallet password
- You'll be given a seed phrase - **SAVE THIS SECURELY!**
- The seed phrase is needed to recover your wallet

**Initial Sync:**
- LND will sync with the Bitcoin testnet blockchain
- This can take 30 minutes to several hours depending on your connection
- You'll see log messages showing sync progress

### Step 2: Create and Unlock Your Wallet

In a **new PowerShell terminal**, create your wallet:

```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet
.\lncli.exe --network=testnet create
```

Follow the prompts to:
1. Set a wallet password (minimum 8 characters)
2. Optionally set a passphrase for your seed
3. Write down your 24-word seed phrase

For subsequent starts, unlock your wallet:

```powershell
.\lncli.exe --network=testnet unlock
```

### Step 3: Wait for LND to Sync

Check sync status:

```powershell
.\lncli.exe --network=testnet getinfo
```

Look for `"synced_to_chain": true` in the output.

### Step 4: Get Testnet Bitcoin

Your node needs testnet Bitcoin to open channels. Get your wallet address:

```powershell
.\lncli.exe --network=testnet newaddress p2wkh
```

Visit a Bitcoin testnet faucet to get test coins:
- https://testnet-faucet.mempool.co/
- https://coinfaucet.eu/en/btc-testnet/
- https://bitcoinfaucet.uo1.net/

Send at least 0.01 tBTC to your address.

Check your balance:

```powershell
.\lncli.exe --network=testnet walletbalance
```

### Step 5: Get Your Node Information

Get your node's public key and connection info:

```powershell
.\lncli.exe --network=testnet getinfo
```

Note the `identity_pubkey` - you'll need this for the faucet.

### Step 6: Run the Lightning Faucet

In a **third PowerShell terminal**, start the faucet:

```powershell
cd C:\Users\jafil\Documents\GitHub\lightning-faucet

# For local testing (HTTP only on port 8080)
.\lightning-faucet.exe --port=8080 --lnd_ip=localhost --network=bitcoin --net=testnet

# For production with HTTPS (requires a domain):
# .\lightning-faucet.exe --lnd_ip=YOUR_PUBLIC_IP --domain=yourdomain.com --network=bitcoin --net=testnet
```

### Step 7: Access the Faucet

Open your web browser and navigate to:
```
http://localhost:8080
```

## Faucet Configuration Options

The faucet accepts several command-line flags:

- `--lnd_ip`: Public IP address of your LND node (default: "10.0.0.9")
- `--port`: HTTP server port (default: "8080")
- `--nodes`: LND node(s) to connect to (default: "localhost:10009")
- `--network`: Network type - "bitcoin" or "litecoin" (default: "bitcoin")
- `--net`: Bitcoin network - "testnet", "mainnet", or "simnet" (default: "testnet")
- `--domain`: Domain name for HTTPS/TLS (default: "faucet.lightning.community")

## Using the Faucet

1. Users must first connect to your LND node
2. They provide their node's public key
3. They specify the channel amount (min: 50,000 sat, max: 16,777,216 sat)
4. The faucet opens a channel with their node
5. Users can use the channel for Lightning Network payments

## Useful LND Commands

Check node status:
```powershell
.\lncli.exe --network=testnet getinfo
```

List open channels:
```powershell
.\lncli.exe --network=testnet listchannels
```

List pending channels:
```powershell
.\lncli.exe --network=testnet pendingchannels
```

Connect to a peer:
```powershell
.\lncli.exe --network=testnet connect <pubkey>@<host>:<port>
```

Open a channel:
```powershell
.\lncli.exe --network=testnet openchannel <pubkey> <amount>
```

Close a channel:
```powershell
.\lncli.exe --network=testnet closechannel <channel_point>
```

## Troubleshooting

### LND won't start
- Check if the port 9735 (P2P) and 10009 (RPC) are not in use
- Check the logs in `C:\Users\jafil\AppData\Local\Lnd\logs\`

### Faucet can't connect to LND
- Ensure LND is running and synced
- Check that the macaroon file exists at: `C:\Users\jafil\AppData\Local\Lnd\data\chain\bitcoin\testnet\admin.macaroon`
- Verify the TLS certificate exists at: `C:\Users\jafil\AppData\Local\Lnd\tls.cert`

### Wallet won't unlock
- Make sure you're using the correct password
- If you forgot your password, you'll need to restore from your seed phrase

### Not syncing
- Check your internet connection
- Try different neutrino peers in `lnd.conf`
- Check LND logs for errors

## Security Notes

⚠️ **Important:**
- This is for **TESTNET ONLY** - never use testnet seeds/passwords for mainnet
- Keep your seed phrase secure (even for testnet)
- The faucet is designed for public use - expect users to request channels
- For production deployment, use proper HTTPS with a valid domain and SSL certificate

## Stopping the Services

To stop the services gracefully:

1. Stop the faucet: Press `Ctrl+C` in the faucet terminal
2. Stop LND: 
   ```powershell
   .\lncli.exe --network=testnet stop
   ```

## Next Steps

- Monitor your channels through the web interface
- Test receiving and sending payments
- Explore the LND API documentation: https://api.lightning.community/
- Join the Lightning Network community: https://discord.gg/lightningnetwork

## Resources

- LND Documentation: https://docs.lightning.engineering/
- Lightning Network Specification: https://github.com/lightningnetwork/lightning-rfc
- Bitcoin Testnet Explorer: https://blockstream.info/testnet/
- Lightning Network Explorer: https://1ml.com/testnet/

