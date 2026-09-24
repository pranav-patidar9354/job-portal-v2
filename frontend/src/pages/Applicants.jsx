import React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import api from "../services/api";

export default function Applicants() {
  const { jobID } = useParams();
  const [apps, setApps] = useState([]);

  async function load() {
    const { data } = await api.get(`/recruiter/jobs/${jobID}/applicants`);
    setApps(data.applications || []);
  }

  useEffect(() => { load(); }, [jobID]);

  async function update(id, status) {
    await api.patch(`/recruiter/applications/${id}/status`, {status});
    load();
  }

  return (
    <main className="container">
      <p className="eyebrow">RECRUITER</p>
      <h2>Applicants</h2>
      <div className="stack">
        {apps.map(app => (
          <div className="card applicant" key={app.id}>
            <div>
              <h3>{app.candidate?.name}</h3>
              <p className="muted">{app.candidate?.email}</p>
              {app.cover_letter && <p className="description">{app.cover_letter}</p>}
            </div>
            <select value={app.status} onChange={e=>update(app.id,e.target.value)}>
              <option value="applied">Applied</option>
              <option value="reviewing">Reviewing</option>
              <option value="shortlisted">Shortlisted</option>
              <option value="accepted">Accepted</option>
              <option value="rejected">Rejected</option>
            </select>
          </div>
        ))}
        {apps.length === 0 && <div className="card empty">No applicants yet.</div>}
      </div>
    </main>
  );
}
