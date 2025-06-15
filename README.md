# 📊 ABT Corp CSV Analytics Dashboard (Golang)

This Go backend powers a high-performance analytics dashboard that processes large transaction CSV files to deliver actionable business insights in under 10 seconds.

---

## 🔍 Table of Contents

1. [Overview](#overview)  
2. [Quick Start](#quick-start)  
   - [Prerequisites](#prerequisites)  
   - [Installation & Run](#installation--run)  
3. [API Endpoints](#api-endpoints)  
4. [Test & Coverage](#test--coverage)  
5. [Project Structure](#project-structure)  
---

## 📌 Overview

ABT Corp requires a fast, reusable solution for generating key insights from their transaction data:

- **Country-level revenue** (by product) sorted descending  
- **Top 20** most purchased products (+ stock)  
- **Monthly sales volume** visualization  
- **Top 30** regions by revenue and items sold  

This backend reads a multi-million‑row CSV (~5 M+ records) using buffered streaming, then uses Go concurrency (worker pool + goroutines + channels) to preprocess data in parallel. It aggregates results into summarized slices and serves them as JSON via Gin REST endpoints, ensuring memory efficiency and fast response times under 10 seconds. Go’s buffered I/O, pooling, and optimized data structures are used to maintain high throughput.


---

## 🚀 Quick Start

### Prerequisites

#### Backend (Go)
- **Go** version **1.20 or later**
- **Git** (for cloning the repository)
- **CSV data file** with **13 columns**:

#### Frontend (React Dashboard)
- **Node.js** version **16 or later**
- **npm** version ≥ 8

### Installation & Run

### 01. Clone the Repository
```bash
git clone https://github.com/GimhaniHM/Go-Technical-Assessment.git

### 02. Set Up the Backend (Go)
(I) cd  Go-Technical-Assessment/backend

(I) Make sure to place CSV file inside the folder
/cmd/app/data as GO_test_5m.csv

(II) install go dependencies
inside backend folder

go mod tidy

(III) run gin server
cd backend/cmd/app
go run main.go
 then gin server will run on http://localhost:8090

03. Set Up the Frontend (React Dashboard)

(I)cd frontend
(II) install the dependencies
npm install
(III) to run dashboard on your browser
npm start
run it on http://localhost:3000

###API Endpoints

| Endpoint                     | Description                              | Query Params          |
| ---------------------------- | ---------------------------------------- | --------------------- |
| `GET /api/revenue/countries` | Country + Product revenue (paginated)    | `limit`, `offset`     |
| `GET /api/products/top`      | Top N products by purchase count & stock | `limit` (default: 20) |
| `GET /api/sales/monthly`     | Monthly aggregated sales volume          | —                     |
| `GET /api/regions/top`       | Top N regions by revenue & item count    | `limit` (default: 30) |


###Test & Coverage
Run tests with coverage:
cd backend
use cmd terminal

go test ./internal/... -coverprofile=coverage.out
 then run below command to create html coverage report
go tool cover -html=coverage.out -o coverage.html
to open coverage report in a browser 
start coverage.html

### Project Structure
backend/
├── internal/
│   ├── handlers/     # Gin HTTP endpoints
│   ├── services/     # Aggregation & business logic
│   ├── utils/        # CSV parsing helper
│   └── models/       # Data transfer objects
├── main.go           # Server entry point
└── README.md         # This file
frontend/

QORIA-TECHNICAL-ASSESSMENT/
└── backend/
    ├── cmd/
    │   └── app/
    │       └── main.go #
    ├── internal/
    │   ├── handlers/   # Gin HTTP endpoints
    │   │   ├── insght_handler.go   #
    │   │   ├── revenue_handler.go   #
    │   │   └── revenue_handler_test.go   #
    │   ├── models/   # Data transfer objects
    │   │   └── models.go
    │   ├── services/   # Aggregation & business logic
    │   │   ├── aggregator.go   #
    │   │   ├── aggregator_test.go   #
    │   │   └── concurrent_aggregator.go   #
    │   └── utils/   # CSV streaming helper
            |--- csvstream.go   #
            |--- csvstream_test.go   #
    ├── go.mod
    └── go.sum


QORIA-TECHNICAL-ASSESSMENT/
└── frontend/
    ├── node_modules/
    ├── public/
    ├── src/
    │   ├── components/
    │   │   ├── Dashboard.js
    │   │   ├── DataTable.js
    │   │   ├── Pagination.css
    │   │   └── Pagination.js
    │   ├── App.css
    │   ├── App.js
    │   ├── index.css
    │   └── index.js
    ├── .gitignore
    ├── package-lock.json
    ├── package.json
    ├── README.md
    └── yarnlock

