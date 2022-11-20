import { Outlet } from "react-router-dom";
import { useStonkState } from "./model/store";
import { Header } from "./components/header/Header";
import { Footer } from "./components/footer/Footer";
import { useEffect } from "react";

function App() {
  const update = useStonkState((state) => state.updateState);

  useEffect(() => {
    update();
  }, []);

  return (
    <div>
      <Header />
      <Outlet />
      <Footer />
    </div>
  );
}

export default App;
