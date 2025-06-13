import { useEffect, useState } from "react";
import DataTable from './DataTable'

export default function Dashboard() {
    return(
        <div style={{ padding: "20px" }}>
            <h2>Dashboard</h2>
            <DataTable/>
        </div>
    );
}