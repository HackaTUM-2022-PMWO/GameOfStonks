import React, { useEffect } from "react";
import { Card } from "../components/card/Card";
import { Container } from "../components/Container";
import { StonkHistoryList } from "../components/listItems/StonkHistoryList";
import { useStonkState } from "../model/store";

function History() {
  const loadHistory = useStonkState((state) => state.getStonksHistory);
  const stonkInfos = useStonkState((state) => state.stonkInfos);

  useEffect(() => {
    loadHistory();
  }, []);

  console.log(stonkInfos);

  return (
    <Container>
      <h1>History</h1>
      {stonkInfos?.map((info) => (
        <React.Fragment key={info.Name}>
          <h2>{info.Name}</h2>
          <StonkHistoryList stonk={info} />
        </React.Fragment>
      ))}
      <Card>
        {(stonkInfos === undefined || stonkInfos.length === 0) && (
          <h2 className="text-center">No history yet</h2>
        )}
      </Card>
    </Container>
  );
}

export default History;
