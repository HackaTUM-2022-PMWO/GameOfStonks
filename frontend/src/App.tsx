import {Outlet} from "react-router-dom";
import {useStonkState} from "./model/store";
import {Header} from "./components/header/Header";
import {Footer} from "./components/footer/Footer";

function App() {
    const state = useStonkState();

    return (
        <div>
            <Header/>
            <Outlet/>
            <Footer />
        </div>
    );
}

export default App;
