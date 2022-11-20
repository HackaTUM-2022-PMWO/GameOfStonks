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
import { useStonkState } from "../model/store";

function Home() {
  const user = useStonkState((state) => state.currentUser);
  const stonks = useStonkState((state) => state.stonkInfos);

  return (
    <>
      <Card>
        <h2 className="text-right font-bold text-3xl">
          {user?.NetWorth}
          <Currency />
        </h2>
        <div className="max-w max-h">
          {/* <ResponsiveContainer width="100%" height={300}> */}
          <GlobalGraph />
          {/* </ResponsiveContainer> */}
        </div>
      </Card>
      <Card>
        <StonkPositionList stonks={user?.Stonks ?? ({} as any)} />
      </Card>
      <Card>
        <StonkPositionList stonks={user?.Stonks ?? ({} as any)} />
      </Card>
    </>
  );
}

export default Home;
