# ⚡ OPEN LIGHTNING CHANNEL - SIMPLE GUIDE

## 🚨 THE PROBLEM
- ✅ **Bitcoin:** 151,976 sats (on-chain)
- ❌ **Lightning:** 0 sats (no channels)
- **Issue:** Money is on-chain, not in Lightning!

## 🎯 SOLUTION: Open a Lightning Channel

### Step 1: Go to Channels in ThunderHub
1. **In ThunderHub, click "Channels"** in the left sidebar
2. **Look for "Open Channel" or "+" button**

### Step 2: Open Channel to a Public Node
**You need to connect to a public Lightning node first:**

**Option A: Connect to ACINQ (Eclair)**
- **Node ID:** `03864ef025fde8fb587d989186ce6a4a186895ee44a926bfc370e2c366597a3f8f`
- **Host:** `34.239.230.56:9735`

**Option B: Connect to Lightning Labs**
- **Node ID:** `03c2abfa93eacec04721c019644584424aab2ba4dff3ac9bdab4e9c97007491dda`
- **Host:** `34.239.230.56:9735`

### Step 3: Open the Channel
1. **Click "Open Channel"**
2. **Paste the Node ID** (from above)
3. **Enter amount:** 50,000 sats (minimum)
4. **Click "Open Channel"**
5. **Wait for confirmation** (takes a few minutes)

### Step 4: Now You Can Send Lightning Payments!
Once the channel is open:
1. **Go back to "Home" in ThunderHub**
2. **Lightning Available should show ~50,000 sats**
3. **Now you can pay your invoice!**

---

## 🚀 QUICK ALTERNATIVE: Use a Lightning Service

If opening channels is too complicated:

### Option 1: Use Voltage's Built-in Service
1. **Go back to your Voltage dashboard**
2. **Look for "Lightning Services" or "Channel Services"**
3. **Use their automated channel opening**

### Option 2: Use a Lightning Service Provider
1. **Visit:** https://lightning.plus/
2. **Connect your node**
3. **They'll open channels for you**

---

## 🎯 WHY THIS HAPPENS

**Lightning Network Structure:**
- **Bitcoin (on-chain):** Your 151,976 sats
- **Lightning (off-chain):** Requires channels to other nodes
- **Channels:** Connections to other Lightning nodes

**The Process:**
1. **Open channel** to another node (costs Bitcoin)
2. **Money moves** from on-chain to Lightning
3. **Now you can send** Lightning payments!

---

## 🎉 AFTER YOU OPEN A CHANNEL

Once you have Lightning funds:

1. **Go to ThunderHub Home**
2. **Click "X Send" under Lightning**
3. **Paste your invoice:**
   ```
   lntb100u1p50t28ypp59xjrv8tmg2sy7a6qu4f7he8htqusl4wdthnwh94kp08e7dpjy96scqzyssp57pfx0capuy6d53m35fpyus6gc6c7ty7mc7gr4nm8h2h60ua0g06q9qyysgqdqgw3jhxap3mqz9gxqyjw5qrzjqt20ylt5sxnkmhy0u8t59605nxy4vqekehu36pntljgf0q85lxsktlexlm5p24ccx5qqqqlgqqqqqeqqjqpvaa59fqwg0lmc70mw9q4zhq6gjreeg5au8c5l87lxj84kkmy7kxchqwnygjqut3g9t27z8j4r9zvap32yaxkskj3pqlt89wq3frqygq6cdpu4
   ```
4. **Click "Pay"**
5. **Your wallet receives 10,000 sats!**

---

## 🚨 SIMPLE SUMMARY

**Current:** Bitcoin = 151,976 sats, Lightning = 0 sats
**Need:** Open Lightning channel to move Bitcoin → Lightning
**Result:** Lightning = ~50,000 sats, can send payments!

**Go to "Channels" in ThunderHub and open a channel to a public node!** ⚡💰
