import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { StonkGraph } from "../components/graphs/StonkGraph";
import { Spinner } from "../components/spinner/Spinner";
import { useStonkState } from "../model/store";
import { Routes } from "../router/router";
import { StonkInfo } from "../services/vo-stonks";

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
      console.log("error");
      // onError();
      return;
    }

    getStonkInfo(stonkName)
      .then(({ ret: info, ret_1: err }) => {
        if (err) {
          onError();
        }
        setStonk(info);
      })
      .catch();
  }, [navigate, setStonk, getStonkInfo, stonkName]);

  if (!stonk) {
    return <Spinner />;
  }

  return (
    <div>
      <h2>{stonk.ID}</h2>
      <StonkGraph stonk={stonk} />
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
