# ⚡ LNO1 vs LNTB - What's the Difference?

## 🎯 Quick Answer

**You have:** `lno1...` (Lightning Offer - reusable address)  
**You need:** `lntb...` (Lightning Invoice - payment request)  
**To get money from Voltage:** Create an INVOICE, not use your OFFER

---

## 📊 The Difference

### LNO1 = Lightning OFFER (BOLT12)

```
lno1qgsyxjtl6luzd9t3pr62xr7eemp6awnejusgf6gw45q75vcfqqqqqqqgq...
```

**What it is:**
- 🔁 **Reusable payment address**
- 📧 Like an email address for Lightning
- 🆕 Newer technology (BOLT12)
- 🎯 People fetch invoices FROM this

**What it's NOT:**
- ❌ Not a payment request
- ❌ Can't paste it into "Pay" field
- ❌ Not all wallets support it yet (including Voltage)

**Think of it like:**
- Your permanent Venmo username (@yourname)
- People look you up and then create a payment

---

### LNTB = Lightning INVOICE (BOLT11)

```
lntb50u1pn2r3a5pp5abc123def456ghi789jkl...
```

**What it is:**
- 💸 **One-time payment request**
- 🧾 Like a bill or invoice
- 💰 Has specific amount encoded (e.g., 5000 sats)
- ✅ Works everywhere, including Voltage

**What it does:**
- ✅ Can be paid immediately
- ✅ Paste into any "Send/Pay" field
- ✅ Standard Lightning payment

**Think of it like:**
- A specific payment request for $50
- You hand someone this exact bill to pay

---

## 📱 How to Get an INVOICE (LNTB) in Your Wallet

### Visual Guide:

```
Your Wallet App
├── Tab 1: "Receive" or "Request"
│   ├── Enter amount: 5000 sats
│   ├── Click "Create Invoice"
│   └── ✅ Generates: lntb... (THIS IS WHAT YOU NEED!)
│
└── Tab 2: "Share" or "Your Address"
    └── Shows: lno1... (Your permanent offer)
```

### Step-by-Step:

1. **Open your Lightning wallet**
2. **Tap "Receive"** (NOT "Share Address")
3. **Enter amount:** `5000` sats
4. **Tap "Create" or "Generate Invoice"**
5. **Copy the text that starts with `lntb`**

### Common Wallet Locations:

**Phoenix Wallet:**
- Bottom bar → "Receive" → Enter amount → "Create invoice"

**Breez:**
- "+" button → "Receive" → Enter amount → "Create"

**Zeus:**
- "Receive" tab → Enter amount → "Create Invoice"

**Blue Wallet:**
- Lightning wallet → "Receive" → Enter amount

**Wallet of Satoshi:**
- "Receive" → Enter amount → "Create Request"

---

## 🔄 The Process

### What You're Currently Doing (Doesn't Work):

```
Your Offer (lno1...) 
    ↓
Voltage "Pay" field
    ↓
❌ Error or doesn't recognize it
```

### What You NEED to Do:

```
Your Wallet → Create Invoice (enter 5000 sats)
    ↓
Generates: lntb... invoice
    ↓
Copy the lntb... string
    ↓
Voltage → "Pay Invoice" → Paste lntb... → Pay
    ↓
✅ Money arrives in your wallet!
```

---

## 🎯 Simple Test Flow

### Step 1: In YOUR Wallet
1. Tap "Receive"
2. Enter: 5000
3. Tap "Create Invoice"
4. **Copy the `lntb...` string** (not lno1!)

### Step 2: In Voltage
1. Go to https://voltage.cloud/
2. Open your testnet node
3. Click "Send" or "Pay Invoice"
4. **Paste the `lntb...` string**
5. Click "Pay"

### Step 3: Check Your Wallet
- ✅ 5000 sats should arrive instantly!
- ✅ Payment complete!

---

## ❓ FAQ

**Q: Why can't I just use my lno1 offer?**  
A: It's newer tech (BOLT12). Voltage and many wallets don't support paying to offers directly yet.

**Q: Is my lno1 offer useless?**  
A: No! Share it with people who have BOLT12-compatible wallets. They can fetch invoices from it and pay you multiple times.

**Q: How do I know if I have the right one?**  
A: For testnet: Must start with `lntb`  
   For mainnet: Would start with `lnbc`

**Q: The invoice expires?**  
A: Yes, usually after 1 hour. Create a new one if it expires.

**Q: Can I create multiple invoices?**  
A: Yes! Create as many as you want. Each is a separate payment request.

---

## 🎯 Bottom Line

**Your LNO1 offer is GOOD** - it's your permanent Lightning address!

**But to receive payment from Voltage NOW:**
1. Create a regular invoice (`lntb...`)
2. Use the "Receive" function in your wallet
3. Enter a specific amount
4. Copy and give to Voltage

**It's like the difference between:**
- Giving someone your email address (lno1 = they can reach out)
- Sending someone a specific invoice (lntb = pay this exact amount now)

---

**For this test: Use the INVOICE method!** ⚡

Then your permanent lno1 offer is great to share with others!

