# 📊 ABT Corp CSV Analytics Dashboard (Golang)

This Go backend powers a high-performance analytics dashboard that processes large transaction CSV files to deliver actionable business insights in under 10 seconds.

---

## 🔍 Table of Contents

1. [Overview](#overview)  
2. [Quick Start](#quick-start)  
   - [Prerequisites](#prerequisites)  
   - [Installation & Run](#installation--run)  
3. [CSV Input Format](#csv-input-format)  
4. [API Endpoints](#api-endpoints)  
5. [Test & Coverage](#test--coverage)  
6. [Project Structure](#project-structure)  
7. [Extensibility](#extensibility)  
8. [Contributing & License](#contributing--license)  

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
cd csv-analytics-api

go mod tidy
go run main.go -csv=/path/to/data.csv
