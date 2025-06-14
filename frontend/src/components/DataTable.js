import React from 'react';

/**
 * Simple reusable table.
 * @param {{ header: string, accessor: string }[]} columns
 * @param {Object[]} data
 */
export default function DataTable({ columns, data }) {
  return (
    <table className="data-table">
      <thead>
        <tr>
          {columns.map(col => (
            <th key={col.accessor}>{col.header}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {data.map((row, i) => (
          <tr key={i}>
            {columns.map(col => (
              <td key={col.accessor}>{row[col.accessor]}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}