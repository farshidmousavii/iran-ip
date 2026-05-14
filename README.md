# iran-ip-ranges — Iran IPv4/IPv6 Prefix Lists & Fetcher

**فارسی** | [English](#english)

---

## فارسی

### این پروژه چیه؟

ترافیک ایرانی را به‌طور خودکار از مسیر VPN یا پروکسی یا روتر  با استفاده از لیست‌های به‌روز IPv4/IPv6 ایران خارج کنید .

مناسب برای سرورهای VPN، روترها، فایروال‌ها و سرورهای پروکسی.

این پروژه هر ۶ ساعت Prefixهای اعلام شده ایران را از RIPE Stat دریافت، merge و normalize می‌کند و خروجی‌های آماده استفاده در قالب‌های مختلف تولید می‌کند.

> این پروژه صرفاً بر اساس داده‌های RIPE Stat کار می‌کند و دقت کامل Geolocation یا Routing را تضمین نمی‌کند.

---

## دانلود مستقیم فایل‌ها

تمامی فایل‌ها در پوشه `dist/` در دسترس هستند:

- **Plain text:** `dist/raw/ipv4.txt`, `dist/raw/ipv6.txt`
- **Clash / Mihomo:** `dist/clash/iran.yaml`
- **Sing-box:** `dist/sing-box/iran.json`
- **Xray:** `dist/xray/iran.json`
- **NFTables ipset:** `dist/firewall/iran.ipset`
- **NFTables config:** `dist/firewall/iran.nft`
- **OpenWRT:** `dist/openwrt/iran.sh`
- **MikroTik RouterOS:** `dist/routeros/ipv4.rsc`, `dist/routeros/ipv6.rsc`

> فایل‌ها هر روز به‌صورت خودکار در **GitHub Releases** منتشر می‌شوند. می‌توانید همه فایل‌ها را به صورت یکجا از [آخرین Release](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest) دانلود کنید.

### لینک‌های دانلود مستقیم

| فایل | لینک |
|------|------|
| `raw/ipv4.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv4.txt) |
| `raw/ipv6.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv6.txt) |
| `routeros/ipv4.rsc` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc) |
| `routeros/ipv6.rsc` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc) |
| `clash/iran.yaml` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml) |
| `sing-box/iran.json` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json) |
| `xray/iran.json` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json) |
| `firewall/iran.ipset` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset) |
| `firewall/iran.nft` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft) |
| `openwrt/iran.sh` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh) |
| همه فایل‌ها (zip) | [دانلود آخرین Release](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest) |

---

## Quick Start

### دانلود مستقیم با curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv6.txt

# Clash / Mihomo
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml

# Sing-box
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json

# Xray
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json

# NFTables ipset
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset

# NFTables config
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft

# OpenWRT script
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh

# MikroTik
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc
```

### راهنمای تنظیمات آماده

فایل‌های مثال آماده در پوشه `examples/` قرار دارند — هر کدام شامل تنظیمات کامل برای یک ابزار خاص:

| ابزار | مسیر مثال |
|------|-----------|
| Clash / Mihomo | [`examples/clash/config.yaml`](examples/clash/config.yaml) |
| Sing-box | [`examples/sing-box/config.json`](examples/sing-box/config.json) |
| Xray | [`examples/xray/config.json`](examples/xray/config.json) |
| NFTables + ipset | [`examples/nftables/rules.nft`](examples/nftables/rules.nft) |
| OpenWRT (firewall) | [`examples/openwrt/firewall-config.sh`](examples/openwrt/firewall-config.sh) |
| MikroTik RouterOS | [`examples/mikrotik/import-script.rsc`](examples/mikrotik/import-script.rsc) |
| Split Tunnel (iptables) | [`examples/split-tunnel.sh`](examples/split-tunnel.sh) |

---

## دو روش استفاده

### ۱. دانلود مستقیم از GitHub

به وسیله GitHub Actions هر ۶ ساعت فایل‌ها به‌روزرسانی می‌شوند.

**دو روش دانلود:**
- **فایل‌های تکی** — هر فایل به‌صورت جداگانه از پوشه `dist/` در دسترس است
- **بسته کامل (zip)** — همه فایل‌ها به‌صورت یکجا در [GitHub Releases](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest) هر روز منتشر می‌شوند

### ۲. روش Self-hosted Web Server

پروژه را روی سرور خود اجرا کنید.

ویژگی‌ها:

- دریافت خودکار Prefixها در startup
- ارائه فایل‌ها از طریق HTTP
- به‌روزرسانی دوره‌ای در پس‌زمینه
- استفاده از cache روی دیسک هنگام قطعی اینترنت
- استفاده از Health Check داخلی
- دارای Graceful Shutdown
- اجرای non-root در Docker

---

## نصب

```bash
git clone https://github.com/farshidmousavii/iran-ip-ranges.git
cd iran-ip-ranges

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
docker build -t iran-ip-ranges .

