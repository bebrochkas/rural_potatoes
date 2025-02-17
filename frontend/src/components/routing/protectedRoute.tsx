import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "@/contexts/authContext";

const ProtectedRoute = () => {
  const { token } = useAuth();
  return token ? <Outlet /> : <Navigate to="/auth" />;
};

export default ProtectedRoute;
