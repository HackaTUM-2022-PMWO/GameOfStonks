import {Card} from "../card/Card";
import SvgArrowRight from "../../icons/ArrowRight";
import {StonkInfo} from "../../services/vo-stonks";

export function StonkHistoryList(props: {stonk : StonkInfo}){
    return (
        <Card className="mx-0">
            <h2>History</h2>
            <ul>
                {props.stonk.MatchHistory?.map((order, index) => (
                    <li key={index} className="flex ">
                        {order.UserSell} <SvgArrowRight /> {order.UserSell}
                    </li>
                ))}
            </ul>
        </Card>
    )
}