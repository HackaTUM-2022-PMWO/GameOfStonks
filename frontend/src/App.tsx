import {Outlet} from "react-router-dom";
import {useStonkState} from "./model/store";
import {Header} from "./components/header/Header";

function App() {
    const state = useStonkState();

    return (
        <div>
            <Header/>
            <Outlet/>
        </div>
    );
}

export default App;
