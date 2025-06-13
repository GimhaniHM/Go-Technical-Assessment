import * as React from 'react';
import { DataGrid } from '@mui/x-data-grid';

const columns = [
  { field: 'id', headerName: 'ID', width: 90 },
  {
    field: 'region',
    headerName: 'Region',
    width: 150,
  },
  {
    field: 'rainfall',
    headerName: 'Rainfall (mm)',
    width: 150,
  },
  {
    field: 'temperature',
    headerName: 'Temperature (°C)',
    width: 170,
  },
];

const rows = [
  { id: 1, region: 'Colombo', rainfall: 120, temperature: 28 },
  { id: 2, region: 'Kandy', rainfall: 95, temperature: 26 },
  { id: 3, region: 'Galle', rainfall: 110, temperature: 27 },
];

export default function DataTable() {
  return (
    <div style={{ height: 400, width: '100%' }}>
      <DataGrid
        rows={rows}
        columns={columns}
        pageSize={5}
        rowsPerPageOptions={[5]}
      />
    </div>
  );
}