# 📊 ABT Corp CSV Analytics Dashboard (Golang + React)

This repository contains a **Go backend** and **React frontend** for a high-performance analytics dashboard that processes large transaction CSV files (\~5M+ rows) in under 10 seconds, delivering key business insights.

---

## 🔍 Table of Contents

1. [Overview](#overview)
2. [Quick Start](#quick-start)

   * [Prerequisites](#prerequisites)
   * [Backend Setup](#backend-setup)
   * [Frontend Setup](#frontend-setup)
3. [API Endpoints](#api-endpoints)
4. [Testing & Coverage](#testing--coverage)
5. [Project Structure](#project-structure)

---

## 📌 Overview

ABT Corp requires:

* **Country-level Revenue** table (by product), sorted descending
* **Top 20 Products** by purchase count (+ current stock)
* **Monthly Sales Volume** chart
* **Top 30 Regions** by revenue & items sold

This solution:

* **Streams** the CSV via `bufio.Reader` + `encoding/csv`
* Uses a **worker pool** (goroutines + channels) to parse & aggregate in parallel
* Builds in-memory maps & converts them to sorted slices
* Exposes REST JSON endpoints via **Gin**
* Frontend built with **React** + **Recharts**, with pagination & responsive charts

---

## 🚀 Quick Start

### Prerequisites

* **Go** ≥ 1.20
* **Node.js** ≥ 16 & **npm** ≥ 8
* A **data CSV** (`.csv`) with **13 columns** in this order:

  ```csv
  transaction_id,transaction_date,user_id,country,region,
  product_id,product_name,category,price,quantity,
  total_price,stock_quantity,added_date
  ```

### Backend Setup
bash
```
# 1. Clone repo
git clone https://github.com/GimhaniHM/Go-Technical-Assessment.git
cd Go-Technical-Assessment/backend

# 2. Place data CSV file inside cmd/app/data/ folder and name it as GO_test_5m.csv

# 3. Install dependencies
go mod tidy

# 4. Run server (defaults: addr=:8090, workers=CPU count)
cd cmd/app
go run main.go
```

**Verify:**

```bash
curl 'http://localhost:8090/api/revenue/countries?limit=5&offset=0'
```

### Frontend Setup

```bash
cd Go-Technical-Assessment/frontend
npm install
npm start
```

Open: `http://localhost:3000`

---

## 🔗 API Endpoints

| Route                    | Method | Query Params                    | Description                                |
| ------------------------ | ------ | ------------------------------- | ------------------------------------------ |
| `/api/revenue/countries` | GET    | `limit` (default 100), `offset` | Country+product revenue table (paginated). |
| `/api/products/top`      | GET    | `limit` (default 20)            | Top N products by purchase count & stock.  |
| `/api/sales/monthly`     | GET    | —                               | Monthly units sold (chronological).        |
| `/api/regions/top`       | GET    | `limit` (default 30)            | Top N regions by revenue & items sold.     |

---

## 🧪 Testing & Coverage

```bash
# Run unit tests & record coverage
cd backend
go test ./internal/... -coverprofile=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html
```

Open `coverage.html` to inspect coverage. Aim for ≥ 80% in services & utils.

---

## 📂 Project Structure

```
backend/
├── cmd/app/main.go             # Entrypoint, CLI flags & HTTP server
├── internal/
│   ├── handlers/
│   │   └── insight_handler.go  # Gin handlers for each endpoint
│   ├── models/models.go        # Data models & JSON DTOs
│   ├── services/concurrent_aggregator.go  # Streaming + concurrent CSV aggregator
│   └── utils/csvstream.go      # Sequential CSV reader with preprocessing
└── go.mod                      # Go module config

frontend/
├── src/
│   ├── components/
│   │   ├── DataTable.js
│   │   ├── Dashboard.js
│   │   └── Pagination.js
│   ├── App.js
│   └── index.js
├── public/
├── package.json
└── README.md                   # (this file)
```

---

