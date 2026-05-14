# iran-ip — Iran IPv4/IPv6 Prefix Lists & Fetcher

**فارسی** | [English](#english)

---

## فارسی

### این پروژه چیه؟

لیست به‌روز Prefixهای IPv4 و IPv6 مربوط به ایران بر اساس داده‌های RIPE Stat.

مناسب برای:

- Clash / Mihomo (proxy rule-provider)
- Sing-box (rule-set)
- Xray / v2fly (routing rules)
- NFTables / ipset (Linux firewall)
- OpenWRT (ipset script)
- MikroTik RouterOS
- Firewall Rules
- Routing Policy
- Traffic Policy / Geo Routing
- Automation Scripts
- Self-hosted IP List Service

پروژه Prefixها را از RIPE Stat دریافت می‌کند، آن‌ها را merge و normalize می‌کند و خروجی‌های آماده استفاده تولید می‌کند.

> این پروژه صرفاً بر اساس داده‌های RIPE Stat کار می‌کند و دقت کامل Geolocation یا Routing را تضمین نمی‌کند.

---

## دانلود مستقیم فایل‌ها

تمامی فایل‌ها در پوشه `dist/` در دسترس هستند:

- **Plain text:** `dist/raw/ipv4.txt`, `dist/raw/ipv6.txt`
- **Clash / Mihomo:** `dist/clash/iran.yaml`
- **Sing-box:** `dist/singbox/iran.json`
- **Xray:** `dist/xray/iran.json`
- **NFTables ipset:** `dist/nftables/iran.ipset`
- **NFTables config:** `dist/nftables/iran.nft`
- **OpenWRT:** `dist/openwrt/iran.sh`
- **MikroTik RouterOS:** `dist/mikrotik/ipv4.rsc`, `dist/mikrotik/ipv6.rsc`

### لینک‌های دانلود مستقیم

| فایل | لینک |
|------|------|
| `raw/ipv4.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv4.txt) |
| `raw/ipv6.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv6.txt) |
| `mikrotik/ipv4.rsc` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc) |
| `mikrotik/ipv6.rsc` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv6.rsc) |
| `clash/iran.yaml` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/clash/iran.yaml) |
| `singbox/iran.json` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/singbox/iran.json) |
| `xray/iran.json` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/xray/iran.json) |
| `nftables/iran.ipset` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.ipset) |
| `nftables/iran.nft` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.nft) |
| `openwrt/iran.sh` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/openwrt/iran.sh) |

---

## Quick Start

### دانلود مستقیم با curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv6.txt

# Clash / Mihomo
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/clash/iran.yaml

# Sing-box
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/singbox/iran.json

# Xray
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/xray/iran.json

# NFTables ipset
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.ipset

# NFTables config
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.nft

# OpenWRT script
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/openwrt/iran.sh

# MikroTik
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv6.rsc
```

### استفاده در نرم‌افزارهای proxy

**Clash / Mihomo** — فایل `iran.yaml` را به عنوان rule-provider تنظیم کنید:

```yaml
rule-providers:
  iran:
    type: http
    behavior: ipcidr
    url: "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/clash/iran.yaml"
    path: ./iran.yaml
    interval: 86400

rules:
  - RULE-SET,iran,DIRECT
  - MATCH,PROXY
```

**Sing-box** — فایل `iran.json` را به عنوان rule-set بارگذاری کنید:

```json
{
  "route": {
    "rule_set": [
      {
        "type": "remote",
        "tag": "iran",
        "format": "source",
        "url": "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/singbox/iran.json"
      }
    ],
    "rules": [
      {
        "rule_set": ["iran"],
        "action": "route",
        "outbound": "direct"
      }
    ]
  }
}
```

**Xray** — فایل `iran.json` را در routing rules استفاده کنید:

```json
{
  "routing": {
    "domainStrategy": "IPIfNonMatch",
    "rules": [
      {
        "type": "field",
        "ip": ["geoip:ir"],
        "outboundTag": "direct"
      }
    ]
  }
}
```

### استفاده در Linux Firewall

**ipset restore:**

```bash
ipset restore < iran.ipset
```

**nftables:**

```nft
include "/etc/nftables/iran.nft"

