import { NONAME } from "dns";
import { Line, LineChart, Tooltip } from "recharts";
import { StonkInfo, User } from "../../services/vo-stonks";

// GraphComponent used to represent progress of all users
export type GraphProps = {
  stonk: StonkInfo;
};

export const StonkGraph = (props: GraphProps) => {
  // TODO: use global state for all users here

  return (
    <LineChart
      width={500}
      height={300}
      data={
        props.stonk.TimeSeries?.map((s) => ({
          time: s.Time,
          value: s.Value,
        })) ?? []
      }
    >
      <Tooltip
        contentStyle={{
          background: "#250049",
          border: 0,
          outline: 0,
        }}
      />
      <Line type="monotone" dataKey="value" stroke="#fff" strokeWidth={5} />
    </LineChart>
  );
};
