import { Routes, Route } from "react-router-dom";
import { AuthPage } from "./pages/auth/authPage";
import { AuthProvider } from "./contexts/authContext";

import { Toaster } from "@/components/ui/toaster";
import ProtectedRoute from "@/components/routing/protectedRoute";
import { HomePage } from "./pages/home/homePage";

function App() {
  return (
    <AuthProvider>
      <div className="w-[100vw] h-[100vh] p-10">
        <Routes>
          <Route path="/auth" element={<AuthPage />} />
          <Route element={<ProtectedRoute />}>
            <Route path="/" element={<HomePage />} />
          </Route>
        </Routes>
        <Toaster />
      </div>
    </AuthProvider>
  );
}

export default App;