chain prerouting {
  type filter hook prerouting priority 0;
  ip daddr @iran-v4 return
  ip6 daddr @iran-v6 return
}
```

### استفاده در OpenWRT

```bash
# کپی اسکریپت به روتر و اجرا
scp openwrt/iran.sh root@192.168.1.1:/etc/iran-ip.sh
ssh root@192.168.1.1 "sh /etc/iran-ip.sh"
```

### استفاده در MikroTik

```rsc
/tool fetch url="https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc" dst-path="ipv4.rsc"
/import ipv4.rsc
```

---

## دو روش استفاده

### ۱. دانلود مستقیم از GitHub

به وسیله GitHub Actions هر ۶ ساعت فایل‌ها را به‌روزرسانی می‌شود.

فقط فایل‌ها را دانلود کنید و در شبکه یا فایروال خود استفاده کنید.

### ۲. Self-hosted Web Server

پروژه را روی سرور خود اجرا کنید.

ویژگی‌ها:

- دریافت خودکار Prefixها در startup
- ارائه فایل‌ها از طریق HTTP
- به‌روزرسانی دوره‌ای در پس‌زمینه
- استفاده از cache روی دیسک هنگام قطعی اینترنت
- Health Check داخلی
- Graceful Shutdown
- اجرای non-root در Docker

---

## نصب

```bash
git clone https://github.com/farshidmousavii/iran-ip.git
cd iran-ip

go run ./cmd/
```

---

## Docker

### Docker Compose (پیشنهادی)

```bash
docker compose up -d
```

فایل‌های تولید شده در پوشه `data/` در دسترس خواهند بود.

### Docker Manual

```bash
docker build -t iran-ip .

docker run -d \
  --name iran-ip \
  -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  -w /app/data \
  iran-ip
```

کانتینر با کاربر non-root اجرا می‌شود و دارای HEALTHCHECK داخلی است.

---

## Flags

| Flag          | Default | Description                           |
| ------------- | ------- | ------------------------------------- |
| `-addr`       | `:8080` | آدرس وب سرور                          |
| `-refresh`    | `6h`    | فاصله به‌روزرسانی خودکار              |
| `-fetch-only` | `false` | فقط دریافت فایل‌ها بدون اجرای وب سرور |

---

## نحوه اجرا

### حالت پیش‌فرض

```bash
go run ./cmd/
```

### پورت دلخواه

```bash
go run ./cmd/ -addr :9090
```

### به‌روزرسانی هر ۱ ساعت

```bash
go run ./cmd/ -refresh 1h
```

### فقط دریافت فایل‌ها

```bash
go run ./cmd/ -fetch-only
```

---

## Web Endpoints

| Endpoint                    | Description                     |
| --------------------------- | ------------------------------- |
| `GET /health`               | وضعیت سلامت سرویس               |
| `GET /ipv4.txt`             | نمایش لیست IPv4                 |
| `GET /ipv6.txt`             | نمایش لیست IPv6                 |
| `GET /ipv4.rsc`             | دانلود اسکریپت MikroTik IPv4    |
| `GET /ipv6.rsc`             | دانلود اسکریپت MikroTik IPv6    |
| `GET /clash/iran.yaml`      | دانلود rule-provider Clash      |
| `GET /singbox/iran.json`    | دانلود rule-set Sing-box        |
| `GET /xray/iran.json`       | دانلود routing rules Xray       |
| `GET /nftables/iran.ipset`  | دانلود اسکریپت ipset restore    |
| `GET /nftables/iran.nft`    | دانلود کانفیگ nftables          |
| `GET /openwrt/iran.sh`      | دانلود اسکریپت OpenWRT          |

تمام endpointها دارای:

```text
Cache-Control: public, max-age=21600
```

هستند.

---

## Health Check

اندپوینت `/health` خروجی JSON برمی‌گرداند:

### در حال initialization

```json
{ "status": "initializing" }
```

HTTP Status:

```text
503
```

### خطا در آخرین دریافت

```json
{ "status": "stale", "last_fetch": "...", "last_error": "..." }
```

HTTP Status:

```text
503
```

### وضعیت سالم

```json
{ "status": "ok", "last_fetch": "..." }
```

HTTP Status:

```text
200
```

---

## MikroTik Import

### Import مستقیم

```rsc
/import ipv4.rsc
/import ipv6.rsc
```

### دریافت مستقیم از سرور یا GitHub

```rsc
:local fileName "IP.rsc"
:local url "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc"

