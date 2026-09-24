import React, { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import api from "../services/api";
import { useAuth } from "../context/AuthContext";

export default function JobDetails() {
  const { id } = useParams();
  const { user } = useAuth();

  const [job, setJob] = useState(null);
  const [message, setMessage] = useState("");
  const [coverLetter, setCoverLetter] = useState("");
  const [isSaved, setIsSaved] = useState(false);

  useEffect(() => {
    loadJob();
  }, [id]);

  useEffect(() => {
    if (user?.role === "candidate") {
      checkSavedStatus();
    }
  }, [id, user]);

  async function loadJob() {
    try {
      const { data } = await api.get(`/jobs/${id}`);
      setJob(data.job);
    } catch (err) {
      setMessage(err.response?.data?.error || "Could not load job.");
    }
  }

  async function checkSavedStatus() {
  try {
    const { data } = await api.get("/candidate/saved-jobs");

    const saved = data.saved_jobs?.some(
      (savedJob) => String(savedJob.job_id) === String(id)
    );

    setIsSaved(saved);
  } catch (err) {
    console.log("Could not check saved status.");
  }
}

  async function apply() {
    try {
      await api.post(`/candidate/jobs/${id}/apply`, {
        cover_letter: coverLetter,
      });

      setMessage("Application submitted successfully.");
    } catch (err) {
      setMessage(err.response?.data?.error || "Could not apply.");
    }
  }

  async function toggleSave() {
    try {
      const { data } = await api.post(`/candidate/jobs/${id}/save`);

      setIsSaved(data.saved);

      if (data.saved) {
        setMessage("Job saved.");
      } else {
        setMessage("Job removed from saved jobs.");
      }
    } catch (err) {
      setMessage(
        err.response?.data?.error || "Could not update saved job."
      );
    }
  }

  if (!job) {
    return (
      <main className="container">
        <p>Loading...</p>
      </main>
    );
  }

  return (
    <main className="container narrow">
      <Link to="/jobs" className="back">
        ← Back to jobs
      </Link>

      <article className="card details">
        <span className="tag">{job.job_type}</span>

        <h1>{job.title}</h1>

        <p className="subtitle">
          {job.company} · {job.location}
        </p>

        <h3>
          ₹{Number(job.salary).toLocaleString("en-IN")}
        </h3>

        <hr />

        <h3>Description</h3>

        <p className="description">
          {job.description}
        </p>

        {message && <div className="success">{message}</div>}

        {user?.role === "candidate" && (
          <div className="apply-box">
            <textarea
              placeholder="Optional cover letter"
              value={coverLetter}
              onChange={(e) => setCoverLetter(e.target.value)}
            />

            <div className="row">
              <button className="button" onClick={apply}>
                Apply Now
              </button>

              <button
                className={`button ${isSaved ? "" : "secondary"}`}
                onClick={toggleSave}
              >
                {isSaved ? "Unsave Job" : "Save Job"}
              </button>
            </div>
          </div>
        )}

        {!user && (
          <p className="muted">
            Please <Link to="/login">login</Link> as a candidate to apply.
          </p>
        )}
      </article>
    </main>
  );
}