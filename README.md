# iran-ip — Iran IPv4/IPv6 Prefix Lists & Fetcher

**فارسی** | [English](#english)

---

## فارسی

### این پروژه چیه؟

لیست به‌روز Prefixهای IPv4 و IPv6 مربوط به ایران بر اساس داده‌های RIPE Stat.

مناسب برای:

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

### IPv4

- `ipv4.txt`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.txt`

### IPv6

- `ipv6.txt`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv6.txt`

### MikroTik RouterOS Scripts

- `ipv4.rsc`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.rsc`

- `ipv6.rsc`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv6.rsc`

---

## Quick Start

### دانلود مستقیم با curl

```bash
a) IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv6.txt
```

### استفاده در MikroTik

```rsc
/tool fetch url="https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.rsc" dst-path="ipv4.rsc"
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

| Endpoint        | Description                  |
| --------------- | ---------------------------- |
| `GET /health`   | وضعیت سلامت سرویس            |
| `GET /ipv4.txt` | نمایش لیست IPv4              |
| `GET /ipv6.txt` | نمایش لیست IPv6              |
| `GET /ipv4.rsc` | دانلود اسکریپت MikroTik IPv4 |
| `GET /ipv6.rsc` | دانلود اسکریپت MikroTik IPv6 |

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
:local url "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.rsc"

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
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.txt`

### IPv6

- `ipv6.txt`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv6.txt`

### MikroTik RouterOS Scripts

- `ipv4.rsc`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.rsc`

- `ipv6.rsc`

- `https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv6.rsc`

---

## Quick Start

### Download with curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv6.txt
```

### MikroTik Usage

```rsc
/tool fetch url="https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.rsc" dst-path="ipv4.rsc"
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

| Endpoint        | Description                   |
| --------------- | ----------------------------- |
| `GET /health`   | Service health endpoint       |
| `GET /ipv4.txt` | View IPv4 list                |
| `GET /ipv6.txt` | View IPv6 list                |
| `GET /ipv4.rsc` | Download MikroTik IPv4 script |
| `GET /ipv6.rsc` | Download MikroTik IPv6 script |

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
:local url "https://raw.githubusercontent.com/farshidmousavii/iran-ip/main/ipv4.rsc"

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