/tool fetch url=$url dst-path=$fileName mode=http

:if ([:len [/file find name=$fileName]] = 0) do={
    :log error "Fetch failed - file not found"
    :return
}

:if ([/file get $fileName size] < 10) do={
    :log error "File too small - abort"
    /file remove $fileName
    :return
}

:log info "Importing $fileName"
/import file-name=$fileName
:log info "Import done"

/file remove $fileName
```

---

## Data Source

Data is fetched from:

- RIPE Stat API
- Country Resource List (IR)

Project source:

- [https://stat.ripe.net/](https://stat.ripe.net/)

---

## License

MIT

---

# English

## What is this?

Maintained IPv4 and IPv6 prefix lists for Iran based on RIPE Stat data.

Useful for:

- Clash / Mihomo (proxy rule-provider)
- Sing-box (rule-set)
- Xray / v2fly (routing rules)
- NFTables / ipset (Linux firewall)
- OpenWRT (ipset script)
- MikroTik RouterOS
- Firewall Rules
- Routing Policies
- Traffic Engineering
- Automation Scripts
- Self-hosted IP List Services

The project fetches announced IP prefixes for Iran from RIPE Stat, merges and normalizes them, and generates ready-to-use output files.

> This project relies on RIPE Stat data and does not guarantee perfect geolocation or routing accuracy.

---

## Direct Downloads

### IPv4

- `ipv4.txt`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv4.txt`

### IPv6

- `ipv6.txt`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv6.txt`

### Clash / Mihomo

- `clash/iran.yaml`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/clash/iran.yaml`

### Sing-box

- `singbox/iran.json`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/singbox/iran.json`

### Xray

- `xray/iran.json`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/xray/iran.json`

### NFTables

- `nftables/iran.ipset`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.ipset`

- `nftables/iran.nft`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.nft`

### OpenWRT

- `openwrt/iran.sh`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/openwrt/iran.sh`

### MikroTik RouterOS

- `mikrotik/ipv4.rsc`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc`

- `mikrotik/ipv6.rsc`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv6.rsc`

---

## Quick Start

### Download with curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/raw/ipv6.txt

# Clash / Mihomo
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/clash/iran.yaml

# Sing-box
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/singbox/iran.json

# Xray
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/xray/iran.json

# NFTables ipset
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.ipset

# NFTables config
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/nftables/iran.nft

# OpenWRT script
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/openwrt/iran.sh

# MikroTik
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv6.rsc
```

### Proxy Software Usage

**Clash / Mihomo** — use `iran.yaml` as a rule-provider:

```yaml
rule-providers:
  iran:
    type: http
    behavior: ipcidr
    url: "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/clash/iran.yaml"
    path: ./iran.yaml
    interval: 86400

rules:
  - RULE-SET,iran,DIRECT
  - MATCH,PROXY
```

**Sing-box** — load `iran.json` as a rule-set:

```json
{
  "route": {
    "rule_set": [
      {
        "type": "remote",
        "tag": "iran",
        "format": "source",
        "url": "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/singbox/iran.json"
      }
    ],
    "rules": [
      {
        "rule_set": ["iran"],
        "action": "route",
        "outbound": "direct"
      }
    ]
  }
}
```

**Xray** — use `iran.json` in routing rules:

```json
{
  "routing": {
    "domainStrategy": "IPIfNonMatch",
    "rules": [
      {
        "type": "field",
        "ip": ["geoip:ir"],
        "outboundTag": "direct"
      }
    ]
  }
}
```

### Linux Firewall Usage

**ipset restore:**

```bash
ipset restore < iran.ipset
```

**nftables:**

```nft
include "/etc/nftables/iran.nft"

chain prerouting {
  type filter hook prerouting priority 0;
  ip daddr @iran-v4 return
  ip6 daddr @iran-v6 return
}
```

