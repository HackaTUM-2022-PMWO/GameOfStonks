import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Spinner } from "../components/spinner/Spinner";
import { useStonkState } from "../model/store";
import { Routes } from "../router/router";
import { StonkInfo } from "../services/vo-stonks";

function Detail() {
  const store = useStonkState();
  const navigate = useNavigate();
  let { stonkName } = useParams();
  const [stonk, setStonk] = useState<StonkInfo | undefined>();

  useEffect(() => {
    // uknown stonk or not set
    if (!stonkName) {
      navigate(Routes.Home);
      return;
    }

    store.getStonkInfo(stonkName).then(({ ret: info, ret_1: err }) => {
      if (err) {
        navigate(Routes.Home);
        return;
      }
      setStonk(info);
    });
  });

  if (!stonk) {
    return <Spinner />;
  }

  return (
    <div>
      {stonk?.ID}
      <ul>
        {stonk.Orders?.map((order, index) => (
          <li key={index}>
            {order.User} {order.Quantity}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Detail;
