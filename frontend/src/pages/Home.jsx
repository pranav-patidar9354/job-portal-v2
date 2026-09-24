import React from "react";
import { Link } from "react-router-dom";

export default function Home() {
  return (
    <main className="hero">
      <div>
        <p className="eyebrow">JOB PORTAL V2</p>
        <h1>Find work.<br /><span>Build your career.</span></h1>
        <p className="hero-text">A modern platform connecting candidates with recruiters.</p>
        <div className="hero-actions">
          <Link className="button" to="/jobs">Explore Jobs</Link>
          <Link className="button secondary" to="/register">Create Account</Link>
        </div>
      </div>
    </main>
  );
}
