import {useEffect, useState} from "react";
import {useNavigate, useParams} from "react-router-dom";
import {Button} from "../components/buttons/Button";
import {RouterButton} from "../components/buttons/RouterButton";
import {Container} from "../components/Container";
import {Currency} from "../components/Currency";
import {Input} from "../components/inputs/Input";
import SvgX from "../icons/X";
import {useStonkState} from "../model/store";
import {getStonkUrl, Routes} from "../router/router";
import {OrderType, StonkName} from "../services/vo-stonks";

function Trade() {
    const [price, setPrice] = useState(22);
    const [qty, setQty] = useState(1);
    const {mode, stonkName} = useParams();
    const navigate = useNavigate();
    const placeOrder = useStonkState((state) => state.placeOrder);

    // sanity check
    useEffect(() => {
        if (!mode || !stonkName) {
            navigate(Routes.Home);
            return;
        }
    }, []);

    return (
        <Container>
            <div className="flex flex-col justify-center items-center gap-4 min-h-screen">
                <div className="flex flex-col items-start gap-5">
                    <h1>{stonkName}</h1>
                    <p className="flex justify-center items-center text-center">
                        <Input
                            label={<span className="text-lg">price</span>}
                            value={price}
                            onChange={(e) =>
                                setPrice(e.currentTarget.valueAsNumber as unknown as number)
                            }
                            className="text-6xl bg-transparent min-w-0 border-l-0 border-t-0 border-r-0 border-b-4 text-white"
                            type="number"
                        />
                        <span className="text-6xl">
            <Currency/>
          </span>
                    </p>
                    <SvgX/>
                    <div className="flex items-center gap-5">
                        <Input
                            label={<span className="text-lg">quantity</span>}
                            value={qty}
                            onChange={(e) =>
                                setQty(e.currentTarget.valueAsNumber as unknown as number)
                            }
                            className="text-6xl bg-transparent min-w-0 border-l-0 border-t-0 border-r-0 border-b-4 text-white"
                            type="number"
                        />
                    </div>
                    <div className="flex gap-4 items-stretch">
                        <RouterButton
                            route={getStonkUrl(stonkName!) as Routes}
                            className="block"
                        >
                            Cancel
                        </RouterButton>
                        <Button
                            onClick={() =>
                                placeOrder({
                                    OrderType: mode as OrderType,
                                    Price: price,
                                    Quantity: qty,
                                    Stonk: stonkName as StonkName,
                                }).finally(() => {
                                    navigate(getStonkUrl(stonkName as StonkName));
                                })
                            }
                            className="block"
                        >
                            {mode === OrderType.Buy ? "Buy" : mode == OrderType.Sell ? "Sell" : "Delete"}
                        </Button>
                    </div>
                </div>
            </div>
        </Container>
    );
}

export default Trade;
