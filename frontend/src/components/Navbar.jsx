import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export default function Navbar() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  return (
    <nav className="nav">
      <Link className="brand" to="/">JobPortal<span>V2</span></Link>
      <div className="nav-links">
        <Link to="/jobs">Jobs</Link>
        {user?.role === "candidate" && <Link to="/applications">Applications</Link>}
        {user?.role === "candidate" && <Link to="/saved">Saved</Link>}
        {user?.role === "recruiter" && <Link to="/recruiter">Dashboard</Link>}
        {user ? (
          <button className="button small" onClick={() => { logout(); navigate("/login"); }}>Logout</button>
        ) : (
          <>
            <Link to="/login">Login</Link>
            <Link className="button small" to="/register">Register</Link>
          </>
        )}
      </div>
    </nav>
  );
}
