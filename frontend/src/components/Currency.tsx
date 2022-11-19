export const Currency = () => {
    const coin = require(`./../assets/coin.png`);
    return(
        <img src={coin} alt={"$$"} className="px-2 inline w-10" />
    )
};