### OpenWRT Usage

```bash
# Copy script to router and run
scp openwrt/iran.sh root@192.168.1.1:/etc/iran-ip.sh
ssh root@192.168.1.1 "sh /etc/iran-ip.sh"
```

### MikroTik Usage

```rsc
/tool fetch url="https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc" dst-path="ipv4.rsc"
/import ipv4.rsc
```

---

## Two Usage Modes

### 1. Download from GitHub

GitHub Actions automatically refreshes the files every 6 hours.

Simply download the files and use them in your firewall, router, or automation setup.

### 2. Self-hosted Web Server

Run the project on your own server.

Features:

- Fetches prefixes on startup
- Serves files via HTTP
- Background auto-refresh
- Disk cache fallback
- Built-in health checks
- Graceful shutdown
- Non-root Docker runtime

---

## Installation

```bash
git clone https://github.com/farshidmousavii/iran-ip.git
cd iran-ip

go run ./cmd/
```

---

## Docker

### Docker Compose (recommended)

```bash
docker compose up -d
```

Generated files will be available in the `data/` directory.

### Docker

```bash
docker build -t iran-ip .

docker run -d \
  --name iran-ip \
  -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  -w /app/data \
  iran-ip
```

Container runs as non-root user and includes a built-in health check.

---

## Flags

| Flag          | Default | Description               |
| ------------- | ------- | ------------------------- |
| `-addr`       | `:8080` | Web server listen address |
| `-refresh`    | `6h`    | Auto-refresh interval     |
| `-fetch-only` | `false` | Fetch files and exit      |

---

## Usage

### Default Mode

```bash
go run ./cmd/
```

### Custom Listen Address

```bash
go run ./cmd/ -addr :9090
```

### Refresh Every Hour

```bash
go run ./cmd/ -refresh 1h
```

### Fetch-only Mode

```bash
go run ./cmd/ -fetch-only
```

---

## Web Endpoints

| Endpoint                    | Description                       |
| --------------------------- | --------------------------------- |
| `GET /health`               | Service health endpoint           |
| `GET /ipv4.txt`             | View IPv4 list                    |
| `GET /ipv6.txt`             | View IPv6 list                    |
| `GET /ipv4.rsc`             | Download MikroTik IPv4 script     |
| `GET /ipv6.rsc`             | Download MikroTik IPv6 script     |
| `GET /clash/iran.yaml`      | Download Clash rule-provider      |
| `GET /singbox/iran.json`    | Download Sing-box rule-set        |
| `GET /xray/iran.json`       | Download Xray routing rules       |
| `GET /nftables/iran.ipset`  | Download ipset restore script     |
| `GET /nftables/iran.nft`    | Download nftables config          |
| `GET /openwrt/iran.sh`      | Download OpenWRT script           |

All file endpoints include:

```text
Cache-Control: public, max-age=21600
```

---

## Health Check

The `/health` endpoint returns JSON:

### Initializing

```json
{ "status": "initializing" }
```

HTTP Status:

```text
503
```

### Last Fetch Failed

```json
{ "status": "stale", "last_fetch": "...", "last_error": "..." }
```

HTTP Status:

```text
503
```

### Healthy

```json
{ "status": "ok", "last_fetch": "..." }
```

HTTP Status:

```text
200
```

---

## MikroTik Import

### Direct Import

```rsc
/import ipv4.rsc
/import ipv6.rsc
```

### Fetch Directly from GitHub or Self-hosted Server

```rsc
:local fileName "IP.rsc"
:local url "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/dist/mikrotik/ipv4.rsc"

/tool fetch url=$url dst-path=$fileName mode=http

:if ([:len [/file find name=$fileName]] = 0) do={
    :log error "Fetch failed - file not found"
    :return
}

:if ([/file get $fileName size] < 10) do={
    :log error "File too small - abort"
    /file remove $fileName
    :return
}

:log info "Importing $fileName"
/import file-name=$fileName
:log info "Import done"

/file remove $fileName
```

---

## Data Source

Data is fetched from:

- RIPE Stat API
- Country Resource List (IR)

Project source:

- [https://stat.ripe.net/](https://stat.ripe.net/)

---

## License

MIT
