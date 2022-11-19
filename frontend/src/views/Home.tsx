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
import { Currency } from "../components/Currency";
import { GlobalGraph } from "../components/graphs/GlobalGraph";
import { PlayerListItem } from "../components/listItems/PlayerListItem";
import { StonkPositionList } from "../components/listItems/StonkPositionList";
import { Plus } from "../icons";

function Home() {
  return (
    <>
      <Card>
        <h2 className="text-right font-bold text-3xl">
          +200
          <Currency />
        </h2>
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
