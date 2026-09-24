import React from "react";
import { useEffect, useState } from "react";
import api from "../services/api";

export default function Applications() {
  const [apps, setApps] = useState([]);

  async function load() {
    const { data } = await api.get("/candidate/applications");
    setApps(data.applications || []);
  }

  useEffect(() => { load(); }, []);

  async function withdraw(id) {
    await api.delete(`/candidate/applications/${id}`);
    load();
  }

  return (
    <main className="container">
      <p className="eyebrow">CANDIDATE</p>
      <h2>My Applications</h2>
      <div className="stack">
        {apps.map(app => (
          <div className="card application" key={app.id}>
            <div>
              <h3>{app.job?.title}</h3>
              <p className="muted">{app.job?.company} · {app.job?.location}</p>
            </div>
            <span className={`status ${app.status}`}>{app.status}</span>
            <button className="button danger small" onClick={()=>withdraw(app.id)}>Withdraw</button>
          </div>
        ))}
        {apps.length === 0 && <div className="card empty">You haven't applied to any jobs yet.</div>}
      </div>
    </main>
  );
}
