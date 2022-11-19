function StartStocks() {
    const item = require('./../assets/paperclip.png') // TODO: replace with dynamic icon image

    return (
        <div className="flex flex-col justify-center items-center h-screen">
            <h1>Your first STONKS:</h1>
            <img className="w-32 h-32" src={item} alt={"item"}/>
            <p>5x Paperclip (PACP)</p>
        </div>
    );
}

export default StartStocks;
