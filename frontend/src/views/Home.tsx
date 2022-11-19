import {
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import { Card } from "../components/card/Card";
import { GlobalGraph } from "../components/graphs/GlobalGraph";
import { PlayerListItem } from "../components/listItems/PlayerListItem";
import { StonkPositionList } from "../components/listItems/StonkPositionList";

function Home() {
  return (
    <>
      <Card>
        {/* <ResponsiveContainer width="100%" height="100%"> */}
        <GlobalGraph />
        {/* </ResponsiveContainer> */}
      </Card>
      <Card>
        <StonkPositionList stonks={[]} />
      </Card>
    </>
  );
}

export default Home;
