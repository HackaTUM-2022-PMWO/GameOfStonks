import { useState } from "react";
import { useParams } from "react-router-dom";
import { Container } from "../components/Container";
import { Currency } from "../components/Currency";
import { CurrencyDisplay } from "./Detail";

function Trade() {
  const [price, setPrice] = useState(0);
  const { mode, stonkName } = useParams();

  console.log({ mode, stonkName });

  return (
    <Container>
      <h1>
        <CurrencyDisplay value={price} />
      </h1>
    </Container>
  );
}

export default Trade;
