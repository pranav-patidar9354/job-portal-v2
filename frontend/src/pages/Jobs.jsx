import React from "react";
import { useEffect, useState } from "react";
import { Link, useSearchParams } from "react-router-dom";
import api from "../services/api";

export default function Jobs() {
  const [params, setParams] = useSearchParams();
  const [jobs, setJobs] = useState([]);
  const [pagination, setPagination] = useState({});
  const [loading, setLoading] = useState(true);

  const [form, setForm] = useState({
    search: params.get("search") || "",
    location: params.get("location") || "",
    company: params.get("company") || "",
    sort: params.get("sort") || "newest",
  });

  async function load() {
    setLoading(true);
    try {
      const { data } = await api.get("/jobs", {
        params: Object.fromEntries(params),
      });
      setJobs(data.jobs || []);
      setPagination(data.pagination || {});
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => { load(); }, [params.toString()]);

  function search(e) {
    e.preventDefault();
    const next = {};
    Object.entries(form).forEach(([key, value]) => {
      if (value) next[key] = value;
    });
    next.page = "1";
    next.limit = "9";
    setParams(next);
  }

  function changePage(page) {
    const next = Object.fromEntries(params);
    next.page = String(page);
    setParams(next);
  }

  return (
    <main className="container">
      <div className="page-head">
        <p className="eyebrow">OPPORTUNITIES</p>
        <h2>Find your next job</h2>
      </div>

      <form className="search-bar card" onSubmit={search}>
        <input placeholder="Search title or description" value={form.search} onChange={e=>setForm({...form,search:e.target.value})} />
        <input placeholder="Location" value={form.location} onChange={e=>setForm({...form,location:e.target.value})} />
        <input placeholder="Company" value={form.company} onChange={e=>setForm({...form,company:e.target.value})} />
        <select value={form.sort} onChange={e=>setForm({...form,sort:e.target.value})}>
          <option value="newest">Newest</option>
          <option value="salary_desc">Highest salary</option>
          <option value="salary_asc">Lowest salary</option>
          <option value="oldest">Oldest</option>
        </select>
        <button className="button">Search</button>
      </form>

      {loading ? <p>Loading jobs...</p> : (
        <div className="grid">
          {jobs.map(job => (
            <Link className="card job-card" key={job.id} to={`/jobs/${job.id}`}>
              <span className="tag">{job.job_type}</span>
              <h3>{job.title}</h3>
              <p className="muted">{job.company} · {job.location}</p>
              <strong>₹{Number(job.salary).toLocaleString("en-IN")}</strong>
              <p className="muted clamp">{job.description}</p>
            </Link>
          ))}
        </div>
      )}

      {!loading && jobs.length === 0 && <div className="card empty">No jobs found.</div>}

      {pagination.total_pages > 1 && (
        <div className="pagination">
          <button disabled={pagination.page <= 1} onClick={()=>changePage(pagination.page-1)}>Previous</button>
          <span>Page {pagination.page} of {pagination.total_pages}</span>
          <button disabled={pagination.page >= pagination.total_pages} onClick={()=>changePage(pagination.page+1)}>Next</button>
        </div>
      )}
    </main>
  );
}
