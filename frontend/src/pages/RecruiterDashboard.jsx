import React from "react";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import api from "../services/api";

const empty = {
  title:"", description:"", company:"", location:"",
  salary:"", job_type:"full-time"
};

export default function RecruiterDashboard() {
  const [jobs, setJobs] = useState([]);
  const [form, setForm] = useState(empty);
  const [editing, setEditing] = useState(null);
  const [message, setMessage] = useState("");

  async function load() {
    const { data } = await api.get("/recruiter/jobs");
    setJobs(data.jobs || []);
  }

  useEffect(() => { load(); }, []);

  async function submit(e) {
    e.preventDefault();
    setMessage("");
    try {
      const payload = {...form, salary:Number(form.salary)};
      if (editing) {
        await api.put(`/recruiter/jobs/${editing}`, payload);
      } else {
        await api.post("/recruiter/jobs", payload);
      }
      setForm(empty);
      setEditing(null);
      setMessage("Saved successfully.");
      load();
    } catch (err) {
      setMessage(err.response?.data?.error || "Could not save job.");
    }
  }

  function edit(job) {
    setEditing(job.id);
    setForm({
      title:job.title, description:job.description, company:job.company,
      location:job.location, salary:job.salary, job_type:job.job_type
    });
    window.scrollTo({top:0, behavior:"smooth"});
  }

  async function remove(id) {
    if (!window.confirm("Delete this job?")) return;
    await api.delete(`/recruiter/jobs/${id}`);
    load();
  }

  return (
    <main className="container">
      <p className="eyebrow">RECRUITER</p>
      <h2>Recruiter Dashboard</h2>

      <form className="card form-card" onSubmit={submit}>
        <h3>{editing ? "Edit Job" : "Create Job"}</h3>
        {message && <div className="success">{message}</div>}

        <div className="form-grid">
          <input placeholder="Job title" value={form.title} onChange={e=>setForm({...form,title:e.target.value})} required />
          <input placeholder="Company" value={form.company} onChange={e=>setForm({...form,company:e.target.value})} required />
          <input placeholder="Location" value={form.location} onChange={e=>setForm({...form,location:e.target.value})} required />
          <input type="number" placeholder="Salary" value={form.salary} onChange={e=>setForm({...form,salary:e.target.value})} required />
          <select value={form.job_type} onChange={e=>setForm({...form,job_type:e.target.value})}>
            <option value="full-time">Full-time</option>
            <option value="part-time">Part-time</option>
            <option value="internship">Internship</option>
            <option value="contract">Contract</option>
          </select>
        </div>

        <textarea placeholder="Job description" value={form.description} onChange={e=>setForm({...form,description:e.target.value})} required />

        <div className="row">
          <button className="button">{editing ? "Update Job" : "Create Job"}</button>
          {editing && (
            <button type="button" className="button secondary" onClick={()=>{setEditing(null);setForm(empty);}}>
              Cancel
            </button>
          )}
        </div>
      </form>

      <h3>Your Jobs</h3>
      <div className="stack">
        {jobs.map(job => (
          <div className="card application" key={job.id}>
            <div>
              <h3>{job.title}</h3>
              <p className="muted">{job.company} · {job.location}</p>
            </div>
            <div className="row">
              <Link className="button small secondary" to={`/recruiter/jobs/${job.id}/applicants`}>Applicants</Link>
              <button className="button small secondary" onClick={()=>edit(job)}>Edit</button>
              <button className="button small danger" onClick={()=>remove(job.id)}>Delete</button>
            </div>
          </div>
        ))}
      </div>
    </main>
  );
}
