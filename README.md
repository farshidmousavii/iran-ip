# iran-ip — Iran IPv4/IPv6 Address List Fetcher

[فارسی](#فارسی) | **English**

---

## فارسی

### این پروژه چیه؟

ابزاری که تمام ساب‌نت‌های IPv4 و IPv6 اعلام‌شده برای ایران (IR) رو از API رایپ (RIPE Stat) دریافت می‌کنه، ادغام و نرمال‌سازی می‌کنه و فایل‌های خروجی تولید می‌کنه:

- **`ipv4.txt`** / **`ipv6.txt`** — لیست تمیز ساب‌نت‌های IPv4/IPv6
- **`ipv4.rsc`** / **`ipv6.rsc`** — اسکریپت آماده برای MikroTik RouterOS

### دو روش استفاده

#### ۱. دانلود از گیت‌هاب

GitHub Action هر ۶ ساعت لیست IP ها رو به‌روز می‌کنه. فقط کافیه فایل‌ها رو از ریپازیتوری دانلود کنید و توی شبکه‌تون استفاده کنید.

#### ۲. وب سرور اختصاصی

پروژه رو روی سرور خودتون اجرا کنید. در ابتدا IP ها رو دریافت می‌کنه و از طریق HTTP ارائه می‌ده. یک جاب در پس‌زمینه به صورت خودکار داده‌ها رو به‌روز می‌کنه. اگه اینترنت در دسترس نباشه، از فایل‌های کش روی دیسک استفاده می‌کنه.

### نصب

```bash
git clone https://github.com/farshidmousavii/iran-ip.git
cd iran-ip
go run ./cmd/
```

### داکر

```bash
# ساخت و اجرا با Docker Compose (پیشنهادی)
docker compose up -d
# فایل‌های خروجی توی پوشه data/ در دسترس هستن

# یا با Docker مستقیم
docker build -t iran-ip .

docker run -d --name iran-ip -p 8080:8080 -v $(pwd)/data:/app/data -w /app/data iran-ip
```

کانتینر با کاربر غیر-root اجرا می‌شه و HEALTHCHECK داخلی داره. فایل‌های تولید شده (`ipv4.txt`, `ipv6.txt`, `ipv4.rsc`, `ipv6.rsc`) در پوشه `data/` در مسیر پروژه قابل دسترسی هستن.

### پرچم‌ها (Flags)

| پرچم | پیش‌فرض | توضیحات |
|---|---|---|
| `-addr` | `:8080` | آدرس پورت وب سرور |
| `-refresh` | `6h` | فاصله به‌روزرسانی خودکار (مثلاً `1h30m`) |
| `-fetch-only` | `false` | فقط دریافت IP و ایجاد فایل، بدون وب سرور |

### نحوه اجرا

```bash
# دریافت IP و اجرای وب سرور (حالت پیش‌فرض)
go run ./cmd/

# آدرس پورت دلخواه
go run ./cmd/ -addr :9090

# به‌روزرسانی هر ۱ ساعت به جای ۶ ساعت
go run ./cmd/ -refresh 1h

# فقط دریافت IP (بدون وب سرور)
go run ./cmd/ -fetch-only
```

### آدرس‌های وب سرور

| آدرس | توضیحات |
|---|---|
| `GET /health` | بررسی سلامت سرور (JSON: status, last_fetch, last_error) |
| `GET /ipv4.txt` | نمایش لیست IPv4 در مرورگر |
| `GET /ipv6.txt` | نمایش لیست IPv6 در مرورگر |
| `GET /ipv4.rsc` | دانلود فایل اسکریپت میکروتیک IPv4 |
| `GET /ipv6.rsc` | دانلود فایل اسکریپت میکروتیک IPv6 |

### بررسی سلامت

اندپوینت `/health` اطلاعات سلامت سرور رو به صورت JSON برمی‌گردونه:

- `{"status":"initializing"}` (503) — هنوز هیچ دریافتی انجام نشده
- `{"status":"stale","last_fetch":"...","last_error":"..."}` (503) — آخرین دریافت با خطا مواجه شده
- `{"status":"ok","last_fetch":"..."}` (200) — همه چیز خوبه

### خاموش شدن امن

سرور سیگنال‌های SIGINT/SIGTERM رو مدیریت می‌کنه و بعد از اتمام درخواست‌های در حال اجرا به صورت تمیز خاموش می‌شه.

### وارد کردن در میکروتیک

```
/import ipv4.rsc
/import ipv6.rsc
```

یا دریافت مستقیم از سرور اختصاصی خودتون از داخل RouterOS:

```rsc
:local fileName "IP.rsc"
:local url "http://YOUR_SERVER:8080/ipv4.rsc"

/tool fetch url=$url dst-path=$fileName mode=http

:if ([:len [/file find name=$fileName]] = 0) do={
    :log error "دریافت ناموفق - فایل پیدا نشد"
    :return
}

:if ([/file get $fileName size] < 10) do={
    :log error "فایل خیلی کوچیکه - لغو"
    /file remove $fileName
    :return
}

:log info "در حال導入 $fileName"
/import file-name=$fileName
:log info "تمام شد"

/file remove $fileName
```
---

## English

### What is this?

A tool that fetches all announced IPv4 and IPv6 subnets for Iran (IR) from RIPE Stat API, merges and normalizes them, and generates output files:

- **`ipv4.txt`** / **`ipv6.txt`** — clean lists of IPv4/IPv6 subnets (CIDR notation)
- **`ipv4.rsc`** / **`ipv6.rsc`** — MikroTik RouterOS address-list scripts ready to import

### Two usage modes

#### 1. Download from GitHub

The GitHub Action automatically fetches and updates the IP lists every 6 hours. Just download the files from the repository and use them in your network.

#### 2. Self-hosted web server

Run the project on your own server. It fetches IPs on startup and serves them via HTTP endpoints. A background job refreshes the data automatically. If internet is unavailable, it falls back to the cached files on disk.

### Installation

```bash
git clone https://github.com/farshidmousavii/iran-ip.git
cd iran-ip
go run ./cmd/
```

### Docker

```bash
# Build and run with Docker Compose (recommended)
docker compose up -d
# Generated files are available in the data/ directory

# Or with plain Docker
docker build -t iran-ip .

docker run -d --name iran-ip -p 8080:8080 -v $(pwd)/data:/app/data -w /app/data iran-ip
```

Container runs as non-root user with built-in health check. Generated files (`ipv4.txt`, `ipv6.txt`, `ipv4.rsc`, `ipv6.rsc`) are accessible from the `data/` directory on the host.

### Flags

| Flag | Default | Description |
|---|---|---|
| `-addr` | `:8080` | Web server listen address |
| `-refresh` | `6h` | Auto-refresh interval (e.g., `1h30m`, `30m`) |
| `-fetch-only` | `false` | Fetch IPs, write files, and exit |

### Usage

```bash
# Fetch IPs and start web server (default)
go run ./cmd/

# Custom listen address
go run ./cmd/ -addr :9090

# Refresh every hour instead of 6 hours
go run ./cmd/ -refresh 1h

# Fetch-only mode (no web server)
go run ./cmd/ -fetch-only
```

### Web Endpoints

| Endpoint | Description |
|---|---|
| `GET /health` | Health check (JSON: status, last_fetch, last_error) |
| `GET /ipv4.txt` | View IPv4 list in browser |
| `GET /ipv6.txt` | View IPv6 list in browser |
| `GET /ipv4.rsc` | Download IPv4 MikroTik script |
| `GET /ipv6.rsc` | Download IPv6 MikroTik script |

All file endpoints include `Cache-Control: public, max-age=21600` headers.

### Health Check

The `/health` endpoint returns JSON with status:

- `{"status":"initializing"}` (503) — no fetch completed yet
- `{"status":"stale","last_fetch":"...","last_error":"..."}` (503) — last fetch failed
- `{"status":"ok","last_fetch":"..."}` (200) — everything good

### Graceful Shutdown

The server handles SIGINT/SIGTERM, finishes in-flight requests, and shuts down cleanly.

### MikroTik Import

```
/import ipv4.rsc
/import ipv6.rsc
```

Or fetch from your self-hosted server directly from RouterOS:

```rsc
:local fileName "IP.rsc"
:local url "http://YOUR_SERVER:8080/ipv4.rsc"

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