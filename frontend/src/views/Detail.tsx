import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Button } from "../components/buttons/Button";
import { RouterButton } from "../components/buttons/RouterButton";
import { Card } from "../components/card/Card";
import { Container } from "../components/Container";
import { Currency } from "../components/Currency";
import { StonkGraph } from "../components/graphs/StonkGraph";
import { SpinnerOverlay } from "../components/spinner/Spinner";
import SvgMinus from "../icons/Minus";
import SvgPlus from "../icons/Plus";
import { useStonkState } from "../model/store";
import { getTradeUrl, Routes } from "../router/router";
import { Order, StonkInfo, StonkName } from "../services/vo-stonks";
import SvgEdit from "../icons/Edit";
import { StonkHistoryList } from "../components/listItems/StonkHistoryList";
import { StonksAssetsMatch } from "../assets/StonksAssetsMatch";
import { PlayerTag } from "../components/listItems/PlayerListItem";
import SvgTrash from "../icons/Trash";

const formatter = new Intl.NumberFormat("en-US", {
  minimumFractionDigits: 2,
});

export const OrderList = (props: { orders: Order[]; editable?: boolean }) => {
  const { editOrder, deleteOrder } = useStonkState((state) => ({
    editOrder: state.updateOrder,
    deleteOrder: state.deleteOrder,
  }));

  return (
    <ul>
      {props.orders.map((order, index) => (
        <li
          className="py-5 flex gap-2 justify-between items-center border-b-foreground border-b-2 last:border-b-0"
          key={index}
        >
          <span className="flex gap-2 items-center">
            <PlayerTag idx={index} value={order.UserName} />
            ordered <b> {order.Quantity} </b>
            for
            <span>
              {order.Price}
              <Currency size={2} />
              each
            </span>
          </span>
          {props.editable && (
            <span className="flex gap-3 opacity-50">
              <button onClick={() => deleteOrder(order)}>
                <SvgTrash />
              </button>
              <button onClick={() => editOrder(order)}>
                <SvgEdit />
              </button>
            </span>
          )}
        </li>
      ))}
    </ul>
  );
};

export const StonkImage = (props: { size: number; stonk: StonkName }) => {
  const img = StonksAssetsMatch.filter(
    (elem) => elem.stonk === props.stonk
  )?.[0]?.img;

  return (
    <>
      {img && (
        <img width={props.size} height={props.size} src={img} alt={"image"} />
      )}
    </>
  );
};

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
    const onError = () => {
      navigate(Routes.Home);
      return;
    };

    // uknown stonk or not set
    if (!stonkName) {
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
    return <SpinnerOverlay />;
  }
  return (
    <Container>
      <div className="flex ml-7 items-center justify-start ">
        <StonkImage stonk={stonkName as StonkName} size={50} />
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
      <div className="flex items-center gap-5 mx-7 min-w-full">
        <RouterButton
          className="py-4 min-w-full px-8 bg-red-500 text-red-900"
          route={getTradeUrl(stonkName!, "sell") as Routes}
        >
          <div className="flex items-center  gap-5">
            <SvgMinus />
            Sell
          </div>
        </RouterButton>
        <RouterButton
          className="py-4 px-8 min-w-full bg-green-500 text-green-900"
          route={getTradeUrl(stonkName!, "buy") as Routes}
        >
          <div className="flex items-center gap-5">
            <SvgPlus /> Buy
          </div>
        </RouterButton>
      </div>
      {stonk.UserOrders && stonk.UserOrders.length > 0 && (
        <Card headline="Pending Orders" className="mx-0">
          <h2>Your Orders</h2>
          <OrderList orders={stonk.UserOrders} editable />
          <br />
          <br />
          <h2>All Orders</h2>
          <OrderList orders={stonk.Orders ?? []} />
        </Card>
      )}

      {stonk.MatchHistory && stonk.MatchHistory.length > 0 && (
        <StonkHistoryList stonk={stonk} />
      )}
    </Container>
  );
}

export default Detail;
