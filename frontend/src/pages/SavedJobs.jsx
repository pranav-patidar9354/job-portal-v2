import React from "react";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import api from "../services/api";

export default function SavedJobs() {
  const [saved, setSaved] = useState([]);

  useEffect(() => {
    api.get("/candidate/saved-jobs").then(({data}) => setSaved(data.saved_jobs || []));
  }, []);

  return (
    <main className="container">
      <p className="eyebrow">CANDIDATE</p>
      <h2>Saved Jobs</h2>
      <div className="grid">
        {saved.map(item => (
          <Link className="card job-card" to={`/jobs/${item.job.id}`} key={item.id}>
            <span className="tag">{item.job.job_type}</span>
            <h3>{item.job.title}</h3>
            <p className="muted">{item.job.company} · {item.job.location}</p>
            <strong>₹{Number(item.job.salary).toLocaleString("en-IN")}</strong>
          </Link>
        ))}
        {saved.length === 0 && <div className="card empty">No saved jobs.</div>}
      </div>
    </main>
  );
}
