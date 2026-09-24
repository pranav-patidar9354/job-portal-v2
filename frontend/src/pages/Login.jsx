import React from "react";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export default function Login() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState("");

  async function submit(e) {
    e.preventDefault();
    setError("");
    try {
      const data = await login(form.email, form.password);
      navigate(data.user.role === "recruiter" ? "/recruiter" : "/jobs");
    } catch (err) {
      setError(err.response?.data?.error || "Login failed");
    }
  }

  return (
    <main className="auth-page">
      <form className="card auth-card" onSubmit={submit}>
        <h2>Welcome back</h2>
        <p className="muted">Login to your account.</p>
        {error && <div className="error">{error}</div>}
        <input type="email" placeholder="Email" value={form.email} onChange={e => setForm({...form,email:e.target.value})} required />
        <input type="password" placeholder="Password" value={form.password} onChange={e => setForm({...form,password:e.target.value})} required />
        <button className="button" type="submit">Login</button>
        <p className="muted">New here? <Link to="/register">Create an account</Link></p>
      </form>
    </main>
  );
}
