import { Outlet } from "react-router-dom";
import { useStonkState } from "./model/store";

function App() {
  const state = useStonkState();

  return (
    <div>
      <Outlet />
    </div>
  );
}

export default App;
