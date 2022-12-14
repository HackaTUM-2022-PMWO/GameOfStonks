import { Outlet } from "react-router-dom";
import { useStonkState } from "./model/store";
import { Header } from "./components/header/Header";
import { Footer } from "./components/footer/Footer";
import { useEffect } from "react";
import { SpinnerOverlay } from "./components/spinner/Spinner";

function App() {
  const isLoading = useStonkState((state) => state.loading);

  return (
    <div>
      <Header />
      <Outlet />
      <Footer />
      {isLoading && <SpinnerOverlay />}
    </div>
  );
}

export default App;
