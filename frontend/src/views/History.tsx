import React, { useEffect } from "react";
import { useStonkState } from "../model/store";

function History() {
  const loadHistory = useStonkState((state) => state.getStonksHistory);
  const stonkInfos = useStonkState((state) => state.stonkInfos);

  useEffect(() => {
    loadHistory();
  }, []);

  console.log(stonkInfos);

  return (
    <div>
      History
      {stonkInfos?.map((info) => (
        <React.Fragment key={info.Name}>
          <h2>{info.Name}</h2>
          <ul>
            {info.MatchHistory?.map((m) => (
              <li key={m.TimeStamp}>
                {m.UserBuy} {m.UserSell} {m.Quantity}
              </li>
            ))}
          </ul>
        </React.Fragment>
      ))}
    </div>
  );
}

export default History;
