# 💰 MOVE MONEY FROM NODE TO LNBITS WALLET

## 🚨 The Problem
- ✅ **Voltage Node:** Has 152K sats
- ❌ **LNBits Wallet:** Shows 0 sats
- **Issue:** Money is in the node, not in the LNBits wallet!

## 🎯 Solution: Create Invoice in LNBits, Pay from Node

### Step 1: Create Invoice in LNBits
1. **In your LNBits wallet** (the one showing 0 sats)
2. **Click "CREATE INVOICE"** (purple button)
3. **Enter amount:** 10,000 sats
4. **Enter memo:** "Test payment"
5. **Click "Create Invoice"**
6. **Copy the invoice** (starts with `lntb...`)

### Step 2: Pay Invoice from Voltage Node
1. **Go back to your Voltage dashboard**
2. **Click on "LNBits"** in the Applications section
3. **Look for "Send" or "Pay Invoice"**
4. **Paste the invoice you just created**
5. **Click "Pay"**

### Step 3: Check LNBits Wallet
1. **Go back to your LNBits wallet**
2. **Refresh the page**
3. **You should now see 10,000 sats!**

---

## 🔄 Alternative: Use ThunderHub

If LNBits doesn't work, try ThunderHub:

### Step 1: Open ThunderHub
1. **In your Voltage dashboard**
2. **Click "ThunderHub"** in Applications
3. **Wait for it to load**

### Step 2: Create Invoice in ThunderHub
1. **Look for "Receive" or "Create Invoice"**
2. **Enter amount:** 10,000 sats
3. **Create the invoice**
4. **Copy the invoice**

### Step 3: Pay from ThunderHub
1. **Look for "Send" or "Pay Invoice"**
2. **Paste the invoice**
3. **Pay it**

---

## 🎯 Why This Happens

**Voltage Node Structure:**
- **Node Balance:** 152K sats (on-chain Bitcoin)
- **LNBits Wallet:** 0 sats (Lightning wallet)
- **Need:** Move money from node to Lightning wallet

**The Process:**
1. **Create invoice** in Lightning wallet (LNBits)
2. **Pay invoice** from node balance
3. **Money moves** from on-chain to Lightning
4. **Now you can send** Lightning payments!

---

## 🚀 Once You Have Money in LNBits

After you get 10,000 sats in your LNBits wallet:

### To Send to Your Wallet App:
1. **Click "PASTE REQUEST"** in LNBits
2. **Paste your wallet's invoice:**
   ```
   lntb100u1p50t28ypp59xjrv8tmg2sy7a6qu4f7he8htqusl4wdthnwh94kp08e7dpjy96scqzyssp57pfx0capuy6d53m35fpyus6gc6c7ty7mc7gr4nm8h2h60ua0g06q9qyysgqdqgw3jhxap3mqz9gxqyjw5qrzjqt20ylt5sxnkmhy0u8t59605nxy4vqekehu36pntljgf0q85lxsktlexlm5p24ccx5qqqqlgqqqqqeqqjqpvaa59fqwg0lmc70mw9q4zhq6gjreeg5au8c5l87lxj84kkmy7kxchqwnygjqut3g9t27z8j4r9zvap32yaxkskj3pqlt89wq3frqygq6cdpu4
   ```
3. **Click "Pay"**
4. **Your wallet app receives 10,000 sats!**

---

## 🎉 Summary

**Current Situation:**
- Node: 152K sats ✅
- LNBits: 0 sats ❌
- Wallet App: Waiting for payment ❌

**What You Need to Do:**
1. **Create invoice in LNBits** (10K sats)
2. **Pay it from your node** (using ThunderHub or LNBits)
3. **Now LNBits has 10K sats**
4. **Send to your wallet app**

**Result:**
- Node: 142K sats ✅
- LNBits: 10K sats ✅
- Wallet App: 10K sats ✅

**The money is there - you just need to move it from the node to the Lightning wallet first!** ⚡💰
