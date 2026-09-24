import React from "react";
import { Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import ProtectedRoute from "./components/ProtectedRoute";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Jobs from "./pages/Jobs";
import JobDetails from "./pages/JobDetails";
import Applications from "./pages/Applications";
import SavedJobs from "./pages/SavedJobs";
import RecruiterDashboard from "./pages/RecruiterDashboard";
import Applicants from "./pages/Applicants";

export default function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/jobs" element={<Jobs />} />
        <Route path="/jobs/:id" element={<JobDetails />} />

        <Route path="/applications" element={
          <ProtectedRoute role="candidate"><Applications /></ProtectedRoute>
        } />

        <Route path="/saved" element={
          <ProtectedRoute role="candidate"><SavedJobs /></ProtectedRoute>
        } />

        <Route path="/recruiter" element={
          <ProtectedRoute role="recruiter"><RecruiterDashboard /></ProtectedRoute>
        } />

        <Route path="/recruiter/jobs/:jobID/applicants" element={
          <ProtectedRoute role="recruiter"><Applicants /></ProtectedRoute>
        } />
      </Routes>
    </>
  );
}
