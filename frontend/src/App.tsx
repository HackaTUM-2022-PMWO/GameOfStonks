import {Routes} from "./router/router";
import {RouterButton} from "./components/RouterButton";
import {Outlet} from "react-router-dom";

function App() {
    return (
        <div>
            <RouterButton route={Routes.Trade} text={"Trade here"}/>
            <Outlet/>
        </div>
    );
}

export default App;
