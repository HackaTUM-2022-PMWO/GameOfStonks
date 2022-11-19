import { BarChart, Line, LineChart, Tooltip } from "recharts";
import { useStonkState } from "../../model/store";
import { User } from "../../services/vo-stonks";
import { PlayerListItem } from "../listItems/PlayerListItem";

// GraphComponent used to represent progress of all users
export type GraphProps = {};

// FIXME: @botastic  with actual data
const data = [
  {
    name: "Page A",
    uv: 4000,
    pv: 2400,
    amt: 2400,
  },
  {
    name: "Page B",
    uv: 3000,
    pv: 1398,
    amt: 2210,
  },
  {
    name: "Page C",
    uv: 2000,
    pv: 9800,
    amt: 2290,
  },
  {
    name: "Page D",
    uv: 2780,
    pv: 3908,
    amt: 2000,
  },
  {
    name: "Page E",
    uv: 1890,
    pv: 4800,
    amt: 2181,
  },
  {
    name: "Page F",
    uv: 2390,
    pv: 3800,
    amt: 2500,
  },
  {
    name: "Page G",
    uv: 3490,
    pv: 4300,
    amt: 2100,
  },
];

export const GlobalGraph = (props: GraphProps) => {
  const users = useStonkState((state) => state.sessionUsers);
  // TODO: use global state for all users here

  return (
    <>
      <LineChart data={data}>
        <Tooltip />
        <Line type="monotone" dataKey="uv" stroke="#D043AC" strokeWidth={5} />
        <Line type="monotone" dataKey="pv" stroke="#fff" strokeWidth={5} />
      </LineChart>
      <ul className="flex">
        {users?.map((user, index) => (
          <PlayerListItem
            idx={index}
            key={user.Name + index}
            value={user.Name}
          />
        ))}
      </ul>
    </>
  );
};
