import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Button } from "../components/buttons/Button";
import { RouterButton } from "../components/buttons/RouterButton";
import { Card } from "../components/card/Card";
import { Container } from "../components/Container";
import { Currency } from "../components/Currency";
import { StonkGraph } from "../components/graphs/StonkGraph";
import { Spinner } from "../components/spinner/Spinner";
import SvgArrowRight from "../icons/ArrowRight";
import SvgMinus from "../icons/Minus";
import SvgPlus from "../icons/Plus";
import { useStonkState } from "../model/store";
import { getTradeUrl, Routes } from "../router/router";
import { StonkInfo, StonkName } from "../services/vo-stonks";
import SvgEdit from "../icons/Edit";
import { StonkHistoryList } from "../components/listItems/StonkHistoryList";
import { StonksAssetsMatch } from "../assets/StonksAssetsMatch";

const formatter = new Intl.NumberFormat("en-IN", {
  maximumSignificantDigits: 3,
  maximumFractionDigits: 2,
  minimumFractionDigits: 2,
});

export const CurrencyDisplay = (props: { value: number }) => {
  return (
    <span className="inline-flex items-center gap-2">
      {formatter.format(props.value)} <Currency />
    </span>
  );
};

function Detail() {
  const getStonkInfo = useStonkState((state) => state.getStonkInfo);
  const navigate = useNavigate();
  const { stonkName } = useParams();
  const [stonk, setStonk] = useState<StonkInfo | undefined>();

  useEffect(() => {
    console.log("render", stonkName);
    const onError = () => {
      navigate(Routes.Home);
      return;
    };

    // uknown stonk or not set
    if (!stonkName) {
      console.log("error");
      onError();
      return;
    }

    getStonkInfo(stonkName as StonkName)
      .then(({ ret: info, ret_1: err }) => {
        if (err) {
          onError();
        }
        setStonk(info);
      })
      .catch();
  }, []);

  if (!stonk) {
    return <Spinner />;
  }

  const img = StonksAssetsMatch.filter((elem) => elem.stonk === stonkName)?.[0]
    .img;

  return (
    <Container>
      <div className="flex items-center justify-start ">
        {img && <img className="w-36 h-36" src={img} alt={"image"} />}
        <h1>{stonkName}</h1>
      </div>

      <Card className="mx-0">
        <h2 className="text-3xl text-right">
          <CurrencyDisplay
            value={
              stonk.TimeSeries?.[stonk.TimeSeries?.length - 1].Value ?? 0.0
            }
          />
        </h2>
        <StonkGraph stonk={stonk} />
      </Card>
      <Card className="mx-0">
        <h2>Pending Orders</h2>
        <ul>
          {stonk.UserOrders?.map((order, index) => (
            <li className="py-5 border-b-foreground border-b-2" key={index}>
              @{order.UserName} ordered {order.Quantity}
            </li>
          ))}
          {stonk.Orders?.map((order, index) => (
            <li key={index} className="py-5 border-b-foreground border-b-2">
              {order.UserName} {order.Quantity}
            </li>
          ))}
        </ul>
      </Card>
      {stonk.MatchHistory && stonk.MatchHistory.length > 0 && (
        <StonkHistoryList stonk={stonk} />
      )}
      <div className="flex justify-evenly">
        <RouterButton
          className="py-4 px-8"
          route={getTradeUrl(stonkName!, "sell") as Routes}
        >
          <div className="flex items-center gap-4">
            <SvgMinus />
            Sell
          </div>
        </RouterButton>
        <RouterButton
          className="py-4 px-8"
          route={getTradeUrl(stonkName!, "buy") as Routes}
        >
          <div className="flex items-center gap-4">
            <SvgPlus /> Buy
          </div>
        </RouterButton>
      </div>
    </Container>
  );
}

export default Detail;