docker run -d \
  --name iran-ip-ranges \
  -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  -w /app/data \
  iran-ip-ranges
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
| `GET /ipv4.rsc`             | دانلود اسکریپت RouterOS IPv4    |
| `GET /ipv6.rsc`             | دانلود اسکریپت RouterOS IPv6    |
| `GET /clash/iran.yaml`      | دانلود rule-provider Clash      |
| `GET /sing-box/iran.json`    | دانلود rule-set Sing-box        |
| `GET /xray/iran.json`       | دانلود routing rules Xray       |
| `GET /firewall/iran.ipset`  | دانلود اسکریپت ipset restore    |
| `GET /firewall/iran.nft`    | دانلود کانفیگ nftables          |
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

## منبع داده ها

داده ها از منابع زیر دریافت میگردند :

- RIPE Stat API
- Country Resource List (IR)

---

## مجوز

MIT

---

# English

## What is this?

Automatically route Iranian traffic outside your VPN or proxy using continuously updated Iran CIDR ranges.

Designed for VPN servers, routers, firewalls, and proxy servers.

This project fetches announced IP prefixes for Iran from RIPE Stat every 6 hours, merges and normalizes them, and generates ready-to-use output files in multiple formats.

> This project relies on RIPE Stat data and does not guarantee perfect geolocation or routing accuracy.

---

## Direct Downloads

All output files are available under the `dist/` directory.

> Daily **GitHub Releases** are also published — download everything as a single archive from the [latest release](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest).

### Clash / Mihomo

- `clash/iran.yaml`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml`

### Sing-box

- `sing-box/iran.json`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json`

### Xray

- `xray/iran.json`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json`

### NFTables

- `firewall/iran.ipset`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset`

- `firewall/iran.nft`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft`

### OpenWRT

- `openwrt/iran.sh`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh`

### MikroTik RouterOS

- `routeros/ipv4.rsc`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc`

- `routeros/ipv6.rsc`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc`

---

## Quick Start

### Download with curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv6.txt

# Clash / Mihomo
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml

# Sing-box
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json

# Xray
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json

# NFTables ipset
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset

# NFTables config
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft

# OpenWRT script
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh

# MikroTik
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc
```

### Example Configurations

Ready-to-use example configs are available in the `examples/` directory:

| Tool | Example Path |
|------|-------------|
| Clash / Mihomo | [`examples/clash/config.yaml`](examples/clash/config.yaml) |
| Sing-box | [`examples/sing-box/config.json`](examples/sing-box/config.json) |
| Xray | [`examples/xray/config.json`](examples/xray/config.json) |
| NFTables + ipset | [`examples/nftables/rules.nft`](examples/nftables/rules.nft) |
| OpenWRT (firewall) | [`examples/openwrt/firewall-config.sh`](examples/openwrt/firewall-config.sh) |
| MikroTik RouterOS | [`examples/mikrotik/import-script.rsc`](examples/mikrotik/import-script.rsc) |
| Split Tunnel (iptables) | [`examples/split-tunnel.sh`](examples/split-tunnel.sh) |

---

## Two Usage Modes

### 1. Download from GitHub

GitHub Actions automatically refreshes the files every 6 hours.

**Two download options:**
- **Individual files** — each file available directly from the `dist/` folder
- **Full archive (zip)** — all files bundled in daily [GitHub Releases](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest)

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
git clone https://github.com/farshidmousavii/iran-ip-ranges.git
cd iran-ip-ranges

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
docker build -t iran-ip-ranges .

docker run -d \
  --name iran-ip-ranges \
  -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  -w /app/data \
  iran-ip-ranges
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
| `GET /ipv4.rsc`             | Download RouterOS IPv4 script     |
| `GET /ipv6.rsc`             | Download RouterOS IPv6 script     |
| `GET /clash/iran.yaml`      | Download Clash rule-provider      |
| `GET /sing-box/iran.json`    | Download Sing-box rule-set        |
| `GET /xray/iran.json`       | Download Xray routing rules       |
| `GET /firewall/iran.ipset`  | Download ipset restore script     |
| `GET /firewall/iran.nft`    | Download nftables config          |
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

## Data Source

Data is fetched from:

- RIPE Stat API
- Country Resource List (IR)

---

## License

MIT
