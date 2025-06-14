import React from 'react';
import './Pagination.css';

/**
 * Pagination controls for table data
 * @param {number} currentPage - current page index (1-based)
 * @param {number} totalPages - total number of pages
 * @param {(page: number) => void} onPageChange - callback to change page
 */
export default function Pagination({ currentPage, totalPages, onPageChange }) {
  return (
    <div className="pagination">
      <button
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage <= 1}
      >
        Previous
      </button>
      <span>
        Page {currentPage} of {totalPages}
      </span>
      <button
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage >= totalPages}
      >
        Next
      </button>
    </div>
  );
}