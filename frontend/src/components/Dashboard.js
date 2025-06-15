import React, { useEffect, useState } from 'react';
import DataTable from './DataTable';
import Pagination from './Pagination';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  Legend,
} from 'recharts';

export default function Dashboard() {
  const pageSize = 10;
  const [page, setPage] = useState(1);
  const [countryData, setCountryData] = useState([]);
  const [totalCount, setTotalCount] = useState(0);

  const [topProducts, setTopProducts] = useState([]);
  const [monthlySales, setMonthlySales] = useState([]);
  const [topRegions, setTopRegions] = useState([]);

  useEffect(() => {
    fetch(`/api/revenue/countries?limit=${pageSize}&offset=${(page - 1) * pageSize}`)
      .then(res => res.json())
      .then(data => {
        setCountryData(data.data || []);
        setTotalCount(data.total || 0);
      });
  }, [page]);

  useEffect(() => {
    fetch('/api/products/top?limit=20')
      .then(res => res.json())
      .then(data => setTopProducts(data));

    fetch('/api/sales/monthly')
      .then(res => res.json())
      .then(data => setMonthlySales(data));

    fetch('/api/regions/top?limit=30')
      .then(res => res.json())
      .then(data => setTopRegions(data));
  }, []);

  const totalPages = Math.ceil(totalCount / pageSize);

  return (
    <div className="dashboard">
      <h1>ABT Dashboard</h1>

      <section>
        <h2>Country-Level Revenue</h2>
        <DataTable
          columns={[
            { header: 'Country', accessor: 'country' },
            { header: 'Product', accessor: 'product_name' },
            { header: 'Total Revenue', accessor: 'total_revenue' },
            { header: '# Transactions', accessor: 'transaction_count' },
          ]}
          data={countryData}
        />
        {totalPages > 1 && (
          <Pagination
            currentPage={page}
            totalPages={totalPages}
            onPageChange={setPage}
          />
        )}
      </section>

      {/* Top 20 Products Chart - increased height for readability */}
      <section style={{ height: 500, marginTop: '2rem' }}>
        <h2>Top 20 Products by Purchase Count</h2>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={topProducts}
            layout="vertical"
            barSize={18}
            barGap={6}
            barCategoryGap="15%"
            margin={{ top: 20, right: 30, left: 140, bottom: 20 }}
          >
            <CartesianGrid stroke="#e0e0e0" strokeDasharray="4 4" />
            <XAxis type="number" tick={{ fontSize: 12 }} />
            <YAxis dataKey="product_name" type="category" tick={{ fontSize: 12 }} width={200} />
            <Tooltip wrapperStyle={{ fontSize: '0.9rem' }} />
            <Legend iconType="circle" />
            <Bar dataKey="purchase_count" name="Purchases" fill="#8884d8" />
            <Bar dataKey="stock_quantity" name="Stock Qty" fill="#82ca9d" />
          </BarChart>
        </ResponsiveContainer>
      </section>

      {/* Monthly Sales Volume Chart */}
      <section style={{ height: 300, marginTop: '2rem' }}>
        <h2>Monthly Sales Volume</h2>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={monthlySales}
            margin={{ top: 20, right: 30, left: 20, bottom: 20 }}
            barSize={20}
          >
            <CartesianGrid stroke="#e0e0e0" strokeDasharray="4 4" />
            <XAxis dataKey="month" tick={{ fontSize: 12 }} />
            <YAxis tick={{ fontSize: 12 }} />
            <Tooltip wrapperStyle={{ fontSize: '0.9rem' }} />
            <Bar dataKey="sales_volume" name="Units Sold" fill="#413ea0" />
          </BarChart>
        </ResponsiveContainer>
      </section>

      {/* Top 30 Regions Chart - extended height for clarity */}
      <section style={{ height: 600, marginTop: '2rem' }}>
        <h2>Top 30 Regions by Revenue</h2>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={topRegions}
            layout="vertical"
            barSize={14}
            barGap={4}
            barCategoryGap="20%"
            margin={{ top: 20, right: 30, left: 180, bottom: 20 }}
          >
            <CartesianGrid stroke="#e0e0e0" strokeDasharray="4 4" />
            <XAxis type="number" tick={{ fontSize: 12 }} />
            <YAxis dataKey="region" type="category" tick={{ fontSize: 12 }} width={220} />
            <Tooltip wrapperStyle={{ fontSize: '0.9rem' }} />
            <Legend iconType="circle" />
            <Bar dataKey="total_revenue" name="Revenue" fill="#82ca9d" />
            <Bar dataKey="items_sold" name="Items Sold" fill="#ffc658" />
          </BarChart>
        </ResponsiveContainer>
      </section>
    </div>
  );
}