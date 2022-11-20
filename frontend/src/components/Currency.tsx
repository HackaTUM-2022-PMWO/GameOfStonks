import coin from "./../assets/coin.png";

export const Currency = (props: { size?: number }) => {
  return (
    <img
      src={coin}
      alt={"$$"}
      width={props.size}
      height={props.size}
      className="px-2 inline w-10"
    />
  );
};
