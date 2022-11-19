import { Line, LineChart, Tooltip } from "recharts";
import { StonkInfo, User } from "../../services/vo-stonks";

// GraphComponent used to represent progress of all users
export type GraphProps = {
  stonk: StonkInfo;
};

export const StonkGraph = (props: GraphProps) => {
  // TODO: use global state for all users here

  return (
    <LineChart width={500} height={300} data={[]}>
      <Tooltip />
      <Line type="monotone" dataKey="pv" stroke="#fff" strokeWidth={5} />
    </LineChart>
  );
};
