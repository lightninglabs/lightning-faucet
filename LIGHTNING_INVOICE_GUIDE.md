# Lightning Invoice Guide - How to Receive Payments

## 🎯 Your Goal
You want to receive Lightning payments in your wallet app and see the funds.

## ✅ Your Current Setup
- **Wallet Address (On-chain):** `tb1qehcjwfcqh7yre3r9xk7qsukayqkgxakcucvdda`
- **Node Public Key:** `03edb07655547936eeae2db87ba02b928914ae134dfa650af7655b65401bca2fbb`
- **Node Alias:** `lightning-faucet-node`
- **Network:** Bitcoin Testnet

## 📝 Creating Lightning Invoices

### Basic Commands:

```powershell
# Create an invoice for 10,000 sats (0.0001 tBTC)
.\lncli.exe --network=testnet addinvoice --amt 10000

# Create an invoice with a memo/description
.\lncli.exe --network=testnet addinvoice --amt 5000 --memo "Payment for coffee"

# Create an invoice with custom expiry (1 hour = 3600 seconds)
.\lncli.exe --network=testnet addinvoice --amt 1000 --expiry 3600
```

### What You'll Get:

When you create an invoice, you'll receive:
```json
{
    "r_hash": "abc123...",  // Payment hash
    "payment_request": "lntb10u1...",  // The actual invoice (starts with lntb for testnet)
    "add_index": "1"
}
```

The `payment_request` is what you share with the person paying you. It looks like:
```
lntb10u1pn...  (for testnet)
lnbc10u1pn...  (for mainnet)
```

## ⚠️ CRITICAL REQUIREMENT: You Need Inbound Liquidity!

### The Problem:
**You currently have 0 channels = You CANNOT receive Lightning payments yet!**

### Why?
Lightning payments work like water pipes:
- To SEND money: You need a channel with outbound capacity (money on your side)
- To RECEIVE money: You need inbound capacity (empty space on your side where money can flow in)

### How to Get Inbound Liquidity:

#### Option 1: Someone Opens a Channel TO You
Have another Lightning node open a channel to you:
```powershell
# They would run on their node:
lncli openchannel --node_key=03edb07655547936eeae2db87ba02b928914ae134dfa650af7655b65401bca2fbb --local_amt=100000
```

#### Option 2: You Open a Channel and Spend Some
1. You open a channel with 100,000 sats
2. You spend 50,000 sats through that channel
3. Now you have 50,000 sats of inbound liquidity (space for others to pay you)

#### Option 3: Use a Service (Testnet)
For testnet, you can:
- Use the faucet you just set up to open channels
- Connect to other testnet nodes and request channels

## 📊 Checking Your Status

### Check if you can receive payments:
```powershell
# List all channels
.\lncli.exe --network=testnet listchannels

# Check channel balance
.\lncli.exe --network=testnet channelbalance
```

Look for:
- `remote_balance`: This is your INBOUND capacity (how much you can receive)
- `local_balance`: This is your OUTBOUND capacity (how much you can send)

### Check all invoices:
```powershell
# List all invoices
.\lncli.exe --network=testnet listinvoices

# List only pending invoices
.\lncli.exe --network=testnet listinvoices --pending_only

# Check a specific invoice by payment hash
.\lncli.exe --network=testnet lookupinvoice <r_hash>
```

## 💰 Where Will Funds Appear?

### Lightning Balance vs On-Chain Balance

You have TWO separate balances:

#### 1. On-Chain Balance (Regular Bitcoin Wallet)
```powershell
.\lncli.exe --network=testnet walletbalance
```
- This is Bitcoin in your regular wallet
- Receives funds sent to: `tb1qehcjwfcqh7yre3r9xk7qsukayqkgxakcucvdda`
- Shows up as `confirmed_balance` and `unconfirmed_balance`

#### 2. Lightning Channel Balance
```powershell
.\lncli.exe --network=testnet channelbalance
```
- This is Bitcoin locked in Lightning channels
- Receives funds from Lightning invoices
- Shows up as `balance` (your side) and `pending_open_balance`

### Important:
**Lightning payments go to your channel balance, NOT your wallet balance!**

To move Lightning funds to your on-chain wallet:
```powershell
# Close a channel (funds return to on-chain wallet)
.\lncli.exe --network=testnet closechannel <channel_point>
```

## 🔄 Complete Workflow Example

### Step 1: Wait for Node to Sync
```powershell
.\lncli.exe --network=testnet getinfo
# Wait until "synced_to_chain": true
```

### Step 2: Fund Your On-Chain Wallet
```powershell
# Get your address
.\lncli.exe --network=testnet newaddress p2wkh

# Go to https://testnet-faucet.mempool.co/ and send tBTC

# Check balance
.\lncli.exe --network=testnet walletbalance
```

### Step 3: Open a Channel (or have someone open one to you)
```powershell
# Connect to a peer first
.\lncli.exe --network=testnet connect <pubkey>@<host>:<port>

# Open channel
.\lncli.exe --network=testnet openchannel --node_key=<pubkey> --local_amt=100000

# Wait for confirmation
.\lncli.exe --network=testnet pendingchannels
```

### Step 4: Create an Invoice
```powershell
.\lncli.exe --network=testnet addinvoice --amt 1000 --memo "Test payment"
```

### Step 5: Share the Invoice
Copy the `payment_request` field and give it to the payer.

### Step 6: Check Payment
```powershell
# List invoices to see if it's paid
.\lncli.exe --network=testnet listinvoices

# Check channel balance (should increase)
.\lncli.exe --network=testnet channelbalance
```

## 🚀 Quick Commands Reference

```powershell
# Node status
.\lncli.exe --network=testnet getinfo

# Wallet (on-chain) balance
.\lncli.exe --network=testnet walletbalance

# Lightning (channel) balance
.\lncli.exe --network=testnet channelbalance

# Create invoice for 10,000 sats
.\lncli.exe --network=testnet addinvoice --amt 10000

# List all invoices
.\lncli.exe --network=testnet listinvoices

# List channels
.\lncli.exe --network=testnet listchannels

# List peers
.\lncli.exe --network=testnet listpeers

# Connect to peer
.\lncli.exe --network=testnet connect <pubkey>@<host>:<port>

# Open channel (need on-chain funds first!)
.\lncli.exe --network=testnet openchannel --node_key=<pubkey> --local_amt=100000

# Close channel
.\lncli.exe --network=testnet closechannel <funding_txid> <output_index>
```

## 🎯 What You Need To Do Now

1. **Wait for LND to sync** (check with `.\lncli.exe --network=testnet getinfo`)
2. **Fund your on-chain wallet** (send tBTC to `tb1qehcjwfcqh7yre3r9xk7qsukayqkgxakcucvdda`)
3. **Open channels or get channels opened to you**
4. **THEN you can receive Lightning payments via invoices**

## ❓ FAQ

**Q: I created an invoice, but no one can pay it. Why?**
A: You probably don't have inbound liquidity (channels where others can send you money).

**Q: Where do Lightning payments show up?**
A: In your channel balance, not your wallet balance. Check with `channelbalance`.

**Q: How do I convert Lightning balance to on-chain balance?**
A: Close the channel. Funds will return to your on-chain wallet.

**Q: Can I receive payments without opening channels?**
A: No, you need at least one channel with inbound capacity.

**Q: My invoice expired, what happens?**
A: Nothing. Create a new one. Default expiry is 1 hour.

---

**Remember: Lightning Network requires channels to work. No channels = no Lightning payments!**

