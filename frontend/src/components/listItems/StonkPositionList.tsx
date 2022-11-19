import { StonkInfo } from "../../services/vo-stonks";

export type StonkPositionListProps = { stonks: StonkInfo[] };

// TODO: verify stonk position datatype with @bosastic
export const StonkPositionList = (props: StonkPositionListProps) => {
  return (
    <>
      <h2>Stonks</h2>
      <ul className="list-none">
        {props.stonks.map((stonk) => (
          <li key={stonk.ID} className="p-4 border-t-1">
            {stonk.ID}
          </li>
        ))}
      </ul>
    </>
  );
};
