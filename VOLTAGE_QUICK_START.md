# Using Your Voltage Testnet Node

## Your Voltage Node Details

- **Node Public Key:** `028741391bfb60c91d72558a25b80fa71b0382a1d829b770f20d47496faa078efc`
- **Clearnet:** `028741391bfb60c91d72558a25b80fa71b0382a1d829b770f20d47496faa078efc@54.214.32.132:20249`
- **API Endpoint:** `app-test-net.t.voltageapp.io:8080`

## ✅ Why Use Voltage Now?

Your Voltage node:
- ✅ Is already synced
- ✅ Has satoshis in it
- ✅ Probably has channels already
- ✅ Works immediately!

Your local node:
- ❌ Not synced yet (DNS issues)
- ❌ No channels
- ❌ Needs setup time

## 🚀 How to Create Invoices on Voltage

### Method 1: Web Dashboard (Easiest)

1. Go to https://voltage.cloud/
2. Log into your account
3. Navigate to your testnet node dashboard
4. Click "Receive" or "Create Invoice"
5. Enter amount (in sats) and optional memo
6. Copy the invoice (starts with `lntb...`)
7. Share it with someone to receive payment!

### Method 2: Voltage API

You'll need your macaroon for authentication (find it in Voltage dashboard).

```bash
# Create an invoice for 1000 sats
curl -X POST https://app-test-net.t.voltageapp.io:8080/v1/invoices \
  -H "Grpc-Metadata-macaroon: YOUR_ADMIN_MACAROON_HERE" \
  -d '{
    "value": "1000",
    "memo": "Payment for services"
  }'
```

### Method 3: lncli with Voltage Credentials

If you have Voltage's connection credentials:

```powershell
# You'd need to configure lncli to connect to Voltage
# This requires their macaroon and TLS cert

lncli --rpcserver=app-test-net.t.voltageapp.io:8080 \
      --macaroonpath=path/to/admin.macaroon \
      --tlscertpath=path/to/tls.cert \
      --network=testnet \
      addinvoice --amt 1000
```

## 📊 Checking Voltage Node Status

### Check Balance
```bash
# Via API
curl https://app-test-net.t.voltageapp.io:8080/v1/balance/channels \
  -H "Grpc-Metadata-macaroon: YOUR_MACAROON"
```

### List Channels
```bash
curl https://app-test-net.t.voltageapp.io:8080/v1/channels \
  -H "Grpc-Metadata-macaroon: YOUR_MACAROON"
```

### List Invoices
```bash
curl https://app-test-net.t.voltageapp.io:8080/v1/invoices \
  -H "Grpc-Metadata-macaroon: YOUR_MACAROON"
```

## 🔗 Connecting Local Node to Voltage (Once Local Node Works)

When your local LND node is synced and working:

### Step 1: Connect to Voltage
```powershell
.\lncli.exe --network=testnet connect 028741391bfb60c91d72558a25b80fa71b0382a1d829b770f20d47496faa078efc@54.214.32.132:20249
```

### Step 2: Verify Connection
```powershell
.\lncli.exe --network=testnet listpeers
```

### Step 3: Open Channel FROM Local TO Voltage
This gives you outbound capacity (ability to send to Voltage):
```powershell
# Make sure you have on-chain funds first!
.\lncli.exe --network=testnet openchannel \
  --node_key=028741391bfb60c91d72558a25b80fa71b0382a1d829b770f20d47496faa078efc \
  --local_amt=100000
```

### Step 4: Open Channel FROM Voltage TO Local
This gives you inbound capacity (ability to receive):
- **Problem:** Your local node at `127.0.0.1` is not publicly accessible
- **Solution:** You'd need to:
  - Set up port forwarding (port 9735)
  - Or use a tunnel service like ngrok
  - Or just use Voltage directly!

## 💡 Recommended Approach

**For Now:**
1. ✅ Use Voltage web dashboard to create invoices
2. ✅ Test receiving payments there
3. ✅ Everything works immediately!

**For Later (Learning):**
1. Wait for local LND to sync (may take hours with DNS issues)
2. Fund local wallet with testnet BTC
3. Connect local node to Voltage
4. Open channels between them
5. Experiment with both nodes

## 🎯 Quick Test: Create Your First Invoice

1. Go to Voltage dashboard
2. Find your testnet node
3. Click "Receive" or "Lightning" → "Receive"
4. Enter amount: `1000` sats
5. Memo: `My first Lightning invoice`
6. Click "Create Invoice"
7. Copy the invoice string (looks like `lntb10u1...`)

Now you can share that invoice with anyone on Lightning testnet and they can pay you!

## 📱 Testing Your Invoice

To test if someone paid your invoice:
1. Create an invoice on Voltage
2. Pay it from another Lightning wallet (or another testnet faucet)
3. Check your Voltage balance - it should increase!

## ⚠️ Local Node Status

Your local LND node has issues:
- DNS resolution failing for neutrino peers
- Not syncing to blockchain
- May need 30-60 minutes to resolve

**Recommendation:** Use Voltage now, fix local node later for learning purposes.

## 🔧 Fixing Local Node (Advanced)

The local node can't connect to neutrino peers. Options:

1. **Wait it out** - Sometimes takes 30+ min to find peers
2. **Check firewall** - Make sure LND can make outbound connections
3. **Try different peers** - Edit `lnd.conf` with different neutrino nodes
4. **Use btcd/bitcoind** - Instead of neutrino mode (more complex)

## 📚 Resources

- [Voltage Docs](https://docs.voltage.cloud/)
- [LND API Reference](https://lightning.engineering/api-docs/api/lnd/)
- [Lightning Network Spec](https://github.com/lightningnetwork/lightning-rfc)

---

**Bottom Line:** Your Voltage node works NOW. Use it! The local node is for learning/testing later.

